package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"

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
		Cmd      string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		User     rqAuth.Users  `json:"user" form:"user" query:"user" long:"user" msg:"user"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	TenantAdminDashboardOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Staffs []rqAuth.StaffWithInvitation `json:"staffs" form:"staffs" query:"staffs" long:"staffs" msg:"staffs"`
	}
)

const (
	TenantAdminDashboardAction = `tenantAdmin/dashboard`

	ErrTenantAdminDashboardUnauthorized   = `unauthorized user`
	ErrTenantAdminDashboardTenantNotFound = `tenant admin not found`
)


var TenantAdminDashboardMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mAuth.Id,
			Label:     "ID",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
		},
		{
			Name: mAuth.Email,
			Label: "Email",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeEmail,
		},
		{
			Name: mAuth.FullName,
			Label: "Full Name",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name: mAuth.Role,
			Label: "Role",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeCombobox,
			Ref:  []string{
				UserSegment, DataEntrySegment, TenantAdminSegment, ReportViewerSegment,
			},
		},
		{
			Name: mAuth.InvitationState,
			Label: "Status",
			ReadOnly: true,
		},
	},
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

	if in.WithMeta {
		out.Meta = &TenantAdminDashboardMeta
	}

	staffs := user.FindUsersByTenant(tenant.TenantCode)

	out.Staffs = staffs
	return
}
