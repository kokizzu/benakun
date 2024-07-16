package domain

import (
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file DataEntryTemplates.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type DataEntryTemplates.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type DataEntryTemplates.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type DataEntryTemplates.go
//go:generate farify doublequote --file DataEntryTemplates.go

type (
	DataEntryTemplatesIn struct {
		RequestCommon
	}
	DataEntryTemplatesOut struct {
		ResponseCommon
	}
)

const (
	DataEntryTemplatesAction = `dataEntry/templates`

	ErrDataEntryTemplatesUnauthorized = `unauthorized user`
)

func (d *Domain) DataEntryTemplates(in *DataEntryTemplatesIn) (out DataEntryTemplatesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrDataEntryTemplatesUnauthorized)
		return
	}

	return
}
