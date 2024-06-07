package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminRestoreCoaChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminRestoreCoaChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminRestoreCoaChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminRestoreCoaChild.go
//go:generate farify doublequote --file TenantAdminRestoreCoaChild.go

// TODO:HABIBI make file naming consistent with url requesting
// TODO:HABIBI delete and restore should be one endpoint, or use old style a=restore/a=delete

type (
	TenantAdminRestoreCoaChildIn struct {
		RequestCommon
		Id uint64 `json:"id" form:"id" query:"id" long:"id" msg:"id"`
	}
	TenantAdminRestoreCoaChildOut struct {
		ResponseCommon
		Coas *[]rqFinance.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
	}
)

const (
	TenantAdminRestoreCoaChildAction = `tenantAdmin/restoreCoaChild`

	ErrTenantAdminRestoreCoaChildUnauthorized      = `unauthorized user to restore coa`
	ErrTenantAdminRestoreCoaChildTenantNotFound    = `tenant admin not found to restore coa`
	ErrTenantAdminRestoreCoaChildCoaParentNotFound = `coa parent not found to restore coa`
	ErrTenantAdminRestoreCoaChildCoaChildNotFound  = `coa child not found to restore coa`
	ErrTenantAdminRestoreCoaChildFailed            = `failed to restore coa`
	ErrTenantAdminRestoreCoaChildNotDeleted        = `coa is not deleted`
)

func (d *Domain) TenantAdminRestoreCoaChild(in *TenantAdminRestoreCoaChildIn) (out TenantAdminRestoreCoaChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminRestoreCoaChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
		out.SetError(400, ErrTenantAdminRestoreCoaChildTenantNotFound)
		return
	}

	child := wcFinance.NewCoaMutator(d.AuthOltp)
	child.Id = in.Id
	if !child.FindById() {
		out.SetError(400, ErrTenantAdminRestoreCoaChildTenantNotFound)
		return
	}

	if child.DeletedAt == 0 {
		out.SetError(400, ErrTenantAdminRestoreCoaChildNotDeleted)
		return
	}

	child.SetDeletedAt(0)

	if !child.DoUpdateById() {
		child.HaveMutation()
		out.SetError(400, ErrTenantAdminRestoreCoaChildFailed)
		return
	}

	// retrieve owned coa
	coa := wcFinance.NewCoaMutator(d.AuthOltp)
	coas := coa.FindCoasByTenant(tenant.TenantCode)

	out.Coas = &coas

	return
}
