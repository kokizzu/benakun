package domain

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
	"strings"
	"unicode"

	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/wcFinance"

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
		Company rqAuth.Orgs `json:"company" form:"company" query:"company" long:"company" msg:"company"`
	}

	UserCreateCompanyOut struct {
		ResponseCommon
		Ok      bool         `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`
		Company *rqAuth.Orgs `json:"company" form:"company" query:"company" long:"company" msg:"company"`
	}
)

const (
	UserCreateCompanyAction = `user/createCompany`

	ErrUserCreateCompanyUserNotFound       		= `user not found`
	ErrUserCreateCompanyAlreadyAdded       		= `company already exist`
	ErrUserCreateCompanyTenantCodeInvalid  		= `invalid tenant code, must be letters only`
	ErrUserCreateCompanyCoaParentNotFound	 		= `coa parent not found when creating default coa levels`
	ErrUserCreateCompanyFailedSaveDefaultCoa	= `save default coa failed`
	ErrUserCreateCompanyFailedUpdateCoaParent = `failed to update coa parent`
	ErrUserCreateCompanyAlreadyHaveCompany		= `already have company`
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
		out.SetError(400, ErrUserCreateCompanyUserNotFound)
		return
	}

	if user.TenantCode != `` {
		out.SetError(400, ErrUserCreateCompanyAlreadyHaveCompany)
		return
	}

	comp := rqAuth.NewOrgs(d.AuthOltp)
	if comp.FindCompanyByTenantCode(in.Company.TenantCode) {
		out.SetError(400, ErrUserCreateCompanyAlreadyAdded)
		return
	}

	if !isLetter(in.Company.TenantCode) {
		out.SetError(400, ErrUserCreateCompanyTenantCodeInvalid)
		return
	}

	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	org.SetName(in.Company.Name)
	org.SetTenantCode(fmt.Sprintf("%s-%s", in.Company.TenantCode, generate4RandomNumber()))
	org.SetHeadTitle(in.Company.HeadTitle)
	org.SetOrgType(mAuth.OrgTypeCompany)

	if !org.DoUpsertById() {
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

func insertCoaLevel(ta *Tt.Adapter,
	name, tenantCode string,
	level float64, parentId uint64,
) (uint64, error) {

	coa := wcFinance.NewCoaMutator(ta)

	coa.SetName(name)
	coa.SetTenantCode(tenantCode)
	coa.SetLevel(level)

	if parentId > 0 {
		coa.SetParentId(parentId)
	}

	if !coa.DoInsert() {
		return 0, errors.New(ErrUserCreateCompanyFailedSaveDefaultCoa)
	}

	return coa.Id, nil
}

func generateCoaLevels(ta *Tt.Adapter, tenantCode string) error {
	for lv, vl := range mFinance.CoaLevelDefaultList {
		parentId, err := insertCoaLevel(ta, vl.Name, tenantCode, S.ToF(lv), 0)
		if err != nil {
			return err
		}

		if len(vl.ChildrenNames) > 0 {
			var children = []any{}

			for _, v := range vl.ChildrenNames {
				childId, err := insertCoaLevel(ta, v, tenantCode, S.ToF(lv), parentId)
				if err != nil {
					return err
				}

				children = append(children, childId)
			}

			if len(children) > 0 {
				parent := wcFinance.NewCoaMutator(ta)
				parent.Id = parentId

				if !parent.FindById() {
					return errors.New(ErrUserCreateCompanyCoaParentNotFound)
				}
				
				parent.SetChildren(children)
				if !parent.DoUpdateById() {
					return errors.New(ErrUserCreateCompanyFailedUpdateCoaParent)
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
