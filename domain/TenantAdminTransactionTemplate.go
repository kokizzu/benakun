package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
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
		WithMeta		bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    		zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		TransactionTemplate *rqFinance.TransactionTemplate `json:"transactionTemplate" form:"transactionTemplate" query:"transactionTemplate" long:"transactionTemplate" msg:"transactionTemplate"`
	}
	TenantAdminTransactionTemplateOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		TransactionTemplates [][]any `json:"transactionTemplates" form:"transactionTemplates" query:"transactionTemplates" long:"transactionTemplates" msg:"transactionTemplates"`
		TransactionTemplate *rqFinance.TransactionTemplate `json:"transactionTemplate" form:"transactionTemplate" query:"transactionTemplate" long:"transactionTemplate" msg:"transactionTemplate"`
	}
)

const (
	TenantAdminTransactionTemplateAction = `tenantAdmin/transactionTemplate`

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
			Name: mFinance.Id,
			Label: "ID",
			DataType: zCrud.DataTypeInt,
			ReadOnly: true,
		},
		{
			Name: mFinance.Name,
			Label: "Nama / Name",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name: mFinance.Color,
			Label: "Warna / Color",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name: mFinance.ImageURL,
			Label: "Gambar / Image",
		},
		{
			Name:      mFinance.CreatedAt,
			Label:     `Dibuat pada / Created at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mFinance.UpdatedAt,
			Label:     `Diperbarui pada / Updated at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mFinance.DeletedAt,
			Label:     `Dihapus pada / Deleted at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
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

	r := rqFinance.NewTransactionTemplate(d.AuthOltp)
	r.TenantCode = user.TenantCode
	out.TransactionTemplates = r.FindByPagination(&TenantAdminTransactionTemplateMeta, &in.Pager, &out.Pager)
	
	return
}
