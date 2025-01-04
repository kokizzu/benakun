package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"testing"

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
