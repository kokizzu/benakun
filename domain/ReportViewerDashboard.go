package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file ReportViewerDashboard.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type ReportViewerDashboard.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type ReportViewerDashboard.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type ReportViewerDashboard.go
//go:generate farify doublequote --file ReportViewerDashboard.go

type (
	ReportViewerDashboardIn struct {
		RequestCommon
	}
	ReportViewerDashboardOut struct {
		ResponseCommon
	}
)

const (
	ReportViewerDashboardAction = `reportViewer/dashboard`
)

func (d *Domain) ReportViewerDashboard(in *ReportViewerDashboardIn) (out ReportViewerDashboardOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	return
}
