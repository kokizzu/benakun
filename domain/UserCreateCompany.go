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

	ErrUserCreateCompanyUserNotFound          = `user not found`
	ErrUserCreateCompanyAlreadyAdded          = `company already exist`
	ErrUserCreateCompanyTenantCodeInvalid     = `invalid tenant code, must be letters only`
	ErrUserCreateCompanyCoaParentNotFound     = `coa parent not found when creating default coa levels`
	ErrUserCreateCompanyFailedSaveDefaultCoa  = `save default coa failed`
	ErrUserCreateCompanyFailedUpdateCoaParent = `failed to update coa parent`
	ErrUserCreateCompanyAlreadyHaveCompany    = `already have company`
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

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.SetTenantCode(fmt.Sprintf("%s-%s", in.Company.TenantCode, generate4RandomNumber()))
	tenant.SetCreatedAt(in.UnixNow())
	tenant.SetCreatedBy(sess.UserId)
	tenant.SetUpdatedAt(in.UnixNow())
	tenant.SetUpdatedBy(sess.UserId)
	if !tenant.DoInsert() {
		out.SetError(400, ErrUserCreateCompanyAlreadyAdded)
	}

	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	org.SetName(in.Company.Name)
	org.SetTenantCode(tenant.TenantCode)
	org.SetHeadTitle(in.Company.HeadTitle)
	org.SetOrgType(mAuth.OrgTypeCompany)

	if !org.DoUpsertById() {
		out.SetError(400, ErrUserCreateCompanyAlreadyAdded)
		return
	}

	user.SetRole(TenantAdminSegment)
	user.SetTenantCode(org.TenantCode)
	if !user.DoUpdateById() {
		out.SetError(400, ErrUserCreateCompanyUserNotFound)
		return
	}

	err := generateDefaultCoa(d.AuthOltp, tenant.TenantCode)
	if err != nil {
		out.SetError(400, err.Error())
		return
	}

	out.Company = &org.Orgs

	hostmap[generateTenantSubdomain(tenant.TenantCode)] = TenantHost{
		TenantCode: tenant.TenantCode,
		OrgId:      org.Id,
	}

	return
}

func insertDefaultCoa(ta *Tt.Adapter, coaDefault mFinance.CoaDefault, tenantCode string, parentId uint64) (uint64, error) {
	coa := wcFinance.NewCoaMutator(ta)
	coa.SetName(coaDefault.Name)
	coa.SetLabel(coaDefault.Label)
	coa.SetTenantCode(tenantCode)

	if parentId > 0 {
		coa.SetParentId(parentId)
	}

	if !coa.DoUpsertById() {
		return 0, errors.New(ErrUserCreateCompanyFailedSaveDefaultCoa)
	}

	if len(coaDefault.Children) > 0 {
		var children []any
		for _, childCoaDefault := range coaDefault.Children {
			childId, err := insertDefaultCoa(ta, childCoaDefault, tenantCode, coa.Id)
			if err != nil {
				return 0, err
			}

			children = append(children, childId)
		}

		coa.SetChildren(children)
		if !coa.DoUpsertById() {
			return 0, errors.New(ErrUserCreateCompanyFailedSaveDefaultCoa)
		}
	}

	switch coaDefault.Label {
	case mFinance.LabelProducts:
		tenant := wcAuth.NewTenantsMutator(ta)
		tenant.TenantCode = tenantCode
		tenant.SetProductsCoaId(coa.Id)
		if !tenant.DoUpdateByTenantCode() {
			return 0, errors.New(ErrUserCreateCompanyFailedSaveDefaultCoa)
		}
	}

	return coa.Id, nil
}

func generateDefaultCoa(ta *Tt.Adapter, tenantCode string) error {
	coaDefaults := mFinance.GetCoaDefaults()
	for _, coaDefault := range coaDefaults {
		if _, err := insertDefaultCoa(ta, coaDefault, tenantCode, 0); err != nil {
			return err
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
