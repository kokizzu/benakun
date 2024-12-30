package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBudget/rqBudget"
	"benakun/model/mBusiness"
	"benakun/model/mBusiness/rqBusiness"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
	"benakun/model/zCrud"
	"errors"
	"testing"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kpango/fastime"
	"github.com/stretchr/testify/assert"
)

func TestProduct(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		coaParent := wcFinance.NewCoaMutator(d.AuthOltp)
		coaParent.SetTenantCode(testSuperAdminTenantCode)
		coaParent.SetName(`Product`)
		coaParent.SetCreatedAt(fastime.UnixNow())
		coaParent.SetUpdatedAt(fastime.UnixNow())
		isInsertCoaParent := coaParent.DoInsert()
		assert.True(t, isInsertCoaParent, `failed to insert coa parent`)

		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.TenantCode = testSuperAdminTenantCode
		tenant.SetProductsCoaId(coaParent.Id)
		isUpdateTenantProductsCoa := tenant.DoUpdateByTenantCode()
		assert.True(t, isUpdateTenantProductsCoa, `failed to update tenant products coa`)

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
	assert.True(t, tenant.FindByTenantCode(), `failed to find tenant`)

	t.Run(`insertCoaParentsMustSucceed`, func(t *testing.T) {
		t.Log(`Insert CoA Parent 01`)
		coaParent01 := wcFinance.NewCoaMutator(d.AuthOltp)
		coaParent01.SetName(`parent01`)
		coaParent01.SetTenantCode(tenant.TenantCode)
		isInsertCoaParent01 := coaParent01.DoInsert()
		assert.True(t, isInsertCoaParent01, `failed to insert coa parent01`)

		t.Log(`Insert CoA Parent 02`)
		coaParent02 := wcFinance.NewCoaMutator(d.AuthOltp)
		coaParent02.SetName(`parent02`)
		coaParent02.SetTenantCode(tenant.TenantCode)
		isInsertCoaParent02 := coaParent02.DoInsert()
		assert.True(t, isInsertCoaParent02, `failed to insert coa parent02`)

		t.Run(`moveCoaMustSucceed`, func(t *testing.T) {
			t.Log(`Insert CoA Child for Coa Parent 01`)
			coaChild := wcFinance.NewCoaMutator(d.AuthOltp)
			coaChild.SetName(`child01`)
			coaChild.SetParentId(coaParent01.Id)
			coaChild.SetTenantCode(tenant.TenantCode)
			isInsertCoaChild01 := coaChild.DoInsert()
			assert.True(t, isInsertCoaChild01, `failed to insert coa child01`)

			t.Log(`Update CoA Parent 01 children`)
			coaParent01.SetChildren([]any{coaChild.Id})
			isUpdateCoaParent01 := coaParent01.DoUpdateById()
			assert.True(t, isUpdateCoaParent01, `failed to update coa parent01`)

			t.Log(`Move CoA Child to Coa Parent 02`)
			coaChild.SetParentId(coaParent02.Id)
			isUpdateCoaChild01 := coaChild.DoUpdateById()
			assert.True(t, isUpdateCoaChild01, `failed to update coa child01`)

			t.Log(`Update CoA Parent 02 children`)
			coaParent02.SetChildren([]any{coaChild.Id})
			isUpdateCoaParent02 := coaParent02.DoUpdateById()
			assert.True(t, isUpdateCoaParent02, `failed to update coa parent02`)

			t.Log(`Remove CoA Child from Coa Parent 01`)
			coaParent01.SetChildren([]any{})
			isUpdateCoaParent01 = coaParent01.DoUpdateById()
			assert.True(t, isUpdateCoaParent01, `failed to update coa parent01`)
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

		if !coa.DoInsert() {
			return 0, errors.New(`failed to insert coa: "` + coa.Name + `"`)
		}

		return coa.Id, nil
	}

	t.Run(`syncMustSucceed`, func(t *testing.T) {
		coaProductId, err := insertCoa(d.AuthOltp, `Product`)
		assert.NoError(t, err, `insert product coa`)

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
		assert.Empty(t, out.Error, `error sync coa`)
	})
}

func TestBankAccounts(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	coaBank := wcFinance.NewCoaMutator(d.AuthOltp)
	coaBank.SetTenantCode(testSuperAdminTenantCode)
	coaBank.SetName(`Bank`)
	coaBank.SetCreatedAt(fastime.UnixNow())
	coaBank.SetUpdatedAt(fastime.UnixNow())
	isInsertCoaBank := coaBank.DoInsert()
	assert.True(t, isInsertCoaBank, `failed to insert coa bank`)

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = testSuperAdminTenantCode
	tenant.SetBanksCoaId(coaBank.Id)
	isUpdateTenantBanksCoa := tenant.DoUpdateByTenantCode()
	assert.True(t, isUpdateTenantBanksCoa, `failed to update tenant banks coa`)

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
		assert.Empty(t, out.Error, `failed to insert bank account`)

		t.Log(`Bank account inserted`)
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
		assert.Empty(t, out.Error, `failed to update bank account`)

		t.Run(`deleteMustSucceed`, func(t *testing.T) {
			in := TenantAdminBankAccountsIn{
				RequestCommon: testAdminRequestCommon(TenantAdminBankAccountsAction),
				Cmd:           zCrud.CmdDelete,
				Account: &rqBudget.BankAccounts{
					Id: 1,
				},
			}

			out := d.TenantAdminBankAccounts(&in)
			assert.Empty(t, out.Error, `failed to delete bank account`)

			t.Run(`restoreMustSucceed`, func(t *testing.T) {
				in := TenantAdminBankAccountsIn{
					RequestCommon: testAdminRequestCommon(TenantAdminBankAccountsAction),
					Cmd:           zCrud.CmdRestore,
					Account: &rqBudget.BankAccounts{
						Id: 1,
					},
				}

				out := d.TenantAdminBankAccounts(&in)
				assert.Empty(t, out.Error, `failed to restore bank account`)
			})
		})
	})
}

func TestLocation(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := TenantAdminLocationsIn{
			RequestCommon: testAdminRequestCommon(TenantAdminLocationsAction),
			Cmd:           zCrud.CmdUpsert,
			Location: rqBusiness.Locations{
				Name:         `Location 01`,
				Country:      `Country 01`,
				StateProvice: `State 01`,
				Address:      `Address 01`,
				Subdistrict:  `Subdistrict 01`,
				CityRegency:  `City 01`,
				Village:      `Village 01`,
				Lat:          1,
				Lng:          1,
				RwBanjar:     `RW 01`,
				RtNeigb:      `RT 01`,
			},
		}

		out := d.TenantAdminLocations(&in)
		assert.Empty(t, out.Error, `failed to insert location`)

		t.Run(`updateMustSucceed`, func(t *testing.T) {
			in := TenantAdminLocationsIn{
				RequestCommon: testAdminRequestCommon(TenantAdminLocationsAction),
				Cmd:           zCrud.CmdUpsert,
				Location: rqBusiness.Locations{
					Id:           out.Location.Id,
					Name:         `Location XD`,
					Country:      `Country XD`,
					StateProvice: `State XD`,
					Address:      `Address XD`,
					Subdistrict:  `Subdistrict XD`,
					CityRegency:  `City XD`,
					Village:      `Village XD`,
					Lat:          1,
					Lng:          1,
					RwBanjar:     `RW XD`,
					RtNeigb:      `RT XD`,
				},
			}

			out := d.TenantAdminLocations(&in)
			assert.Empty(t, out.Error, `failed to update location`)

			t.Run(`deleteMustSucceed`, func(t *testing.T) {
				in := TenantAdminLocationsIn{
					RequestCommon: testAdminRequestCommon(TenantAdminLocationsAction),
					Cmd:           zCrud.CmdDelete,
					Location: rqBusiness.Locations{
						Id: out.Location.Id,
					},
				}

				out := d.TenantAdminLocations(&in)
				assert.Empty(t, out.Error, `failed to delete location`)

				t.Run(`restoreMustSucceed`, func(t *testing.T) {
					in := TenantAdminLocationsIn{
						RequestCommon: testAdminRequestCommon(TenantAdminLocationsAction),
						Cmd:           zCrud.CmdRestore,
						Location: rqBusiness.Locations{
							Id: out.Location.Id,
						},
					}

					out := d.TenantAdminLocations(&in)
					assert.Empty(t, out.Error, `failed to restore location`)
				})
			})
		})
	})
}

func TestTransactionTemplate(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := TenantAdminTransactionTemplateIn{
			RequestCommon: testAdminRequestCommon(TenantAdminTransactionTemplateAction),
			Cmd:           zCrud.CmdUpsert,
			TransactionTemplate: &rqFinance.TransactionTemplate{
				Name:     `Transaction Template 01`,
				Color:    `#000000`,
				ImageURL: `https://example.com/image.png`,
			},
		}

		out := d.TenantAdminTransactionTemplate(&in)
		assert.Empty(t, out.Error, `failed to insert transaction template`)

		t.Run(`updateMustSucceed`, func(t *testing.T) {
			in := TenantAdminTransactionTemplateIn{
				RequestCommon: testAdminRequestCommon(TenantAdminTransactionTemplateAction),
				Cmd:           zCrud.CmdUpsert,
				TransactionTemplate: &rqFinance.TransactionTemplate{
					Id:       out.TransactionTemplate.Id,
					Name:     `Transaction Template XD`,
					Color:    `#000000`,
					ImageURL: `https://example.com/image.png`,
				},
			}

			out := d.TenantAdminTransactionTemplate(&in)
			assert.Empty(t, out.Error, `failed to update transaction template`)

			t.Run(`deleteMustSucceed`, func(t *testing.T) {
				in := TenantAdminTransactionTemplateIn{
					RequestCommon: testAdminRequestCommon(TenantAdminTransactionTemplateAction),
					Cmd:           zCrud.CmdDelete,
					TransactionTemplate: &rqFinance.TransactionTemplate{
						Id: out.TransactionTemplate.Id,
					},
				}

				out := d.TenantAdminTransactionTemplate(&in)
				assert.Empty(t, out.Error, `failed to delete transaction template`)

				t.Run(`restoreMustSucceed`, func(t *testing.T) {
					in := TenantAdminTransactionTemplateIn{
						RequestCommon: testAdminRequestCommon(TenantAdminTransactionTemplateAction),
						Cmd:           zCrud.CmdRestore,
						TransactionTemplate: &rqFinance.TransactionTemplate{
							Id: out.TransactionTemplate.Id,
						},
					}

					out := d.TenantAdminTransactionTemplate(&in)
					assert.Empty(t, out.Error, `failed to restore transaction template`)
				})
			})
		})
	})
}
