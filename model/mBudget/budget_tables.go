package mBudget

import (
	"github.com/kokizzu/gotro/D/Tt"
)

// TODO: make this table per tenant (run ony migration per tenant

func ValidPlanType(pType string) bool {
	switch pType {
		case `vision`:
			return true
		case `mission`:
			return true
		case `program`:
			return true
		case `activity`:
			return true
		default:
			return false
	}
}

const (
	TablePlans Tt.TableName = `plans`

	Id       = `id`
	PlanType = `planType`
	// visions defines the "what the endgame" what we want to achieve qualitatively
	// mission defines the "why" to achieve the vision quantitatively, must be specific, measurable, accountable, result-oriented, time-bound
	// program defines the when, where, how to support the mission
	// activity defines the who, how much cost, how long, how often to support the program
	// each activity will have disbursement to make advance payment or anything needed

	/*
		example:

		Vision: To be the leading software development company in the region,
			providing innovative and high-quality solutions to our clients.

		Mission: To provide our clients with the best software solutions that meet their needs and exceed their expectations.
			We strive to achieve this by leveraging the latest technologies,
			hiring the best talent, and maintaining a culture of excellence.
			Our goal is to achieve a customer satisfaction rate of 95% by the end of 2024.

		Programs:
			Social Media Marketing Brand Awareness
			System Information Development and Customization
			Trainings

		Activity:
			- marketing
			- analysis
			- design
			- development with fully documented
			- automated testing
			- automated deployment
			- maintenance (server cost)
			- support (customer service)
			- operational cost (office, electricity, internet, etc)
	*/

	ParentId    = `parentId`
	CreatedAt   = `createdAt`
	CreatedBy   = `createdBy`
	UpdatedAt   = `updatedAt`
	UpdatedBy   = `updatedBy`
	DeletedAt   = `deletedAt`
	Title       = `title`
	Description = `description`
	OrgId       = `orgId`
	PerYear     = `perYear`
	BudgetIDR   = `budgetIDR`
	BudgetUSD   = `budgetUSD`
	BudgetEUR   = `budgetEUR`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TablePlans: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{PlanType, Tt.String},
			{ParentId, Tt.Unsigned},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{Title, Tt.String},
			{Description, Tt.String},
			{OrgId, Tt.Unsigned},
			{PerYear, Tt.Integer},
			{BudgetIDR, Tt.Integer},
			{BudgetUSD, Tt.Integer},
			{BudgetEUR, Tt.Integer},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
	},
}
