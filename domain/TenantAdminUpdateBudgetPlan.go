package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBudget"
	"benakun/model/mBudget/rqBudget"
	"benakun/model/mBudget/wcBudget"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminUpdateBudgetPlan.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminUpdateBudgetPlan.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminUpdateBudgetPlan.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminUpdateBudgetPlan.go
//go:generate farify doublequote --file TenantAdminUpdateBudgetPlan.go

type (
	TenantAdminUpdateBudgetPlanIn struct {
		RequestCommon
		Id uint64 `json:"id" form:"id" query:"id" long:"id" msg:"id"`
		PlanType string `json:"planType" form:"planType" query:"planType" long:"planType" msg:"planType"`
		Title     string `json:"title" form:"title" query:"title" long:"title" msg:"title"`
		Description string `json:"description" form:"description" query:"description" long:"description" msg:"description"`
		PerYear int64 `json:"perYear" form:"perYear" query:"perYear" long:"perYear" msg:"perYear"`
		BudgetIDR int64 `json:"budgetIDR" form:"budgetIDR" query:"budgetIDR" long:"budgetIDR" msg:"budgetIDR"`
		BudgetUSD int64 `json:"budgetUSD" form:"budgetUSD" query:"budgetUSD" long:"budgetUSD" msg:"budgetUSD"`
		BudgetEUR int64 `json:"budgetEUR" form:"budgetEUR" query:"budgetEUR" long:"budgetEUR" msg:"budgetEUR"`
	}
	TenantAdminUpdateBudgetPlanOut struct {
		ResponseCommon
		Plans *[]rqBudget.Plans `json:"plans" form:"plans" query:"plans" long:"plans" msg:"plans"`
	}
)

const (
	TenantAdminUpdateBudgetPlanAction = `tenantAdmin/updateBudgetPlan`

	ErrTenantAdminUpdateBudgetPlanUnauthorized = `unauthorized user to update plan`
	ErrTenantAdminUpdateBudgetPlanTenantNotFound = `tenant admin not found to update plan`
	ErrTenantAdminUpdateBudgetPlanOrgNotFound = `organization not found to update plan`
	ErrTenantAdminUpdateBudgetPlanInvalidOrg = `invalid organization type to update plan`
	ErrTenantAdminUpdateBudgetPlanInvalidPlanType = `invalid plan type`
	ErrTenantAdminUpdateBudgetPlanFailed = `failed to update plan`
	ErrTenantAdminUpdateBudgetPlanNotFound = `budget plan not found`
)

func (d *Domain) TenantAdminUpdateBudgetPlan(in *TenantAdminUpdateBudgetPlanIn) (out TenantAdminUpdateBudgetPlanOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminUpdateBudgetPlanUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminUpdateBudgetPlanTenantNotFound)
		return
	}

	plan := wcBudget.NewPlansMutator(d.AuthOltp)
	plan.Id = in.Id
	if !plan.FindById() {
		out.SetError(400, ErrTenantAdminUpdateBudgetPlanNotFound)
		return
	}

	if in.Description != plan.Description {
		plan.Description = in.Description
	}

	// vision and mission cannot edit fields other than description
	if !(plan.PlanType == mBudget.PlanTypeVision || plan.PlanType == mBudget.PlanTypeMission) {
		if in.Title != plan.Title {
			plan.Title = in.Title
		}
		if in.PerYear != plan.PerYear {
			plan.PerYear = in.PerYear
		}
		if in.BudgetIDR != plan.BudgetIDR {
			plan.BudgetIDR = in.BudgetIDR
		}
		if in.BudgetUSD != plan.BudgetUSD {
			plan.BudgetUSD = in.BudgetUSD
		}
		if in.BudgetEUR != plan.BudgetEUR {
			plan.BudgetEUR = in.BudgetEUR
		}
	}

	if !plan.DoUpdateById() {
		out.SetError(400, ErrTenantAdminUpdateBudgetPlanFailed)
		return
	}

	plans := wcBudget.NewPlansMutator(d.AuthOltp)
	toPlans := plans.FindPlansByOrg(plan.OrgId)

	out.Plans = &toPlans

	return
}
