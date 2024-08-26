package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file ReportViewerGeneralLedger.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type ReportViewerGeneralLedger.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type ReportViewerGeneralLedger.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type ReportViewerGeneralLedger.go
//go:generate farify doublequote --file ReportViewerDashboard.go

type (
	ReportViewerGeneralLedgerIn struct {
		RequestCommon
	}
	ReportViewerGeneralLedgerOut struct {
		ResponseCommon
	}
)

const (
	ReportViewerGeneralLedgerAction = `reportViewer/generalLedger`
)

func (d *Domain) ReportViewerGeneralLedger(in *ReportViewerGeneralLedgerIn) (out ReportViewerGeneralLedgerOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	return
}
