package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminDeleteOrganizationChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminDeleteOrganizationChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminDeleteOrganizationChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminDeleteOrganizationChild.go
//go:generate farify doublequote --file TenantAdminDeleteOrganizationChild.go

type (
	TenantAdminDeleteOrganizationChildIn struct {
		RequestCommon
		Id         uint64      `json:"id" form:"id" query:"id" long:"id" msg:"id"`
	}
	TenantAdminDeleteOrganizationChildOut struct {
		ResponseCommon
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminDeleteOrganizationChildAction = `tenantAdmin/deleteOrganizationChild`

	ErrTenantAdminDeleteOrganizationChildUnauthorized      = `unauthorized user to delete organization`
	ErrTenantAdminDeleteOrganizationChildTenantNotFound    = `tenant admin not found to delete organization`
	ErrTenantAdminDeleteOrganizationChildOrgParentNotFound = `cannot found parent to delete organization`
	ErrTenantAdminDeleteOrganizationChildOrgChildNotFound  = `organization child not found to delete organization`
	ErrTenantAdminDeleteOrganizationChildFailed						 = `failed to delete organization`
	ErrTenantAdminDeleteOrganizationChildOrgAlreadyDeleted = `organization is already deleted`
)

func (d *Domain) TenantAdminDeleteOrganizationChild(in *TenantAdminDeleteOrganizationChildIn) (out TenantAdminDeleteOrganizationChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminDeleteOrganizationChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminDeleteOrganizationChildTenantNotFound)
		return
	}

	child := wcAuth.NewOrgsMutator(d.AuthOltp)
	child.Id = in.Id
	if !child.FindById() {
		out.SetError(400, ErrTenantAdminDeleteOrganizationChildOrgChildNotFound)
		return
	}

	if child.DeletedAt != 0 {
		out.SetError(400, ErrTenantAdminDeleteOrganizationChildOrgAlreadyDeleted)
		return
	}

	child.SetDeletedAt(in.UnixNow())

	if !child.DoUpdateById() {
		child.HaveMutation()
		out.SetError(400, ErrTenantAdminDeleteOrganizationChildFailed)
		return
	}

	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	orgs := org.FindOrgsByTenant(tenant.TenantCode)

	out.Orgs = &orgs

	return
}

