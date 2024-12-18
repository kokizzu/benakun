package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mBudget/rqBudget"
	"benakun/model/mBusiness"
	"benakun/model/mBusiness/rqBusiness"
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
	"benakun/model/zCrud"
	"errors"
	"testing"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/X"
	"github.com/kpango/fastime"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := TenantAdminProductsIn{
			RequestCommon: testAdminRequestCommon(TenantAdminProductsAction),
			Cmd:           zCrud.CmdUpsert,
			Product: &rqBusiness.Products{
				Name:    `Cheeseburger`,
				Detail:  `Daging sapi, keju, acar, bawang, mustard, dan saus tomat dalam roti berwijen.`,
				CogsIDR: 40000,
				Rule:    mBusiness.RuleTypeFIFO,
				Kind:    mBusiness.KindTypeGOODS,
			},
		}

		out := d.TenantAdminProducts(&in)
		assert.Empty(t, out.Error)
		assert.NotNil(t, out.Product)

		t.Run(`editMustSucceed`, func(t *testing.T) {
			in := TenantAdminProductsIn{
				RequestCommon: testAdminRequestCommon(TenantAdminProductsAction),
				Cmd:           zCrud.CmdUpsert,
				Product: &rqBusiness.Products{
					Id:      out.Product.Id,
					Name:    out.Product.Name + ` [edited]`,
					Detail:  out.Product.Detail + ` [edited]`,
					CogsIDR: out.Product.CogsIDR,
					Rule:    mBusiness.RuleTypeFIFO,
					Kind:    mBusiness.KindTypeGOODS,
				},
			}

			out := d.TenantAdminProducts(&in)
			assert.Empty(t, out.Error)
			assert.NotNil(t, out.Product)
		})

		t.Run(`deleteMustSucceed`, func(t *testing.T) {
			in := TenantAdminProductsIn{
				RequestCommon: testAdminRequestCommon(TenantAdminProductsAction),
				Cmd:           zCrud.CmdDelete,
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
				Cmd:           zCrud.CmdRestore,
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

	type coaTreeNode struct {
		coaId    uint64
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
				return 0, errors.New(`failed to insert coa: "` + coa.Name + `"`)
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
					return 0, errors.New(`failed to update child of coa: "` + coa.Name + `"`)
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
				coaId:    coa.Id,
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

		var insertCoa = func(ta *Tt.Adapter, c rqFinance.Coa) (id uint64, err error) {
			coa := wcFinance.NewCoaMutator(ta)
			coa.SetTenantCode(c.TenantCode)
			coa.SetName(c.Name)
			coa.SetParentId(c.ParentId)
			coa.SetChildren(c.Children)
			coa.SetCreatedAt(fastime.UnixNow())
			coa.SetUpdatedAt(fastime.UnixNow())

			if !coa.DoUpsertById() {
				return 0, errors.New(`failed to insert coa: "` + coa.Name + `"`)
			}
			return coa.Id, nil
		}

		t.Run(`moveCoaMustSucceed`, func(t *testing.T) {
			t.Log(`insert parent 01`)
			parent01_id, err := insertCoa(d.AuthOltp, rqFinance.Coa{
				TenantCode: tenant.TenantCode,
				Name:       `parent01`,
			})

			assert.Nil(t, err, `insert parent 01 coa`)

			t.Log(`insert parent 02`)
			parent02_id, err := insertCoa(d.AuthOltp, rqFinance.Coa{
				TenantCode: tenant.TenantCode,
				Name:       `parent02`,
			})

			assert.Nil(t, err, `insert parent 02 coa`)

			t.Log(`insert child 01`)
			child01_id, err := insertCoa(d.AuthOltp, rqFinance.Coa{
				TenantCode: tenant.TenantCode,
				Name:       `child01`,
				ParentId:   parent01_id,
			})

			assert.Nil(t, err, `insert child 01 coa`)

			t.Log(`update parent01's children`)
			parent01 := wcFinance.NewCoaMutator(d.AuthOltp)
			parent01.Id = parent01_id
			if !parent01.FindById() {
				t.Error(`parent 01 coa not found`)
			}
			parent01.SetChildren([]any{child01_id})
			if !parent01.DoUpsertById() {
				t.Error(`failed to update parent 01 coa`)
			}

			t.Log(`move child01 to parent02`)
			child01 := wcFinance.NewCoaMutator(d.AuthOltp)
			child01.Id = child01_id
			if !child01.FindById() {
				t.Error(`child 01 coa not found`)
			}
			child01.SetParentId(parent02_id)
			if !child01.DoUpsertById() {
				t.Error(`failed to move child 01 to parent 02`)
			}

			t.Log(`update parent02's children`)
			parent02 := wcFinance.NewCoaMutator(d.AuthOltp)
			parent02.Id = parent02_id
			if !parent02.FindById() {
				t.Error(`parent 02 coa not found`)
			}
			parent02.SetChildren([]any{child01_id})
			if !parent02.DoUpsertById() {
				t.Error(`failed to update parent 02 coa`)
			}
		})
	})
}

func TestSyncCoa(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	var insertCoa = func(ta *Tt.Adapter, name string) (id uint64, err error) {
		coa := wcFinance.NewCoaMutator(ta)
		coa.SetTenantCode(testSuperAdminTenantCode)
		coa.SetName(name)
		coa.SetCreatedAt(fastime.UnixNow())
		coa.SetUpdatedAt(fastime.UnixNow())

		if !coa.DoUpsertById() {
			return 0, errors.New(`failed to insert coa: "` + coa.Name + `"`)
		}

		return coa.Id, nil
	}

	t.Run(`syncMustSucceed`, func(t *testing.T) {
		coaProductId, err := insertCoa(d.AuthOltp, `Product`)
		assert.Nil(t, err, `insert product coa`)

		coaSupplierId, err := insertCoa(d.AuthOltp, `Supplier`)
		assert.Nil(t, err, `insert supplier coa`)

		coaCustomerId, err := insertCoa(d.AuthOltp, `Customer`)
		assert.Nil(t, err, `insert customer coa`)

		coaStaffId, err := insertCoa(d.AuthOltp, `Staff`)
		assert.Nil(t, err, `insert staff coa`)

		coaBankId, err := insertCoa(d.AuthOltp, `Bank`)
		assert.Nil(t, err, `insert bank coa`)

		in := TenantAdminSyncCoaIn{
			RequestCommon: testAdminRequestCommon(TenantAdminSyncCoaAction),
			Tenant: &rqAuth.Tenants{
				ProductsCoaId:  coaProductId,
				SuppliersCoaId: coaSupplierId,
				CustomersCoaId: coaCustomerId,
				StaffsCoaId:    coaStaffId,
				BanksCoaId:     coaBankId,
			},
		}

		out := d.TenantAdminSyncCoa(&in)
		assert.Nil(t, out.Error, `sync coa`)
	})
}

func TestBankAccounts(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := TenantAdminBankAccountsIn{
			RequestCommon: testAdminRequestCommon(TenantAdminBankAccountsAction),
			Cmd:           zCrud.CmdUpsert,
			Account: &rqBudget.BankAccounts{
				TenantCode:  testSuperAdminTenantCode,
				AccountName: `Account 01`,
				BankName:    `Bank 01`,
				Name:        `Account for bank 01`,
			},
		}

		out := d.TenantAdminBankAccounts(&in)
		assert.Nil(t, out.Error, `insert bank account`)
	})

	t.Run(`updateMustSucceed`, func(t *testing.T) {
		in := TenantAdminBankAccountsIn{
			RequestCommon: testAdminRequestCommon(TenantAdminBankAccountsAction),
			Cmd:           zCrud.CmdUpsert,
			Account: &rqBudget.BankAccounts{
				Id:          1,
				TenantCode:  testSuperAdminTenantCode,
				AccountName: `Account 01`,
				BankName:    `Bank 01`,
				Name:        `Account for bank 01`,
			},
		}

		out := d.TenantAdminBankAccounts(&in)
		assert.Nil(t, out.Error, `update bank account`)

		t.Run(`deleteMustSucceed`, func(t *testing.T) {
			in := TenantAdminBankAccountsIn{
				RequestCommon: testAdminRequestCommon(TenantAdminBankAccountsAction),
				Cmd:           zCrud.CmdDelete,
				Account: &rqBudget.BankAccounts{
					Id: 1,
				},
			}

			out := d.TenantAdminBankAccounts(&in)
			assert.Nil(t, out.Error, `delete bank account`)

			t.Run(`restoreMustSucceed`, func(t *testing.T) {
				in := TenantAdminBankAccountsIn{
					RequestCommon: testAdminRequestCommon(TenantAdminBankAccountsAction),
					Cmd:           zCrud.CmdRestore,
					Account: &rqBudget.BankAccounts{
						Id: 1,
					},
				}

				out := d.TenantAdminBankAccounts(&in)
				assert.Nil(t, out.Error, `restore bank account`)
			})
		})
	})
}
