package domain

import (
	"benakun/model/mFinance/rqFinance"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateCoaChild(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	TestCreateCompany(t)

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := TenantAdminUpsertCoaChildIn{
			RequestCommon: testAdminRequestCommon(TenantAdminUpsertCoaChildAction),
			Coa: rqFinance.Coa{
				Name: `Company`,
				ParentId: 1,
			},
		}

		out := d.TenantAdminUpsertCoaChild(&in)
		assert.Empty(t, out.Error)
		t.Log(`COAS`, out.Coas)
	})
}

// TODO: Unit test CoA
// Unit test budgeting
// Locations
// Products