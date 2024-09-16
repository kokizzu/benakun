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
		Cmd      string               `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		StartDate string `json:"startDate" form:"startDate" query:"startDate" long:"startDate" msg:"startDate"`
		EndDate string `json:"endDate" form:"endDate" query:"endDate" long:"endDate" msg:"endDate"`
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

	if !mFinance.IsValidDate(in.StartDate) {
		out.SetError(400, ErrReportViewerTrialBalanceInvalidDate)
		return
	}

	if !mFinance.IsValidDate(in.EndDate) {
		out.SetError(400, ErrReportViewerTrialBalanceInvalidDate)
		return
	}

	trxJournal := rqFinance.NewTransactionJournal(d.AuthOltp)
	trxJournal.TenantCode = tenantCode

	out.TransactionJournals = trxJournal.FindTrxJournalsByDateByTenant(
		in.StartDate, in.EndDate,
	)

	return
}
