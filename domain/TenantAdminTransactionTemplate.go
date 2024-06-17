package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminTransactionTemplate.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminTransactionTemplate.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminTransactionTemplate.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminTransactionTemplate.go
//go:generate farify doublequote --file TenantAdminTransactionTemplate.go

type (
	TenantAdminTransactionTemplateIn struct {
		RequestCommon
		Cmd      		string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		StaffEmail	string				`json:"staffEmail" form:"staffEmail" query:"staffEmail" long:"staffEmail" msg:"staffEmail"`
		TenantCode  string 				`json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
		Role 				string 				`json:"role" form:"role" query:"role" long:"role" msg:"role"`
		WithMeta		bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    		zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	TenantAdminTransactionTemplateOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Staffs [][]any `json:"staffs" form:"staffs" query:"staffs" long:"staffs" msg:"staffs"`
		StaffsForm *[]rqAuth.Staff `json:"staffsForm" form:"staffsForm" query:"staffsForm" long:"staffsForm" msg:"staffsForm"`
	}
)

const (
	TenantAdminTransactionTemplateAction = `tenantAdmin/dashboard`

	ErrTenantAdminTransactionTemplateUnauthorized   = `unauthorized user`
	ErrTenantAdminTransactionTemplateTenantNotFound = `tenant admin not found`
	ErrTenantAdminTransactionTemplateStaffEmailRequired = `staff email is required`
	ErrTenantAdminTransactionTemplateUserNotFound = `user not found` 
	ErrTenantAdminTransactionTemplateInvalidStaff = `invalid staff`
	ErrTenantAdminTransactionTemplateEmptyState = `failed to modify staff, state is empty`
	ErrTenantAdminTransactionTemplateFailed = `failed to update staff`
	ErrTenantAdminTransactionTemplateNotTenant = `cannot invite user if not tenant`
	ErrTenantAdminTransactionTemplateInvalidRole = `invalid staff role to modify`
)

var TenantAdminTransactionTemplateMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mAuth.Id,
			Label:     "ID",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly: true,
		},
	},
}

func (d *Domain) TenantAdminTransactionTemplate(in *TenantAdminTransactionTemplateIn) (out TenantAdminTransactionTemplateOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminTransactionTemplateUnauthorized)
		return
	}

	if in.WithMeta {
		out.Meta = &TenantAdminTransactionTemplateMeta
	}

	return
}
