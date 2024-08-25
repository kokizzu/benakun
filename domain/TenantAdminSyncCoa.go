package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance/rqFinance"

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
		Tenant		*rqAuth.Tenants `json:"tenant" form:"tenant" query:"tenant" long:"tenant" msg:"tenant"`
	}
)

const (
	TenantAdminSyncCoaAction = `tenantAdmin/syncCoa`

	ErrTenantAdminSyncCoaUnauthorized   			= `unauthorized user`
	ErrTenantAdminSyncCoaTenantNotFound				= `tenant not found`
	ErrTenantAdminSyncCoaNotFound							= `coa not found`
	ErrTenantAdminSyncCoaFailed								= `failed to sync coa`
	ErrTenantAdminSyncCoaSameCoa							= `coa cannot be same`
	ErrTenantAdminSyncCoaProductsCoaNotFound	= `products coa not found`
	ErrTenantAdminSyncCoaSuppliersCoaNotFound	= `suppliers coa not found`
	ErrTenantAdminSyncCoaCustomersCoaNotFound	= `customers coa not found`
	ErrTenantAdminSyncCoaCustomersRecCoaNotFound	= `customers receivables coa not found`
	ErrTenantAdminSyncCoaStaffsCoaNotFound		= `staffs coa not found`
	ErrTenantAdminSyncCoaBanksCoaNotFound			= `banks coa not found`
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

	if !mAuth.IsCoaDifferent(
		in.Tenant.ProductsCoaId, in.Tenant.SuppliersCoaId, in.Tenant.CustomersCoaId,
		in.Tenant.StaffsCoaId, in.Tenant.BanksCoaId,
	) {
		out.SetError(400, ErrTenantAdminSyncCoaSameCoa)
		return
	}

	if in.Tenant.ProductsCoaId != tenant.ProductsCoaId {
		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.Id = in.Tenant.ProductsCoaId
		if !coa.FindById() {
			out.SetError(400, ErrTenantAdminSyncCoaProductsCoaNotFound)
		}

		tenant.SetProductsCoaId(in.Tenant.ProductsCoaId)
	}

	if in.Tenant.SuppliersCoaId != tenant.SuppliersCoaId {
		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.Id = in.Tenant.SuppliersCoaId
		if !coa.FindById() {
			out.SetError(400, ErrTenantAdminSyncCoaSuppliersCoaNotFound)
		}

		tenant.SetSuppliersCoaId(in.Tenant.SuppliersCoaId)
	}

	if in.Tenant.CustomersCoaId != tenant.CustomersCoaId {
		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.Id = in.Tenant.CustomersCoaId
		if !coa.FindById() {
			out.SetError(400, ErrTenantAdminSyncCoaCustomersCoaNotFound)
		}

		tenant.SetCustomersCoaId(in.Tenant.CustomersCoaId)
	}

	if in.Tenant.CustomerReceivablesCoaId != tenant.CustomerReceivablesCoaId {
		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.Id = in.Tenant.CustomersCoaId
		if !coa.FindById() {
			out.SetError(400, ErrTenantAdminSyncCoaCustomersRecCoaNotFound)
		}

		tenant.SetCustomerReceivablesCoaId(in.Tenant.CustomersCoaId)
	}

	if in.Tenant.StaffsCoaId != tenant.StaffsCoaId {
		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.Id = in.Tenant.StaffsCoaId
		if !coa.FindById() {
			out.SetError(400, ErrTenantAdminSyncCoaStaffsCoaNotFound)
		}

		tenant.SetStaffsCoaId(in.Tenant.StaffsCoaId)
	}

	if in.Tenant.BanksCoaId != tenant.BanksCoaId {
		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.Id = in.Tenant.BanksCoaId
		if !coa.FindById() {
			out.SetError(400, ErrTenantAdminSyncCoaBanksCoaNotFound)
		}
		
		tenant.SetBanksCoaId(in.Tenant.BanksCoaId)
	}

	if !tenant.DoUpdateByTenantCode() {
		out.SetError(400, ErrTenantAdminSyncCoaFailed)
		return
	}

	out.Tenant = &tenant.Tenants

	return
}
