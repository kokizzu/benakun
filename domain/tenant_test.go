package domain

import (
	"benakun/model/mBusiness"
	"benakun/model/mBusiness/rqBusiness"
	"benakun/model/zCrud"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := TenantAdminProductsIn{
			RequestCommon: testAdminRequestCommon(TenantAdminProductsAction),
			Cmd: zCrud.CmdUpsert,
			Product: &rqBusiness.Products{
				Name: `Cheeseburger`,
				Detail: `Daging sapi, keju, acar, bawang, mustard, dan saus tomat dalam roti berwijen.`,
				CogsIDR: 40000,
				Rule: mBusiness.RuleTypeFIFO,
				Kind: mBusiness.KindTypeGOODS,
			},
		}

		out := d.TenantAdminProducts(&in)
		assert.Empty(t, out.Error)
		assert.NotNil(t, out.Product)

		t.Run(`editMustSucceed`, func(t *testing.T) {
			in := TenantAdminProductsIn{
				RequestCommon: testAdminRequestCommon(TenantAdminProductsAction),
				Cmd: zCrud.CmdUpsert,
				Product: &rqBusiness.Products{
					Id: out.Product.Id,
					Name: out.Product.Name + ` [edited]`,
					Detail: out.Product.Detail + ` [edited]`,
					CogsIDR: out.Product.CogsIDR,
					Rule: mBusiness.RuleTypeFIFO,
					Kind: mBusiness.KindTypeGOODS,
				},
			}
	
			out := d.TenantAdminProducts(&in)
			assert.Empty(t, out.Error)
			assert.NotNil(t, out.Product)
		})

		t.Run(`deleteMustSucceed`, func(t *testing.T) {
			in := TenantAdminProductsIn{
				RequestCommon: testAdminRequestCommon(TenantAdminProductsAction),
				Cmd: zCrud.CmdDelete,
				Product: &rqBusiness.Products{
					Id: out.Product.Id,
				},
			}

			out := d.TenantAdminProducts(&in)
			assert.Empty(t, out.Error)
			assert.NotNil(t, out.Product)
		})

		t.Run(`restoreMustSucceed`, func(t *testing.T) {
			in := TenantAdminProductsIn{
				RequestCommon: testAdminRequestCommon(TenantAdminProductsAction),
				Cmd: zCrud.CmdRestore,
				Product: &rqBusiness.Products{
					Id: out.Product.Id,
				},
			}

			out := d.TenantAdminProducts(&in)
			assert.Empty(t, out.Error)
			assert.NotNil(t, out.Product)
		})
	})
}

func TestMoveCoA(t *testing.T) {
	
}