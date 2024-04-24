package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBudget"
	"benakun/model/mBudget/rqBudget"
	"benakun/model/mBudget/wcBudget"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminCreatePlan.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminCreatePlan.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminCreatePlan.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminCreatePlan.go
//go:generate farify doublequote --file TenantAdminCreatePlan.go

type (
	TenantAdminCreatePlanIn struct {
		RequestCommon
		PlanType string `json:"planType" form:"planType" query:"planType" long:"planType" msg:"planType"`
		Title     string `json:"title" form:"title" query:"title" long:"title" msg:"title"`
		Description string `json:"description" form:"description" query:"description" long:"description" msg:"description"`
		ParentId uint64 `json:"parentId" form:"parentId" query:"parentId" long:"parentId" msg:"parentId"`
		OrgId uint64 `json:"orgId" form:"orgId" query:"orgId" long:"orgId" msg:"orgId"`
		PerYear int64 `json:"perYear" form:"perYear" query:"perYear" long:"perYear" msg:"perYear"`
		BudgetIDR int64 `json:"bugetIDR" form:"bugetIDR" query:"bugetIDR" long:"bugetIDR" msg:"bugetIDR"`
		BudgetUSD int64 `json:"budgetUSD" form:"budgetUSD" query:"budgetUSD" long:"budgetUSD" msg:"budgetUSD"`
		BudgetEUR int64 `json:"budgetEUR" form:"budgetEUR" query:"budgetEUR" long:"budgetEUR" msg:"budgetEUR"`
	}
	TenantAdminCreatePlanOut struct {
		ResponseCommon
		Plans *[]rqBudget.Plans `json:"plans" form:"plans" query:"plans" long:"plans" msg:"plans"`
	}
)

const (
	TenantAdminCreatePlanAction = `tenantAdmin/createCreatePlan`

	ErrTenantAdminCreatePlanUnauthorized = `unauthorized user to create plan`
	ErrTenantAdminCreatePlanTenantNotFound = `tenant admin not found to create plan`
	ErrTenantAdminCreatePlanOrgNotFound = `organization not found to create plan`
	ErrTenantAdminCreatePlanInvalidOrg = `invalid organization type to create plan`
	ErrTenantAdminCreatePlanParentPlanNotFound = `parent plan not found, don't add parent instead`
	ErrTenantAdminCreatePlanInvalidPlanType = `invalid plan type`
	ErrTenantAdminCreatePlanFailed = `failed to create plan`
)

func (d *Domain) TenantAdminCreatePlan(in *TenantAdminCreatePlanIn) (out TenantAdminCreatePlanOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminCreatePlanUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminCreatePlanTenantNotFound)
		return
	}

	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	org.Id = in.OrgId
	if !org.FindById() {
		out.SetError(400, ErrTenantAdminCreatePlanOrgNotFound)
		return
	}

	if org.OrgType == mAuth.OrgTypeJob {
		out.SetError(400, ErrTenantAdminCreatePlanInvalidOrg)
		return
	}

	if in.ParentId != 0 {
		planParent := wcBudget.NewPlansMutator(d.AuthOltp)
		planParent.Id = in.ParentId
		if !planParent.FindById() {
			out.SetError(400, ErrTenantAdminCreatePlanParentPlanNotFound)
			return
		}
	}

	if !mBudget.ValidPlanType(in.PlanType) {
		out.SetError(400, ErrTenantAdminCreatePlanInvalidPlanType)
		return
	}

	plan := wcBudget.NewPlansMutator(d.AuthOltp)
	plan.PlanType = in.PlanType
	plan.Title = in.Title
	plan.Description = in.Description
	plan.ParentId = in.ParentId
	plan.OrgId = in.OrgId
	plan.PerYear = in.PerYear
	plan.BudgetIDR = in.BudgetIDR
	plan.BudgetUSD = in.BudgetUSD
	plan.BudgetEUR = in.BudgetEUR

	if !plan.DoInsert() {
		out.SetError(400, ErrTenantAdminCreatePlanFailed)
		return
	}

	plans := wcBudget.NewPlansMutator(d.AuthOltp)
	toPlans := plans.FindPlansByOrg(in.OrgId)

	out.Plans = &toPlans

	return
}
