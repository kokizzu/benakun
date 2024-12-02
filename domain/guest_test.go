package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`loginMustSucceed`, func(t *testing.T) {
		in := GuestLoginIn{
			RequestCommon: testGuestRequestCommon(GuestLoginAction),
			Email: "kaito1412@proton.me",
			Password: "kaito12345678",
		}

		out := d.GuestLogin(&in)
		assert.Empty(t, out.Error)
		assert.NotNil(t, out.User)

		t.Log(`user:`, out.User)
	})
}