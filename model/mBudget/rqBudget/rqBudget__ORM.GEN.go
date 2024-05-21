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

// BankAccounts DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqBudget__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqBudget__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqBudget__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqBudget__ORM.GEN.go
type BankAccounts struct {
	Adapter             *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id                  uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	CreatedAt           int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy           uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt           int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy           uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt           int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy           uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy          uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	Name                string      `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	ParentBankAccountId uint64      `json:"parentBankAccountId,string" form:"parentBankAccountId" query:"parentBankAccountId" long:"parentBankAccountId" msg:"parentBankAccountId"`
	ChildBankAccountId  uint64      `json:"childBankAccountId,string" form:"childBankAccountId" query:"childBankAccountId" long:"childBankAccountId" msg:"childBankAccountId"`
	AccountNumber       int64       `json:"accountNumber" form:"accountNumber" query:"accountNumber" long:"accountNumber" msg:"accountNumber"`
	BankName            string      `json:"bankName" form:"bankName" query:"bankName" long:"bankName" msg:"bankName"`
	AccountName         string      `json:"accountName" form:"accountName" query:"accountName" long:"accountName" msg:"accountName"`
	IsProfitCenter      bool        `json:"isProfitCenter" form:"isProfitCenter" query:"isProfitCenter" long:"isProfitCenter" msg:"isProfitCenter"`
	IsCostCenter        bool        `json:"isCostCenter" form:"isCostCenter" query:"isCostCenter" long:"isCostCenter" msg:"isCostCenter"`
	StaffId             uint64      `json:"staffId,string" form:"staffId" query:"staffId" long:"staffId" msg:"staffId"`
}

// NewBankAccounts create new ORM reader/query object
func NewBankAccounts(adapter *Tt.Adapter) *BankAccounts {
	return &BankAccounts{Adapter: adapter}
}

// SpaceName returns full package and table name
func (b *BankAccounts) SpaceName() string { //nolint:dupl false positive
	return string(mBudget.TableBankAccounts) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (b *BankAccounts) SqlTableName() string { //nolint:dupl false positive
	return `"bankAccounts"`
}

func (b *BankAccounts) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (b *BankAccounts) FindById() bool { //nolint:dupl false positive
	res, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{b.Id})
	if L.IsError(err, `BankAccounts.FindById failed: `+b.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		b.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (b *BankAccounts) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "name"
	, "parentBankAccountId"
	, "childBankAccountId"
	, "accountNumber"
	, "bankName"
	, "accountName"
	, "isProfitCenter "
	, "isCostCenter"
	, "staffId"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (b *BankAccounts) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "name"
	, "parentBankAccountId"
	, "childBankAccountId"
	, "accountNumber"
	, "bankName"
	, "accountName"
	, "isProfitCenter "
	, "isCostCenter"
	, "staffId"
	`
}

// ToUpdateArray generate slice of update command
func (b *BankAccounts) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, b.Id},
		A.X{`=`, 1, b.CreatedAt},
		A.X{`=`, 2, b.CreatedBy},
		A.X{`=`, 3, b.UpdatedAt},
		A.X{`=`, 4, b.UpdatedBy},
		A.X{`=`, 5, b.DeletedAt},
		A.X{`=`, 6, b.DeletedBy},
		A.X{`=`, 7, b.RestoredBy},
		A.X{`=`, 8, b.Name},
		A.X{`=`, 9, b.ParentBankAccountId},
		A.X{`=`, 10, b.ChildBankAccountId},
		A.X{`=`, 11, b.AccountNumber},
		A.X{`=`, 12, b.BankName},
		A.X{`=`, 13, b.AccountName},
		A.X{`=`, 14, b.IsProfitCenter},
		A.X{`=`, 15, b.IsCostCenter},
		A.X{`=`, 16, b.StaffId},
	}
}

// IdxId return name of the index
func (b *BankAccounts) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (b *BankAccounts) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxCreatedAt return name of the index
func (b *BankAccounts) IdxCreatedAt() int { //nolint:dupl false positive
	return 1
}

// SqlCreatedAt return name of the column being indexed
func (b *BankAccounts) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (b *BankAccounts) IdxCreatedBy() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedBy return name of the column being indexed
func (b *BankAccounts) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (b *BankAccounts) IdxUpdatedAt() int { //nolint:dupl false positive
	return 3
}

// SqlUpdatedAt return name of the column being indexed
func (b *BankAccounts) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (b *BankAccounts) IdxUpdatedBy() int { //nolint:dupl false positive
	return 4
}

// SqlUpdatedBy return name of the column being indexed
func (b *BankAccounts) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (b *BankAccounts) IdxDeletedAt() int { //nolint:dupl false positive
	return 5
}

// SqlDeletedAt return name of the column being indexed
func (b *BankAccounts) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (b *BankAccounts) IdxDeletedBy() int { //nolint:dupl false positive
	return 6
}

// SqlDeletedBy return name of the column being indexed
func (b *BankAccounts) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (b *BankAccounts) IdxRestoredBy() int { //nolint:dupl false positive
	return 7
}

// SqlRestoredBy return name of the column being indexed
func (b *BankAccounts) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxName return name of the index
func (b *BankAccounts) IdxName() int { //nolint:dupl false positive
	return 8
}

// SqlName return name of the column being indexed
func (b *BankAccounts) SqlName() string { //nolint:dupl false positive
	return `"name"`
}

// IdxParentBankAccountId return name of the index
func (b *BankAccounts) IdxParentBankAccountId() int { //nolint:dupl false positive
	return 9
}

// SqlParentBankAccountId return name of the column being indexed
func (b *BankAccounts) SqlParentBankAccountId() string { //nolint:dupl false positive
	return `"parentBankAccountId"`
}

// IdxChildBankAccountId return name of the index
func (b *BankAccounts) IdxChildBankAccountId() int { //nolint:dupl false positive
	return 10
}

// SqlChildBankAccountId return name of the column being indexed
func (b *BankAccounts) SqlChildBankAccountId() string { //nolint:dupl false positive
	return `"childBankAccountId"`
}

// IdxAccountNumber return name of the index
func (b *BankAccounts) IdxAccountNumber() int { //nolint:dupl false positive
	return 11
}

// SqlAccountNumber return name of the column being indexed
func (b *BankAccounts) SqlAccountNumber() string { //nolint:dupl false positive
	return `"accountNumber"`
}

// IdxBankName return name of the index
func (b *BankAccounts) IdxBankName() int { //nolint:dupl false positive
	return 12
}

// SqlBankName return name of the column being indexed
func (b *BankAccounts) SqlBankName() string { //nolint:dupl false positive
	return `"bankName"`
}

// IdxAccountName return name of the index
func (b *BankAccounts) IdxAccountName() int { //nolint:dupl false positive
	return 13
}

// SqlAccountName return name of the column being indexed
func (b *BankAccounts) SqlAccountName() string { //nolint:dupl false positive
	return `"accountName"`
}

// IdxIsProfitCenter return name of the index
func (b *BankAccounts) IdxIsProfitCenter() int { //nolint:dupl false positive
	return 14
}

// SqlIsProfitCenter return name of the column being indexed
func (b *BankAccounts) SqlIsProfitCenter() string { //nolint:dupl false positive
	return `"isProfitCenter "`
}

// IdxIsCostCenter return name of the index
func (b *BankAccounts) IdxIsCostCenter() int { //nolint:dupl false positive
	return 15
}

// SqlIsCostCenter return name of the column being indexed
func (b *BankAccounts) SqlIsCostCenter() string { //nolint:dupl false positive
	return `"isCostCenter"`
}

// IdxStaffId return name of the index
func (b *BankAccounts) IdxStaffId() int { //nolint:dupl false positive
	return 16
}

// SqlStaffId return name of the column being indexed
func (b *BankAccounts) SqlStaffId() string { //nolint:dupl false positive
	return `"staffId"`
}

// ToArray receiver fields to slice
func (b *BankAccounts) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if b.Id != 0 {
		id = b.Id
	}
	return A.X{
		id,
		b.CreatedAt,           // 1
		b.CreatedBy,           // 2
		b.UpdatedAt,           // 3
		b.UpdatedBy,           // 4
		b.DeletedAt,           // 5
		b.DeletedBy,           // 6
		b.RestoredBy,          // 7
		b.Name,                // 8
		b.ParentBankAccountId, // 9
		b.ChildBankAccountId,  // 10
		b.AccountNumber,       // 11
		b.BankName,            // 12
		b.AccountName,         // 13
		b.IsProfitCenter,      // 14
		b.IsCostCenter,        // 15
		b.StaffId,             // 16
	}
}

// FromArray convert slice to receiver fields
func (b *BankAccounts) FromArray(a A.X) *BankAccounts { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.CreatedAt = X.ToI(a[1])
	b.CreatedBy = X.ToU(a[2])
	b.UpdatedAt = X.ToI(a[3])
	b.UpdatedBy = X.ToU(a[4])
	b.DeletedAt = X.ToI(a[5])
	b.DeletedBy = X.ToU(a[6])
	b.RestoredBy = X.ToU(a[7])
	b.Name = X.ToS(a[8])
	b.ParentBankAccountId = X.ToU(a[9])
	b.ChildBankAccountId = X.ToU(a[10])
	b.AccountNumber = X.ToI(a[11])
	b.BankName = X.ToS(a[12])
	b.AccountName = X.ToS(a[13])
	b.IsProfitCenter = X.ToBool(a[14])
	b.IsCostCenter = X.ToBool(a[15])
	b.StaffId = X.ToU(a[16])
	return b
}

// FromUncensoredArray convert slice to receiver fields
func (b *BankAccounts) FromUncensoredArray(a A.X) *BankAccounts { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.CreatedAt = X.ToI(a[1])
	b.CreatedBy = X.ToU(a[2])
	b.UpdatedAt = X.ToI(a[3])
	b.UpdatedBy = X.ToU(a[4])
	b.DeletedAt = X.ToI(a[5])
	b.DeletedBy = X.ToU(a[6])
	b.RestoredBy = X.ToU(a[7])
	b.Name = X.ToS(a[8])
	b.ParentBankAccountId = X.ToU(a[9])
	b.ChildBankAccountId = X.ToU(a[10])
	b.AccountNumber = X.ToI(a[11])
	b.BankName = X.ToS(a[12])
	b.AccountName = X.ToS(a[13])
	b.IsProfitCenter = X.ToBool(a[14])
	b.IsCostCenter = X.ToBool(a[15])
	b.StaffId = X.ToU(a[16])
	return b
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (b *BankAccounts) FindOffsetLimit(offset, limit uint32, idx string) []BankAccounts { //nolint:dupl false positive
	var rows []BankAccounts
	res, err := b.Adapter.Select(b.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `BankAccounts.FindOffsetLimit failed: `+b.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := BankAccounts{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (b *BankAccounts) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := b.Adapter.Select(b.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `BankAccounts.FindOffsetLimit failed: `+b.SpaceName()) {
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
func (b *BankAccounts) Total() int64 { //nolint:dupl false positive
	rows := b.Adapter.CallBoxSpace(b.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// BankAccountsFieldTypeMap returns key value of field name and key
var BankAccountsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:                  Tt.Unsigned,
	`createdAt`:           Tt.Integer,
	`createdBy`:           Tt.Unsigned,
	`updatedAt`:           Tt.Integer,
	`updatedBy`:           Tt.Unsigned,
	`deletedAt`:           Tt.Integer,
	`deletedBy`:           Tt.Unsigned,
	`restoredBy`:          Tt.Unsigned,
	`name`:                Tt.String,
	`parentBankAccountId`: Tt.Unsigned,
	`childBankAccountId`:  Tt.Unsigned,
	`accountNumber`:       Tt.Integer,
	`bankName`:            Tt.String,
	`accountName`:         Tt.String,
	`isProfitCenter `:     Tt.Boolean,
	`isCostCenter`:        Tt.Boolean,
	`staffId`:             Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Plans DAO reader/query struct
type Plans struct {
	Adapter     *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id          uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	PlanType    string      `json:"planType" form:"planType" query:"planType" long:"planType" msg:"planType"`
	ParentId    uint64      `json:"parentId,string" form:"parentId" query:"parentId" long:"parentId" msg:"parentId"`
	CreatedAt   int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy   uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt   int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy   uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt   int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy   uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy  uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	Title       string      `json:"title" form:"title" query:"title" long:"title" msg:"title"`
	Description string      `json:"description" form:"description" query:"description" long:"description" msg:"description"`
	OrgId       uint64      `json:"orgId,string" form:"orgId" query:"orgId" long:"orgId" msg:"orgId"`
	PerYear     int64       `json:"perYear" form:"perYear" query:"perYear" long:"perYear" msg:"perYear"`
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
	, "parentId"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "title"
	, "description"
	, "orgId"
	, "perYear"
	, "budgetIDR"
	, "budgetUSD"
	, "budgetEUR"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *Plans) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "planType"
	, "parentId"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "title"
	, "description"
	, "orgId"
	, "perYear"
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
		A.X{`=`, 2, p.ParentId},
		A.X{`=`, 3, p.CreatedAt},
		A.X{`=`, 4, p.CreatedBy},
		A.X{`=`, 5, p.UpdatedAt},
		A.X{`=`, 6, p.UpdatedBy},
		A.X{`=`, 7, p.DeletedAt},
		A.X{`=`, 8, p.DeletedBy},
		A.X{`=`, 9, p.RestoredBy},
		A.X{`=`, 10, p.Title},
		A.X{`=`, 11, p.Description},
		A.X{`=`, 12, p.OrgId},
		A.X{`=`, 13, p.PerYear},
		A.X{`=`, 14, p.BudgetIDR},
		A.X{`=`, 15, p.BudgetUSD},
		A.X{`=`, 16, p.BudgetEUR},
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

// IdxParentId return name of the index
func (p *Plans) IdxParentId() int { //nolint:dupl false positive
	return 2
}

// SqlParentId return name of the column being indexed
func (p *Plans) SqlParentId() string { //nolint:dupl false positive
	return `"parentId"`
}

// IdxCreatedAt return name of the index
func (p *Plans) IdxCreatedAt() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedAt return name of the column being indexed
func (p *Plans) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (p *Plans) IdxCreatedBy() int { //nolint:dupl false positive
	return 4
}

// SqlCreatedBy return name of the column being indexed
func (p *Plans) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (p *Plans) IdxUpdatedAt() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedAt return name of the column being indexed
func (p *Plans) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (p *Plans) IdxUpdatedBy() int { //nolint:dupl false positive
	return 6
}

// SqlUpdatedBy return name of the column being indexed
func (p *Plans) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (p *Plans) IdxDeletedAt() int { //nolint:dupl false positive
	return 7
}

// SqlDeletedAt return name of the column being indexed
func (p *Plans) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (p *Plans) IdxDeletedBy() int { //nolint:dupl false positive
	return 8
}

// SqlDeletedBy return name of the column being indexed
func (p *Plans) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (p *Plans) IdxRestoredBy() int { //nolint:dupl false positive
	return 9
}

// SqlRestoredBy return name of the column being indexed
func (p *Plans) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxTitle return name of the index
func (p *Plans) IdxTitle() int { //nolint:dupl false positive
	return 10
}

// SqlTitle return name of the column being indexed
func (p *Plans) SqlTitle() string { //nolint:dupl false positive
	return `"title"`
}

// IdxDescription return name of the index
func (p *Plans) IdxDescription() int { //nolint:dupl false positive
	return 11
}

// SqlDescription return name of the column being indexed
func (p *Plans) SqlDescription() string { //nolint:dupl false positive
	return `"description"`
}

// IdxOrgId return name of the index
func (p *Plans) IdxOrgId() int { //nolint:dupl false positive
	return 12
}

// SqlOrgId return name of the column being indexed
func (p *Plans) SqlOrgId() string { //nolint:dupl false positive
	return `"orgId"`
}

// IdxPerYear return name of the index
func (p *Plans) IdxPerYear() int { //nolint:dupl false positive
	return 13
}

// SqlPerYear return name of the column being indexed
func (p *Plans) SqlPerYear() string { //nolint:dupl false positive
	return `"perYear"`
}

// IdxBudgetIDR return name of the index
func (p *Plans) IdxBudgetIDR() int { //nolint:dupl false positive
	return 14
}

// SqlBudgetIDR return name of the column being indexed
func (p *Plans) SqlBudgetIDR() string { //nolint:dupl false positive
	return `"budgetIDR"`
}

// IdxBudgetUSD return name of the index
func (p *Plans) IdxBudgetUSD() int { //nolint:dupl false positive
	return 15
}

// SqlBudgetUSD return name of the column being indexed
func (p *Plans) SqlBudgetUSD() string { //nolint:dupl false positive
	return `"budgetUSD"`
}

// IdxBudgetEUR return name of the index
func (p *Plans) IdxBudgetEUR() int { //nolint:dupl false positive
	return 16
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
		p.ParentId,    // 2
		p.CreatedAt,   // 3
		p.CreatedBy,   // 4
		p.UpdatedAt,   // 5
		p.UpdatedBy,   // 6
		p.DeletedAt,   // 7
		p.DeletedBy,   // 8
		p.RestoredBy,  // 9
		p.Title,       // 10
		p.Description, // 11
		p.OrgId,       // 12
		p.PerYear,     // 13
		p.BudgetIDR,   // 14
		p.BudgetUSD,   // 15
		p.BudgetEUR,   // 16
	}
}

// FromArray convert slice to receiver fields
func (p *Plans) FromArray(a A.X) *Plans { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.PlanType = X.ToS(a[1])
	p.ParentId = X.ToU(a[2])
	p.CreatedAt = X.ToI(a[3])
	p.CreatedBy = X.ToU(a[4])
	p.UpdatedAt = X.ToI(a[5])
	p.UpdatedBy = X.ToU(a[6])
	p.DeletedAt = X.ToI(a[7])
	p.DeletedBy = X.ToU(a[8])
	p.RestoredBy = X.ToU(a[9])
	p.Title = X.ToS(a[10])
	p.Description = X.ToS(a[11])
	p.OrgId = X.ToU(a[12])
	p.PerYear = X.ToI(a[13])
	p.BudgetIDR = X.ToI(a[14])
	p.BudgetUSD = X.ToI(a[15])
	p.BudgetEUR = X.ToI(a[16])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *Plans) FromUncensoredArray(a A.X) *Plans { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.PlanType = X.ToS(a[1])
	p.ParentId = X.ToU(a[2])
	p.CreatedAt = X.ToI(a[3])
	p.CreatedBy = X.ToU(a[4])
	p.UpdatedAt = X.ToI(a[5])
	p.UpdatedBy = X.ToU(a[6])
	p.DeletedAt = X.ToI(a[7])
	p.DeletedBy = X.ToU(a[8])
	p.RestoredBy = X.ToU(a[9])
	p.Title = X.ToS(a[10])
	p.Description = X.ToS(a[11])
	p.OrgId = X.ToU(a[12])
	p.PerYear = X.ToI(a[13])
	p.BudgetIDR = X.ToI(a[14])
	p.BudgetUSD = X.ToI(a[15])
	p.BudgetEUR = X.ToI(a[16])
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
	`parentId`:    Tt.Unsigned,
	`createdAt`:   Tt.Integer,
	`createdBy`:   Tt.Unsigned,
	`updatedAt`:   Tt.Integer,
	`updatedBy`:   Tt.Unsigned,
	`deletedAt`:   Tt.Integer,
	`deletedBy`:   Tt.Unsigned,
	`restoredBy`:  Tt.Unsigned,
	`title`:       Tt.String,
	`description`: Tt.String,
	`orgId`:       Tt.Unsigned,
	`perYear`:     Tt.Integer,
	`budgetIDR`:   Tt.Integer,
	`budgetUSD`:   Tt.Integer,
	`budgetEUR`:   Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
