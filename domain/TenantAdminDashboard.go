package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminDashboard.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminDashboard.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminDashboard.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminDashboard.go
//go:generate farify doublequote --file TenantAdminDashboard.go

type (
	TenantAdminDashboardIn struct {
		RequestCommon
	}
	TenantAdminDashboardOut struct {
		ResponseCommon
		Staffs []rqAuth.StaffWithInvitation `json:"staffs" form:"staffs" query:"staffs" long:"staffs" msg:"staffs"`
	}
)

const (
	TenantAdminDashboardAction = `tenantAdmin/dashboard`

	ErrTenantAdminDashboardUnauthorized   = `unauthorized user`
	ErrTenantAdminDashboardTenantNotFound = `tenant admin not found`
)

func (d *Domain) TenantAdminDashboard(in *TenantAdminDashboardIn) (out TenantAdminDashboardOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminDashboardUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminDashboardTenantNotFound)
		return
	}

	staffs := user.FindUsersByTenant(tenant.TenantCode)

	out.Staffs = staffs
	return
}
