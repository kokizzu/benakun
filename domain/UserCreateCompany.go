package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/wcFinance"
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"unicode"

	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/S"
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

	if in.TenantCode == `` || !isLetter(in.TenantCode){
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
	org.SetTenantCode(fmt.Sprintf("%s-%s", in.TenantCode, generate4RandomNumber()))
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

func insertCoaLevel(ta *Tt.Adapter, tenantCode string, level float64, name string, parentId uint64) (uint64, error) {
	coa := wcFinance.NewCoaMutator(ta)
	coa.SetTenantCode(tenantCode)
	coa.SetLevel(level)
	coa.SetName(name)
	if parentId > 0 {
		coa.SetParentId(parentId)
	}
	if !coa.DoInsert() {
		return 0, errors.New(ErrUserCreateCompanyCoaExist)
	}

	return coa.Id, nil
}

func generateCoaLevels(ta *Tt.Adapter, tenantCode string) error {
	for lv, vl := range mFinance.CoaLevelDefaultList {
		if _, err := insertCoaLevel(ta, tenantCode, float64(S.ToInt(lv)), vl.Name, 0); err != nil {
			return err
		}

		if len(vl.ChildrenNames) > 0 {
			parent := wcFinance.NewCoaMutator(ta)
			parent.TenantCode = tenantCode
			parent.Level = float64(S.ToInt(lv))

			parentId := parent.FindCoaIdByTenantByLevel()
			if parentId == 0 {
				return errors.New(ErrUserCreateCompanyCoaExist)
			}

			var children = []any{}

			for _, v := range vl.ChildrenNames {
				childId, err := insertCoaLevel(ta, tenantCode, float64(S.ToInt(lv)), v, parentId)
				if err != nil {
					return err
				}
				
				children = append(children, childId)
			}

			if len(children) > 0 {
				parent.SetChildren(children)
				parent.Id = parentId
				if !parent.DoUpdateById() {
					return errors.New(ErrUserCreateCompanyCoaExist)
				}
			}
		}
	}

	return nil
}

func generate4RandomNumber() string {
	to4Numbers := make([]int, 0)

	for i := 0; i < 4; i++ {
		rnum, _ := rand.Int(rand.Reader, big.NewInt(9))
		to4Numbers = append(to4Numbers, int(rnum.Int64()))
	}

	strNumbers := make([]string, len(to4Numbers))
	for i, num := range to4Numbers {
		strNumbers[i] = I.ToS(int64(num))
	}

	result := strings.Join(strNumbers, "")

	return result
}

func isLetter(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}