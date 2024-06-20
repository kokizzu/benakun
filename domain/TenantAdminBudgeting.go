package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mBudget"
	"benakun/model/mBudget/rqBudget"
	"benakun/model/mBudget/wcBudget"
	"benakun/model/zCrud"
	"time"

	"github.com/kokizzu/gotro/I"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminBudgeting.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminBudgeting.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminBudgeting.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminBudgeting.go
//go:generate farify doublequote --file TenantAdminBudgeting.go

type (
	TenantAdminBudgetingIn struct {
		RequestCommon
		Cmd     string        	`json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		OrgId 	uint64 					`json:"orgId" form:"orgId" query:"orgId" long:"orgId" msg:"orgId"`
		Plan 		rqBudget.Plans	`json:"plan" form:"plan" query:"plan" long:"plan" msg:"plan"`
	}
	TenantAdminBudgetingOut struct {
		ResponseCommon
		Orgs 	*[]rqAuth.Orgs 		`json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
		Plans *[]rqBudget.Plans `json:"plans" form:"plans" query:"plans" long:"plans" msg:"plans"`
	}
)

const (
	TenantAdminBudgetingAction = `tenantAdmin/budgeting`

	ErrTenantAdminBudgetingUnauthorized 			= `unauthorized user`
	ErrTenantAdminBudgetingTenantNotFound 		= `tenant admin not found`
	ErrTenantAdminBudgetingOrgNotFound				= `organization not found to get budget plans`
	ErrTenantAdminBudgetingPlanNotFound				= `budget plan not found`
	ErrTenantAdminBudgetingParentPlanNotFound	= `could not found parent of budget plan`
	ErrTenantAdminBudgetingOrgPlanNotFound		= `organization not found to modify budget plan`
	ErrTenantAdminBudgetingInvalidPlanType    = `invalid plan type`
	ErrTenantAdminBudgetingSaveFailed					= `tenant admin budgeting save failed`
)

func (d *Domain) TenantAdminBudgeting(in *TenantAdminBudgetingIn) (out TenantAdminBudgetingOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := rqAuth.NewUsers(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminBudgetingUnauthorized)
		return
	}

	tenant := rqAuth.NewTenants(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
		out.SetError(400, ErrTenantAdminBudgetingTenantNotFound)
		return
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		if in.OrgId != 0 {
			org := rqAuth.NewOrgs(d.AuthOltp)
			org.Id = in.OrgId
			if !org.FindById() {
				out.SetError(400, ErrTenantAdminBudgetingOrgNotFound)
				return
			}

			plans := rqBudget.NewPlans(d.AuthOltp)
			toPlans := plans.FindPlansByOrg(org.Id)

			out.Plans = &toPlans
		}
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		plan := wcBudget.NewPlansMutator(d.AuthOltp)
		plan.Id = in.Plan.Id

		org := rqAuth.NewOrgs(d.AuthOltp)
		org.Id = in.Plan.OrgId
		if !org.FindById() {
			out.SetError(400, ErrTenantAdminBudgetingOrgPlanNotFound)
			return
		}

		if plan.Id > 0 {
			if !plan.FindById() {
				out.SetError(400, ErrTenantAdminBudgetingPlanNotFound)
				return
			}

			if plan.PlanType == mBudget.PlanTypeActivity {
				planParent := rqBudget.NewPlans(d.AuthOltp)
				planParent.Id = in.Plan.ParentId
				if !planParent.FindById() {
					out.SetError(400, ErrTenantAdminBudgetingParentPlanNotFound)
					return
				}
			}

			switch in.Cmd {
			case zCrud.CmdUpsert:
				if plan.Title != in.Plan.Title{
					plan.SetTitle(in.Plan.Title)
				}

				if plan.Description != in.Plan.Description {
					plan.SetDescription(in.Plan.Description)
				}

				if len(I.UToS(in.Plan.YearOf)) == 4 {
					if plan.YearOf != in.Plan.YearOf {
						plan.SetYearOf(in.Plan.YearOf)
					}
				}

				if plan.BudgetIDR != in.Plan.BudgetIDR {
					plan.SetBudgetIDR(in.Plan.BudgetIDR)
				}

				if plan.Quantity != in.Plan.Quantity {
					plan.SetQuantity(in.Plan.Quantity)
				}

				if plan.Unit != in.Plan.Unit {
					plan.SetUnit(in.Plan.Unit)
				}
			case zCrud.CmdDelete:
				if !(plan.PlanType == mBudget.PlanTypeMission || plan.PlanType == mBudget.PlanTypeVision) {
					if plan.DeletedAt == 0 {
						plan.SetDeletedAt(in.UnixNow())
						plan.SetDeletedBy(sess.UserId)
					}
				}
			case zCrud.CmdRestore:
				if !(plan.PlanType == mBudget.PlanTypeMission || plan.PlanType == mBudget.PlanTypeVision) {
					if plan.DeletedAt > 0 {
						plan.SetDeletedAt(0)
						plan.SetRestoredBy(sess.UserId)
					}
				}
			}
		} else {
			if !mBudget.IsValidPlanType(in.Plan.PlanType) {
				out.SetError(400, ErrTenantAdminBudgetingInvalidPlanType)
				return
			}

			plan.SetTitle(in.Plan.Title)
			plan.SetDescription(in.Plan.Description)

			if len(I.UToS(in.Plan.YearOf)) != 4 || in.Plan.YearOf == 0 {
				plan.SetYearOf(uint64(time.Now().Year()))
			} else {
				plan.SetYearOf(in.Plan.YearOf)
			}

			plan.SetBudgetIDR(in.Plan.BudgetIDR)
			plan.SetQuantity(in.Plan.Quantity)
			plan.SetUnit(in.Plan.Unit)
			plan.SetCreatedAt(in.UnixNow())
			plan.SetCreatedBy(sess.UserId)
		}

		plan.SetUpdatedAt(in.UnixNow())
		plan.SetUpdatedBy(sess.UserId)

		if !plan.DoUpsertById() {
			out.SetError(500, ErrTenantAdminBudgetingSaveFailed)
			return
		}

		plans := plan.FindPlansByOrg(org.Id)
		out.Plans = &plans
	case zCrud.CmdList:
		org := rqAuth.NewOrgs(d.AuthOltp)
		orgs := org.FindOrgsByTenant(tenant.TenantCode)

		out.Orgs = &orgs
	}

	return
}
