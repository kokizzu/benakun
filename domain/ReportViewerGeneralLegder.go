package domain

import "benakun/model/mFinance/rqFinance"

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file ReportViewerGeneralLedger.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type ReportViewerGeneralLedger.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type ReportViewerGeneralLedger.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type ReportViewerGeneralLedger.go
//go:generate farify doublequote --file ReportViewerDashboard.go

type (
	ReportViewerGeneralLedgerIn struct {
		RequestCommon
		Coa        rqFinance.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
	}
	ReportViewerGeneralLedgerOut struct {
		ResponseCommon
		TransactionJournals []rqFinance.TransactionJournal `json:"transactionJournals" form:"transactionJournals" query:"transactionJournals" long:"transactionJournals" msg:"transactionJournals"`
	}
)

const (
	ReportViewerGeneralLedgerAction = `reportViewer/generalLedger`

	ErrReportViewerGeneralLedgerUnauthorized = `you are unauthorized as report viewer`
	ErrReportViewerGeneralLedgerCoaEmpty = `coa cannot be empty`
	ErrReportViewerGeneralLedgerNotTenantHost = `invalid tenant, must be in valid domain`
)

func (d *Domain) ReportViewerGeneralLedger(in *ReportViewerGeneralLedgerIn) (out ReportViewerGeneralLedgerOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustReportViewer(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrReportViewerGeneralLedgerUnauthorized)
		return
	}

	if !sess.IsReportViewer {
		out.SetError(400, ErrReportViewerGeneralLedgerUnauthorized)
	}

	tenantCode, err := GetTenantCodeByHost(in.Host)
	if err != nil {
		out.SetError(400, ErrReportViewerGeneralLedgerNotTenantHost)
		return
	}

	if in.Coa.Id == 0 {
		out.SetError(400, ErrReportViewerGeneralLedgerCoaEmpty)
	}

	trxJournal := rqFinance.NewTransactionJournal(d.AuthOltp)
	trxJournal.TenantCode = tenantCode
	trxJournal.CoaId = in.Coa.Id
	out.TransactionJournals = trxJournal.FindTrxJournalsByCoaByTenant()
	
	return
}
