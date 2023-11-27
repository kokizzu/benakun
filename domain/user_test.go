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
			"tenantCode": "habiternax",
			"name": "Ternak Linux",
			"headTitle": "Linux dev",
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

		assert.NotZero(t, out.Company.CreatedAt)
	})
}
