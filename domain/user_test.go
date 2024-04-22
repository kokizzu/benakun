package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCompany(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {

		in := UserCreateCompanyIn{
			RequestCommon: testAdminRequestCommon(UserCreateCompanyAction),
			TenantCode:    `coolcompany`,
			CompanyName:   `Cool Company`,
			HeadTitle:     `Mr. Ben`,
		}

		out := d.UserCreateCompany(&in)
		assert.Empty(t, out.Error)
		assert.NotZero(t, out.Company.Id)
		assert.NotZero(t, out.Company.CreatedAt)
	})
}
