package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminUpsertOrganizationChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminUpsertOrganizationChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminUpsertOrganizationChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminUpsertOrganizationChild.go
//go:generate farify doublequote --file TenantAdminUpsertOrganizationChild.go

type (
	TenantAdminUpsertOrganizationChildIn struct {
		RequestCommon
		Org rqAuth.Orgs `json:"org" form:"org" query:"org" long:"org" msg:"org"`
	}
	TenantAdminUpsertOrganizationChildOut struct {
		ResponseCommon
		Org  *rqAuth.Orgs   `json:"org" form:"org" query:"org" long:"org" msg:"org"`
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminUpsertOrganizationChildAction = `tenantAdmin/updateOrganizationChild`

	ErrTenantAdminUpsertOrganizationChildUnauthorized      = `unauthorized user`
	ErrTenantAdminUpsertOrganizationChildTenantNotFound    = `tenant admin not found`
	ErrTenantAdminUpsertOrganizationChildOrgParentNotFound = `organization parent not found`
	ErrTenantAdminUpsertOrganizationChildOrgNotFound       = `organization not found`
	ErrTenantAdminUpsertOrganizationChildFailed            = `failed to update organization`
	ErrTenantAdminUpsertOrganizationChildOrgChildNotFound  = `organization child not found to create organization`
	ErrTenantAdminUpsertOrganizationOrgAlreadyExist        = `organization already exist`
	ErrTenantAdminUpsertOrganizationJobCannotAddChild      = `cannot create child if org type is job`
)

func (d *Domain) TenantAdminUpsertOrganizationChild(in *TenantAdminUpsertOrganizationChildIn) (out TenantAdminUpsertOrganizationChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminUpsertOrganizationChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminUpsertOrganizationChildTenantNotFound)
		return
	}

	parent := wcAuth.NewOrgsMutator(d.AuthOltp)
	parent.Id = in.Org.ParentId
	if !parent.FindById() {
		out.SetError(400, ErrTenantAdminUpsertOrganizationChildFailed)
		return
	}

	child := wcAuth.NewOrgsMutator(d.AuthOltp)

	if in.Org.Id > 0 {
		child.Id = in.Org.Id
		if !child.FindById() {
			out.SetError(400, ErrTenantAdminUpsertOrganizationChildOrgNotFound)
			return
		}
	}

	switch parent.OrgType {
	case mAuth.OrgTypeCompany:
		child.SetOrgType(mAuth.OrgTypeDept)
	case mAuth.OrgTypeDept:
		child.SetOrgType(mAuth.OrgTypeDivision)
	case mAuth.OrgTypeDivision:
		child.SetOrgType(mAuth.OrgTypeJob)
	case mAuth.OrgTypeJob:
		out.SetError(400, ErrTenantAdminUpsertOrganizationJobCannotAddChild)
		return
	default:
		return
	}

	child.SetTenantCode(tenant.TenantCode)
	child.SetCreatedAt(in.UnixNow())
	child.SetCreatedBy(sess.UserId)
	child.SetUpdatedAt(in.UnixNow())
	child.SetUpdatedBy(sess.UserId)

	child.SetAll(in.Org, nil, nil)

	if !child.DoUpsertById() {
		child.HaveMutation()
		out.SetError(400, ErrTenantAdminUpsertOrganizationChildFailed)
		return
	}

	out.Org = &child.Orgs

	if in.Org.Id <= 0 {
		children := parent.Children
		children = append(children, child.Id)
		parent.SetChildren(children)
		parent.SetUpdatedAt(in.UnixNow())
		parent.SetUpdatedBy(sess.UserId)
		if !parent.DoUpdateById() {
			parent.HaveMutation()
			out.SetError(400, ErrTenantAdminUpsertOrganizationChildFailed)
			return
		}
	}

	r := rqAuth.NewOrgs(d.AuthOltp)
	orgs := r.FindOrgsByTenant(tenant.TenantCode)

	out.Orgs = &orgs

	return
}
