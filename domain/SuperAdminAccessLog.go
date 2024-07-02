package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file SuperAdminAccessLog.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type SuperAdminAccessLog.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type SuperAdminAccessLog.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type SuperAdminAccessLog.go
//go:generate farify doublequote --file SuperAdminAccessLog.go

type (
	SuperAdminAccessLogIn struct {
		RequestCommon
	}

	SuperAdminAccessLogOut struct {
		ResponseCommon
	}
)

const (
	SuperAdminAccessLogAction = `superAdmin/accessLog`

	ErrSuperAdminAccessLogUnauthorized = `you are unauthorized to do this operation`
)

func (d *Domain) SuperAdminAccessLog(in *SuperAdminAccessLogIn) (out SuperAdminAccessLogOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustSuperAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrSuperAdminAccessLogUnauthorized)
		return
	}

	return
}
