package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminCoa.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminCoa.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminCoa.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminCoa.go
//go:generate farify doublequote --file TenantAdminCoa.go

type (
	TenantAdminCoaIn struct {
		RequestCommon
	}
	TenantAdminCoaOut struct {
		ResponseCommon
		Coas *[]rqAuth.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
	}
)

const (
	TenantAdminCoaAction = `tenantAdmin/coa`

	ErrTenantAdminCoaUnauthorized   = `unauthorized user for coa`
	ErrTenantAdminCoaTenantNotFound = `tenant admin not found for coa`
)

func (d *Domain) TenantAdminCoa(in *TenantAdminCoaIn) (out TenantAdminCoaOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminCoaUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminCoaTenantNotFound)
		return
	}

	coa := wcAuth.NewCoaMutator(d.AuthOltp)
	coas := coa.FindCoasByTenant(tenant.TenantCode)

	out.Coas = &coas

	return
}
