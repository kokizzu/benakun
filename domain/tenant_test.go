package domain

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

type inviteJoin struct {
	Email string `json:"email"`
}

func TestInviteJoinCompany(t *testing.T) {
	d, closer := testDomain()
	defer closer()
	const py = `{
		"email": "habi@ternaklinux.com"
	}`

	var ij inviteJoin

	t.Run(`insertMustSucceed`, func(t *testing.T) {

		err := json.Unmarshal([]byte(py), &ij)
		assert.NoError(t, err)

		in := TenantAdminInviteUserIn{
			RequestCommon: testAdminRequestCommon(TenantAdminInviteUserAction),
			Email:         ij.Email,
		}

		out := d.TenantAdminInviteUser(&in)
		assert.Empty(t, out.Error)
		assert.NotEmpty(t, out.Message)
	})
}
