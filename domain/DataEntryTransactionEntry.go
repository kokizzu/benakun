package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance/rqFinance"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file DataEntryTransactionEntry.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type DataEntryTransactionEntry.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type DataEntryTransactionEntry.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type DataEntryTransactionEntry.go
//go:generate farify doublequote --file DataEntryTransactionEntry.go

type (
	DataEntryTransactionEntryIn struct {
		RequestCommon
		TransactionTemplate *rqFinance.TransactionTemplate `json:"transactionTemplate" form:"transactionTemplate" query:"transactionTemplate" long:"transactionTemplate" msg:"transactionTemplate"`
		TransactionJournal  *rqFinance.TransactionJournal `json:"transactionJournal" form:"transactionJournal" query:"transactionJournal" long:"transactionJournal" msg:"transactionJournal"`
		BusinessTransaction *rqFinance.BusinessTransaction `json:"businessTransaction" form:"businessTransaction" query:"businessTransaction" long:"businessTransaction" msg:"businessTransaction"`
	}
	DataEntryTransactionEntryOut struct {
		ResponseCommon
		TransactionTplDetails *[]rqFinance.TransactionTplDetail `json:"transactionTplDetails" form:"transactionTplDetails" query:"transactionTplDetails" long:"transactionTplDetails" msg:"transactionTplDetails"`
	}
)

const (
	DataEntryTransactionEntryAction = `dataEntry/transactionEntry`
)

func (d *Domain) DataEntryTransactionEntry(in *DataEntryTransactionEntryIn) (out DataEntryTransactionEntryOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrTenantAdminProductsUnauthorized)
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminProductsUnauthorized)
		return
	}

	trxTplDetail := rqFinance.NewTransactionTplDetail(d.AuthOltp)
	trxTplDetail.ParentId = in.TransactionTemplate.Id
	trxTplDetail.TenantCode = user.TenantCode
	trxTplDetails := trxTplDetail.FindTrxTplDetailsByTenantByTrxTplId()

	out.TransactionTplDetails = &trxTplDetails
	
	return
}
