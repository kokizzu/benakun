package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminUpsertCoaChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminUpsertCoaChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminUpsertCoaChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminUpsertCoaChild.go
//go:generate farify doublequote --file TenantAdminUpsertCoaChild.go

// TODO:HABIBI make file naming consistent with url requesting

type (
	TenantAdminUpsertCoaChildIn struct {
		RequestCommon
		Coa rqFinance.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
	}
	TenantAdminUpsertCoaChildOut struct {
		ResponseCommon
		Coas *[]rqFinance.Coa `json:"coas" form:"coas" query:"coas" long:"coas" msg:"coas"`
	}
)

const (
	TenantAdminUpsertCoaChildAction = `tenantAdmin/createCoaChild`

	ErrTenantAdminUpsertCoaChildUnauthorized      = `unauthorized user to create coa child`
	ErrTenantAdminUpsertCoaChildTenantNotFound    = `tenant admin not found to create coa child`
	ErrTenantAdminUpsertCoaChildCoaParentNotFound = `coa parent not found for this coa child`
	ErrTenantAdminUpsertCoaChildCoaChildNotFound  = `coa child not found when adds child to parent`
	ErrTenantAdminUpsertCoaChildFailedSave        = `coa upsert failed`
)

func (d *Domain) TenantAdminUpsertCoaChild(in *TenantAdminUpsertCoaChildIn) (out TenantAdminUpsertCoaChildOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminUpsertCoaChildUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminUpsertCoaChildTenantNotFound)
		return
	}

	// find parent id
	parent := wcFinance.NewCoaMutator(d.AuthOltp)
	parent.Id = in.Coa.ParentId
	if !parent.FindById() {
		out.SetError(400, ErrTenantAdminUpsertCoaChildCoaParentNotFound)
		return
	}

	// insert coa child
	coa := wcFinance.NewCoaMutator(d.AuthOltp)
	if in.Coa.Id > 0 {
		coa.Id = in.Coa.Id
		if !coa.FindById() {
			out.SetError(400, ErrTenantAdminUpsertCoaChildCoaChildNotFound)
			return
		}
	} else {
		coa.SetCreatedAt(in.UnixNow())
		coa.SetCreatedBy(sess.UserId)
	}

	coa.SetAll(in.Coa, nil, nil)
	coa.SetUpdatedAt(in.UnixNow())
	coa.SetUpdatedBy(sess.UserId)
	coa.SetTenantCode(tenant.TenantCode)

	if !coa.DoUpsertById() {
		coa.HaveMutation()
		out.SetError(400, ErrTenantAdminUpsertCoaChildFailedSave)
		return
	}

	if in.Coa.Id <= 0 {
		// update childrens of parent coa
		children := parent.Children
		children = append(children, coa.Id)
		parent.SetChildren(children)
		parent.SetUpdatedAt(in.UnixNow())
		parent.SetUpdatedBy(sess.UserId)

		if !parent.DoUpdateById() {
			parent.HaveMutation()
			out.SetError(400, ErrTenantAdminUpsertCoaChildCoaParentNotFound)
			return
		}
	}

	// retrieve owned coa
	r := rqFinance.NewCoa(d.AuthOltp)
	coas := r.FindCoasByTenant(tenant.TenantCode)

	out.Coas = &coas

	return
}
