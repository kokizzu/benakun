package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminSyncCoa.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminSyncCoa.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminSyncCoa.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminSyncCoa.go
//go:generate farify doublequote --file TenantAdminSyncCoa.go

type (
	TenantAdminSyncCoaIn struct {
		RequestCommon
		Tenant		*rqAuth.Tenants `json:"tenant" form:"tenant" query:"tenant" long:"tenant" msg:"tenant"`
	}
	TenantAdminSyncCoaOut struct {
		ResponseCommon
	}
)

const (
	TenantAdminSyncCoaAction = `tenantAdmin/syncCoa`

	ErrTenantAdminSyncCoaUnauthorized   = `unauthorized user`
	ErrTenantAdminSyncCoaTenantNotFound	= `tenant not found`
	ErrTenantAdminSyncCoaNotFound				= `coa not found`
	ErrTenantAdminSyncFailed						= `failed to sync coa`
)

func (d *Domain) TenantAdminSyncCoa(in *TenantAdminSyncCoaIn) (out TenantAdminSyncCoaOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrTenantAdminSyncCoaUnauthorized)
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminSyncCoaUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminSyncCoaTenantNotFound)
		return
	}

	if in.Tenant.ProductsCoaId != tenant.ProductsCoaId {
		tenant.SetProductsCoaId(in.Tenant.ProductsCoaId)
	}

	if in.Tenant.SuppliersCoaId != tenant.SuppliersCoaId {
		tenant.SetSuppliersCoaId(in.Tenant.SuppliersCoaId)
	}

	if in.Tenant.CustomersCoaId != tenant.CustomersCoaId {
		tenant.SetCustomersCoaId(in.Tenant.CustomersCoaId)
	}

	if in.Tenant.StaffsCoaId != tenant.StaffsCoaId {
		tenant.SetStaffsCoaId(in.Tenant.StaffsCoaId)
	}

	if in.Tenant.BanksCoaId != tenant.BanksCoaId {
		tenant.SetBanksCoaId(in.Tenant.BanksCoaId)
	}

	if !tenant.DoUpdateByTenantCode() {
		out.SetError(400, ErrTenantAdminSyncFailed)
		return
	}

	return
}
