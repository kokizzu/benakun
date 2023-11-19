package rqBudget

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mBudget"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Plans DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqBudget__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqBudget__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqBudget__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqBudget__ORM.GEN.go
type Plans struct {
	Adapter     *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id          uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	PlanType    string      `json:"planType" form:"planType" query:"planType" long:"planType" msg:"planType"`
	CreatedAt   int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy   uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt   int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy   uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt   int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	Title       string      `json:"title" form:"title" query:"title" long:"title" msg:"title"`
	Description string      `json:"description" form:"description" query:"description" long:"description" msg:"description"`
	OrgId       uint64      `json:"orgId,string" form:"orgId" query:"orgId" long:"orgId" msg:"orgId"`
	BudgetIDR   int64       `json:"budgetIDR" form:"budgetIDR" query:"budgetIDR" long:"budgetIDR" msg:"budgetIDR"`
	BudgetUSD   int64       `json:"budgetUSD" form:"budgetUSD" query:"budgetUSD" long:"budgetUSD" msg:"budgetUSD"`
	BudgetEUR   int64       `json:"budgetEUR" form:"budgetEUR" query:"budgetEUR" long:"budgetEUR" msg:"budgetEUR"`
}

// NewPlans create new ORM reader/query object
func NewPlans(adapter *Tt.Adapter) *Plans {
	return &Plans{Adapter: adapter}
}

// SpaceName returns full package and table name
func (p *Plans) SpaceName() string { //nolint:dupl false positive
	return string(mBudget.TablePlans) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (p *Plans) SqlTableName() string { //nolint:dupl false positive
	return `"plans"`
}

func (p *Plans) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (p *Plans) FindById() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{p.Id})
	if L.IsError(err, `Plans.FindById failed: `+p.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		p.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (p *Plans) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "planType"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "title"
	, "description"
	, "orgId"
	, "budgetIDR"
	, "budgetUSD"
	, "budgetEUR"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *Plans) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "planType"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "title"
	, "description"
	, "orgId"
	, "budgetIDR"
	, "budgetUSD"
	, "budgetEUR"
	`
}

// ToUpdateArray generate slice of update command
func (p *Plans) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, p.Id},
		A.X{`=`, 1, p.PlanType},
		A.X{`=`, 2, p.CreatedAt},
		A.X{`=`, 3, p.CreatedBy},
		A.X{`=`, 4, p.UpdatedAt},
		A.X{`=`, 5, p.UpdatedBy},
		A.X{`=`, 6, p.DeletedAt},
		A.X{`=`, 7, p.Title},
		A.X{`=`, 8, p.Description},
		A.X{`=`, 9, p.OrgId},
		A.X{`=`, 10, p.BudgetIDR},
		A.X{`=`, 11, p.BudgetUSD},
		A.X{`=`, 12, p.BudgetEUR},
	}
}

// IdxId return name of the index
func (p *Plans) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (p *Plans) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxPlanType return name of the index
func (p *Plans) IdxPlanType() int { //nolint:dupl false positive
	return 1
}

// SqlPlanType return name of the column being indexed
func (p *Plans) SqlPlanType() string { //nolint:dupl false positive
	return `"planType"`
}

// IdxCreatedAt return name of the index
func (p *Plans) IdxCreatedAt() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedAt return name of the column being indexed
func (p *Plans) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (p *Plans) IdxCreatedBy() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedBy return name of the column being indexed
func (p *Plans) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (p *Plans) IdxUpdatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlUpdatedAt return name of the column being indexed
func (p *Plans) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (p *Plans) IdxUpdatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedBy return name of the column being indexed
func (p *Plans) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (p *Plans) IdxDeletedAt() int { //nolint:dupl false positive
	return 6
}

// SqlDeletedAt return name of the column being indexed
func (p *Plans) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxTitle return name of the index
func (p *Plans) IdxTitle() int { //nolint:dupl false positive
	return 7
}

// SqlTitle return name of the column being indexed
func (p *Plans) SqlTitle() string { //nolint:dupl false positive
	return `"title"`
}

// IdxDescription return name of the index
func (p *Plans) IdxDescription() int { //nolint:dupl false positive
	return 8
}

// SqlDescription return name of the column being indexed
func (p *Plans) SqlDescription() string { //nolint:dupl false positive
	return `"description"`
}

// IdxOrgId return name of the index
func (p *Plans) IdxOrgId() int { //nolint:dupl false positive
	return 9
}

// SqlOrgId return name of the column being indexed
func (p *Plans) SqlOrgId() string { //nolint:dupl false positive
	return `"orgId"`
}

// IdxBudgetIDR return name of the index
func (p *Plans) IdxBudgetIDR() int { //nolint:dupl false positive
	return 10
}

// SqlBudgetIDR return name of the column being indexed
func (p *Plans) SqlBudgetIDR() string { //nolint:dupl false positive
	return `"budgetIDR"`
}

// IdxBudgetUSD return name of the index
func (p *Plans) IdxBudgetUSD() int { //nolint:dupl false positive
	return 11
}

// SqlBudgetUSD return name of the column being indexed
func (p *Plans) SqlBudgetUSD() string { //nolint:dupl false positive
	return `"budgetUSD"`
}

// IdxBudgetEUR return name of the index
func (p *Plans) IdxBudgetEUR() int { //nolint:dupl false positive
	return 12
}

// SqlBudgetEUR return name of the column being indexed
func (p *Plans) SqlBudgetEUR() string { //nolint:dupl false positive
	return `"budgetEUR"`
}

// ToArray receiver fields to slice
func (p *Plans) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.PlanType,    // 1
		p.CreatedAt,   // 2
		p.CreatedBy,   // 3
		p.UpdatedAt,   // 4
		p.UpdatedBy,   // 5
		p.DeletedAt,   // 6
		p.Title,       // 7
		p.Description, // 8
		p.OrgId,       // 9
		p.BudgetIDR,   // 10
		p.BudgetUSD,   // 11
		p.BudgetEUR,   // 12
	}
}

// FromArray convert slice to receiver fields
func (p *Plans) FromArray(a A.X) *Plans { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.PlanType = X.ToS(a[1])
	p.CreatedAt = X.ToI(a[2])
	p.CreatedBy = X.ToU(a[3])
	p.UpdatedAt = X.ToI(a[4])
	p.UpdatedBy = X.ToU(a[5])
	p.DeletedAt = X.ToI(a[6])
	p.Title = X.ToS(a[7])
	p.Description = X.ToS(a[8])
	p.OrgId = X.ToU(a[9])
	p.BudgetIDR = X.ToI(a[10])
	p.BudgetUSD = X.ToI(a[11])
	p.BudgetEUR = X.ToI(a[12])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *Plans) FromUncensoredArray(a A.X) *Plans { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.PlanType = X.ToS(a[1])
	p.CreatedAt = X.ToI(a[2])
	p.CreatedBy = X.ToU(a[3])
	p.UpdatedAt = X.ToI(a[4])
	p.UpdatedBy = X.ToU(a[5])
	p.DeletedAt = X.ToI(a[6])
	p.Title = X.ToS(a[7])
	p.Description = X.ToS(a[8])
	p.OrgId = X.ToU(a[9])
	p.BudgetIDR = X.ToI(a[10])
	p.BudgetUSD = X.ToI(a[11])
	p.BudgetEUR = X.ToI(a[12])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *Plans) FindOffsetLimit(offset, limit uint32, idx string) []Plans { //nolint:dupl false positive
	var rows []Plans
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Plans.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Plans{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *Plans) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Plans.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows, Tt.QueryMetaFrom(res, err)
	}
	tuples := res.Tuples()
	rows = make([]A.X, len(tuples))
	for z, row := range tuples {
		rows[z] = row
	}
	return rows, Tt.QueryMetaFrom(res, nil)
}

// Total count number of rows
func (p *Plans) Total() int64 { //nolint:dupl false positive
	rows := p.Adapter.CallBoxSpace(p.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// PlansFieldTypeMap returns key value of field name and key
var PlansFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:          Tt.Unsigned,
	`planType`:    Tt.String,
	`createdAt`:   Tt.Integer,
	`createdBy`:   Tt.Unsigned,
	`updatedAt`:   Tt.Integer,
	`updatedBy`:   Tt.Unsigned,
	`deletedAt`:   Tt.Integer,
	`title`:       Tt.String,
	`description`: Tt.String,
	`orgId`:       Tt.Unsigned,
	`budgetIDR`:   Tt.Integer,
	`budgetUSD`:   Tt.Integer,
	`budgetEUR`:   Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
