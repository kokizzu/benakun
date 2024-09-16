package domain

import "benakun/model/mFinance/rqFinance"

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file ReportViewerLossIncomeStatements.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type ReportViewerLossIncomeStatements.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type ReportViewerLossIncomeStatements.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type ReportViewerLossIncomeStatements.go
//go:generate farify doublequote --file ReportViewerDashboard.go

type (
	ReportViewerLossIncomeStatementsIn struct {
		RequestCommon
		Coa        rqFinance.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
	}
	ReportViewerLossIncomeStatementsOut struct {
		ResponseCommon
		TransactionJournals []rqFinance.TransactionJournal `json:"transactionJournals" form:"transactionJournals" query:"transactionJournals" long:"transactionJournals" msg:"transactionJournals"`
	}
)

const (
	ReportViewerLossIncomeStatementsAction = `reportViewer/lossIncomeStatements`

	ErrReportViewerLossIncomeStatementsUnauthorized = `you are unauthorized as report viewer`
)

func (d *Domain) ReportViewerLossIncomeStatements(in *ReportViewerLossIncomeStatementsIn) (out ReportViewerLossIncomeStatementsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustReportViewer(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrReportViewerLossIncomeStatementsUnauthorized)
		return
	}

	if !sess.IsReportViewer {
		out.SetError(400, ErrReportViewerLossIncomeStatementsUnauthorized)
	}
	return
}
