package domain

import (
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
		Staffs []Staff `json:"staffs" form:"staffs" query:"staffs" long:"staffs" msg:"staffs"`
	}
)

const (
	TenantAdminDashboardAction = `tenantAdmin/dashboard`

	ErrTenantAdminDashboardUnauthorized   = `unauthorized user`
	ErrTenantAdminDashboardTenantNotFound = `tenant admin not found`
)

type Staff struct {
	Id              string `json:"id" form:"id" query:"id" long:"id" msg:"id"`
	Email           string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	FullName        string `json:"fullName" form:"fullName" query:"fullName" long:"fullName" msg:"fullName"`
	InvitationState string `json:"invitationState" form:"invitationState" query:"invitationState" long:"invitationState" msg:"invitationState"`
	Role            string `json:"role" form:"role" query:"role" long:"role" msg:"role"`
}

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

	resp := user.FindUsersByTenant(tenant.TenantCode)
	var staffs []Staff
	if len(resp) > 0 {
		for _, stf := range resp {
			if len(stf) >= 5 {
				st := Staff{
					Id:              stf[0].(string),
					Email:           stf[1].(string),
					FullName:        stf[2].(string),
					InvitationState: stf[3].(string),
					Role:            stf[4].(string),
				}
				staffs = append(staffs, st)
			}
		}
		out.Staffs = staffs
	} else {
		out.Staffs = []Staff{}
	}
	return
}
