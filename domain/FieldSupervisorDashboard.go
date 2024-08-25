package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file FieldSupervisor.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type FieldSupervisor.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type FieldSupervisor.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type FieldSupervisor.go
//go:generate farify doublequote --file FieldSupervisor.go

type (
	FieldSupervisorIn struct {
		RequestCommon
	}
	FieldSupervisorOut struct {
		ResponseCommon
	}
)

const (
	FieldSupervisorAction = `fieldSupervisor/dashboard`
)

func (d *Domain) FieldSupervisor(in *FieldSupervisorIn) (out FieldSupervisorOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	return
}
