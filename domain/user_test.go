package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"fmt"
	"testing"

	"github.com/kokizzu/lexid"
	"github.com/kpango/fastime"
	"github.com/stretchr/testify/assert"
)

func TestCreateCompany(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = testSuperAdminEmail
	isUserExist := user.FindByEmail()
	assert.True(t, isUserExist, `user must exist`)

	user.SetTenantCode(``)
	isUserUpdated := user.DoUpdateById()
	assert.True(t, isUserUpdated, `user must be updated`)

	t.Run(`createMustSucceed`, func(t *testing.T) {
		in := UserCreateCompanyIn{
			RequestCommon: testAdminRequestCommon(UserCreateCompanyAction),
			Company: rqAuth.Orgs{
				TenantCode: `coolcompany`,
				Name:       `Cool Company`,
				HeadTitle:  `Mr. Ben`,
			},
		}

		out := d.UserCreateCompany(&in)
		assert.Empty(t, out.Error, `failed to create company`)
	})
}

func TestUserUpdateProfile(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	in := UserUpdateProfileIn{
		UserName: `kaito1412`,
		FullName: `Kaito Kuroba`,
		Email:    `kaito1412@proton.me`,
		Country:  `JP`,
		Language: `jp`,
	}

	out := d.UserUpdateProfile(&in)
	assert.Empty(t, out.Error, `failed to update profile`)
}

func TestUserChangePassword(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	const oldPassword = `kaito12345678`
	const newPassword = `kaito12345678_xd`

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = testSuperAdminEmail
	assert.True(t, user.FindByEmail(), `user not exist`)

	user.SetPassword(oldPassword)
	user.SetEncryptedPassword(oldPassword, fastime.UnixNow())
	user.SetUpdatedAt(fastime.UnixNow())
	assert.True(t, user.DoUpdateByEmail(), `failed to update user`)

	t.Run(`changeMustSucceed`, func(t *testing.T) {
		in := UserChangePasswordIn{
			RequestCommon: testAdminRequestCommon(UserChangePasswordAction),
			OldPass:       oldPassword,
			NewPass:       newPassword,
		}

		out := d.UserChangePassword(&in)
		assert.Empty(t, out.Error, `failed to change password`)
	})

	t.Run(`changeMustFailed`, func(t *testing.T) {
		in := UserChangePasswordIn{
			RequestCommon: testAdminRequestCommon(UserChangePasswordAction),
			OldPass:       oldPassword,
			NewPass:       oldPassword,
		}

		out := d.UserChangePassword(&in)
		assert.NotEmpty(t, out.Error, `password changed, it must be failed`)
	})
}

func TestUserLogout(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`logoutMustSucceed`, func(t *testing.T) {
		in := UserLogoutIn{
			RequestCommon: testAdminRequestCommon(UserLogoutAction),
		}

		out := d.UserLogout(&in)
		assert.Empty(t, out.Error, `failed to logout`)
	})

	t.Run(`logoutMustFailed`, func(t *testing.T) {
		in := UserLogoutIn{
			RequestCommon: RequestCommon{
				SessionToken: `invalidSessionToken`,
				RequestId:    lexid.ID(),
				IpAddress:    "127.0.2.1",
				Debug:        true,
				Host:         fmt.Sprintf("http://%s:1234", testSuperAdminTenantCode),
				Action:       UserLogoutAction,
				Lat:          -1,
				Long:         -2,
				now:          testTime.Unix(),
				start:        testTime,
			},
		}

		out := d.UserLogout(&in)
		assert.NotEmpty(t, out.Error, `logout succeeded, it must be failed`)
	})
}

func TestUserSessionKill(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`sessionKillMustSucceed`, func(t *testing.T) {
		in := UserSessionKillIn{
			RequestCommon: testAdminRequestCommon(UserSessionKillAction),
		}

		in.SessionTokenHash = in.RequestCommon.SessionToken

		out := d.UserSessionKill(&in)
		assert.Empty(t, out.Error, `failed to kill session`)
	})

	t.Run(`sessionKillMustFailed`, func(t *testing.T) {
		in := UserSessionKillIn{
			RequestCommon: RequestCommon{
				SessionToken: `invalidSessionToken`,
				RequestId:    lexid.ID(),
				IpAddress:    "127.0.2.1",
				Debug:        true,
				Host:         fmt.Sprintf("http://%s:1234", testSuperAdminTenantCode),
				Action:       UserSessionKillAction,
				Lat:          -1,
				Long:         -2,
				now:          testTime.Unix(),
				start:        testTime,
			},
			SessionTokenHash: `invalidSessionTokenHash`,
		}

		out := d.UserSessionKill(&in)
		assert.NotEmpty(t, out.Error, `session killed, it must be failed`)
	})
}
