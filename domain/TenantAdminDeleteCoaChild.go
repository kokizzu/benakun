package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminDeleteCoaChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminDeleteCoaChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminDeleteCoaChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminDeleteCoaChild.go
//go:generate farify doublequote --file TenantAdminDeleteCoaChild.go

type (
	TenantAdminDeleteCoaChildIn struct {
		RequestCommon
		Id         uint64      `json:"id" form:"id" query:"id" long:"id" msg:"id"`
	}
	TenantAdminDeleteCoaChildOut struct {
		ResponseCommon
		Coas *[]rqAuth.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
	}
)

const (
	TenantAdminDeleteCoaChildAction = `tenantAdmin/deleteCoaChild`

	ErrTenantAdminDeleteCoaChildUnauthorized      = `unauthorized user`
	ErrTenantAdminDeleteCoaChildTenantNotFound    = `tenant admin not found`
	ErrTenantAdminDeleteCoaChildCoaParentNotFound = `coa parent not found`
	ErrTenantAdminDeleteCoaChildCoaChildNotFound  = `coa child not found`
	ErrTenantAdminDeleteCoaChildFailed = `failed to delete coa child`
	ErrTenantAdminDeleteCoaChildCoaChildHaveChildren = `cannot delete if have children`
)

func (d *Domain) TenantAdminDeleteCoaChild(in *TenantAdminDeleteCoaChildIn) (out TenantAdminDeleteCoaChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminDeleteCoaChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminDeleteCoaChildTenantNotFound)
		return
	}

	child := wcAuth.NewCoaMutator(d.AuthOltp)
	child.Id = in.Id
	if !child.FindById() {
		out.SetError(400, ErrTenantAdminDeleteCoaChildTenantNotFound)
		return
	}

	if len(child.Children) > 0 {
		out.SetError(400, ErrTenantAdminDeleteCoaChildCoaChildHaveChildren)
		return
	}

	// delete coa child
	child.SetDeletedAt(in.UnixNow())
	child.SetUpdatedAt(in.UnixNow())
	child.SetUpdatedBy(sess.UserId)
	
	if !child.DoUpdateById() {
		child.HaveMutation()
		out.SetError(400, ErrTenantAdminDeleteCoaChildFailed)
		return
	}

	// retrieve owned coa
	coa := wcAuth.NewCoaMutator(d.AuthOltp)
	coas := coa.FindCoasByTenant(tenant.TenantCode)

	out.Coas = &coas

	return
}
