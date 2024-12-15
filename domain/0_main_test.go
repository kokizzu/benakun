package domain

import (
	"context"
	"database/sql"
	"fmt"
	"net"
	"os"
	"testing"
	"time"

	"github.com/kokizzu/gotro/D"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/id64"
	"github.com/kokizzu/lexid"
	"github.com/kpango/fastime"
	"github.com/ory/dockertest/v3"
	"github.com/tarantool/go-tarantool/v2"
	"golang.org/x/sync/errgroup"

	"benakun/conf"
	"benakun/model"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/xMailer"
)

// create dockertest instance

var testTt *Tt.Adapter
var testCh *Ch.Adapter
var testMailer xMailer.Mailer
var testTime = fastime.Now()
var testSuperAdminSessionToken string
var testAdmin *rqAuth.Users

const (
	testSuperAdminEmail      = `admin@localhost`
	testSuperAdminPassword   = `adminPass`
	testSuperAdminUserName   = `admin1`
	testSuperAdminTenantCode = `admin-1234`
)

func TestMain(m *testing.M) {
	if os.Getenv(`USE_COMPOSE`) != `` {
		// use local docker compose
		conf.LoadEnv()

		var err error
		eg := errgroup.Group{}
		eg.Go(func() error {
			chConf := conf.EnvClickhouse()
			testCh, err = chConf.Connect()
			return err
		})
		eg.Go(func() error {
			ttConf := conf.EnvTarantool()
			testTt, err = ttConf.Connect()
			return err
		})
		eg.Go(func() error {
			mhConf := conf.EnvMailhog()
			mailer, err := xMailer.NewMailhog(mhConf)
			testMailer = xMailer.Mailer{
				SendMailFunc: mailer.SendEmail,
			}
			return err
		})
		err = eg.Wait()
		L.PanicIf(err, `eg.Wait`)

	} else {
		// setup dockertest instance
		dockerPool := D.InitDockerTest("")
		dockerPool.Pool.MaxWait = 5 * time.Minute
		defer dockerPool.Cleanup()

		eg := errgroup.Group{}

		// attach tarantool
		eg.Go(func() error {
			tdt := &Tt.TtDockerTest{
				User:     "testT",
				Password: "passT",
			}
			img := tdt.ImageVersion(dockerPool, ``)
			dockerPool.Spawn(img, func(res *dockertest.Resource) error {
				t, err := tdt.ConnectCheck(res)
				time.Sleep(1 * time.Second)
				if err != nil {
					return err
				}
				testTt = &Tt.Adapter{
					Connection: t,
					Reconnect: func() *tarantool.Connection {
						t, err := tdt.ConnectCheck(res)
						L.IsError(err, `tdt.ConnectCheck`)
						return t
					},
				}
				return nil
			})
			return nil
		})

		// attach clickhouse
		eg.Go(func() error {
			cdt := &Ch.ChDockerTest{
				User:     "testC",
				Password: "passC",
				Database: "default",
			}
			img := cdt.ImageLatest(dockerPool)
			dockerPool.Spawn(img, func(res *dockertest.Resource) error {
				c, err := cdt.ConnectCheck(res)
				time.Sleep(1 * time.Second)
				if err != nil {
					return err
				}
				testCh = &Ch.Adapter{
					DB: c,
					Reconnect: func() *sql.DB {
						c, err := cdt.ConnectCheck(res)
						L.IsError(err, `cdt.ConnectCheck`)
						return c
					},
				}
				return nil
			})
			return nil
		})

		// mailer
		eg.Go(func() error {
			mailhogConf := conf.MailhogConf{
				MailhogHost: `localhost`,
				MailhogPort: 1025,
			}
			mailhogPort := fmt.Sprint(mailhogConf.MailhogPort)
			dockerPool.Spawn(&dockertest.RunOptions{
				Name:       `dockertest-mailhog-` + dockerPool.Uniq,
				Repository: "mailhog/mailhog",
				Tag:        `latest`,
				NetworkID:  dockerPool.Network.ID,
			}, func(res *dockertest.Resource) error {
				_, err := net.Dial("tcp", res.GetHostPort(mailhogPort+"/tcp"))
				if err != nil {
					return err
				}
				mailHog, err := xMailer.NewMailhog(mailhogConf)
				L.PanicIf(err, `xMailer.NewMailhog`)
				testMailer.SendMailFunc = mailHog.SendEmail
				return nil
			})
			return nil
		})

		err := eg.Wait()
		L.PanicIf(err, `eg.Wait`)
	}

	// run migration
	model.RunMigration(testTt, testCh, testTt, testTt, testTt)

	// run tests
	m.Run()

	// teardown dockertest instance
}

func testDomain() (*Domain, func()) {
	log := conf.InitLogger()

	d := &Domain{
		AuthOltp: testTt,
		AuthOlap: testCh,

		StorOltp: testTt,

		Mailer:  xMailer.Mailer{SendMailFunc: testMailer.SendMailFunc},
		IsBgSvc: false,

		Log: log,

		Superadmins: M.SB{testSuperAdminEmail: true},
	}
	d.InitTimedBuffer()

	// create admin
	admin := wcAuth.NewUsersMutator(testTt)
	admin.Email = testSuperAdminEmail
	admin.SetEncryptedPassword(testSuperAdminPassword, fastime.UnixNow())
	admin.SecretCode = id64.SID() + S.RandomCB63(1)
	admin.TenantCode = testSuperAdminTenantCode
	if !admin.FindByEmail() {
		admin.DoInsert()
	}
	testAdmin = &admin.Users
	testAdmin.Adapter = nil // prevent modification

	// create tenant admin
	tenant := wcAuth.NewTenantsMutator(testTt)
	tenant.TenantCode = testSuperAdminTenantCode
	if !tenant.FindByTenantCode() {
		tenant.DoInsert()
	}

	// create session
	session := wcAuth.NewSessionsMutator(testTt)
	session.UserId = admin.Id
	sess := &Session{
		UserId:        admin.Id,
		ExpiredAt:     testTime.AddDate(0, 0, conf.CookieDays).Unix(),
		Email:         admin.Email,
		TenantCode:    testSuperAdminTenantCode,
		IsTenantAdmin: true,
	}
	testSuperAdminSessionToken = sess.Encrypt(``) // empty user agent to simplify testing
	session.SessionToken = testSuperAdminSessionToken
	session.ExpiredAt = sess.ExpiredAt
	if !session.FindBySessionToken() {
		session.DoInsert()
	}

	return d, func() {
		go d.authLogs.Close()
		d.WaitTimedBufferFinalFlush()
	}
}

func testAdminRequestCommon(action string) RequestCommon {
	return RequestCommon{
		TracerContext: context.Background(),
		RequestId:     lexid.ID(),
		SessionToken:  testSuperAdminSessionToken,
		UserAgent:     "",
		IpAddress:     "127.0.2.1",
		Debug:         true,
		Host:          "localhost:1234",
		Action:        action,
		Lat:           -1,
		Long:          -2,
		now:           testTime.Unix(),
		start:         testTime,
	}
}

func testGuestRequestCommon(action string) RequestCommon {
	return RequestCommon{
		TracerContext: context.Background(),
		RequestId:     lexid.ID(),
		SessionToken:  "",
		UserAgent:     "",
		IpAddress:     "127.0.2.1",
		Debug:         true,
		Host:          "localhost:1234",
		Action:        action,
		Lat:           -1,
		Long:          -2,
		now:           testTime.Unix(),
		start:         testTime,
	}
}
