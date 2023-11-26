package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"crypto/rand"
	"fmt"
	"math/big"

	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserCreateCompany.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserCreateCompany.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserCreateCompany.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserCreateCompany.go
//go:generate farify doublequote --file UserCreateCompany.go

type (
	UserCreateCompanyIn struct {
		RequestCommon
		Company        rqAuth.Orgs `json:"company" form:"company" query:"company" long:"company" msg:"company"`
		UserTenantCode string      `json:"userTenantCode" form:"userTenantCode" query:"userTenantCode" long:"userTenantCode" msg:"userTenantCode"`
	}

	UserCreateCompanyOut struct {
		ResponseCommon
		Ok      bool         `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`
		Company *rqAuth.Orgs `json:"company" form:"company" query:"company" long:"company" msg:"company"`
	}
)

const (
	UserCreateCompanyAction = `user/createCompany`

	ErrCreateCompanyUserNotFound = `user not found`
)

func (d *Domain) UserCreateCompany(in *UserCreateCompanyIn) (out UserCreateCompanyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrCreateCompanyUserNotFound)
		return
	}

	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	org.SetAll(in.Company, M.SB{}, M.SB{})
	org.SetTenantCode(fmt.Sprintf("%s_%d", in.UserTenantCode, generateRandomNumber()))

	org.Adapter = nil
	out.Company = &org.Orgs
	return
}

func generateRandomNumber() int {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(9000))
	return int(randomNumber.Int64()) + 1000
}
