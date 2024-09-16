package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mBusiness"
	"benakun/model/mBusiness/rqBusiness"
	"benakun/model/mFinance"
	"benakun/model/mFinance/wcFinance"
	"benakun/model/zCrud"
	"errors"
	"testing"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/X"
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
	d, closer := testDomain()
	defer closer()

	tenant := rqAuth.NewTenants(d.AuthOltp)
	tenant.TenantCode = testSuperAdminTenantCode
	if !tenant.FindByTenantCode() {
		t.Error(`tenant not found`)
		return
	}

	type coaTreeNode struct{
		coaId uint64
		children []coaTreeNode
	}

	coaTrees := []coaTreeNode{}

	t.Run(`insertDefaultCoaMustSucceed`, func(t *testing.T) {
		var insertDefaultCoa func(ta *Tt.Adapter, coaDefault mFinance.CoaDefault, tenantCode string, parentId uint64) (uint64, error)
		insertDefaultCoa = func(ta *Tt.Adapter, coaDefault mFinance.CoaDefault, tenantCode string, parentId uint64) (uint64, error) {
			coa := wcFinance.NewCoaMutator(ta)
			coa.SetName(coaDefault.Name)
			coa.SetLabel(coaDefault.Label)
			coa.SetTenantCode(tenantCode)
		
			if parentId > 0 {
				coa.SetParentId(parentId)
			}
		
			if !coa.DoUpsertById() {
				return 0, errors.New(`failed to insert coa: "`+coa.Name+`"`)
			}
		
			var children []any
			if len(coaDefault.Children) > 0 {
				for _, childCoaDefault := range coaDefault.Children {
					childId, err := insertDefaultCoa(ta, childCoaDefault, tenantCode, coa.Id)
					if err != nil {
						return 0, err
					}
		
					children = append(children, childId)
				}
		
				coa.SetChildren(children)
				if !coa.DoUpsertById() {
					return 0, errors.New(`failed to update child of coa: "`+coa.Name+`"`)
				}				
			}

			toChildren := []coaTreeNode{}
			for _, child := range children {
				chNode := coaTreeNode{
					coaId: X.ToU(child),
				}
				toChildren = append(toChildren, chNode)
			}

			coaTrees = append(coaTrees, coaTreeNode{
				coaId: coa.Id,
				children: toChildren,
			})
		
			return coa.Id, nil
		}

		coaDefaults := mFinance.GetCoaDefaults()
		for _, coaDefault := range coaDefaults {
			_, err := insertDefaultCoa(d.AuthOltp, coaDefault, tenant.TenantCode, 0)
			
			assert.Nil(t, err, `insert default coa`)
		}

		t.Log(`coa default inserted`)

		t.Run(`moveCoaMustSucceed`, func(t *testing.T) {

		})
	})
}