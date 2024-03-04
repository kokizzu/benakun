package domain

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"

	"github.com/kokizzu/gotro/D/Tt"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserCreateCompany.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserCreateCompany.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserCreateCompany.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserCreateCompany.go
//go:generate farify doublequote --file UserCreateCompany.go

type (
	UserCreateCompanyIn struct {
		RequestCommon
		TenantCode  string `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
		CompanyName string `json:"companyName" form:"companyName" query:"companyName" long:"companyName" msg:"companyName"`
		HeadTitle   string `json:"headTitle" form:"headTitle" query:"headTitle" long:"headTitle" msg:"headTitle"`
	}

	UserCreateCompanyOut struct {
		ResponseCommon
		Ok      bool         `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`
		Company *rqAuth.Orgs `json:"company" form:"company" query:"company" long:"company" msg:"company"`
	}
)

const (
	UserCreateCompanyAction = `user/createCompany`

	ErrUserCreateCompanyUserNotFound       = `user not found`
	ErrUserCreateCompanyAlreadyAdded       = `company already exist`
	ErrUserCreateCompanyTenantCodeInvalid  = `tenant code must be valid`
	ErrUserCreateCompanyCompanyNameInvalid = `company name must be valid`
	ErrUserCreateCompanyHeadTitleInvalid   = `head title must be valid`
	ErrUserCreateCompanyCoaExist           = `coa already exist`
)

func (d *Domain) UserCreateCompany(in *UserCreateCompanyIn) (out UserCreateCompanyOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.TenantCode == `` {
		out.SetError(400, ErrUserCreateCompanyTenantCodeInvalid)
		return
	}

	if in.CompanyName == `` {
		out.SetError(400, ErrUserCreateCompanyCompanyNameInvalid)
		return
	}

	if in.HeadTitle == `` {
		out.SetError(400, ErrUserCreateCompanyHeadTitleInvalid)
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrUserCreateCompanyUserNotFound)
		return
	}

	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	org.SetTenantCode(fmt.Sprintf("%s_%d", in.TenantCode, generate4RandomNumber()))
	org.SetHeadTitle(in.HeadTitle)
	org.SetName(in.CompanyName)
	org.SetOrgType(mAuth.OrgTypeCompany)
	org.SetCreatedAt(in.UnixNow())
	org.SetCreatedBy(sess.UserId)
	org.SetUpdatedAt(in.UnixNow())
	org.SetUpdatedBy(sess.UserId)

	if !org.DoInsert() {
		out.SetError(400, ErrUserCreateCompanyAlreadyAdded)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.SetTenantCode(org.TenantCode)
	tenant.SetCreatedAt(in.UnixNow())
	tenant.SetCreatedBy(sess.UserId)
	tenant.SetUpdatedAt(in.UnixNow())
	tenant.SetUpdatedBy(sess.UserId)
	if !tenant.DoInsert() {
		out.SetError(400, ErrUserCreateCompanyAlreadyAdded)
	}

	user.SetRole(TenantAdminSegment)
	user.SetTenantCode(org.TenantCode)
	if !user.DoUpdateById() {
		out.SetError(400, ErrUserCreateCompanyUserNotFound)
		return
	}

	err := generateCoaLevels(d.AuthOltp, org.TenantCode)
	if err != nil {
		out.SetError(400, err.Error())
		return
	}

	out.Company = &org.Orgs
	return
}

func generate4RandomNumber() int {
	randomNumber, _ := rand.Int(rand.Reader, big.NewInt(9000))
	return int(randomNumber.Int64()) + 1000
}

func generateCoaLevels(ao *Tt.Adapter, tenantCode string) error {
	for i := 1; i < 7; i++ {
		name := ``
		if i == 1 {
			name = mAuth.CoaLevel1Name
		} else if i == 2 {
			name = mAuth.CoaLevel2Name
		} else if i == 3 {
			name = mAuth.CoaLevel3Name
		} else if i == 4 {
			name = mAuth.CoaLevel4Name
		} else if i == 5 {
			name = mAuth.CoaLevel5Name
		} else if i == 6 {
			name = mAuth.CoaLevel6Name
		} else if i == 7 {
			name = mAuth.CoaLevel7Name
		}
		coa := wcAuth.NewCoaMutator(ao)
		coa.SetTenantCode(tenantCode)
		coa.SetLevel(float64(i))
		coa.SetName(name)
		if !coa.DoInsert() {
			return errors.New(ErrUserCreateCompanyCoaExist)
		}
	}

	return nil
}
