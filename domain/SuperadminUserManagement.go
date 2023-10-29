package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file SuperAdminUserManagement.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type SuperAdminUserManagement.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type SuperAdminUserManagement.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type SuperAdminUserManagement.go
//go:generate farify doublequote --file SuperAdminUserManagement.go

type (
	SuperAdminUserManagementIn struct {
		RequestCommon
	}
	SuperAdminUserManagementOut struct {
		ResponseCommon
	}
)

const (
	SuperAdminUserManagementAction = `superAdmin/userManagement`
)

func (d *Domain) SuperAdminUserManagement(in *SuperAdminUserManagementIn) (out SuperAdminUserManagementOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	// TODO: do something

	return
}
