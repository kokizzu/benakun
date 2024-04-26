package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminRestoreOrganizationChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminRestoreOrganizationChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminRestoreOrganizationChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminRestoreOrganizationChild.go
//go:generate farify doublequote --file TenantAdminRestoreOrganizationChild.go

type (
	TenantAdminRestoreOrganizationChildIn struct {
		RequestCommon
		Id         uint64      `json:"id" form:"id" query:"id" long:"id" msg:"id"`
	}
	TenantAdminRestoreOrganizationChildOut struct {
		ResponseCommon
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminRestoreOrganizationChildAction = `tenantAdmin/restoreOrganizationChild`

	ErrTenantAdminRestoreOrganizationChildUnauthorized      = `unauthorized user to restore organization`
	ErrTenantAdminRestoreOrganizationChildTenantNotFound    = `tenant admin not found to restore organization`
	ErrTenantAdminRestoreOrganizationChildOrgParentNotFound = `cannot found parent to restore organization`
	ErrTenantAdminRestoreOrganizationChildOrgChildNotFound  = `organization child not found to restore organization`
	ErrTenantAdminRestoreOrganizationChildFailed						= `failed to restore organization`
	ErrTenantAdminRestoreOrgChildNotDeleted	 								= `organization is not deleted`
)

func (d *Domain) TenantAdminRestoreOrganizationChild(in *TenantAdminRestoreOrganizationChildIn) (out TenantAdminRestoreOrganizationChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminRestoreOrganizationChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminRestoreOrganizationChildTenantNotFound)
		return
	}

	child := wcAuth.NewOrgsMutator(d.AuthOltp)
	child.Id = in.Id
	if !child.FindById() {
		out.SetError(400, ErrTenantAdminRestoreOrganizationChildOrgChildNotFound)
		return
	}

	if child.DeletedAt == 0 {
		out.SetError(400, ErrTenantAdminRestoreOrgChildNotDeleted)
		return
	}

	child.SetDeletedAt(0)

	if !child.DoUpdateById() {
		child.HaveMutation()
		out.SetError(400, ErrTenantAdminRestoreOrganizationChildFailed)
		return
	}

	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	orgs := org.FindOrgsByTenant(tenant.TenantCode)

	out.Orgs = &orgs

	return
}

