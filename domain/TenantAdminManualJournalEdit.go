package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminManualJournalEdit.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminManualJournalEdit.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminManualJournalEdit.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminManualJournalEdit.go
//go:generate farify doublequote --file TenantAdminManualJournalEdit.go

type (
	TenantAdminManualJournalEditIn struct {
		RequestCommon
		Cmd      string               `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		TransactionJournal  rqFinance.TransactionJournal `json:"transactionJournal" form:"transactionJournal" query:"transactionJournal" long:"transactionJournal" msg:"transactionJournal"`
	}
	TenantAdminManualJournalEditOut struct {
		ResponseCommon
		TransactionJournal  *rqFinance.TransactionJournal `json:"transactionJournal" form:"transactionJournal" query:"transactionJournal" long:"transactionJournal" msg:"transactionJournal"`
	}
)

const (
	TenantAdminManualJournalEditAction = `tenantAdmin/manualJournal/edit`

	ErrTenantAdminManualJournalEditUnauthorized 								= `unauthorized user`
	ErrTenantAdminManualJournalEditNotTenant    								= `must be tenant admin to do this operation`
	ErrTenantAdminManualJournalEditJournalNotFound							= `transaction journal not found`
)

func (d *Domain) TenantAdminManualJournalEdit(in *TenantAdminManualJournalEditIn) (out TenantAdminManualJournalEditOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := rqAuth.NewUsers(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminManualJournalEditUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminProductsNotTenant)
		return
	}
	
	switch in.Cmd {
	case zCrud.CmdForm:
		trxJournal := rqFinance.NewTransactionJournal(d.AuthOltp)
		trxJournal.Id = in.TransactionJournal.Id
		if !trxJournal.FindById() {
			out.SetError(400, ErrTenantAdminManualJournalEditJournalNotFound)
			return
		}
		out.TransactionJournal = trxJournal
	case zCrud.CmdUpsert:

	}

	return
}
