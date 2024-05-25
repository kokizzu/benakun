package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminBudgeting.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminBudgeting.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminBudgeting.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminBudgeting.go
//go:generate farify doublequote --file TenantAdminBudgeting.go

type (
	TenantAdminBudgetingIn struct {
		RequestCommon
	}
	TenantAdminBudgetingOut struct {
		ResponseCommon
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminBudgetingAction = `tenantAdmin/budgeting`
)

func (d *Domain) TenantAdminBudgeting(in *TenantAdminBudgetingIn) (out TenantAdminBudgetingOut) {
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
