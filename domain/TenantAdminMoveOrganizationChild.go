package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminMoveOrganizationChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminMoveOrganizationChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminMoveOrganizationChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminMoveOrganizationChild.go
//go:generate farify doublequote --file TenantAdminMoveOrganizationChild.go

// TODO:HABIBI make file naming consistent with url requesting

type (
	TenantAdminMoveOrganizationChildIn struct {
		RequestCommon
		Id         uint64 `json:"id" form:"id" query:"id" long:"id" msg:"id"`
		MoveToIdx  int    `json:"moveToIdx" form:"moveToIdx" query:"moveToIdx" long:"moveToIdx" msg:"moveToIdx"`
		ToParentId uint64 `json:"toParentId" form:"toParentId" query:"toParentId" long:"toParentId" msg:"toParentId"`
	}

	TenantAdminMoveOrganizationChildOut struct {
		ResponseCommon
		Org  *rqAuth.Orgs   `json:"org" form:"org" query:"org" long:"org" msg:"org"`
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminMoveOrganizationChildAction = `tenantAdmin/moveOrganizationChild`

	ErrTenantAdminMoveOrganizationChildUnauthorized       = `unauthorized user to move this organization`
	ErrTenantAdminMoveOrganizationChildTenantNotFound     = `cannot move organization if you are not a tenant admin`
	ErrTenantAdminMoveOrganizationChildOrgNotFound        = `invalid organization to move`
	ErrTenantAdminMoveOrganizationChildParentNotFound     = `parent not found to move this organization`
	ErrTenantAdminMoveOrganizationChildToParentNotFound   = `cannot found parent to move this organization`
	ErrTenantAdminMoveOrganizationChildShouldNotCompany   = `cannot move organization if type is company`
	ErrTenantAdminMoveOrganizationChildFailedMoveChildren = `failed to move organization child`
	ErrTenantAdminMoveOrganizationChildMustSameParentType = `cannot move to other parent if different organization type`
	ErrTenantAdminMoveOrganizationChildFailedUpdateChild  = `failed to update organization child`
)

func (d *Domain) TenantAdminMoveOrganizationChild(in *TenantAdminMoveOrganizationChildIn) (out TenantAdminMoveOrganizationChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminMoveOrganizationChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
		out.SetError(400, ErrTenantAdminMoveOrganizationChildTenantNotFound)
		return
	}

	child := wcAuth.NewOrgsMutator(d.AuthOltp)
	child.Id = in.Id
	if !child.FindById() {
		out.SetError(400, ErrTenantAdminMoveOrganizationChildOrgNotFound)
		return
	}

	parent := wcAuth.NewOrgsMutator(d.AuthOltp)
	parent.Id = child.ParentId
	if !parent.FindById() {
		out.SetError(400, ErrTenantAdminMoveOrganizationChildParentNotFound)
		return
	}

	// if organization move to the same parent
	if parent.Id == in.ToParentId {
		if len(parent.Children) >= 2 {
			children, err := moveChildToIndex(parent.Children, in.Id, in.MoveToIdx)
			if err != nil {
				out.SetError(400, ErrTenantAdminMoveOrganizationChildOrgNotFound)
				return
			}

			parent.SetChildren(children)
			if !parent.DoUpdateById() {
				parent.HaveMutation()
				out.SetError(400, ErrTenantAdminMoveOrganizationChildFailedMoveChildren)
				return
			}
		}

		out.Org = &child.Orgs

		org := wcAuth.NewOrgsMutator(d.AuthOltp)
		orgs := org.FindOrgsByTenant(tenant.TenantCode)
		out.Orgs = &orgs

		return
	}

	// if organization move to different parent
	toParent := wcAuth.NewOrgsMutator(d.AuthOltp)
	toParent.Id = in.ToParentId
	if !toParent.FindById() {
		out.SetError(400, ErrTenantAdminMoveOrganizationChildToParentNotFound)
		return
	}

	// organization should move to the same parent type
	switch child.OrgType {
	case mAuth.OrgTypeCompany:
		out.SetError(400, ErrTenantAdminMoveOrganizationChildShouldNotCompany)
		return
	case mAuth.OrgTypeDept:
		if toParent.OrgType != mAuth.OrgTypeCompany {
			out.SetError(400, ErrTenantAdminMoveOrganizationChildMustSameParentType)
			return
		}
	case mAuth.OrgTypeDivision:
		if toParent.OrgType != mAuth.OrgTypeDept {
			out.SetError(400, ErrTenantAdminMoveOrganizationChildMustSameParentType)
			return
		}
	case mAuth.OrgTypeJob:
		if toParent.OrgType != mAuth.OrgTypeDivision {
			out.SetError(400, ErrTenantAdminMoveOrganizationChildMustSameParentType)
			return
		}
	}

	// operation if organization move to other parent

	// add organization to parent destination
	children := insertChildToIndex(toParent.Children, in.Id, in.MoveToIdx)

	child.SetParentId(in.ToParentId)
	if !child.DoUpdateById() {
		out.SetError(400, ErrTenantAdminMoveOrganizationChildFailedUpdateChild)
		return
	}

	toParent.SetChildren(children)
	if !toParent.DoUpdateById() {
		toParent.HaveMutation()
		out.SetError(400, ErrTenantAdminMoveOrganizationChildFailedMoveChildren)
		return
	}

	// remove organization from previous parent
	children, err := removeChild(parent.Children, in.Id)
	if err != nil {
		out.SetError(400, ErrTenantAdminMoveOrganizationChildOrgNotFound)
		return
	}
	parent.SetChildren(children)
	if !parent.DoUpdateById() {
		parent.HaveMutation()
		out.SetError(400, ErrTenantAdminMoveOrganizationChildFailedMoveChildren)
		return
	}

	out.Org = &child.Orgs

	org := rqAuth.NewOrgs(d.AuthOltp)
	orgs := org.FindOrgsByTenant(tenant.TenantCode)
	
	out.Orgs = &orgs

	return
}
