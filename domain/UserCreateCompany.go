package domain

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"unicode"

	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBudget/wcBudget"
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

	ErrUserCreateCompanyUserNotFound                = `user not found`
	ErrUserCreateCompanyAlreadyAdded                = `company already exist`
	ErrUserCreateCompanyTenantCodeInvalid           = `invalid tenant code, must be letters only`
	ErrUserCreateCompanyCoaParentNotFound           = `coa parent not found when creating default coa levels`
	ErrUserCreateCompanyFailedSaveDefaultCoa        = `save default coa failed`
	ErrUserCreateCompanyFailedUpdateCoaParent       = `failed to update coa parent`
	ErrUserCreateCompanyAlreadyHaveCompany          = `already have company`
	ErrUserCreateCompanyFailedSaveTrxTemplate       = `failed to save transaction template`
	ErrUserCreateCompanyFailedSaveTrxTemplateDetail = `failed to save transaction template detail`
	ErrUserCreateCompanyFailedSaveBankAccount       = `failed to save default bank account`
	ErrUserCreateCompanyFailedGenerateColor         = `failed to generate color for transaction template`
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

	if !org.DoUpsert() {
		out.SetError(400, ErrUserCreateCompanyAlreadyAdded)
		return
	}

	if !sess.IsSuperAdmin {
		user.SetRole(TenantAdminSegment)
	}
	user.SetTenantCode(org.TenantCode)
	if !user.DoUpdateById() {
		out.SetError(400, ErrUserCreateCompanyUserNotFound)
		return
	}

	bankAccount := wcBudget.NewBankAccountsMutator(d.AuthOltp)
	bankAccount.TenantCode = tenant.TenantCode
	bankAccount.AccountName = `Rekening Utama`
	bankAccount.AccountNumber = 0000000
	if !bankAccount.DoInsert() {
		out.SetError(400, ErrUserCreateCompanyFailedSaveBankAccount)
		return
	}

	trxTemplateMaps := make(map[string]uint64)

	for _, tplName := range mFinance.TransactionTemplatesDefault {
		trxTemplate := wcFinance.NewTransactionTemplateMutator(d.AuthOltp)
		trxTemplate.SetName(tplName)

		h, s, l, err := RANDOMHSL()
		if err != nil {
			out.SetError(400, ErrUserCreateCompanyFailedGenerateColor)
			return
		}

		hexColor := `#` + HSL2HEX(h, s, l)
		trxTemplate.SetColor(hexColor)
		trxTemplate.SetTenantCode(tenant.TenantCode)
		trxTemplate.SetCreatedAt(in.UnixNow())
		trxTemplate.SetCreatedBy(sess.UserId)
		trxTemplate.SetUpdatedAt(in.UnixNow())
		trxTemplate.SetUpdatedBy(sess.UserId)
		if !trxTemplate.DoInsert() {
			out.SetError(400, ErrUserCreateCompanyFailedSaveTrxTemplate)
			return
		}

		trxTemplateMaps[trxTemplate.Name] = trxTemplate.Id
	}

	coaMaps := make(map[string]uint64)

	var insertDefaultCoa func(ta *Tt.Adapter, coaDefault mFinance.CoaDefault, tenantCode string, parentId uint64) (uint64, error)
	insertDefaultCoa = func(ta *Tt.Adapter, coaDefault mFinance.CoaDefault, tenantCode string, parentId uint64) (uint64, error) {
		coa := wcFinance.NewCoaMutator(ta)
		coa.SetName(coaDefault.Name)
		coa.SetLabel(coaDefault.Label)
		coa.SetTenantCode(tenantCode)
		coa.SetEditable(false)

		if parentId > 0 {
			coa.SetParentId(parentId)
		}

		if !coa.DoUpdateById() {
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
			if !coa.DoUpdateById() {
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
		case mFinance.LabelSuppliers:
			tenant := wcAuth.NewTenantsMutator(ta)
			tenant.TenantCode = tenantCode
			tenant.SetSuppliersCoaId(coa.Id)
			if !tenant.DoUpdateByTenantCode() {
				return 0, errors.New(ErrUserCreateCompanyFailedSaveDefaultCoa)
			}
		case mFinance.LabelStaff:
			tenant := wcAuth.NewTenantsMutator(ta)
			tenant.TenantCode = tenantCode
			tenant.SetStaffsCoaId(coa.Id)
			if !tenant.DoUpdateByTenantCode() {
				return 0, errors.New(ErrUserCreateCompanyFailedSaveDefaultCoa)
			}
		case mFinance.LabelBankAccount:
			tenant := wcAuth.NewTenantsMutator(ta)
			tenant.TenantCode = tenantCode
			tenant.SetBanksCoaId(coa.Id)
			if !tenant.DoUpdateByTenantCode() {
				return 0, errors.New(ErrUserCreateCompanyFailedSaveDefaultCoa)
			}
		case mFinance.LabelCustomer:
			tenant := wcAuth.NewTenantsMutator(ta)
			tenant.TenantCode = tenantCode
			tenant.SetCustomersCoaId(coa.Id)
			if !tenant.DoUpdateByTenantCode() {
				return 0, errors.New(ErrUserCreateCompanyFailedSaveDefaultCoa)
			}
		case mFinance.LabelCustomerReceivables:
			tenant := wcAuth.NewTenantsMutator(ta)
			tenant.TenantCode = tenantCode
			tenant.SetCustomerReceivablesCoaId(coa.Id)
			if !tenant.DoUpdateByTenantCode() {
				return 0, errors.New(ErrUserCreateCompanyFailedSaveDefaultCoa)
			}
		case mFinance.LabelFunders:
			tenant := wcAuth.NewTenantsMutator(ta)
			tenant.TenantCode = tenantCode
			tenant.SetFundersCoaId(coa.Id)
			if !tenant.DoUpdateByTenantCode() {
				return 0, errors.New(ErrUserCreateCompanyFailedSaveDefaultCoa)
			}
		}

		coaMaps[coa.Name] = coa.Id
		return coa.Id, nil
	}

	coaDefaults := mFinance.GetCoaDefaults()
	for _, coaDefault := range coaDefaults {
		if _, err := insertDefaultCoa(d.AuthOltp, coaDefault, tenant.TenantCode, 0); err != nil {
			out.SetError(400, err.Error())
			return
		}
	}

	for tplName, tplDetails := range mFinance.TransactionTemplateDetailsMap {
		for _, td := range tplDetails {
			if id, exist := coaMaps[td.CoaName]; exist {
				trxTplDetail := wcFinance.NewTransactionTplDetailMutator(d.AuthOltp)
				trxTplDetail.SetCoaId(id)
				trxTplDetail.SetIsDebit(td.IsDebit)
				if len(td.Attributes) > 0 {
					trxTplDetail.SetAttributes(td.Attributes)
				}
				trxTplDetail.SetParentId(trxTemplateMaps[tplName])
				trxTplDetail.SetTenantCode(tenant.TenantCode)
				trxTplDetail.SetCreatedAt(in.UnixNow())
				trxTplDetail.SetCreatedBy(sess.UserId)
				trxTplDetail.SetUpdatedAt(in.UnixNow())
				trxTplDetail.SetUpdatedBy(sess.UserId)
				if !trxTplDetail.DoInsert() {
					out.SetError(400, ErrUserCreateCompanyFailedSaveTrxTemplateDetail)
					return
				}
			}
		}
	}

	out.Company = &org.Orgs

	hostmap[generateTenantSubdomain(tenant.TenantCode)] = TenantHost{
		TenantCode: tenant.TenantCode,
		OrgId:      org.Id,
	}

	return
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

func RANDOMHSL() (int, int, int, error) {
	var MAX_VALUE int64 = 360

	hueBigInt, err := rand.Int(rand.Reader, big.NewInt(MAX_VALUE))
	if err != nil {
		return 0, 0, 0, err
	}
	return int(hueBigInt.Int64()), 82, 43, nil
}

func HSL2RGB(h, s, l int) (int, int, int) {
	var hf, sf, lf, vf, minf, svf, sixf, fractf, vsfractf, rf, gf, bf float64

	hf = math.Max(math.Min(float64(int(h)), 360), 0) / 360
	sf = math.Max(math.Min(float64(int(s)), 100), 0) / 100
	lf = math.Max(math.Min(float64(int(l)), 100), 0) / 100

	if lf <= 0.5 {
		vf = lf * (1 + sf)
	} else {
		vf = lf + sf - sf - 1*sf
	}
	if vf == 0 {
		return int(0), int(0), int(0)
	}
	minf = 2*lf - vf
	svf = (vf - minf) / vf
	hf = 6 * hf
	sixf = float64(int(hf))
	fractf = hf - sixf
	vsfractf = vf * svf * fractf
	switch sixf {
	case 1:
		rf = vf - vsfractf
		gf = vf
		bf = minf
	case 2:
		rf = minf
		gf = vf
		bf = minf + vsfractf
	case 3:
		rf = minf
		gf = vf - vsfractf
		bf = vf
	case 4:
		rf = minf + vsfractf
		gf = minf
		bf = vf
	case 5:
		rf = vf
		gf = minf
		bf = vf - vsfractf
	default:
		rf = vf
		gf = minf + vsfractf
		bf = minf
	}

	return int(rf * 255), int(gf * 255), int(bf * 255)
}

func RGB2HEX(r, g, b int) string {
	var hexr, hexg, hexb string
	r = int(math.Max(math.Min(float64(r), 255), 0))
	g = int(math.Max(math.Min(float64(g), 255), 0))
	b = int(math.Max(math.Min(float64(b), 255), 0))

	hexr = strconv.FormatInt(int64(r), 16)
	if r < 16 {
		hexr = "0" + hexr
	}
	hexg = strconv.FormatInt(int64(g), 16)
	if g < 16 {
		hexg = "0" + hexg
	}
	hexb = strconv.FormatInt(int64(b), 16)
	if b < 16 {
		hexb = "0" + hexb
	}

	return hexr + hexg + hexb
}

func HSL2HEX(h, s, l int) string {
	r, g, b := HSL2RGB(h, s, l)
	return RGB2HEX(r, g, b)
}
