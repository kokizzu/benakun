package domain

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
	}
)

const (
	TenantAdminBudgetingAction = `tenantAdmin/budgeting`
)

func (d *Domain) TenantAdminBudgeting(in *TenantAdminBudgetingIn) (out TenantAdminBudgetingOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	return
}
