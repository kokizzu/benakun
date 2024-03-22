package domain

// TODO
// move organization child to other parent
// make sure it only works to move under the same company
// organization type cannot be change,
// that means organization child cannot be change to higher level
// e.g. JobA cannot be change to DivB

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

type (
	TenantAdminMoveOrganizationChildIn struct {
		RequestCommon
		Id uint64 `json:"id" form:"id" query:"id" long:"id" msg:"id"`
		MoveTo string `json:"moveTo" form:"moveTo" query:"moveTo" long:"moveTo" msg:"moveTo"`
		ToParentId uint64 `json:"ToParentId" form:"ToParentId" query:"ToParentId" long:"ToParentId" msg:"ToParentId"`
	}

	TenantAdminMoveOrganizationChildOut struct {
		ResponseCommon
		Org *rqAuth.Orgs `json:"org" form:"org" query:"org" long:"org" msg:"org"`
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminMoveOrganizationChildAction = `tenantAdmin/moveOrganizationChild`

	ErrTenantAdminMoveOrganizationChildUnauthorized      = `unauthorized user to move this organization`
	ErrTenantAdminMoveOrganizationChildTenantNotFound    = `cannot move organization if you are not a tenant admin`
	ErrTenantAdminMoveOrganizationChildOrgNotFound = `invalid organization to move`
	ErrTenantAdminMoveOrganizationChildParentNotFound = `parent not found to move this organization`
	ErrTenantAdminMoveOrganizationChildToParentNotFound = `cannot found parent to move this organization`
	ErrTenantAdminMoveOrganizationChildShouldNotCompany = `cannot move organization if type is company`
	ErrTenantAdminMoveOrganizationChildFailedMoveChildren = `failed to move organization child`
	ErrTenantAdminMoveOrganizationChildMustSameParentType = `cannot move to other parent if different organization type`
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
	if !tenant.FindByTenantCode() {
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

	toParent := wcAuth.NewOrgsMutator(d.AuthOltp)
	toParent.Id = in.ToParentId
	if !toParent.FindById() {
		out.SetError(400, ErrTenantAdminMoveOrganizationChildParentNotFound)
		return
	}

	// default to first
	if !(in.MoveTo == `first` || in.MoveTo == `end`) {
		in.MoveTo = `first`
	}

	switch child.OrgType {
	case mAuth.OrgTypeCompany:
		out.SetError(400, ErrTenantAdminMoveOrganizationChildShouldNotCompany)
		return
	case mAuth.OrgTypeDept:
		switch in.MoveTo {
		case `first`:
			moveChildToFirst(&parent.Children, child.Id)
			parent.SetChildren(parent.Children)
			if !parent.DoUpdateById() {
				parent.HaveMutation()
				out.SetError(400, ErrTenantAdminMoveOrganizationChildFailedMoveChildren)
				return
			}
		case `end`:
			moveChildToEnd(&parent.Children, child.Id)
			parent.SetChildren(parent.Children)
			if !parent.DoUpdateById() {
				parent.HaveMutation()
				out.SetError(400, ErrTenantAdminMoveOrganizationChildFailedMoveChildren)
				return
			}
		}
	case mAuth.OrgTypeDivision:
		if toParent.OrgType != mAuth.OrgTypeDept {
			out.SetError(400, ErrTenantAdminMoveOrganizationChildMustSameParentType)
			return
		}

		switch in.MoveTo {
		case `first`:
			if parent.Id != toParent.Id {
				// TODO remove element
			}
			moveChildToFirst(&parent.Children, child.Id)
			parent.SetChildren(parent.Children)
		case `end`:
			moveChildToEnd(&parent.Children, child.Id)
			parent.SetChildren(parent.Children)
		}

		if !parent.DoUpdateById() {
			parent.HaveMutation()
			out.SetError(400, ErrTenantAdminMoveOrganizationChildFailedMoveChildren)
			return
		}
	}

	return
}

func moveChildToFirst(children *[]any, elm any) {
	elmIdx := 0

	for i, v := range *children {
		if v == elm {
			elmIdx = i
		}
	}

	temp := (*children)[elmIdx]

	for i := elmIdx; i > 0; i-- {
		(*children)[i] = (*children)[i-1]
	}

	(*children)[0] = temp
}

func moveChildToEnd(children *[]any, elm any) {
	elmIdx := 0

	for i, v := range *children {
		if v == elm {
			elmIdx = i
		}
	}

	element := (*children)[elmIdx]
	*children = append((*children)[:elmIdx], (*children)[elmIdx+1:]...)

	*children = append(*children, element)
}

// TODO
// function removeElement
// function addChildToFirst
// function addChildToEnd