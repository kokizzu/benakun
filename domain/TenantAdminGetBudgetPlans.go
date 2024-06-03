package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBudget/rqBudget"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminGetBudgetPlans.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminGetBudgetPlans.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminGetBudgetPlans.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminGetBudgetPlans.go
//go:generate farify doublequote --file TenantAdminGetBudgetPlans.go

type (
	TenantAdminGetBudgetPlansIn struct {
		RequestCommon
		OrgId uint64 `json:"orgId" form:"orgId" query:"orgId" long:"orgId" msg:"orgId"`
	}
	TenantAdminGetBudgetPlansOut struct {
		ResponseCommon
		Plans *[]rqBudget.Plans `json:"plans" form:"plans" query:"plans" long:"plans" msg:"plans"`
	}
)

const (
	TenantAdminGetBudgetPlansAction = `tenantAdmin/getBudgetPlans`

	ErrTenantAdminGetBudgetPlanUnauthorized = `unauthorized user to get budget plans`
	ErrTenantAdminGetBudgetPlanTenantNotFound = `tenant admin not found to get budget plans`
	ErrTenantAdminGetBudgetPlanOrgNotFound = `organization not found to get budget plans`
	ErrTenantAdminGetBudgetPlanFailed = `failed to get budget plans`
)

func (d *Domain) TenantAdminGetBudgetPlans(in *TenantAdminGetBudgetPlansIn) (out TenantAdminGetBudgetPlansOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminGetBudgetPlanUnauthorized)
		return
	}

	tenant := rqAuth.NewTenants(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminGetBudgetPlanTenantNotFound)
		return
	}

	org := rqAuth.NewOrgs(d.AuthOltp)
	org.Id = in.OrgId
	if !org.FindById() {
		out.SetError(400, ErrTenantAdminGetBudgetPlanOrgNotFound)
		return
	}

	plans := rqBudget.NewPlans(d.AuthOltp)
	toPlans := plans.FindPlansByOrg(in.OrgId)

	out.Plans = &toPlans

	return
}
