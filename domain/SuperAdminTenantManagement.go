package domain

import (
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"

	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file SuperAdminTenantManagement.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type SuperAdminTenantManagement.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type SuperAdminTenantManagement.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type SuperAdminTenantManagement.go
//go:generate farify doublequote --file SuperAdminTenantManagement.go

type (
	SuperAdminTenantManagementIn struct {
		RequestCommon
		Cmd      string         `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		Tenant   rqAuth.Tenants `json:"tenant" form:"tenant" query:"tenant" long:"tenant" msg:"tenant"`
		WithMeta bool           `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn  `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	SuperAdminTenantManagementOut struct {
		ResponseCommon
		Pager  zCrud.PagerOut  `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta   *zCrud.Meta     `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Tenant *rqAuth.Tenants `json:"tenant" form:"tenant" query:"tenant" long:"tenant" msg:"tenant"`
		// listing
		Tenants [][]any `json:"tenants" form:"tenants" query:"tenants" long:"tenants" msg:"tenants"`
	}
)

const (
	SuperAdminTenantManagementAction = `superAdmin/tenantManagement`
	ErrTenantIdNotFound              = `tenant id not found`
	ErrTenantSaveFailed              = `tenant save failed`
)

var SuperAdminTenantManagementMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mAuth.Id,
			Label:     "ID",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
		},
		{
			Name:      mAuth.TenantCode,
			Label:     "Tenant Code",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeCombobox,
		},
		{
			Name:      mAuth.CreatedAt,
			Label:     `Created At`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mAuth.UpdatedAt,
			Label:     `Updated At`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mAuth.DeletedAt,
			Label:     `Deleted At`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
	},
}

func (d *Domain) SuperAdminTenantManagement(in *SuperAdminTenantManagementIn) (out SuperAdminTenantManagementOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.refId = in.Tenant.Id

	sess := d.MustSuperAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &SuperAdminTenantManagementMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		if in.Tenant.Id <= 0 {
			out.Meta = &SuperAdminTenantManagementMeta
		}

		tenant := rqAuth.NewTenants(d.AuthOltp)
		tenant.Id = in.Tenant.Id
		if !tenant.FindById() {
			out.SetError(400, ErrTenantIdNotFound)
			return
		}
		out.Tenant = tenant
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.Id = in.Tenant.Id
		createTables := false
		if tenant.Id > 0 {
			if !tenant.FindById() {
				out.SetError(400, ErrTenantIdNotFound)
				return
			}

			tenant.SetAll(in.Tenant, M.SB{}, M.SB{})

			if in.Cmd == zCrud.CmdDelete {
				if tenant.DeletedAt == 0 {
					tenant.SetDeletedAt(in.UnixNow())
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if tenant.DeletedAt > 0 {
					tenant.SetDeletedAt(0)
				}
				L.Print(`tenant.DeletedAt`, tenant.DeletedAt)
			}
		} else {
			tenant.SetCreatedAt(in.UnixNow())
			createTables = true
		}

		if tenant.HaveMutation() {
			tenant.SetUpdatedAt(in.UnixNow())
			tenant.SetUpdatedBy(sess.UserId)
			if tenant.Id == 0 {
				tenant.SetCreatedAt(in.UnixNow())
			}
		}

		if !tenant.DoUpsert() {
			out.SetError(500, ErrTenantSaveFailed)
		}

		out.Tenant = &tenant.Tenants
		out.Tenant.Adapter = nil

		L.Print(`tenant.DeletedAt`, tenant.DeletedAt)
		if createTables {
			// TODO: migrate tables for this tenant
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		t := rqAuth.NewTenants(d.AuthOltp)
		out.Tenants = t.FindByPagination(&SuperAdminTenantManagementMeta, &in.Pager, &out.Pager)

		// TODO: decorate with how many users, orgs, coa, etc that per tenant
	}
	return
}
