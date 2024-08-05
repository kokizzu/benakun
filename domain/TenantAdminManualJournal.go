package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminManualJournal.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminManualJournal.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminManualJournal.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminManualJournal.go
//go:generate farify doublequote --file TenantAdminManualJournal.go

type (
	TenantAdminManualJournalIn struct {
		RequestCommon
		Cmd      string               `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool                 `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn        `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		TransactionJournal  rqFinance.TransactionJournal `json:"transactionJournal" form:"transactionJournal" query:"transactionJournal" long:"transactionJournal" msg:"transactionJournal"`
	}
	TenantAdminManualJournalOut struct {
		ResponseCommon
		Pager    						zCrud.PagerOut       `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta     						*zCrud.Meta          `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		TransactionJournal  *rqFinance.TransactionJournal `json:"transactionJournal" form:"transactionJournal" query:"transactionJournal" long:"transactionJournal" msg:"transactionJournal"`
		TransactionJournals [][]any              `json:"transactionJournals" form:"transactionJournals" query:"transactionJournals" long:"transactionJournals" msg:"transactionJournals"`
	}
)

const (
	TenantAdminManualJournalAction = `tenantAdmin/manualJournal`
)

var TenantAdminManualJournalMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:     mFinance.Id,
			Label:    "ID",
			DataType: zCrud.DataTypeInt,
			ReadOnly: true,
		},
		{
			Name: mFinance.CoaId,
			Label: "CoA (Chart of Account)",
			DataType: zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name: mFinance.DebitIDR,
			Label: "Debit (IDR)",
			DataType: zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name: mFinance.CreditIDR,
			Label: "Credit (IDR)",
			DataType: zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name: mFinance.Description,
			Label: "Deskripsi / Description",
			DataType: zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name: mFinance.DetailObj,
			Label: "Detail",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
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

func (d *Domain) TenantAdminManualJournal(in *TenantAdminManualJournalIn) (out TenantAdminManualJournalOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &TenantAdminManualJournalMeta
	}

	user := rqAuth.NewUsers(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminProductsUnauthorized)
		return
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		r := rqFinance.NewTransactionJournal(d.AuthOltp)
		r.TenantCode = user.TenantCode
		out.TransactionJournals = r.FindByPagination(&TenantAdminManualJournalMeta, &in.Pager, &out.Pager)
	}

	return
}
