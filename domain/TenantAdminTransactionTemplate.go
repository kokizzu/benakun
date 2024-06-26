package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
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
	ErrTenantAdminTransactionTemplateUserNotFound = `user not found` 
	ErrTenantAdminTransactionTemplateNotFound = `transaction template not found`
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
			InputType: zCrud.InputTypeColor,
		},
		{
			Name: mFinance.ImageURL,
			Label: "Gambar / Image",
			ReadOnly: true,
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

	switch in.Cmd {
	case zCrud.CmdForm:
		if in.TransactionTemplate.Id <= 0 {
			out.Meta = &TenantAdminTransactionTemplateMeta
			return
		}

		trxTemplate := rqFinance.NewTransactionTemplate(d.AuthOltp)
		trxTemplate.Id = in.TransactionTemplate.Id
		if !trxTemplate.FindById() {
			out.SetError(400, ErrTenantAdminTransactionTemplateNotFound)
			return
		}

		out.TransactionTemplate = trxTemplate
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.TenantCode = user.TenantCode
		if !tenant.FindByTenantCode() {
			out.SetError(400, ErrTenantAdminTransactionTemplateTenantNotFound)
			return
		}

		trxTemplate := wcFinance.NewTransactionTemplateMutator(d.AuthOltp)
		trxTemplate.Id = in.TransactionTemplate.Id
		if trxTemplate.Id > 0 {
			if !trxTemplate.FindById() {
				out.SetError(400, ErrTenantAdminTransactionTemplateNotFound)
				return
			}
			
			if in.Cmd == zCrud.CmdDelete {
				if trxTemplate.DeletedAt == 0 {
					trxTemplate.SetDeletedAt(in.UnixNow())
					trxTemplate.SetDeletedBy(sess.UserId)
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if trxTemplate.DeletedAt > 0 {
					trxTemplate.SetDeletedAt(0)
					trxTemplate.SetRestoredBy(sess.UserId)
				}
			}
		} else {
			trxTemplate.SetTenantCode(user.TenantCode)
		}

		if in.TransactionTemplate.Name != `` &&  in.TransactionTemplate.Name != trxTemplate.Name {
			trxTemplate.SetName(in.TransactionTemplate.Name)
		}

		if in.TransactionTemplate.Color != `` && in.TransactionTemplate.Color != trxTemplate.Color {
			trxTemplate.SetColor(in.TransactionTemplate.Color)
		}

		if trxTemplate.HaveMutation() {
			trxTemplate.SetUpdatedAt(in.UnixNow())
			trxTemplate.SetUpdatedBy(sess.UserId)
			if trxTemplate.Id == 0 {
				trxTemplate.SetCreatedAt(in.UnixNow())
				trxTemplate.SetCreatedBy(sess.UserId)
			}
		}

		if !trxTemplate.DoUpsertById() {
			out.SetError(500, ErrTenantAdminProductsSaveFailed)
		}

		out.TransactionTemplate = &trxTemplate.TransactionTemplate

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		r := rqFinance.NewTransactionTemplate(d.AuthOltp)
		r.TenantCode = user.TenantCode
		out.TransactionTemplates = r.FindByPagination(&TenantAdminTransactionTemplateMeta, &in.Pager, &out.Pager)
	}

	r := rqFinance.NewTransactionTemplate(d.AuthOltp)
	r.TenantCode = user.TenantCode
	out.TransactionTemplates = r.FindByPagination(&TenantAdminTransactionTemplateMeta, &in.Pager, &out.Pager)
	
	return
}
