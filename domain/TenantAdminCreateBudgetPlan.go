package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBudget"
	"benakun/model/mBudget/rqBudget"
	"benakun/model/mBudget/wcBudget"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminCreateBudgetPlan.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminCreateBudgetPlan.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminCreateBudgetPlan.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminCreateBudgetPlan.go
//go:generate farify doublequote --file TenantAdminCreateBudgetPlan.go

type (
	TenantAdminCreateBudgetPlanIn struct {
		RequestCommon
		PlanType string `json:"planType" form:"planType" query:"planType" long:"planType" msg:"planType"`
		Title     string `json:"title" form:"title" query:"title" long:"title" msg:"title"`
		Description string `json:"description" form:"description" query:"description" long:"description" msg:"description"`
		ParentId uint64 `json:"parentId" form:"parentId" query:"parentId" long:"parentId" msg:"parentId"`
		OrgId uint64 `json:"orgId" form:"orgId" query:"orgId" long:"orgId" msg:"orgId"`
		PerYear int64 `json:"perYear" form:"perYear" query:"perYear" long:"perYear" msg:"perYear"`
		BudgetIDR int64 `json:"budgetIDR" form:"budgetIDR" query:"budgetIDR" long:"budgetIDR" msg:"budgetIDR"`
		BudgetUSD int64 `json:"budgetUSD" form:"budgetUSD" query:"budgetUSD" long:"budgetUSD" msg:"budgetUSD"`
		BudgetEUR int64 `json:"budgetEUR" form:"budgetEUR" query:"budgetEUR" long:"budgetEUR" msg:"budgetEUR"`
	}
	TenantAdminCreateBudgetPlanOut struct {
		ResponseCommon
		Plans *[]rqBudget.Plans `json:"plans" form:"plans" query:"plans" long:"plans" msg:"plans"`
	}
)

const (
	TenantAdminCreateBudgetPlanAction = `tenantAdmin/createBudgetPlan`

	ErrTenantAdminCreateBudgetPlanUnauthorized = `unauthorized user to create plan`
	ErrTenantAdminCreateBudgetPlanTenantNotFound = `tenant admin not found to create plan`
	ErrTenantAdminCreateBudgetPlanOrgNotFound = `organization not found to create plan`
	ErrTenantAdminCreateBudgetPlanInvalidOrg = `invalid organization type to create plan`
	ErrTenantAdminCreateBudgetPlanParentPlanNotFound = `parent plan not found, don't add parent instead`
	ErrTenantAdminCreateBudgetPlanInvalidPlanType = `invalid plan type`
	ErrTenantAdminCreateBudgetPlanFailed = `failed to create plan`
	ErrTenantAdminCreateBudgetPlanVisionExist = `vision already exist`
	ErrTenantAdminCreateBudgetPlanMissionExist = `mission already exist`
)

func (d *Domain) TenantAdminCreateBudgetPlan(in *TenantAdminCreateBudgetPlanIn) (out TenantAdminCreateBudgetPlanOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminCreateBudgetPlanUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminCreateBudgetPlanTenantNotFound)
		return
	}

	org := wcAuth.NewOrgsMutator(d.AuthOltp)
	org.Id = in.OrgId
	if !org.FindById() {
		out.SetError(400, ErrTenantAdminCreateBudgetPlanOrgNotFound)
		return
	}

	if org.OrgType == mAuth.OrgTypeJob {
		out.SetError(400, ErrTenantAdminCreateBudgetPlanInvalidOrg)
		return
	}

	if !mBudget.ValidPlanType(in.PlanType) {
		out.SetError(400, ErrTenantAdminCreateBudgetPlanInvalidPlanType)
		return
	}

	if in.PlanType == mBudget.PlanTypeActivity {
		planParent := wcBudget.NewPlansMutator(d.AuthOltp)
		planParent.Id = in.ParentId
		if !planParent.FindById() {
			out.SetError(400, ErrTenantAdminCreateBudgetPlanParentPlanNotFound)
			return
		}
	}

	plan := wcBudget.NewPlansMutator(d.AuthOltp)

	plan.PlanType = in.PlanType
	plan.OrgId = in.OrgId
	
	switch in.PlanType {
	case mBudget.PlanTypeVision:
		if plan.IsVisionExist() {
			out.SetError(400, ErrTenantAdminCreateBudgetPlanVisionExist)
			return
		}
		plan.Title = mBudget.TitleVision
	case mBudget.PlanTypeMission:
		if plan.IsMissionExist() {
			out.SetError(400, ErrTenantAdminCreateBudgetPlanMissionExist)
			return
		}
		plan.Title = mBudget.TitleMission
	default:
		plan.Title = in.Title
	}

	plan.Description = in.Description
	plan.ParentId = in.ParentId
	plan.PerYear = in.PerYear
	plan.BudgetIDR = in.BudgetIDR
	plan.BudgetUSD = in.BudgetUSD
	plan.BudgetEUR = in.BudgetEUR
	plan.CreatedAt = in.UnixNow()
	plan.UpdatedAt = in.UnixNow()

	if !plan.DoInsert() {
		out.SetError(400, ErrTenantAdminCreateBudgetPlanFailed)
		return
	}

	plans := wcBudget.NewPlansMutator(d.AuthOltp)
	toPlans := plans.FindPlansByOrg(in.OrgId)

	out.Plans = &toPlans

	return
}
