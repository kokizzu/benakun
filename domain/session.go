package domain

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
	"strings"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/lexid"
	"github.com/kpango/fastime"
	"github.com/mojura/enkodo"
	"github.com/segmentio/fasthash/fnv1a"
	"github.com/zeebo/xxh3"

	"benakun/conf"
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

type Session struct {
	UserId     uint64
	ExpiredAt  int64 // in seconds
	Email      string
	TenantCode string
	Role       string

	// not saved but retrieved from SUPERADMIN_EMAILS env
	IsSuperAdmin      bool
	IsTenantAdmin     bool
	IsDataEntry       bool
	IsReportViewer    bool
	IsFieldSupervisor bool

	Segments M.SB
}

// list of first segment of url path, if empty then only /guest segment
const (
	SuperAdminSegment      = `superAdmin`
	TenantAdminSegment     = `tenantAdmin`
	DataEntrySegment       = `dataEntry`
	ReportViewerSegment    = `reportViewer`
	FieldSupervisorSegment = `fieldSupervisor`
	GuestSegment           = `guest` // any user that not yet login
	UserSegment            = `user`  // any user that already login
)

func (s *Session) MarshalEnkodo(enc *enkodo.Encoder) (err error) {
	_ = enc.Uint64(s.UserId)
	_ = enc.Int64(s.ExpiredAt)
	_ = enc.String(s.Email)
	return
}

func (s *Session) UnmarshalEnkodo(dec *enkodo.Decoder) (err error) {
	if s.UserId, err = dec.Uint64(); err != nil {
		return
	}
	if s.ExpiredAt, err = dec.Int64(); err != nil {
		return
	}
	if s.Email, err = dec.String(); err != nil {
		return
	}
	return
}

func createHash(key1, key2 string) string {
	res := xxh3.HashString128(key1 + conf.PROJECT_NAME + key2) // PROJECT_NAME = salt, if you change this, all token will be invalidated
	const x = 256
	return string([]byte{
		byte(res.Hi >> (64 - 8) % x),
		byte(res.Hi >> (64 - 16) % x),
		byte(res.Hi >> (64 - 24) % x),
		byte(res.Hi >> (64 - 32) % x),
		byte(res.Hi >> (64 - 40) % x),
		byte(res.Hi >> (64 - 48) % x),
		byte(res.Hi >> (64 - 56) % x),
		byte(res.Hi >> (64 - 64) % x), // nolint: staticcheck
		byte(res.Lo >> (64 - 8) % x),
		byte(res.Lo >> (64 - 16) % x),
		byte(res.Lo >> (64 - 24) % x),
		byte(res.Lo >> (64 - 32) % x),
		byte(res.Lo >> (64 - 40) % x),
		byte(res.Lo >> (64 - 48) % x),
		byte(res.Lo >> (64 - 56) % x),
		byte(res.Lo >> (64 - 64) % x), // nolint: staticcheck
	})
}

const TokenSeparator = `|`

func (s *Session) Encrypt(userAgent string) string {
	key1 := lexid.NanoID()
	key2 := S.EncodeCB63(fnv1a.HashString64(userAgent), 1)
	block, err := aes.NewCipher([]byte(createHash(key1, key2)))
	L.PanicIf(err, `aes.NewCipher`)
	gcm, err := cipher.NewGCM(block)
	L.PanicIf(err, `cipher.NewGCM`)
	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	L.PanicIf(err, `io.ReadFull`)
	buffer := bytes.NewBuffer(nil)
	w := enkodo.NewWriter(buffer)
	err = w.Encode(s)
	L.PanicIf(err, `w.Encode`)
	cipherText := gcm.Seal(nonce, nonce, buffer.Bytes(), nil)
	return key1 + TokenSeparator + hex.EncodeToString(cipherText) + TokenSeparator + key2
}

func (s *Session) Decrypt(sessionToken, userAgent string) bool {
	strs := strings.Split(sessionToken, TokenSeparator)
	tokenLen := len(strs)
	if tokenLen != 3 {
		L.Print(`sessionToken length mismatch: ` + I.ToStr(tokenLen) + ` value: ` + sessionToken)
		return false
	}
	uaHash := S.EncodeCB63(fnv1a.HashString64(userAgent), 1)
	if strs[2] != uaHash {
		L.Print(`userAgent mismatch: ` + strs[2])
		return false
	}
	data, err := hex.DecodeString(strs[1])
	if L.IsError(err, `hex.DecodeString`) {
		return false
	}
	key := []byte(createHash(strs[0], strs[2]))
	block, err := aes.NewCipher(key)
	if L.IsError(err, `aes.NewCipher`) {
		return false
	}
	gcm, err := cipher.NewGCM(block)
	if L.IsError(err, `cipher.NewGCM`) {
		return false
	}
	nonceSize := gcm.NonceSize()
	if len(data) < nonceSize {
		L.Print(`len(data) < nonceSize`)
		return false
	}
	nonce, cipherText := data[:nonceSize], data[nonceSize:]
	plainText, err := gcm.Open(nil, nonce, cipherText, nil)
	if L.IsError(err, `gcm.Open`) {
		return false
	}
	err = enkodo.Unmarshal(plainText, s)
	return !L.IsError(err, `enkodo.Unmarshal`)
}

func (d *Domain) CreateSession(userId uint64, email, userAgent, ip string) (*wcAuth.SessionsMutator, *Session) {
	session := wcAuth.NewSessionsMutator(d.AuthOltp)
	session.UserId = userId
	session.Device = userAgent
	session.LoginIPs = ip
	sess := &Session{
		UserId:    userId,
		ExpiredAt: fastime.Now().AddDate(0, 0, conf.CookieDays).Unix(),
		Email:     email,
	}
	session.SessionToken = sess.Encrypt(userAgent)
	session.ExpiredAt = sess.ExpiredAt
	d.segmentsFromSession(sess)
	return session, sess
}

func (d *Domain) ExpireSession(token string, out *ResponseCommon) int64 {
	session := wcAuth.NewSessionsMutator(d.AuthOltp)
	session.SessionToken = token
	now := fastime.UnixNow()
	if session.FindBySessionToken() {
		out.SessionToken = conf.CookieLogoutValue
		if session.ExpiredAt > now {
			session.SetExpiredAt(now)
			if !session.DoUpdateBySessionToken() {
				out.SetError(500, ErrUserSessionRemovalFailed)
				return 0
			}
			return now
		}
		return session.ExpiredAt
	}
	return 0
}

const (
	ErrSessionTokenEmpty     = `sessionToken empty`
	ErrSessionTokenInvalid   = `sessionToken invalid`
	ErrSessionTokenExpired   = `sessionToken expired`
	ErrSessionTokenNotFound  = `sessionToken not found`
	ErrSessionTokenLoggedOut = `sessionToken already logged out`

	ErrSegmentNotAllowed = `session segment not allowed`

	ErrSessionUserNotSuperAdmin      = `session email is not superadmin`
	ErrSessionUserNotTenantAdmin     = `session user is not tenant admin`
	ErrSessionUserNotDataEntry       = `session user is not data entry`
	ErrSessionUserNotFieldSupervisor = `session user is not field supervisor`
	ErrSessionUserNotReportViewer    = `session user is not report viewer`
)

func (d *Domain) MustLogin(in RequestCommon, out *ResponseCommon) (res *Session) {
	// TODO: modify to not re-decode session token
	if in.SessionToken == `` {
		out.SetError(498, ErrSessionTokenEmpty)
		return nil
	}
	defer func() {
		if res == nil {
			// force user to clear cookie
			out.SessionToken = conf.CookieLogoutValue
		}
	}()
	sess := &Session{}
	if !sess.Decrypt(in.SessionToken, in.UserAgent) {
		out.SetError(498, ErrSessionTokenInvalid)
		return nil
	}
	now := fastime.UnixNow()
	if sess.ExpiredAt < now {
		out.SetError(498, ErrSessionTokenExpired)
		return nil
	}

	session := rqAuth.NewSessions(d.AuthOltp)
	session.SessionToken = in.SessionToken
	if !session.FindBySessionToken() {
		out.SetError(498, ErrSessionTokenNotFound)
		return nil
	}
	if session.ExpiredAt <= now {
		out.SetError(498, ErrSessionTokenLoggedOut)
		return nil
	}

	user := rqAuth.NewUsers(d.AuthOltp)
	user.Id = session.UserId
	if !user.FindById() {
		sess.Role = GuestSegment
		out.SetError(498, ErrUserIdNotFound)
		return nil
	} else {
		sess.Role = UserSegment
	}
	sess.TenantCode = user.TenantCode

	if !sess.IsSuperAdmin {
		tCode, _ := GetTenantCodeByHost(in.Host)
		mapState, err := mAuth.ToInvitationStateMap(user.InvitationState)
		if err == nil {
			sess.Role = mapState.GetRoleByTenantCode(tCode)
		}
	}

	if sess.TenantCode != `` {
		tCode, _ := GetTenantCodeByHost(in.Host)
		if tCode == sess.TenantCode {
			sess.Role = TenantAdminSegment
		} else {
			sess.Role = UserSegment
		}
	}

	sess.IsSuperAdmin = d.Superadmins[sess.Email]

	sess.Segments = M.SB{}

	switch sess.Role {
	case TenantAdminSegment:
		sess.Segments[TenantAdminSegment] = true
		sess.Segments[ReportViewerSegment] = true
		sess.Segments[DataEntrySegment] = true
		sess.Segments[FieldSupervisorSegment] = true
		sess.Segments[UserSegment] = true
		sess.Segments[GuestSegment] = true

		sess.Role = TenantAdminSegment
	case DataEntrySegment:
		sess.Segments[DataEntrySegment] = true
		sess.Segments[UserSegment] = true
		sess.Segments[GuestSegment] = true

		sess.Role = DataEntrySegment
	case ReportViewerSegment:
		sess.Segments[ReportViewerSegment] = true
		sess.Segments[UserSegment] = true
		sess.Segments[GuestSegment] = true

		sess.Role = ReportViewerSegment
	case FieldSupervisorSegment:
		sess.Segments[FieldSupervisorSegment] = true
		sess.Segments[UserSegment] = true
		sess.Segments[GuestSegment] = true

		sess.Role = FieldSupervisorSegment
	case UserSegment:
		sess.Segments[GuestSegment] = true
		sess.Segments[UserSegment] = true

		sess.Role = UserSegment
	case GuestSegment:
		sess.Segments[GuestSegment] = true
	}

	if sess.IsSuperAdmin {
		sess.Segments[SuperAdminSegment] = true
		sess.Segments[TenantAdminSegment] = true
		sess.Segments[ReportViewerSegment] = true
		sess.Segments[DataEntrySegment] = true
		sess.Segments[FieldSupervisorSegment] = true
		sess.Segments[UserSegment] = true
		sess.Segments[GuestSegment] = true

		sess.Role = SuperAdminSegment
	}

	if !sess.Segments[UserSegment] && !sess.Segments[SuperAdminSegment] {
		out.SetError(403, ErrSegmentNotAllowed)
		return nil
	}

	out.actor = sess.UserId
	return sess
}

func (d *Domain) MustSuperAdmin(in RequestCommon, out *ResponseCommon) (sess *Session) {
	sess = d.MustLogin(in, out)
	if sess == nil {
		return nil
	}
	if !sess.IsSuperAdmin {
		out.SetError(403, ErrSessionUserNotSuperAdmin)
		return nil
	}
	sess.IsSuperAdmin = true
	return sess
}

func (d *Domain) MustTenantAdmin(in RequestCommon, out *ResponseCommon) (sess *Session) {
	sess = d.MustLogin(in, out)
	if sess == nil {
		return nil
	}

	if sess.TenantCode == `` {
		if !sess.IsSuperAdmin {
			out.SetError(403, ErrSessionUserNotTenantAdmin)
			return nil
		}
	} else {
		if !sess.IsSuperAdmin {
			tenantCode, _ := GetTenantCodeByHost(in.Host)
			if sess.TenantCode != tenantCode {
				out.SetError(403, ErrSessionUserNotTenantAdmin)
				return nil
			}
			tenant := rqAuth.NewTenants(d.AuthOltp)
			tenant.TenantCode = sess.TenantCode
			if !tenant.FindByTenantCode() {
				out.SetError(403, ErrSessionUserNotTenantAdmin)
				return nil
			}

			if tenant.DeletedAt > 0 {
				out.SetError(403, ErrSessionUserNotTenantAdmin)
				return nil
			}
		}

		sess.Segments[TenantAdminSegment] = true
	}

	sess.IsTenantAdmin = true
	return sess
}

func (d *Domain) MustDataEntry(in RequestCommon, out *ResponseCommon) (sess *Session) {
	sess = d.MustLogin(in, out)
	if sess == nil {
		return nil
	}

	if _, ok := sess.Segments[DataEntrySegment]; !ok {
		out.SetError(403, ErrSessionUserNotDataEntry)
		return nil
	}

	sess.IsDataEntry = true
	return sess
}

func (d *Domain) MustReportViewer(in RequestCommon, out *ResponseCommon) (sess *Session) {
	sess = d.MustLogin(in, out)
	if sess == nil {
		return nil
	}

	if _, ok := sess.Segments[ReportViewerSegment]; !ok {
		out.SetError(403, ErrSessionUserNotReportViewer)
		return nil
	}

	sess.IsReportViewer = true
	return sess
}

func (d *Domain) MustFieldSupervisor(in RequestCommon, out *ResponseCommon) (sess *Session) {
	sess = d.MustLogin(in, out)
	if sess == nil {
		return nil
	}

	if _, ok := sess.Segments[FieldSupervisorSegment]; !ok {
		out.SetError(403, ErrSessionUserNotFieldSupervisor)
		return nil
	}

	sess.IsFieldSupervisor = true
	return sess
}

func TryDecodeSession(in RequestCommon) (res Session) {
	if in.SessionToken == `` {
		return
	}
	sess := &Session{}
	if !sess.Decrypt(in.SessionToken, in.UserAgent) {
		return
	}
	now := fastime.UnixNow()
	if sess.ExpiredAt < now {
		return
	}
	res = *sess
	return
}
