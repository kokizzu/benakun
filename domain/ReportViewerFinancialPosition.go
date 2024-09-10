package domain

import "benakun/model/mFinance/rqFinance"

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file ReportViewerFinancialPosition.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type ReportViewerFinancialPosition.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type ReportViewerFinancialPosition.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type ReportViewerFinancialPosition.go
//go:generate farify doublequote --file ReportViewerDashboard.go

type (
	ReportViewerFinancialPositionIn struct {
		RequestCommon
		Coa        rqFinance.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
	}
	ReportViewerFinancialPositionOut struct {
		ResponseCommon
		TransactionJournals []rqFinance.TransactionJournal `json:"transactionJournals" form:"transactionJournals" query:"transactionJournals" long:"transactionJournals" msg:"transactionJournals"`
	}
)

const (
	ReportViewerFinancialPositionAction = `reportViewer/financialPosition`

	ErrReportViewerFinancialPositionUnauthorized = `you are unauthorized as report viewer`
)

func (d *Domain) ReportViewerFinancialPosition(in *ReportViewerFinancialPositionIn) (out ReportViewerFinancialPositionOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustReportViewer(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrReportViewerFinancialPositionUnauthorized)
		return
	}

	if !sess.IsReportViewer {
		out.SetError(400, ErrReportViewerFinancialPositionUnauthorized)
	}
	return
}
