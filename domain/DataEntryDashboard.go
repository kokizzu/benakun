package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file DataEntryDashboard.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type DataEntryDashboard.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type DataEntryDashboard.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type DataEntryDashboard.go
//go:generate farify doublequote --file DataEntryDashboard.go

type (
	DataEntryDashboardIn struct {
		RequestCommon
	}
	DataEntryDashboardOut struct {
		ResponseCommon
	}
)

const (
	DataEntryDashboardAction = `dataEntry/dashboard`
)

func (d *Domain) DataEntryDashboard(in *DataEntryDashboardIn) (out DataEntryDashboardOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	return
}
