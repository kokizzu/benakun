package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminCreateCoaChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminCreateCoaChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminCreateCoaChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminCreateCoaChild.go
//go:generate farify doublequote --file TenantAdminCreateCoaChild.go

type (
	TenantAdminCreateCoaChildIn struct {
		RequestCommon
		Name     string `json:"name" form:"name" query:"name" long:"name" msg:"name"`
		ParentId uint64 `json:"parentId,string" form:"parentId" query:"parentId" long:"parentId" msg:"parentId"`
	}
	TenantAdminCreateCoaChildOut struct {
		ResponseCommon
		Coas *[]rqAuth.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
	}
)

const (
	TenantAdminCreateCoaChildAction = `tenantAdmin/createCoaChild`

	ErrTenantAdminCreateCoaChildUnauthorized      = `unauthorized user`
	ErrTenantAdminCreateCoaChildTenantNotFound    = `tenant admin not found`
	ErrTenantAdminCreateCoaChildCoaParentNotFound = `coa parent not found`
	ErrTenantAdminCreateCoaChildCoaChildNotFound  = `coa child not found`
)

func (d *Domain) TenantAdminCreateCoaChild(in *TenantAdminCreateCoaChildIn) (out TenantAdminCreateCoaChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminCreateCoaChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminCreateCoaChildTenantNotFound)
		return
	}

	// find parent id
	parent := wcAuth.NewCoaMutator(d.AuthOltp)
	parent.Id = in.ParentId
	if !parent.FindById() {
		out.SetError(400, ErrTenantAdminCreateCoaChildCoaParentNotFound)
		return
	}

	// insert coa child
	child := wcAuth.NewCoaMutator(d.AuthOltp)
	child.SetTenantCode(tenant.TenantCode)
	child.SetLevel(parent.Level)
	child.SetName(in.Name)
	child.SetParentId(parent.Id)
	if !child.DoInsert() {
		out.SetError(400, ErrTenantAdminCreateCoaChildCoaParentNotFound)
		return
	}

	// find child id
	fchild := wcAuth.NewCoaMutator(d.AuthOltp)
	fchild.ParentId = parent.Id
	fchild.Name = in.Name
	childId := fchild.FindCoaChildIdByParentIdByName()
	if childId == 0 {
		out.SetError(400, ErrTenantAdminCreateCoaChildCoaChildNotFound)
		return
	}

	// update childrens of parent coa
	children := parent.Children
	children = append(children, childId)
	parent.SetChildren(children)
	if !parent.DoUpdateById() {
		out.SetError(400, ErrTenantAdminCreateCoaChildCoaParentNotFound)
		return
	}

	// retrieve owned coa
	coa := wcAuth.NewCoaMutator(d.AuthOltp)
	coas := coa.FindCoasByTenant(tenant.TenantCode)

	out.Coas = &coas

	return
}
