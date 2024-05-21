package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminOrganization.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminOrganization.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminOrganization.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminOrganization.go
//go:generate farify doublequote --file TenantAdminOrganization.go

type (
	TenantAdminOrganizationIn struct {
		RequestCommon
	}

	TenantAdminOrganizationOut struct {
		ResponseCommon
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminOrganizationAction = `tenantAdmin/organization`

	ErrTenantAdminOrganizationUnauthorized                 = `unauthorized user`
	ErrTenantAdminOrganizationTenantNotFound               = `tenant admin not found`
	ErrTenantAdminOrganizationOldParentNotFound            = `old parent of the organization cannot be found`
	ErrTenantAdminOrganizationNewParentNotFound            = `new parent of the organization cannot be found`
	ErrTenantAdminOrganizationUpdatedOldParentChildren     = `cannot update childs of old parent`
	ErrTenantAdminOrganizationUpdatedNewParentChildren     = `cannot update childs of new parent`
	ErrTenantAdminOrganizationUpdatedParentForOrganization = `cannot update parent for organization`
	ErrTenantAdminOrganizationNotFound                     = `cannot find the organization`
	ErrTenantAdminOrganizationChildsNotFound               = `cannot find childs for update`
	ErrTenantAdminOrganizationUpdatedChilds                = `cannot update childs`
)

func (d *Domain) TenantAdminOrganization(in *TenantAdminOrganizationIn) (out TenantAdminOrganizationOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminCoaUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
		out.SetError(400, ErrTenantAdminCoaTenantNotFound)
		return
	}

	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	orgs := org.FindOrgsByTenant(tenant.TenantCode)

	out.Orgs = &orgs

	return
}