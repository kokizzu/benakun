package domain

import (
	"benakun/model/mAuth/rqAuth"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompany(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := UserCreateCompanyIn{
			RequestCommon: testAdminRequestCommon(UserCreateCompanyAction),
			Company: rqAuth.Orgs{
				TenantCode: `coolcompany`,
				Name:       `Cool Company`,
				HeadTitle:  `Mr. Ben`,
			},
		}

		out := d.UserCreateCompany(&in)
		assert.Empty(t, out.Error)
		assert.NotZero(t, out.Company.Id)
		assert.NotZero(t, out.Company.CreatedAt)
	})
}
