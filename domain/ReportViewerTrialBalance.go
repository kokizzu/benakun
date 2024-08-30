package domain

import (
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file ReportViewerTrialBalance.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type ReportViewerTrialBalance.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type ReportViewerTrialBalance.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type ReportViewerTrialBalance.go
//go:generate farify doublequote --file ReportViewerDashboard.go

type (
	ReportViewerTrialBalanceIn struct {
		RequestCommon
		TransactionJournal rqFinance.TransactionJournal `json:"transactionJournal" form:"transactionJournal" query:"transactionJournal" long:"transactionJournal" msg:"transactionJournal"`
	}
	ReportViewerTrialBalanceOut struct {
		ResponseCommon
		TransactionJournals []rqFinance.TransactionJournal `json:"transactionJournals" form:"transactionJournals" query:"transactionJournals" long:"transactionJournals" msg:"transactionJournals"`
	}
)

const (
	ReportViewerTrialBalanceAction = `reportViewer/trialBalance`

	ErrReportViewerTrialBalanceUnauthorized = `you are unauthorized as report viewer`
	ErrReportViewerTrialBalanceInvalidDate 								= `invalid date format, must be use format "2006-01-02"`
	ErrReportViewerTrialBalanceNotTenantHost = `invalid tenant, must be in valid domain`
)

func (d *Domain) ReportViewerTrialBalance(in *ReportViewerTrialBalanceIn) (out ReportViewerTrialBalanceOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustReportViewer(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrDataEntryTransactionEntryUnauthorized)
		return
	}

	if !sess.IsReportViewer {
		out.SetError(400, ErrReportViewerTrialBalanceUnauthorized)
	}

	tenantCode, err := GetTenantCodeByHost(in.Host)
	if err != nil {
		out.SetError(400, ErrReportViewerTrialBalanceNotTenantHost)
		return
	}

	if !mFinance.IsValidDate(in.TransactionJournal.Date) {
		out.SetError(400, ErrReportViewerTrialBalanceInvalidDate)
		return
	}

	trxJournal := rqFinance.NewTransactionJournal(d.AuthOltp)
	trxJournal.TenantCode = tenantCode
	trxJournal.Date = in.TransactionJournal.Date

	out.TransactionJournals = trxJournal.FindTrxJournalsByDateByTenant()

	return
}
