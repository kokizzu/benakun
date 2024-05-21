package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminUpdateCoaChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminUpdateCoaChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminUpdateCoaChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminUpdateCoaChild.go
//go:generate farify doublequote --file TenantAdminUpdateCoaChild.go

type (
	TenantAdminUpdateCoaChildIn struct {
		RequestCommon
		Id         uint64      `json:"id" form:"id" query:"id" long:"id" msg:"id"`
		Name     	string `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	}
	TenantAdminUpdateCoaChildOut struct {
		ResponseCommon
		Coas *[]rqAuth.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
	}
)

const (
	TenantAdminUpdateCoaChildAction = `tenantAdmin/updateCoaChild`

	ErrTenantAdminUpdateCoaChildUnauthorized      = `unauthorized user to update coa child`
	ErrTenantAdminUpdateCoaChildTenantNotFound    = `tenant admin not found to update coa child`
	ErrTenantAdminUpdateCoaChildCoaChildNotFound  = `coa child not found, make sure the id is valid`
	ErrTenantAdminUpdateCoaChildFailed = `failed to update coa child`
)

func (d *Domain) TenantAdminUpdateCoaChild(in *TenantAdminUpdateCoaChildIn) (out TenantAdminUpdateCoaChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminUpdateCoaChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
		out.SetError(400, ErrTenantAdminUpdateCoaChildTenantNotFound)
		return
	}

	// update coa child
	child := wcAuth.NewCoaMutator(d.AuthOltp)
	child.Id = in.Id
	if !child.FindById() {
		out.SetError(400, ErrTenantAdminUpdateCoaChildCoaChildNotFound)
		return
	}

	child.SetName(in.Name)
	child.SetUpdatedAt(in.UnixNow())
	child.SetUpdatedBy(sess.UserId)
	
	if !child.DoUpdateById() {
		child.HaveMutation()
		out.SetError(400, ErrTenantAdminUpdateCoaChildFailed)
		return
	}

	// retrieve owned coa
	coa := wcAuth.NewCoaMutator(d.AuthOltp)
	coas := coa.FindCoasByTenant(tenant.TenantCode)

	out.Coas = &coas

	return
}
