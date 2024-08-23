package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
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
		CoaId 	 uint64 `json:"coaId" form:"coaId" query:"coaId" long:"coaId" msg:"coaId"`
		TransactionTplId 	 uint64 `json:"transactionTplId" form:"transactionTplId" query:"transactionTplId" long:"transactionTplId" msg:"transactionTplId"`
		TransactionJournal  rqFinance.TransactionJournal `json:"transactionJournal" form:"transactionJournal" query:"transactionJournal" long:"transactionJournal" msg:"transactionJournal"`
		BusinessTransaction rqFinance.BusinessTransaction `json:"businessTransaction" form:"businessTransaction" query:"businessTransaction" long:"businessTransaction" msg:"businessTransaction"`
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

	ErrTenantAdminManualJournalUnauthorized 								= `unauthorized user`
	ErrTenantAdminManualJournalNotTenant    								= `must be tenant admin to do this operation`
	ErrTenantAdminManualJournalCoANotFound									= `coa not found to entry this journal`
	ErrTenantAdminManualJournalTransactionTemplateNotFound	= `transaction template not found to entry this journal` 
	ErrTenantAdminManualJournalSaveFailed										= `failed to save transaction journal`
	ErrTenantAdminManualJournalUpdateFailed									= `failed to update transaction journal`
	ErrTenantAdminManualJournalSaveFailedBusinessTrx				= `failed to save business transaction`
	ErrTenantAdminManualJournalNotFound											= `transaction journal not found`
	ErrTenantAdminManualJournalInvalidDetailObject 					= `invalid format for detail object`
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
			InputType: zCrud.InputTypeCombobox,
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
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeTextArea,
		},
		{
			Name: mFinance.Date,
			Label: "Tanggal / Date",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name: mFinance.DetailObj,
			Label: "Detail",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeTextArea,
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
		out.SetError(400, ErrTenantAdminManualJournalUnauthorized)
		return
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		trxJournal := rqFinance.NewTransactionJournal(d.AuthOltp)
		trxJournal.Id = in.TransactionJournal.Id
		if !trxJournal.FindById() {
			out.SetError(400, ErrTenantAdminManualJournalNotFound)
			return
		}
		
		out.TransactionJournal = trxJournal
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.TenantCode = user.TenantCode
		if !tenant.FindByTenantCode() {
			out.SetError(400, ErrTenantAdminProductsNotTenant)
			return
		}

		if in.Cmd == zCrud.CmdUpsert {
			if in.TransactionJournal.Id == 0 {
				coa := rqFinance.NewCoa(d.AuthOltp)
				coa.Id = in.CoaId
				if !coa.FindById() {
					out.SetError(400, ErrTenantAdminManualJournalCoANotFound)
					return
				}

				trxTemplate := rqFinance.NewTransactionTemplate(d.AuthOltp)
				trxTemplate.Id = in.TransactionTplId
				if !trxTemplate.FindById() {
					out.SetError(400, ErrTenantAdminManualJournalTransactionTemplateNotFound)
					return
				}

				trxJournal := wcFinance.NewTransactionJournalMutator(d.AuthOltp)
				trxJournal.SetTenantCode(tenant.TenantCode)
				trxJournal.SetCoaId(coa.Id)
				trxJournal.SetTransactionTemplateId(trxTemplate.Id)
				trxJournal.SetCreditIDR(in.TransactionJournal.CreditIDR)
				trxJournal.SetDebitIDR(in.TransactionJournal.DebitIDR)
				trxJournal.SetDescriptions(in.TransactionJournal.Descriptions)
				trxJournal.SetDate(in.TransactionJournal.Date)
				trxJournal.SetDetailObj(in.TransactionJournal.DetailObj)
				trxJournal.SetCreatedAt(in.UnixNow())
				trxJournal.SetCreatedBy(sess.UserId)
				trxJournal.SetUpdatedAt(in.UnixNow())
				trxJournal.SetUpdatedBy(sess.UserId)
				if !trxJournal.DoInsert() {
					out.SetError(400, ErrTenantAdminManualJournalSaveFailed)
					return
				}
				
				businessTrx := wcFinance.NewBusinessTransactionMutator(d.AuthOltp)
				businessTrx.SetTenantCode(tenant.TenantCode)
				businessTrx.SetStartDate(in.BusinessTransaction.StartDate)
				businessTrx.SetEndDate(in.BusinessTransaction.EndDate)
				businessTrx.SetTransactionTemplateId(trxTemplate.Id)
				businessTrx.SetCreatedAt(in.UnixNow())
				businessTrx.SetCreatedBy(sess.UserId)
				businessTrx.SetUpdatedAt(in.UnixNow())
				businessTrx.SetUpdatedBy(sess.UserId)

				if !businessTrx.DoInsert() {
					out.SetError(400, ErrTenantAdminManualJournalSaveFailedBusinessTrx)
					return
				}
			} else {
				trxJournal := wcFinance.NewTransactionJournalMutator(d.AuthOltp)
				trxJournal.Id = in.TransactionJournal.Id
				if !trxJournal.FindById() {
					out.SetError(400, ErrTenantAdminManualJournalNotFound)
					return
				}

				if in.TransactionJournal.DebitIDR != 0 && in.TransactionJournal.DebitIDR != trxJournal.DebitIDR{
					trxJournal.SetDebitIDR(in.TransactionJournal.DebitIDR)
				}

				if in.TransactionJournal.CreditIDR != 0 && in.TransactionJournal.CreditIDR != trxJournal.CreditIDR {
					trxJournal.SetCreditIDR(in.TransactionJournal.CreditIDR)
				}

				if in.TransactionJournal.Descriptions != `` && in.TransactionJournal.Descriptions != trxJournal.Descriptions {
					trxJournal.SetDescriptions(in.TransactionJournal.Descriptions)
				}

				if in.TransactionJournal.Date != `` && in.TransactionJournal.Date != trxJournal.Date {
					if mFinance.IsValidDate(in.TransactionJournal.Date) {
						trxJournal.SetDate(in.TransactionJournal.Date)
					}
				}

				if in.TransactionJournal.DetailObj != `` && in.TransactionJournal.DetailObj != trxJournal.DetailObj {
					if mFinance.IsValidDetailObject(in.TransactionJournal.DetailObj) {
						trxJournal.SetDetailObj(in.TransactionJournal.DetailObj)
					} else {
						out.SetError(400, ErrTenantAdminManualJournalInvalidDetailObject)
						return
					}
				}

				trxJournal.SetUpdatedAt(in.UnixNow())
				if !trxJournal.DoUpdateById() {
					out.SetError(400, ErrTenantAdminManualJournalUpdateFailed)
					return
				}
			}
		}

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
