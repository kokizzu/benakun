package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"fmt"
	"time"

	"github.com/kokizzu/gotro/X"
	"github.com/rs/zerolog/log"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserJoinCompany.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserJoinCompany.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserJoinCompany.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserJoinCompany.go
//go:generate farify doublequote --file UserJoinCompany.go

type (
	UserJoinCompanyIn struct {
		RequestCommon
		Company string `json:"company" form:"company" query:"company" long:"company" msg:"company"`
	}

	UserJoinCompanyOut struct {
		ResponseCommon
		Ok      bool         `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`
		Company *rqAuth.Orgs `json:"company" form:"company" query:"company" long:"company" msg:"company"`
	}
)

const (
	ErrUserJoinCompanyUserNotFound    = `user not found`
	ErrUserJoinCompanyCompanyNotFound = `company not found`
)

func (d *Domain) UserJoinCompany(in *UserJoinCompanyIn) (out UserJoinCompanyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.Company == `` {
		out.SetError(400, ErrUserJoinCompanyCompanyNotFound)
		return
	}

	log.Print(`Company = `, in.Company)

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrUserCreateCompanyUserNotFound)
		return
	}

	org := rqAuth.NewOrgs(d.AuthOltp)
	org.Name = in.Company
	if !org.FindByCompanyName() {
		out.SetError(400, ErrUserJoinCompanyCompanyNotFound)
		return
	}

	currentTime := time.Now()
	state := fmt.Sprintf("tenant:[%s]:accepted:%v", org.TenantCode, currentTime.Format("2006/01/02"))
	user.SetInvitationState(state)
	user.SetInvitedAt(X.ToS(in.UnixNow()))

	out.Company = org
	return
}
