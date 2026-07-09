package domain

import (
	"testing"

	"github.com/kokizzu/gotro/X"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`loginMustSucceed`, func(t *testing.T) {
		in := GuestLoginIn{
			RequestCommon: testGuestRequestCommon(GuestLoginAction),
			Email:         testSuperAdminEmail,
			Password:      testSuperAdminPassword,
		}

		out := d.GuestLogin(&in)
		assert.Empty(t, out.Error)
		assert.NotNil(t, out.User)

		t.Log(`User.Email:`, out.User.Email)
		t.Log(`User.FullName:`, out.User.FullName)
		t.Log(`User.TenantCode:`, out.User.TenantCode)
		t.Log(`User.CreatedAt:`, out.User.CreatedAt)
		t.Log(`User.UpdatedAt:`, out.User.UpdatedAt)
	})

	t.Run(`loginMustFail`, func(t *testing.T) {
		in := GuestLoginIn{
			RequestCommon: testGuestRequestCommon(GuestLoginAction),
			Email:         "testemail@notemail.com",
			Password:      "oijrishhsrg222",
		}

		out := d.GuestLogin(&in)
		assert.NotEmpty(t, out.Error)
		assert.Nil(t, out.User)
	})
}

func TestRegister(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`registerMustSucceed`, func(t *testing.T) {
		registeredEmail := "kaito1412@proton.me"
		registeredPass := "kaito12345678"

		in := GuestRegisterIn{
			RequestCommon: testGuestRequestCommon(GuestRegisterAction),
			Email:         registeredEmail,
			Password:      registeredPass,
		}

		out := d.GuestRegister(&in)
		assert.Empty(t, out.Error)
		assert.NotNil(t, out.User)

		t.Log(`User: `, X.ToJsonPretty(out.User))

		t.Run(`registerMustFail`, func(t *testing.T) {
			in := GuestRegisterIn{
				RequestCommon: testGuestRequestCommon(GuestRegisterAction),
				Email:         registeredEmail,
				Password:      registeredPass,
			}

			out := d.GuestRegister(&in)
			assert.NotEmpty(t, out.Error)
			assert.Nil(t, out.User)
		})
	})
}
