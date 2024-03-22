package domain

import (
	"testing"

	"github.com/goccy/go-json"
	"github.com/stretchr/testify/assert"
)

type inviteJoin struct {
	Email string `json:"email"`
}

func TestInviteJoinCompany(t *testing.T) {
	d, closer := testDomain()
	defer closer()
	const py = `{
		"email": "habi@ternaklinux.com"
	}`

	var ij inviteJoin

	t.Run(`insertMustSucceed`, func(t *testing.T) {

		err := json.Unmarshal([]byte(py), &ij)
		assert.NoError(t, err)

		in := TenantAdminInviteUserIn{
			RequestCommon: testAdminRequestCommon(TenantAdminInviteUserAction),
			Email:         ij.Email,
		}

		out := d.TenantAdminInviteUser(&in)
		assert.Empty(t, out.Error)
		assert.NotEmpty(t, out.Message)
	})
}

func TestCreateCoaChild(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := TenantAdminCreateCoaChildIn{
			RequestCommon: testAdminRequestCommon(TenantAdminCreateCoaChildAction),
			Name: "makan",
			ParentId: 1,
		}

		out := d.TenantAdminCreateCoaChild(&in)
		assert.Empty(t, out.Error)
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
		if assert.Empty(t, out.Error) {
			t.Run(`createDepartmentMustSucceed`, func(t *testing.T) {
				in := TenantAdminCreateOrganizationChildIn{
					RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
					Name: `Department A`,
					HeadTitle: `Habibi Dep`,
					ParentId: out.Company.Id,
				}

				out := d.TenantAdminCreateOrganizationChild(&in)
				assert.Empty(t, out.Error)
				if out.Org != nil {
					t.Run(`createDevisionMustSucceed`, func(t *testing.T) {
						in := TenantAdminCreateOrganizationChildIn{
							RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
							Name: `Division A`,
							HeadTitle: `Habibi Div`,
							ParentId: out.Org.Id,
						}

						out := d.TenantAdminCreateOrganizationChild(&in)
						assert.Empty(t, out.Error)
						if out.Org != nil {
							t.Run(`createJobMustSucceed`, func(t *testing.T) {
								in := TenantAdminCreateOrganizationChildIn{
									RequestCommon: testAdminRequestCommon(TenantAdminCreateOrganizationChildAction),
									Name: `Job A`,
									HeadTitle: `Habibi Job`,
									ParentId: out.Org.Id,
								}
		
								out := d.TenantAdminCreateOrganizationChild(&in)
								assert.Empty(t, out.Error)
							})
						}
					})
				}
			})
		}
	})
}