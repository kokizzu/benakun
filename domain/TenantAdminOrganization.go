package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminOrganization.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminOrganization.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminOrganization.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminOrganization.go
//go:generate farify doublequote --file TenantAdminOrganization.go

type (
	TenantAdminOrganizationIn struct {
		RequestCommon
	}
	TenantAdminOrganizationOut struct {
		ResponseCommon
	}
)

const (
	TenantAdminOrganizationAction = `tenantAdmin/organization`
)

func (d *Domain) TenantAdminOrganization(in *TenantAdminOrganizationIn) (out TenantAdminOrganizationOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	return
}
