package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserManagement(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := SuperAdminUserManagementIn{
			RequestCommon: testAdminRequestCommon(SuperAdminUserManagementAction),
			Cmd:           zCrud.CmdUpsert,
			TenantAdmin:   testSuperAdminTenantCode,
			User: rqAuth.Users{
				Email:    `kaitokuroba1412@gmail.com`,
				FullName: `Kaito Kuroba`,
				Role:     mAuth.RoleReportViewer,
			},
		}

		out := d.SuperAdminUserManagement(&in)
		assert.Empty(t, out.Error, `failed to insert user`)

		t.Run(`updateMustSucceed`, func(t *testing.T) {
			in := SuperAdminUserManagementIn{
				RequestCommon: testAdminRequestCommon(SuperAdminUserManagementAction),
				Cmd:           zCrud.CmdUpsert,
				TenantAdmin:   testSuperAdminTenantCode,
				User: rqAuth.Users{
					Id:       out.User.Id,
					Email:    `kaitokuroba1412@gmail.co.xd`,
					FullName: `Kaito Kuroba XD`,
					Role:     mAuth.RoleReportViewer,
				},
			}

			out := d.SuperAdminUserManagement(&in)
			assert.Empty(t, out.Error, `failed to update user`)

			t.Run(`deleteMustSucceed`, func(t *testing.T) {
				in := SuperAdminUserManagementIn{
					RequestCommon: testAdminRequestCommon(SuperAdminUserManagementAction),
					Cmd:           zCrud.CmdDelete,
					TenantAdmin:   testSuperAdminTenantCode,
					User: rqAuth.Users{
						Id: out.User.Id,
					},
				}

				out := d.SuperAdminUserManagement(&in)
				assert.Empty(t, out.Error, `failed to delete user`)
			})
		})
	})
}

func TestTenantManagement(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	tenantCode := fmt.Sprintf("test-%s", generate4RandomNumber())

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		user := wcAuth.NewUsersMutator(d.AuthOltp)
		user.SetEmail(`kaitokuroba1412@gmail.com`)
		user.SetFullName(`Kaito Kuroba`)
		user.SetRole(mAuth.RoleTenantAdmin)
		user.SetTenantCode(tenantCode)
		assert.True(t, user.DoInsert(), `failed to insert user`)

		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.SetTenantCode(tenantCode)
		assert.True(t, tenant.DoInsert(), `failed to insert tenant`)

		org := wcAuth.NewOrgsMutator(d.AuthOltp)
		org.SetTenantCode(tenantCode)
		org.SetName(`Magician`)
		org.SetHeadTitle(`Kaito Kuroba`)
		assert.True(t, org.DoInsert(), `failed to insert organization`)

		t.Run(`deleteMustSucceed`, func(t *testing.T) {
			in := SuperAdminTenantManagementIn{
				Cmd: zCrud.CmdDelete,
				Tenant: rqAuth.Tenants{
					Id: tenant.Id,
				},
			}

			out := d.SuperAdminTenantManagement(&in)
			assert.Empty(t, out.Error, `failed to delete tenant`)

			t.Run(`restoreMustSucceed`, func(t *testing.T) {
				in := SuperAdminTenantManagementIn{
					Cmd: zCrud.CmdRestore,
					Tenant: rqAuth.Tenants{
						Id: tenant.Id,
					},
				}

				out := d.SuperAdminTenantManagement(&in)
				assert.Empty(t, out.Error, `failed to restore tenant`)
			})
		})
	})
}
