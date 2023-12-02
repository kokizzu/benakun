package domain

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

type createCompany struct {
	TenantCode  string `json:"tenantCode"`
	CompanyName string `json:"companyName"`
	HeadTitle   string `json:"headTitle"`
}

func TestCreateCompany(t *testing.T) {
	d, closer := testDomain()
	defer closer()
	const orgJson = `{
		"tenantCode": "coolcom",
		"companyName": "Cool Company",
		"headTitle": "It's an AI company"
	}`

	var c createCompany

	t.Run(`insertMustSucceed`, func(t *testing.T) {

		err := json.Unmarshal([]byte(orgJson), &c)
		assert.NoError(t, err)

		in := UserCreateCompanyIn{
			RequestCommon: testAdminRequestCommon(UserCreateCompanyAction),
			TenantCode:    c.TenantCode,
			CompanyName:   c.CompanyName,
			HeadTitle:     c.HeadTitle,
		}

		out := d.UserCreateCompany(&in)
		assert.Empty(t, out.Error)
		assert.NotZero(t, out.Company.Id)
		assert.NotZero(t, out.Company.CreatedAt)
		assert.Equal(t, out.Company.Id, out.Company.CreatedBy)
		assert.NotEmpty(t, in.TenantCode)
		assert.NotEmpty(t, in.CompanyName)
		assert.NotEmpty(t, in.HeadTitle)
		assert.NotEmpty(t, out.Company.TenantCode)
	})
}
