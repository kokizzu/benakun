package domain

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"

	"benakun/model/mAuth/rqAuth"
)

func TestCreateCompany(t *testing.T) {
	d, closer := testDomain()
	defer closer()
	const orgJson = `{
		"company": {
			"tenantCode": "coolcom",
			"name": "Cool Company",
			"headTitle": "It's an AI company",
		}
	}`

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := UserCreateCompanyIn{
			RequestCommon: testAdminRequestCommon(UserCreateCompanyAction),
			Company:       rqAuth.Orgs{},
		}

		err := json.Unmarshal([]byte(orgJson), &in.Company)
		assert.NoError(t, err)

		out := d.UserCreateCompany(&in)
		assert.Empty(t, out.Error)
		assert.NotZero(t, out.Company.Id)
		in.Company.Id = out.Company.Id
		assert.NotZero(t, out.Company.CreatedAt)
		in.Company.CreatedAt = out.Company.CreatedAt
		assert.Equal(t, out.Company.Id, out.Company.CreatedBy)
		in.Company.Id = out.Company.CreatedBy
		assert.NotEmpty(t, in.Company.TenantCode)
		assert.NotEmpty(t, in.Company.Name)
		assert.NotEmpty(t, in.Company.HeadTitle)
	})
}
