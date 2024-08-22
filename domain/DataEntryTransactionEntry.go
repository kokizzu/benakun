package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
	"benakun/model/zCrud"

	"github.com/goccy/go-json"
	"github.com/kokizzu/gotro/S"
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
		TransactionTplId 	 uint64 `json:"transactionTplId" form:"transactionTplId" query:"transactionTplId" long:"transactionTplId" msg:"transactionTplId"`
		TransactionTplDetailsId 	 []uint64 `json:"transactionTplDetailsId" form:"transactionTplDetailsId" query:"transactionTplDetailsId" long:"transactionTplDetailsId" msg:"transactionTplDetailsId"`
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

	ErrDataEntryTransactionEntryUnauthorized 								= `you are unauthorized as data entry`
	ErrDataEntryTransactionEntryUserNotFound 								= `user not found`
	ErrDataEntryTransactionEntryTenantNotFound 							= `tenant not found`
	ErrDataEntryTransactionEntryNotDataEntry 								= `must be data entry to do this operation`
	ErrDataEntryTransactionEntryInvalidUserRole 						= `invalid role to this tenant`
	ErrDataEntryTransactionEntryCoaNotFound 								= `coa not found to journaling this transaction`
	ErrDataEntryTransactionEntryCoaIsEmpty 									= `coa id cannot be empty`
	ErrDataEntryTransactionEntryInvalidCoaChild 						= `invalid coa child`
	ErrDataEntryTransactionEntryInvalidDate 								= `invalid date format, must be use format "2006-01-02"`
	ErrDataEntryTransactionEntryInvalidDetailObject 				= `invalid format for detail object`
	ErrDataEntryTransactionEntrySaveFailed 									= `failed to save transaction journal`
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

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert:
		trxTemplate := rqFinance.NewTransactionTemplate(d.AuthOltp)
		trxTemplate.Id = in.TransactionTplId
		if !trxTemplate.FindById() {
			out.SetError(400, ErrDataEntryTransactionEntryTransactionTemplateNotFound)
			return
		}

		if len(in.TransactionJournals) > 0 {
			for idx, v := range in.TransactionJournals {
				trxTplDetail := rqFinance.NewTransactionTplDetail(d.AuthOltp)
				trxTplDetail.Id = in.TransactionTplDetailsId[idx]
				if !trxTplDetail.FindById() {
					out.SetError(400, `todo error`)
					return
				}

				trxJournal := wcFinance.NewTransactionJournalMutator(d.AuthOltp)
				trxJournal.SetTenantCode(tenant.TenantCode)
				
				if v.CoaId > 0 {
					coaChild := rqFinance.NewCoa(d.AuthOltp)
					coaChild.Id = v.CoaId
					if !coaChild.FindById() {
						out.SetError(400, ErrDataEntryTransactionEntryCoaNotFound)
						return
					}

					// if coaChild.ParentId != trxTplDetail.CoaId {
					// 	out.SetError(400, ErrDataEntryTransactionEntryInvalidCoaChild)
					// 	return
					// }

					trxJournal.SetCoaId(coaChild.Id)
				} else {
					trxJournal.SetCoaId(trxTplDetail.CoaId)
				}

				trxJournal.SetDescriptions(v.Descriptions)

				if !mFinance.IsValidDate(v.Date) {
					out.SetError(400, ErrDataEntryTransactionEntryInvalidDate)
					return
				}
				trxJournal.SetDate(v.Date)

				parsedAttributes := trxTplDetail.ParseAttributes()
				var creditDebitIDR int64 = 0

				if v.DetailObj != `` && parsedAttributes.IsSales {
					var trxJournalDetailObj mFinance.TransactionJournalDetailObject
					err := json.Unmarshal([]byte(v.DetailObj), &trxJournalDetailObj)
					if err != nil || S.ToI(trxJournalDetailObj.SalesPriceIDR) == 0{
						out.SetError(400, ErrDataEntryTransactionEntryInvalidDetailObject)
						return
					}

					creditDebitIDR = trxJournalDetailObj.SalesCount * S.ToI(trxJournalDetailObj.SalesPriceIDR)
					trxJournal.SetDetailObj(v.DetailObj)
				}

				if parsedAttributes.IsSales {
					if trxTplDetail.IsDebit {
						trxJournal.SetDebitIDR(creditDebitIDR)
					} else {
						trxJournal.SetCreditIDR(creditDebitIDR)
					}
				} else {
					trxJournal.SetDebitIDR(v.DebitIDR)
					trxJournal.SetCreditIDR(v.CreditIDR)
				}

				trxJournal.SetTransactionTemplateId(trxTemplate.Id)
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