package rqBudget

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mBudget"

	"github.com/tarantool/go-tarantool/v2"

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
	TenantCode          string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
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
	IsFunder            bool        `json:"isFunder" form:"isFunder" query:"isFunder" long:"isFunder" msg:"isFunder"`
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
	res, err := b.Adapter.RetryDo(
		tarantool.NewSelectRequest(b.SpaceName()).
			Index(b.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(b.Id)}),
	)
	if L.IsError(err, `BankAccounts.FindById failed: `+b.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			b.FromArray(row)
			return true
		}
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (b *BankAccounts) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
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
	, "isFunder"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (b *BankAccounts) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
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
	, "isFunder"
	`
}

// ToUpdateArray generate slice of update command
func (b *BankAccounts) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, b.Id).
		Assign(1, b.TenantCode).
		Assign(2, b.CreatedAt).
		Assign(3, b.CreatedBy).
		Assign(4, b.UpdatedAt).
		Assign(5, b.UpdatedBy).
		Assign(6, b.DeletedAt).
		Assign(7, b.DeletedBy).
		Assign(8, b.RestoredBy).
		Assign(9, b.Name).
		Assign(10, b.ParentBankAccountId).
		Assign(11, b.ChildBankAccountId).
		Assign(12, b.AccountNumber).
		Assign(13, b.BankName).
		Assign(14, b.AccountName).
		Assign(15, b.IsProfitCenter).
		Assign(16, b.IsCostCenter).
		Assign(17, b.StaffId).
		Assign(18, b.IsFunder)
}

// IdxId return name of the index
func (b *BankAccounts) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (b *BankAccounts) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (b *BankAccounts) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (b *BankAccounts) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxCreatedAt return name of the index
func (b *BankAccounts) IdxCreatedAt() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedAt return name of the column being indexed
func (b *BankAccounts) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (b *BankAccounts) IdxCreatedBy() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedBy return name of the column being indexed
func (b *BankAccounts) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (b *BankAccounts) IdxUpdatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlUpdatedAt return name of the column being indexed
func (b *BankAccounts) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (b *BankAccounts) IdxUpdatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedBy return name of the column being indexed
func (b *BankAccounts) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (b *BankAccounts) IdxDeletedAt() int { //nolint:dupl false positive
	return 6
}

// SqlDeletedAt return name of the column being indexed
func (b *BankAccounts) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (b *BankAccounts) IdxDeletedBy() int { //nolint:dupl false positive
	return 7
}

// SqlDeletedBy return name of the column being indexed
func (b *BankAccounts) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (b *BankAccounts) IdxRestoredBy() int { //nolint:dupl false positive
	return 8
}

// SqlRestoredBy return name of the column being indexed
func (b *BankAccounts) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxName return name of the index
func (b *BankAccounts) IdxName() int { //nolint:dupl false positive
	return 9
}

// SqlName return name of the column being indexed
func (b *BankAccounts) SqlName() string { //nolint:dupl false positive
	return `"name"`
}

// IdxParentBankAccountId return name of the index
func (b *BankAccounts) IdxParentBankAccountId() int { //nolint:dupl false positive
	return 10
}

// SqlParentBankAccountId return name of the column being indexed
func (b *BankAccounts) SqlParentBankAccountId() string { //nolint:dupl false positive
	return `"parentBankAccountId"`
}

// IdxChildBankAccountId return name of the index
func (b *BankAccounts) IdxChildBankAccountId() int { //nolint:dupl false positive
	return 11
}

// SqlChildBankAccountId return name of the column being indexed
func (b *BankAccounts) SqlChildBankAccountId() string { //nolint:dupl false positive
	return `"childBankAccountId"`
}

// IdxAccountNumber return name of the index
func (b *BankAccounts) IdxAccountNumber() int { //nolint:dupl false positive
	return 12
}

// SqlAccountNumber return name of the column being indexed
func (b *BankAccounts) SqlAccountNumber() string { //nolint:dupl false positive
	return `"accountNumber"`
}

// IdxBankName return name of the index
func (b *BankAccounts) IdxBankName() int { //nolint:dupl false positive
	return 13
}

// SqlBankName return name of the column being indexed
func (b *BankAccounts) SqlBankName() string { //nolint:dupl false positive
	return `"bankName"`
}

// IdxAccountName return name of the index
func (b *BankAccounts) IdxAccountName() int { //nolint:dupl false positive
	return 14
}

// SqlAccountName return name of the column being indexed
func (b *BankAccounts) SqlAccountName() string { //nolint:dupl false positive
	return `"accountName"`
}

// IdxIsProfitCenter return name of the index
func (b *BankAccounts) IdxIsProfitCenter() int { //nolint:dupl false positive
	return 15
}

// SqlIsProfitCenter return name of the column being indexed
func (b *BankAccounts) SqlIsProfitCenter() string { //nolint:dupl false positive
	return `"isProfitCenter "`
}

// IdxIsCostCenter return name of the index
func (b *BankAccounts) IdxIsCostCenter() int { //nolint:dupl false positive
	return 16
}

// SqlIsCostCenter return name of the column being indexed
func (b *BankAccounts) SqlIsCostCenter() string { //nolint:dupl false positive
	return `"isCostCenter"`
}

// IdxStaffId return name of the index
func (b *BankAccounts) IdxStaffId() int { //nolint:dupl false positive
	return 17
}

// SqlStaffId return name of the column being indexed
func (b *BankAccounts) SqlStaffId() string { //nolint:dupl false positive
	return `"staffId"`
}

// IdxIsFunder return name of the index
func (b *BankAccounts) IdxIsFunder() int { //nolint:dupl false positive
	return 18
}

// SqlIsFunder return name of the column being indexed
func (b *BankAccounts) SqlIsFunder() string { //nolint:dupl false positive
	return `"isFunder"`
}

// ToArray receiver fields to slice
func (b *BankAccounts) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if b.Id != 0 {
		id = b.Id
	}
	return A.X{
		id,
		b.TenantCode,          // 1
		b.CreatedAt,           // 2
		b.CreatedBy,           // 3
		b.UpdatedAt,           // 4
		b.UpdatedBy,           // 5
		b.DeletedAt,           // 6
		b.DeletedBy,           // 7
		b.RestoredBy,          // 8
		b.Name,                // 9
		b.ParentBankAccountId, // 10
		b.ChildBankAccountId,  // 11
		b.AccountNumber,       // 12
		b.BankName,            // 13
		b.AccountName,         // 14
		b.IsProfitCenter,      // 15
		b.IsCostCenter,        // 16
		b.StaffId,             // 17
		b.IsFunder,            // 18
	}
}

// FromArray convert slice to receiver fields
func (b *BankAccounts) FromArray(a A.X) *BankAccounts { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.TenantCode = X.ToS(a[1])
	b.CreatedAt = X.ToI(a[2])
	b.CreatedBy = X.ToU(a[3])
	b.UpdatedAt = X.ToI(a[4])
	b.UpdatedBy = X.ToU(a[5])
	b.DeletedAt = X.ToI(a[6])
	b.DeletedBy = X.ToU(a[7])
	b.RestoredBy = X.ToU(a[8])
	b.Name = X.ToS(a[9])
	b.ParentBankAccountId = X.ToU(a[10])
	b.ChildBankAccountId = X.ToU(a[11])
	b.AccountNumber = X.ToI(a[12])
	b.BankName = X.ToS(a[13])
	b.AccountName = X.ToS(a[14])
	b.IsProfitCenter = X.ToBool(a[15])
	b.IsCostCenter = X.ToBool(a[16])
	b.StaffId = X.ToU(a[17])
	b.IsFunder = X.ToBool(a[18])
	return b
}

// FromUncensoredArray convert slice to receiver fields
func (b *BankAccounts) FromUncensoredArray(a A.X) *BankAccounts { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.TenantCode = X.ToS(a[1])
	b.CreatedAt = X.ToI(a[2])
	b.CreatedBy = X.ToU(a[3])
	b.UpdatedAt = X.ToI(a[4])
	b.UpdatedBy = X.ToU(a[5])
	b.DeletedAt = X.ToI(a[6])
	b.DeletedBy = X.ToU(a[7])
	b.RestoredBy = X.ToU(a[8])
	b.Name = X.ToS(a[9])
	b.ParentBankAccountId = X.ToU(a[10])
	b.ChildBankAccountId = X.ToU(a[11])
	b.AccountNumber = X.ToI(a[12])
	b.BankName = X.ToS(a[13])
	b.AccountName = X.ToS(a[14])
	b.IsProfitCenter = X.ToBool(a[15])
	b.IsCostCenter = X.ToBool(a[16])
	b.StaffId = X.ToU(a[17])
	b.IsFunder = X.ToBool(a[18])
	return b
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (b *BankAccounts) FindOffsetLimit(offset, limit uint32, idx string) []BankAccounts { //nolint:dupl false positive
	var rows []BankAccounts
	res, err := b.Adapter.RetryDo(
		tarantool.NewSelectRequest(b.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `BankAccounts.FindOffsetLimit failed: `+b.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := BankAccounts{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (b *BankAccounts) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := b.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(b.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `BankAccounts.FindOffsetLimit failed: `+b.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `BankAccounts.FindOffsetLimit failed: `+b.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	rows = make([]A.X, len(res))
	for _, row := range res {
		row, ok := row.([]any)
		if ok {
			rows = append(rows, row)
		}
	}
	return rows, Tt.QueryMetaFrom(resp, nil)
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
	`tenantCode`:          Tt.String,
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
	`isFunder`:            Tt.Boolean,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Plans DAO reader/query struct
type Plans struct {
	Adapter     *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id          uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantCode  string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
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
	YearOf      uint64      `json:"yearOf" form:"yearOf" query:"yearOf" long:"yearOf" msg:"yearOf"`
	BudgetIDR   uint64      `json:"budgetIDR,string" form:"budgetIDR" query:"budgetIDR" long:"budgetIDR" msg:"budgetIDR"`
	BudgetUSD   uint64      `json:"budgetUSD" form:"budgetUSD" query:"budgetUSD" long:"budgetUSD" msg:"budgetUSD"`
	Quantity    uint64      `json:"quantity" form:"quantity" query:"quantity" long:"quantity" msg:"quantity"`
	Unit        string      `json:"unit" form:"unit" query:"unit" long:"unit" msg:"unit"`
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
	res, err := p.Adapter.RetryDo(
		tarantool.NewSelectRequest(p.SpaceName()).
			Index(p.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(p.Id)}),
	)
	if L.IsError(err, `Plans.FindById failed: `+p.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			p.FromArray(row)
			return true
		}
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (p *Plans) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
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
	, "yearOf"
	, "budgetIDR"
	, "budgetUSD"
	, "quantity"
	, "unit"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *Plans) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
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
	, "yearOf"
	, "budgetIDR"
	, "budgetUSD"
	, "quantity"
	, "unit"
	`
}

// ToUpdateArray generate slice of update command
func (p *Plans) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, p.Id).
		Assign(1, p.TenantCode).
		Assign(2, p.PlanType).
		Assign(3, p.ParentId).
		Assign(4, p.CreatedAt).
		Assign(5, p.CreatedBy).
		Assign(6, p.UpdatedAt).
		Assign(7, p.UpdatedBy).
		Assign(8, p.DeletedAt).
		Assign(9, p.DeletedBy).
		Assign(10, p.RestoredBy).
		Assign(11, p.Title).
		Assign(12, p.Description).
		Assign(13, p.OrgId).
		Assign(14, p.YearOf).
		Assign(15, p.BudgetIDR).
		Assign(16, p.BudgetUSD).
		Assign(17, p.Quantity).
		Assign(18, p.Unit)
}

// IdxId return name of the index
func (p *Plans) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (p *Plans) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (p *Plans) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (p *Plans) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxPlanType return name of the index
func (p *Plans) IdxPlanType() int { //nolint:dupl false positive
	return 2
}

// SqlPlanType return name of the column being indexed
func (p *Plans) SqlPlanType() string { //nolint:dupl false positive
	return `"planType"`
}

// IdxParentId return name of the index
func (p *Plans) IdxParentId() int { //nolint:dupl false positive
	return 3
}

// SqlParentId return name of the column being indexed
func (p *Plans) SqlParentId() string { //nolint:dupl false positive
	return `"parentId"`
}

// IdxCreatedAt return name of the index
func (p *Plans) IdxCreatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlCreatedAt return name of the column being indexed
func (p *Plans) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (p *Plans) IdxCreatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlCreatedBy return name of the column being indexed
func (p *Plans) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (p *Plans) IdxUpdatedAt() int { //nolint:dupl false positive
	return 6
}

// SqlUpdatedAt return name of the column being indexed
func (p *Plans) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (p *Plans) IdxUpdatedBy() int { //nolint:dupl false positive
	return 7
}

// SqlUpdatedBy return name of the column being indexed
func (p *Plans) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (p *Plans) IdxDeletedAt() int { //nolint:dupl false positive
	return 8
}

// SqlDeletedAt return name of the column being indexed
func (p *Plans) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (p *Plans) IdxDeletedBy() int { //nolint:dupl false positive
	return 9
}

// SqlDeletedBy return name of the column being indexed
func (p *Plans) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (p *Plans) IdxRestoredBy() int { //nolint:dupl false positive
	return 10
}

// SqlRestoredBy return name of the column being indexed
func (p *Plans) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxTitle return name of the index
func (p *Plans) IdxTitle() int { //nolint:dupl false positive
	return 11
}

// SqlTitle return name of the column being indexed
func (p *Plans) SqlTitle() string { //nolint:dupl false positive
	return `"title"`
}

// IdxDescription return name of the index
func (p *Plans) IdxDescription() int { //nolint:dupl false positive
	return 12
}

// SqlDescription return name of the column being indexed
func (p *Plans) SqlDescription() string { //nolint:dupl false positive
	return `"description"`
}

// IdxOrgId return name of the index
func (p *Plans) IdxOrgId() int { //nolint:dupl false positive
	return 13
}

// SqlOrgId return name of the column being indexed
func (p *Plans) SqlOrgId() string { //nolint:dupl false positive
	return `"orgId"`
}

// IdxYearOf return name of the index
func (p *Plans) IdxYearOf() int { //nolint:dupl false positive
	return 14
}

// SqlYearOf return name of the column being indexed
func (p *Plans) SqlYearOf() string { //nolint:dupl false positive
	return `"yearOf"`
}

// IdxBudgetIDR return name of the index
func (p *Plans) IdxBudgetIDR() int { //nolint:dupl false positive
	return 15
}

// SqlBudgetIDR return name of the column being indexed
func (p *Plans) SqlBudgetIDR() string { //nolint:dupl false positive
	return `"budgetIDR"`
}

// IdxBudgetUSD return name of the index
func (p *Plans) IdxBudgetUSD() int { //nolint:dupl false positive
	return 16
}

// SqlBudgetUSD return name of the column being indexed
func (p *Plans) SqlBudgetUSD() string { //nolint:dupl false positive
	return `"budgetUSD"`
}

// IdxQuantity return name of the index
func (p *Plans) IdxQuantity() int { //nolint:dupl false positive
	return 17
}

// SqlQuantity return name of the column being indexed
func (p *Plans) SqlQuantity() string { //nolint:dupl false positive
	return `"quantity"`
}

// IdxUnit return name of the index
func (p *Plans) IdxUnit() int { //nolint:dupl false positive
	return 18
}

// SqlUnit return name of the column being indexed
func (p *Plans) SqlUnit() string { //nolint:dupl false positive
	return `"unit"`
}

// ToArray receiver fields to slice
func (p *Plans) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.TenantCode,  // 1
		p.PlanType,    // 2
		p.ParentId,    // 3
		p.CreatedAt,   // 4
		p.CreatedBy,   // 5
		p.UpdatedAt,   // 6
		p.UpdatedBy,   // 7
		p.DeletedAt,   // 8
		p.DeletedBy,   // 9
		p.RestoredBy,  // 10
		p.Title,       // 11
		p.Description, // 12
		p.OrgId,       // 13
		p.YearOf,      // 14
		p.BudgetIDR,   // 15
		p.BudgetUSD,   // 16
		p.Quantity,    // 17
		p.Unit,        // 18
	}
}

// FromArray convert slice to receiver fields
func (p *Plans) FromArray(a A.X) *Plans { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.TenantCode = X.ToS(a[1])
	p.PlanType = X.ToS(a[2])
	p.ParentId = X.ToU(a[3])
	p.CreatedAt = X.ToI(a[4])
	p.CreatedBy = X.ToU(a[5])
	p.UpdatedAt = X.ToI(a[6])
	p.UpdatedBy = X.ToU(a[7])
	p.DeletedAt = X.ToI(a[8])
	p.DeletedBy = X.ToU(a[9])
	p.RestoredBy = X.ToU(a[10])
	p.Title = X.ToS(a[11])
	p.Description = X.ToS(a[12])
	p.OrgId = X.ToU(a[13])
	p.YearOf = X.ToU(a[14])
	p.BudgetIDR = X.ToU(a[15])
	p.BudgetUSD = X.ToU(a[16])
	p.Quantity = X.ToU(a[17])
	p.Unit = X.ToS(a[18])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *Plans) FromUncensoredArray(a A.X) *Plans { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.TenantCode = X.ToS(a[1])
	p.PlanType = X.ToS(a[2])
	p.ParentId = X.ToU(a[3])
	p.CreatedAt = X.ToI(a[4])
	p.CreatedBy = X.ToU(a[5])
	p.UpdatedAt = X.ToI(a[6])
	p.UpdatedBy = X.ToU(a[7])
	p.DeletedAt = X.ToI(a[8])
	p.DeletedBy = X.ToU(a[9])
	p.RestoredBy = X.ToU(a[10])
	p.Title = X.ToS(a[11])
	p.Description = X.ToS(a[12])
	p.OrgId = X.ToU(a[13])
	p.YearOf = X.ToU(a[14])
	p.BudgetIDR = X.ToU(a[15])
	p.BudgetUSD = X.ToU(a[16])
	p.Quantity = X.ToU(a[17])
	p.Unit = X.ToS(a[18])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *Plans) FindOffsetLimit(offset, limit uint32, idx string) []Plans { //nolint:dupl false positive
	var rows []Plans
	res, err := p.Adapter.RetryDo(
		tarantool.NewSelectRequest(p.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Plans.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := Plans{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *Plans) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := p.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(p.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Plans.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `Plans.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	rows = make([]A.X, len(res))
	for _, row := range res {
		row, ok := row.([]any)
		if ok {
			rows = append(rows, row)
		}
	}
	return rows, Tt.QueryMetaFrom(resp, nil)
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
	`tenantCode`:  Tt.String,
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
	`yearOf`:      Tt.Unsigned,
	`budgetIDR`:   Tt.Unsigned,
	`budgetUSD`:   Tt.Unsigned,
	`quantity`:    Tt.Unsigned,
	`unit`:        Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
