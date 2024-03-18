package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminCreateOrganizationChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminCreateOrganizationChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminCreateOrganizationChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminCreateOrganizationChild.go
//go:generate farify doublequote --file TenantAdminCreateOrganizationChild.go

type (
	TenantAdminCreateOrganizationChildIn struct {
		RequestCommon
		Name string `json:"name" form:"name" query:"name" long:"name" msg:"name"`
		HeadTitle   string `json:"headTitle" form:"headTitle" query:"headTitle" long:"headTitle" msg:"headTitle"`
		ParentId uint64 `json:"parentId" form:"parentId" query:"parentId" long:"parentId" msg:"parentId"`
		OrgType string `json:"orgType" form:"orgType" query:"orgType" long:"orgType" msg:"orgType"`
	}

	TenantAdminCreateOrganizationChildOut struct {
		ResponseCommon
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminCreateOrganizationChildAction = `tenantAdmin/createOrganizationChild`

	ErrTenantAdminCreateOrganizationChildUnauthorized      = `unauthorized user`
	ErrTenantAdminCreateOrganizationChildTenantNotFound    = `tenant admin not found`
	ErrTenantAdminCreateOrganizationChildOrgParentNotFound = `organization parent not found`
	ErrTenantAdminCreateOrganizationChildOrgChildNotFound  = `organization child not found`
	ErrTenantAdminCreateOrganizationOrgAlreadyExist = `organization already exist`
	ErrTenantAdminCreateOrganizationChildFailed = `create organization child failed`
	ErrTenantAdminCreateOrganizationInvalidOrgType = `invalid type of organization`
	ErrTenantAdminCreateOrganizationJobCannotAddChild = `cannot create child if org type is job`
)

func (d *Domain) TenantAdminCreateOrganizationChild(in *TenantAdminCreateOrganizationChildIn) (out TenantAdminCreateOrganizationChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminCoaUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminCoaTenantNotFound)
		return
	}

	// Find parent
	parent := wcAuth.NewOrgsMutator(d.AuthOltp)
	parent.Id = in.ParentId
	if !parent.FindById() {
		out.SetError(400, ErrTenantAdminCreateOrganizationChildOrgParentNotFound)
		return
	}

	// Job is the lowest level, cannot create child for it
	if parent.OrgType == mAuth.OrgTypeJob {
		out.SetError(400, ErrTenantAdminCreateOrganizationJobCannotAddChild)
		return
	}

	// Create child
	child := wcAuth.NewOrgsMutator(d.AuthOltp)
	// payload should be string for readability
	switch in.OrgType {
	case `company`:
		child.SetOrgType(mAuth.OrgTypeCompany)
	case `department`:
		child.SetOrgType(mAuth.OrgTypeDept)
	case `division`:
		child.SetOrgType(mAuth.OrgTypeDivision)
	case `job`:
		child.SetOrgType(mAuth.OrgTypeJob)
	default:
		out.SetError(400, ErrTenantAdminCreateOrganizationInvalidOrgType)
		return
	}
	child.SetName(in.Name)
	child.SetHeadTitle(in.HeadTitle)
	child.SetTenantCode(tenant.TenantCode)
	child.SetParentId(in.ParentId)
	child.SetCreatedAt(in.UnixNow())
	child.SetCreatedBy(sess.UserId)
	child.SetUpdatedAt(in.UnixNow())
	child.SetUpdatedBy(sess.UserId)

	if !child.DoInsert() {
		out.SetError(400, ErrTenantAdminCreateOrganizationChildFailed)
		return
	}

	// finc the created child to update parent
	fchild := wcAuth.NewOrgsMutator(d.AuthOltp)
	fchild.ParentId = parent.Id
	fchild.Name = in.Name
	childId := fchild.FindOrgChildIdByParentIdByName()
	if childId == 0 {
		out.SetError(400, ErrTenantAdminCreateOrganizationChildOrgChildNotFound)
		return
	}

	// update children for parent
	children := parent.Children
	children = append(children, childId)
	parent.SetChildren(children)
	parent.SetUpdatedAt(in.UnixNow())
	parent.SetUpdatedBy(sess.UserId)
	if !parent.DoUpdateById() {
		parent.HaveMutation()
		out.SetError(400, ErrTenantAdminCreateOrganizationChildFailed)
		return
	}

	// retrieve owned orgs
	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	orgs := org.FindOrgsByTenant(tenant.TenantCode)

	out.Orgs = &orgs

	return
}