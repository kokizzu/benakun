package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
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
		ParentId uint64 `json:"parentId" form:"parentId" query:"parentId" long:"parentId" msg:"parentId"`
	}
	TenantAdminCreateCoaChildOut struct {
		ResponseCommon
		Coas *[]rqFinance.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
	}
)

const (
	TenantAdminCreateCoaChildAction = `tenantAdmin/createCoaChild`

	ErrTenantAdminCreateCoaChildUnauthorized      = `unauthorized user to create coa child`
	ErrTenantAdminCreateCoaChildTenantNotFound    = `tenant admin not found to create coa child`
	ErrTenantAdminCreateCoaChildCoaParentNotFound = `coa parent not found for this coa child`
	ErrTenantAdminCreateCoaChildCoaChildNotFound  = `coa child not found when adds child to parent`
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
	if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
		out.SetError(400, ErrTenantAdminCreateCoaChildTenantNotFound)
		return
	}

	// find parent id
	parent := wcFinance.NewCoaMutator(d.AuthOltp)
	parent.Id = in.ParentId
	if !parent.FindById() {
		out.SetError(400, ErrTenantAdminCreateCoaChildCoaParentNotFound)
		return
	}

	// insert coa child
	child := wcFinance.NewCoaMutator(d.AuthOltp)
	child.SetTenantCode(tenant.TenantCode)
	child.SetLevel(parent.Level)
	child.SetName(in.Name)
	child.SetParentId(parent.Id)
	child.SetCreatedAt(in.UnixNow())
	child.SetCreatedBy(sess.UserId)
	child.SetUpdatedAt(in.UnixNow())
	child.SetUpdatedBy(sess.UserId)

	if !child.DoInsert() {
		out.SetError(400, ErrTenantAdminCreateCoaChildCoaParentNotFound)
		return
	}

	// update childrens of parent coa
	children := parent.Children
	children = append(children, child.Id)
	parent.SetChildren(children)
	parent.SetUpdatedAt(in.UnixNow())
	parent.SetUpdatedBy(sess.UserId)
	
	if !parent.DoUpdateById() {
		parent.HaveMutation()
		out.SetError(400, ErrTenantAdminCreateCoaChildCoaParentNotFound)
		return
	}

	// retrieve owned coa
	coa := wcFinance.NewCoaMutator(d.AuthOltp)
	coas := coa.FindCoasByTenant(tenant.TenantCode)

	out.Coas = &coas

	return
}
