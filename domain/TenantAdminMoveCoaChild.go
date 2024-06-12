package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminMoveCoaChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminMoveCoaChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminMoveCoaChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminMoveCoaChild.go
//go:generate farify doublequote --file TenantAdminMoveCoaChild.go

// TODO:HABIBI make file naming consistent with url requesting

type (
	TenantAdminMoveCoaChildIn struct {
		RequestCommon
		Id         uint64 `json:"id" form:"id" query:"id" long:"id" msg:"id"`
		MoveToIdx  int    `json:"moveToIdx" form:"moveToIdx" query:"moveToIdx" long:"moveToIdx" msg:"moveToIdx"`
		ToParentId uint64 `json:"toParentId" form:"toParentId" query:"toParentId" long:"toParentId" msg:"toParentId"`
	}

	TenantAdminMoveCoaChildOut struct {
		ResponseCommon
		Org  *rqAuth.Orgs   `json:"org" form:"org" query:"org" long:"org" msg:"org"`
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminMoveCoaChildAction = `tenantAdmin/moveCoaChild`

	ErrTenantAdminMoveCoaChildUnauthorized       = `unauthorized user to move this organization`
	ErrTenantAdminMoveCoaChildTenantNotFound     = `cannot move organization if you are not a tenant admin`
	ErrTenantAdminMoveCoaChildOrgNotFound        = `invalid organization to move`
	ErrTenantAdminMoveCoaChildParentNotFound     = `parent not found to move this organization`
	ErrTenantAdminMoveCoaChildToParentNotFound   = `cannot found parent to move this organization`
	ErrTenantAdminMoveCoaChildShouldNotCompany   = `cannot move organization if type is company`
	ErrTenantAdminMoveCoaChildFailedMoveChildren = `failed to move organization child`
	ErrTenantAdminMoveCoaChildMustSameParentType = `cannot move to other parent if different organization type`
	ErrTenantAdminMoveCoaChildFailedUpdateChild  = `failed to update organization child`
)

func (d *Domain) TenantAdminMoveCoaChild(in *TenantAdminMoveCoaChildIn) (out TenantAdminMoveCoaChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminMoveCoaChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
		out.SetError(400, ErrTenantAdminMoveCoaChildTenantNotFound)
		return
	}

	return
}