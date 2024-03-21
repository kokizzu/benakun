package domain

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestCreateCoaChild(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := TenantAdminCreateCoaChildIn{
			RequestCommon: testAdminRequestCommon(TenantAdminCreateCoaChildAction),
			Name: "makan",
			ParentId: 1,
		}

		out := d.TenantAdminCreateCoaChild(&in)
		assert.Empty(t, out.Error)
	})
}

func TestTenantAdminOrganization(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`addDepartmentMustSucceed`, func(t *testing.T) {
		in := TenantAdminCreateOrganizationChildIn{
			RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
			Name: `Department 1`,
			HeadTitle: `Mr. Habi`,
			ParentId: 1,
		}

		out := d.TenantAdminCreateOrganizationChild(&in)
		assert.Empty(t, out.Error)
		require.Empty(t, out.Error)

		if out.Orgs != nil {
			t.Log(`Organization:`, out.Orgs)
		}
	})
}