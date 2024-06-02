package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBudget"
	"benakun/model/mBudget/rqBudget"
	"benakun/model/mBudget/wcBudget"
	"time"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminUpsertBudgetPlan.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminUpsertBudgetPlan.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminUpsertBudgetPlan.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminUpsertBudgetPlan.go
//go:generate farify doublequote --file TenantAdminUpsertBudgetPlan.go

type (
	TenantAdminUpsertBudgetPlanIn struct {
		RequestCommon
		Plan rqBudget.Plans
	}
	TenantAdminUpsertBudgetPlanOut struct {
		ResponseCommon
		Plans *[]rqBudget.Plans `json:"plans" form:"plans" query:"plans" long:"plans" msg:"plans"`
	}
)

const (
	TenantAdminUpsertBudgetPlanAction = `tenantAdmin/upsertBudgetPlan`

	ErrTenantAdminUpsertBudgetPlanUnauthorized       = `unauthorized user to upsert plan`
	ErrTenantAdminUpsertBudgetPlanTenantNotFound     = `tenant admin not found to upsert plan`
	ErrTenantAdminUpsertBudgetPlanOrgNotFound        = `organization not found to upsert plan`
	ErrTenantAdminUpsertBudgetPlanOrgDiffer          = `cannot overwrite different organization's plan`
	ErrTenantAdminUpsertBudgetPlanParentPlanNotFound = `parent plan not found, don't add parent instead`
	ErrTenantAdminUpsertBudgetPlanNotFound           = `plan to update not found`
	ErrTenantAdminUpsertBudgetPlanInvalidPlanType    = `invalid plan type`
	ErrTenantAdminUpsertBudgetPlanFailed             = `failed to upsert plan`
	ErrTenantAdminUpsertBudgetPlanVisionExist        = `vision already exist`
	ErrTenantAdminUpsertBudgetPlanMissionExist       = `mission already exist`
)

func (d *Domain) TenantAdminUpsertBudgetPlan(in *TenantAdminUpsertBudgetPlanIn) (out TenantAdminUpsertBudgetPlanOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminUpsertBudgetPlanUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminUpsertBudgetPlanTenantNotFound)
		return
	}

	// TODO: use it on production
	// if hostmap[in.Host].OrgId <= 0 {
	// 	out.SetError(400, ErrTenantAdminUpsertBudgetPlanOrgNotFound)
	// 	return
	// }

	if !mBudget.IsValidPlanType(in.Plan.PlanType) {
		out.SetError(400, ErrTenantAdminUpsertBudgetPlanInvalidPlanType)
		return
	}

	if in.Plan.PlanType == mBudget.PlanTypeActivity {
		planParent := wcBudget.NewPlansMutator(d.AuthOltp)
		planParent.Id = in.Plan.ParentId
		if !planParent.FindById() {
			out.SetError(400, ErrTenantAdminUpsertBudgetPlanParentPlanNotFound)
			return
		}
	}

	plan := wcBudget.NewPlansMutator(d.AuthOltp)

	if in.Plan.Id > 0 {
		plan.Id = in.Plan.Id
		if !plan.FindById() {
			out.SetError(400, ErrTenantAdminUpsertBudgetPlanNotFound)
			return
		}
	}
	plan.SetAll(in.Plan, nil, nil)
	if in.Plan.YearOf == 0 {
		plan.SetYearOf(uint64(time.Now().Year()))
	}

	differentPlanId := in.Plan.Id != plan.Id
	switch plan.PlanType {
	case mBudget.PlanTypeVision:
		if plan.IsVisionExist() && differentPlanId {
			out.SetError(400, ErrTenantAdminUpsertBudgetPlanVisionExist)
			return
		}
		plan.Title = mBudget.TitleVision
	case mBudget.PlanTypeMission:
		if plan.IsMissionExist() && differentPlanId {
			out.SetError(400, ErrTenantAdminUpsertBudgetPlanMissionExist)
			return
		}
		plan.Title = mBudget.TitleMission
	default:
	}

	plan.UpdatedAt = in.UnixNow()

	if !plan.DoUpsertById() {
		out.SetError(400, ErrTenantAdminUpsertBudgetPlanFailed)
		return
	}

	plans := rqBudget.NewPlans(d.AuthOltp)
	toPlans := plans.FindPlansByOrg(plan.OrgId)

	out.Plans = &toPlans

	return
}
