package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
	"benakun/model/zCrud"
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
		Cmd      string               `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		CoaId 	 uint64 `json:"coaId" form:"coaId" query:"coaId" long:"coaId" msg:"coaId"`
		TransactionTplId 	 uint64 `json:"transactionTplId" form:"transactionTplId" query:"transactionTplId" long:"transactionTplId" msg:"transactionTplId"`
		TransactionJournals []rqFinance.TransactionJournal `json:"transactionJournals" form:"transactionJournals" query:"transactionJournals" long:"transactionJournals" msg:"transactionJournals"`
		BusinessTransaction rqFinance.BusinessTransaction `json:"businessTransaction" form:"businessTransaction" query:"businessTransaction" long:"businessTransaction" msg:"businessTransaction"`
	}
	DataEntryTransactionEntryOut struct {
		ResponseCommon
		CoaChildren map[string]string `json:"coaChildren" form:"coaChildren" query:"coaChildren" long:"coaChildren" msg:"coaChildren"`
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
	ErrDataEntryTransactionEntryCoaIsEmpty = `coa id cannot be empty`
	ErrDataEntryTransactionEntryInvalidCoaChild = `invalid coa child`
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

	switch in.Cmd {
	case zCrud.CmdForm:
		coaChoices := coa.FindCoasChoicesChildByParentByTenant()
		out.CoaChildren = coaChoices
	case zCrud.CmdUpsert:
		trxTemplate := rqFinance.NewTransactionTemplate(d.AuthOltp)
		trxTemplate.Id = in.TransactionTplId
		if !trxTemplate.FindById() {
			out.SetError(400, ErrDataEntryTransactionEntryTransactionTemplateNotFound)
			return
		}

		if len(in.TransactionJournals) > 0 {
			for _, v := range in.TransactionJournals {
				trxJournal := wcFinance.NewTransactionJournalMutator(d.AuthOltp)
				trxJournal.SetTenantCode(tenant.TenantCode)

				if v.CoaId > 0 {
					coaChild := rqFinance.NewCoa(d.AuthOltp)
					coaChild.Id = v.CoaId
					if !coaChild.FindById() {
						out.SetError(400, ErrDataEntryTransactionEntryCoaNotFound)
						return
					}

					if coaChild.ParentId != coa.Id {
						out.SetError(400, ErrDataEntryTransactionEntryInvalidCoaChild)
						return
					}

					trxJournal.SetCoaId(coaChild.Id)
				} else {
					trxJournal.SetCoaId(coa.Id)
				}
				trxJournal.SetDebitIDR(v.DebitIDR)
				trxJournal.SetCreditIDR(v.CreditIDR)
				trxJournal.SetDescriptions(v.Descriptions)

				if !isStrIsoDate(v.Date) {
					out.SetError(400, ErrDataEntryTransactionEntryInvalidDate)
					return
				}
				trxJournal.SetDate(v.Date)

				if v.DetailObj != `` {
					if !mFinance.IsValidDetailObject(v.DetailObj) {
						out.SetError(400, ErrDataEntryTransactionEntryInvalidDetailObject)
						return
					}
				}
				
				trxJournal.SetTransactionTemplateId(trxTemplate.Id)
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
			out.SetError(400, ErrDataEntryTransactionEntrySaveFailed)
			return
		}
	}
	
	return
}

func isStrIsoDate(dateStr string) bool {
	_, err := time.Parse("2006-01-02", dateStr)
	return err == nil
}