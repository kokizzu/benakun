package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminUpdateOrganizationChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminUpdateOrganizationChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminUpdateOrganizationChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminUpdateOrganizationChild.go
//go:generate farify doublequote --file TenantAdminUpdateOrganizationChild.go

type (
	TenantAdminUpdateOrganizationChildIn struct {
		RequestCommon
		Id         	uint64 `json:"id" form:"id" query:"id" long:"id" msg:"id"`
		Name				string `json:"name" form:"name" query:"name" long:"name" msg:"name"`
		HeadTitle   string `json:"headTitle" form:"headTitle" query:"headTitle" long:"headTitle" msg:"headTitle"`
	}
	TenantAdminUpdateOrganizationChildOut struct {
		ResponseCommon
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminUpdateOrganizationChildAction = `tenantAdmin/updateOrganizationChild`

	ErrTenantAdminUpdateOrganizationChildUnauthorized = `unauthorized user`
	ErrTenantAdminUpdateOrganizationChildTenantNotFound = `tenant admin not found`
	ErrTenantAdminUpdateOrganizationChildOrgParentNotFound = `organization parent not found`
	ErrTenantAdminUpdateOrganizationChildOrgNotFound  = `organization not found`
	ErrTenantAdminUpdateOrganizationChildFailed = `failed to update organization`
)

func (d *Domain) TenantAdminUpdateOrganizationChild(in *TenantAdminUpdateOrganizationChildIn) (out TenantAdminUpdateOrganizationChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminUpdateOrganizationChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminUpdateOrganizationChildTenantNotFound)
		return
	}

	child := wcAuth.NewOrgsMutator(d.AuthOltp)
	child.Id = in.Id
	if !child.FindById() {
		out.SetError(400, ErrTenantAdminUpdateOrganizationChildOrgNotFound)
		return
	}

	child.SetName(in.Name)
	child.SetHeadTitle(in.HeadTitle)
	child.SetUpdatedAt(in.UnixNow())
	child.SetUpdatedBy(sess.UserId)

	if !child.DoUpdateById() {
		child.HaveMutation()
		out.SetError(400, ErrTenantAdminUpdateOrganizationChildFailed)
		return
	}

	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	orgs := org.FindOrgsByTenant(tenant.TenantCode)

	out.Orgs = &orgs

	return
}
