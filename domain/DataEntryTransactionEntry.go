package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
	"time"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file DataEntryTransactionEntry.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type DataEntryTransactionEntry.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type DataEntryTransactionEntry.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type DataEntryTransactionEntry.go
//go:generate farify doublequote --file DataEntryTransactionEntry.go

type (
	DataEntryTransactionEntryIn struct {
		RequestCommon
		CoaId uint64 `json:"coaId" form:"coaId" query:"coaId" long:"coaId" msg:"coaId"`
		TransactionJournals []rqFinance.TransactionJournal `json:"transactionJournals" form:"transactionJournals" query:"transactionJournals" long:"transactionJournals" msg:"transactionJournals"`
		BusinessTransaction rqFinance.BusinessTransaction `json:"businessTransaction" form:"businessTransaction" query:"businessTransaction" long:"businessTransaction" msg:"businessTransaction"`
	}
	DataEntryTransactionEntryOut struct {
		ResponseCommon
		
	}
)

const (
	DataEntryTransactionEntryAction = `dataEntry/transactionEntry`

	ErrDataEntryTransactionEntryUnauthorized = `you are unauthorized as data entry`
	ErrDataEntryTransactionEntryUserNotFound = `user not found`
	ErrDataEntryTransactionEntryTenantNotFound = `tenant not found`
	ErrDataEntryTransactionEntryNotDataEntry = `must be data entry to do this operation`
	ErrDataEntryTransactionEntryInvalidUserRole = `invalid role to this tenant`
	ErrDataEntryTransactionEntryCoaNotFound = `coa not found to journaling this transaction`
	ErrDataEntryTransactionEntryInvalidDate = `invalid date format, must be use format "01/02/2006"`
	ErrDataEntryTransactionEntryInvalidDetailObject = `invalid detail format`
	ErrDataEntryTransactionEntrySaveFailed = `failed to save transaction journal`
	ErrDataEntryTransactionEntryTransactionTemplateNotFound = `transaction template not found`
)

func (d *Domain) DataEntryTransactionEntry(in *DataEntryTransactionEntryIn) (out DataEntryTransactionEntryOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrDataEntryTransactionEntryUnauthorized)
		return
	}

	if !sess.IsDataEntry {
		out.SetError(400, ErrDataEntryTransactionEntryNotDataEntry)
		return
	}

	tenantCode, err := GetTenantCodeByHost(in.Host)
	if err != nil {
		out.SetError(400, ErrDataEntryTransactionEntryTenantNotFound)
		return
	}

	user := rqAuth.NewUsers(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrDataEntryTransactionEntryUserNotFound)
		return
	}

	mapState, err := mAuth.ToInvitationStateMap(user.InvitationState)
	if err != nil {
		out.SetError(400, ErrDataEntryTransactionEntryInvalidUserRole)
		return
	}

	if mapState.GetRoleByTenantCode(tenantCode) != mAuth.RoleDataEntry {
		out.SetError(400, ErrDataEntryTransactionEntryNotDataEntry)
		return
	}

	tenant := rqAuth.NewTenants(d.AuthOltp)
	tenant.TenantCode = tenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrDataEntryTransactionEntryTenantNotFound)
		return
	}

	coa := rqFinance.NewCoa(d.AuthOltp)
	coa.Id = in.CoaId
	if !coa.FindById() {
		out.SetError(400, ErrDataEntryTransactionEntryCoaNotFound)
		return
	}

	trxJournal := wcFinance.NewTransactionJournalMutator(d.AuthOltp)

	if len(in.TransactionJournals) > 0 {
		for _, v := range in.TransactionJournals {
			trxJournal.SetTenantCode(tenant.TenantCode)
			trxJournal.SetCoaId(coa.Id)
			trxJournal.SetDebitIDR(v.DebitIDR)
			trxJournal.SetCreditIDR(v.CreditIDR)
			trxJournal.SetDescriptions(v.Descriptions)

			if !isStrIsoDate(v.Date) {
				out.SetError(400, ErrDataEntryTransactionEntryInvalidDate)
				return
			}
			trxJournal.SetDate(v.Date)

			if !mFinance.IsValidDetailObject(v.DetailObj) {
				out.SetError(400, ErrDataEntryTransactionEntryInvalidDetailObject)
				return
			}
			trxJournal.SetDetailObj(v.DetailObj)
			trxJournal.SetCreatedAt(in.UnixNow())
			trxJournal.SetCreatedBy(sess.UserId)
			trxJournal.SetUpdatedAt(in.UnixNow())
			trxJournal.SetUpdatedBy(sess.UserId)

			if !trxJournal.DoInsert() {
				out.SetError(400, ErrDataEntryTransactionEntrySaveFailed)
				return
			}
		}
	}

	trxTemplate := rqFinance.NewTransactionTemplate(d.AuthOltp)
	trxTemplate.Id = in.BusinessTransaction.TransactionTemplateId
	if !trxTemplate.FindById() {
		out.SetError(400, ErrDataEntryTransactionEntryTransactionTemplateNotFound)
		return
	}

	businessTrx := wcFinance.NewBusinessTransactionMutator(d.AuthOltp)
	businessTrx.SetTenantCode(tenant.TenantCode)
	businessTrx.SetStartDate(in.BusinessTransaction.StartDate)
	businessTrx.SetEndDate(in.BusinessTransaction.EndDate)
	businessTrx.SetCreatedAt(in.UnixNow())
	businessTrx.SetCreatedBy(sess.UserId)
	businessTrx.SetUpdatedAt(in.UnixNow())
	businessTrx.SetUpdatedBy(sess.UserId)
	
	return
}

func isStrIsoDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}