package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminManualJournal.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminManualJournal.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminManualJournal.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminManualJournal.go
//go:generate farify doublequote --file TenantAdminManualJournal.go

type (
	TenantAdminManualJournalIn struct {
		RequestCommon
	}
	TenantAdminManualJournalOut struct {
		ResponseCommon
	}
)

const (
	TenantAdminManualJournalAction = `tenantAdmin/manualJournal`
)

func (d *Domain) TenantAdminManualJournal(in *TenantAdminManualJournalIn) (out TenantAdminManualJournalOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustSuperAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}
	return
}
