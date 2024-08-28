package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file ReportViewerTrialBalance.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type ReportViewerTrialBalance.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type ReportViewerTrialBalance.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type ReportViewerTrialBalance.go
//go:generate farify doublequote --file ReportViewerDashboard.go

type (
	ReportViewerTrialBalanceIn struct {
		RequestCommon
	}
	ReportViewerTrialBalanceOut struct {
		ResponseCommon
	}
)

const (
	ReportViewerTrialBalanceAction = `reportViewer/trialBalance`
)

func (d *Domain) ReportViewerTrialBalance(in *ReportViewerTrialBalanceIn) (out ReportViewerTrialBalanceOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	return
}
