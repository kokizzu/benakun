package domain

import (
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file DataEntryTemplate.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type DataEntryTemplate.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type DataEntryTemplate.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type DataEntryTemplate.go
//go:generate farify doublequote --file DataEntryTemplate.go

type (
	DataEntryTemplateIn struct {
		RequestCommon
	}
	DataEntryTemplateOut struct {
		ResponseCommon
	}
)

const (
	DataEntryTemplateAction = `dataEntry/template`

	ErrDataEntryTemplateUnauthorized = `unauthorized user`
)

func (d *Domain) DataEntryTemplate(in *DataEntryTemplateIn) (out DataEntryTemplateOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrDataEntryTemplateUnauthorized)
		return
	}

	return
}
