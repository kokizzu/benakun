package domain

import (
	"benakun/model/mFinance/rqFinance"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestTenantAdminOrganization(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`createCompanyMustSucceed`, func(t *testing.T) {
		in := UserCreateCompanyIn{
			RequestCommon: testAdminRequestCommon(UserCreateCompanyAction),
			TenantCode: `habi`,
			CompanyName: `Benalu Dev`,
			HeadTitle: `Ahmad Habibi`,
		}

		out := d.UserCreateCompany(&in)
		require.Empty(t, out.Error)
		require.NotEmpty(t, out.Company)

		// Organization IDs to move/do operation
		// Make sure channel size is same as how many time it will be use
		ID_divC := make(chan uint64, 2)
		ID_divD := make(chan uint64)
		ID_jobF := make(chan uint64, 6)

		if assert.Empty(t, out.Error) {
			go t.Run(`addDeptAMustSucceed`, func(t *testing.T) {
				in := TenantAdminCreateOrganizationChildIn{
					RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
					Name: `DeptA`,
					HeadTitle: `Mr. DeptA`,
					ParentId: out.Company.Id,
				}

				out := d.TenantAdminCreateOrganizationChild(&in)
				require.Empty(t, out.Error)
				assert.NotEmpty(t, out.Org)
				assert.NotZero(t, out.Org.Id)
				if assert.Empty(t, out.Error) {
					t.Run(`addDeptA2MustError`, func(t *testing.T) {
						in := TenantAdminCreateOrganizationChildIn{
							RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
							Name: `DeptA2`,
							HeadTitle: `Mr. DeptA2`,
							ParentId: 000000000, // error
						}
		
						out := d.TenantAdminCreateOrganizationChild(&in)
						require.NotEmpty(t, out.Error)
						assert.NotEmpty(t, out.Error)
					})

					t.Run(`addDivCMustSucceed`, func(t *testing.T) {
						in := TenantAdminCreateOrganizationChildIn{
							RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
							Name: `DivC`,
							HeadTitle: `Mr. DivC`,
							ParentId: out.Org.Id,
						}

						out := d.TenantAdminCreateOrganizationChild(&in)
						require.Empty(t, out.Error)
						assert.NotEmpty(t, out.Org)
						assert.NotZero(t, out.Org.Id)

						// send it twice because will be use twice
						ID_divC <- out.Org.Id; ID_divC <- out.Org.Id 

						if assert.Empty(t, out.Error) {
							t.Run(`addJobAMustSucceed`, func(t *testing.T) {
								in := TenantAdminCreateOrganizationChildIn{
									RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
									Name: `JobA`,
									HeadTitle: `Mr. JobA`,
									ParentId: out.Org.Id,
								}
		
								out := d.TenantAdminCreateOrganizationChild(&in)
								assert.Empty(t, out.Error)
							})

							t.Run(`moveJobFToDivCMustSucceed`, func(t *testing.T) {
								in := TenantAdminMoveOrganizationChildIn{
									RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
									Id: <-ID_jobF,
									MoveToIdx: 0,
									ToParentId: out.Org.Id,
								}
		
								out := d.TenantAdminMoveOrganizationChild(&in)
								require.Empty(t, out.Error)
								assert.NotEmpty(t, out.Org)
								assert.NotZero(t, out.Org.Id)
							})
						}
					})

					t.Run(`addDivDMustSucceed`, func(t *testing.T) {
						in := TenantAdminCreateOrganizationChildIn{
							RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
							Name: `DivD`,
							HeadTitle: `Mr. DivD`,
							ParentId: out.Org.Id,
						}

						out := d.TenantAdminCreateOrganizationChild(&in)
						require.Empty(t, out.Error)
						assert.NotEmpty(t, out.Org)
						assert.NotZero(t, out.Org.Id)

						ID_divD <- out.Org.Id

						t.Run(`moveDivCToDivDMustError`, func(t *testing.T) {
							in := TenantAdminMoveOrganizationChildIn{
								RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
								Id: <-ID_divC,
								MoveToIdx: 0,
								ToParentId: out.Org.Id,
							}
	
							out := d.TenantAdminMoveOrganizationChild(&in)
							require.NotEmpty(t, out.Error)
							assert.NotEmpty(t, out.Error)
						})
					})
				}
			})

			go t.Run(`addDeptBMustSucceed`, func(t *testing.T) {
				in := TenantAdminCreateOrganizationChildIn{
					RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
					Name: `DeptB`,
					HeadTitle: `Mr. DeptB`,
					ParentId: out.Company.Id,
				}

				out := d.TenantAdminCreateOrganizationChild(&in)
				require.Empty(t, out.Error)
				assert.NotEmpty(t, out.Org)
				assert.NotZero(t, out.Org.Id)
				if assert.Empty(t, out.Error) {
					t.Run(`addDivEMustSucceed`, func(t *testing.T) {
						in := TenantAdminCreateOrganizationChildIn{
							RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
							Name: `DivE`,
							HeadTitle: `Mr. DivE`,
							ParentId: out.Org.Id,
						}

						out := d.TenantAdminCreateOrganizationChild(&in)
						require.Empty(t, out.Error)
						assert.NotEmpty(t, out.Org)
						assert.NotZero(t, out.Org.Id)
						if assert.Empty(t, out.Error) {
							t.Run(`addJobFMustSucceed`, func(t *testing.T) {
								in := TenantAdminCreateOrganizationChildIn{
									RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
									Name: `JobF`,
									HeadTitle: `Mr. JobF`,
									ParentId: out.Org.Id,
								}
		
								out := d.TenantAdminCreateOrganizationChild(&in)
								assert.Empty(t, out.Error)

								ID_jobF <- out.Org.Id // 1
								ID_jobF <- out.Org.Id // 2
								ID_jobF <- out.Org.Id // 3
								ID_jobF <- out.Org.Id // 4
								ID_jobF <- out.Org.Id // 5
								ID_jobF <- out.Org.Id // 6
							})
						}
					})

					t.Run(`moveDivCToDeptBLastPositionMustSucceed`, func(t *testing.T) {
						in := TenantAdminMoveOrganizationChildIn{
							RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
							Id: <-ID_divC,
							MoveToIdx: 1, // move to the last position
							ToParentId: out.Org.Id,
						}

						out := d.TenantAdminMoveOrganizationChild(&in)
						require.Empty(t, out.Error)
						assert.NotEmpty(t, out.Org)
						assert.NotZero(t, out.Org.Id)
					})

					t.Run(`moveDivDToDeptBFirstPositionMustSucceed`, func(t *testing.T) {
						in := TenantAdminMoveOrganizationChildIn{
							RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
							Id: <-ID_divD,
							MoveToIdx: 0, // move to the First position
							ToParentId: out.Org.Id,
						}

						out := d.TenantAdminMoveOrganizationChild(&in)
						require.Empty(t, out.Error)
						assert.NotEmpty(t, out.Org)
						assert.NotZero(t, out.Org.Id)
					})

					t.Run(`moveJobFToDeptBMustError`, func(t *testing.T) {
						in := TenantAdminMoveOrganizationChildIn{
							RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
							Id: <-ID_jobF,
							MoveToIdx: 1,
							ToParentId: out.Org.Id,
						}

						out := d.TenantAdminMoveOrganizationChild(&in)
						require.NotEmpty(t, out.Error)
						assert.NotEmpty(t, out.Error)
					})

					t.Run(`moveJobFToDeptBFirstPositionMustError`, func(t *testing.T) {
						in := TenantAdminMoveOrganizationChildIn{
							RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
							Id: <-ID_jobF,
							MoveToIdx: 0,
							ToParentId: out.Org.Id,
						}

						out := d.TenantAdminMoveOrganizationChild(&in)
						require.NotEmpty(t, out.Error)
						assert.NotEmpty(t, out.Error)
					})
				}
			})

			go t.Run(`moveJobFToCompanyFirstPositionMustError`, func(t *testing.T) {
				in := TenantAdminMoveOrganizationChildIn{
					RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
					Id: <-ID_jobF,
					MoveToIdx: 0,
					ToParentId: out.Company.Id,
				}

				out := d.TenantAdminMoveOrganizationChild(&in)
				require.NotEmpty(t, out.Error)
				assert.NotEmpty(t, out.Error)
			})

			go t.Run(`moveJobFToCompanyLastPositionMustError`, func(t *testing.T) {
				in := TenantAdminMoveOrganizationChildIn{
					RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
					Id: <-ID_jobF,
					MoveToIdx: 2,
					ToParentId: out.Company.Id,
				}

				out := d.TenantAdminMoveOrganizationChild(&in)
				require.NotEmpty(t, out.Error)
				assert.NotEmpty(t, out.Error)
			})

			go t.Run(`moveJobFToCompanyMiddlePositionMustError`, func(t *testing.T) {
				in := TenantAdminMoveOrganizationChildIn{
					RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
					Id: <-ID_jobF,
					MoveToIdx: 1, // It's middle position
					ToParentId: out.Company.Id,
				}

				out := d.TenantAdminMoveOrganizationChild(&in)
				require.NotEmpty(t, out.Error)
				assert.NotEmpty(t, out.Error)
			})
		}
	})
}