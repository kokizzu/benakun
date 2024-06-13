package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
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
		Coa  *rqFinance.Coa   `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
		Coas *[]rqFinance.Coa `json:"coas" form:"coas" query:"coas" long:"coas" msg:"coas"`
	}
)

const (
	TenantAdminMoveCoaChildAction = `tenantAdmin/moveCoaChild`

	ErrTenantAdminMoveCoaChildUnauthorized       = `unauthorized user to move this coa`
	ErrTenantAdminMoveCoaChildTenantNotFound     = `cannot move coa if you are not a tenant admin`
	ErrTenantAdminMoveCoaChildCoaNotFound        = `coa not found`
	ErrTenantAdminMoveCoaChildParentNotFound     = `parent not found to move this coa`
	ErrTenantAdminMoveCoaChildToParentNotFound   = `cannot find parent to move this coa`
	ErrTenantAdminMoveCoaChildFailedMoveChildren = `failed to move coa child`
	ErrTenantAdminMoveCoaChildMustSameParentType = `cannot move to other parent if different coa type`
	ErrTenantAdminMoveCoaChildFailedUpdateChild  = `failed to update coa child`
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
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminMoveCoaChildTenantNotFound)
		return
	}

	child := wcFinance.NewCoaMutator(d.AuthOltp)
	child.Id = in.Id
	if !child.FindById() {
		out.SetError(400, ErrTenantAdminMoveCoaChildCoaNotFound )
		return
	}

	parent := wcFinance.NewCoaMutator(d.AuthOltp)
	parent.Id = child.ParentId
	if !parent.FindById() {
		out.SetError(400, ErrTenantAdminMoveCoaChildParentNotFound)
		return
	}

	if parent.Id == in.ToParentId {
		if len(parent.Children) >= 2 {
			children, err := moveChildToIndex(parent.Children, in.Id, in.MoveToIdx)
			if err != nil {
				out.SetError(400, ErrTenantAdminMoveCoaChildCoaNotFound )
				return
			}

			parent.SetChildren(children)
			if !parent.DoUpdateById() {
				parent.HaveMutation()
				out.SetError(400, ErrTenantAdminMoveCoaChildFailedMoveChildren)
				return
			}
		}

		out.Coa = &child.Coa

		coa := rqFinance.NewCoa(d.AuthOltp)
		coas := coa.FindCoasByTenant(tenant.TenantCode)

		out.Coas = &coas

		return
	}

	toParent := wcFinance.NewCoaMutator(d.AuthOltp)
	toParent.Id = in.ToParentId
	if !toParent.FindById() {
		out.SetError(400, ErrTenantAdminMoveCoaChildToParentNotFound)
		return
	}

	children := insertChildToIndex(toParent.Children, in.Id, in.MoveToIdx)

	child.SetParentId(in.ToParentId)
	if !child.DoUpdateById() {
		child.HaveMutation()
		out.SetError(400, ErrTenantAdminMoveCoaChildFailedUpdateChild)
		return
	}

	toParent.SetChildren(children)
	if !toParent.DoUpdateById() {
		toParent.HaveMutation()
		out.SetError(400, ErrTenantAdminMoveCoaChildFailedMoveChildren)
		return
	}

	children, err := removeChild(parent.Children, in.Id)
	if err != nil {
		out.SetError(400, ErrTenantAdminMoveCoaChildCoaNotFound)
		return
	}

	parent.SetChildren(children)
	if !parent.DoUpdateById() {
		parent.HaveMutation()
		out.SetError(400, ErrTenantAdminMoveCoaChildFailedMoveChildren)
		return
	}

	out.Coa = &child.Coa

	coa := rqFinance.NewCoa(d.AuthOltp)
	coas := coa.FindCoasByTenant(tenant.TenantCode)

	out.Coas = &coas

	return
}