package rqFinance

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mFinance"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// BusinessTransaction DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqFinance__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqFinance__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqFinance__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqFinance__ORM.GEN.go
type BusinessTransaction struct {
	Adapter               *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id                    uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantCode            string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	StartDate             string      `json:"startDate" form:"startDate" query:"startDate" long:"startDate" msg:"startDate"`
	EndDate               string      `json:"endDate" form:"endDate" query:"endDate" long:"endDate" msg:"endDate"`
	CreatedAt             int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy             uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt             int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy             uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt             int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy             uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy            uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	TransactionTemplateId uint64      `json:"transactionTemplateId,string" form:"transactionTemplateId" query:"transactionTemplateId" long:"transactionTemplateId" msg:"transactionTemplateId"`
}

// NewBusinessTransaction create new ORM reader/query object
func NewBusinessTransaction(adapter *Tt.Adapter) *BusinessTransaction {
	return &BusinessTransaction{Adapter: adapter}
}

// SpaceName returns full package and table name
func (b *BusinessTransaction) SpaceName() string { //nolint:dupl false positive
	return string(mFinance.TableBusinessTransaction) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (b *BusinessTransaction) SqlTableName() string { //nolint:dupl false positive
	return `"businessTransaction"`
}

func (b *BusinessTransaction) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (b *BusinessTransaction) FindById() bool { //nolint:dupl false positive
	res, err := b.Adapter.Select(b.SpaceName(), b.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{b.Id})
	if L.IsError(err, `BusinessTransaction.FindById failed: `+b.SpaceName()) {
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
func (b *BusinessTransaction) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "startDate"
	, "endDate"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "transactionTemplateId"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (b *BusinessTransaction) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "startDate"
	, "endDate"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "transactionTemplateId"
	`
}

// ToUpdateArray generate slice of update command
func (b *BusinessTransaction) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, b.Id},
		A.X{`=`, 1, b.TenantCode},
		A.X{`=`, 2, b.StartDate},
		A.X{`=`, 3, b.EndDate},
		A.X{`=`, 4, b.CreatedAt},
		A.X{`=`, 5, b.CreatedBy},
		A.X{`=`, 6, b.UpdatedAt},
		A.X{`=`, 7, b.UpdatedBy},
		A.X{`=`, 8, b.DeletedAt},
		A.X{`=`, 9, b.DeletedBy},
		A.X{`=`, 10, b.RestoredBy},
		A.X{`=`, 11, b.TransactionTemplateId},
	}
}

// IdxId return name of the index
func (b *BusinessTransaction) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (b *BusinessTransaction) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (b *BusinessTransaction) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (b *BusinessTransaction) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxStartDate return name of the index
func (b *BusinessTransaction) IdxStartDate() int { //nolint:dupl false positive
	return 2
}

// SqlStartDate return name of the column being indexed
func (b *BusinessTransaction) SqlStartDate() string { //nolint:dupl false positive
	return `"startDate"`
}

// IdxEndDate return name of the index
func (b *BusinessTransaction) IdxEndDate() int { //nolint:dupl false positive
	return 3
}

// SqlEndDate return name of the column being indexed
func (b *BusinessTransaction) SqlEndDate() string { //nolint:dupl false positive
	return `"endDate"`
}

// IdxCreatedAt return name of the index
func (b *BusinessTransaction) IdxCreatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlCreatedAt return name of the column being indexed
func (b *BusinessTransaction) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (b *BusinessTransaction) IdxCreatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlCreatedBy return name of the column being indexed
func (b *BusinessTransaction) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (b *BusinessTransaction) IdxUpdatedAt() int { //nolint:dupl false positive
	return 6
}

// SqlUpdatedAt return name of the column being indexed
func (b *BusinessTransaction) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (b *BusinessTransaction) IdxUpdatedBy() int { //nolint:dupl false positive
	return 7
}

// SqlUpdatedBy return name of the column being indexed
func (b *BusinessTransaction) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (b *BusinessTransaction) IdxDeletedAt() int { //nolint:dupl false positive
	return 8
}

// SqlDeletedAt return name of the column being indexed
func (b *BusinessTransaction) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (b *BusinessTransaction) IdxDeletedBy() int { //nolint:dupl false positive
	return 9
}

// SqlDeletedBy return name of the column being indexed
func (b *BusinessTransaction) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (b *BusinessTransaction) IdxRestoredBy() int { //nolint:dupl false positive
	return 10
}

// SqlRestoredBy return name of the column being indexed
func (b *BusinessTransaction) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxTransactionTemplateId return name of the index
func (b *BusinessTransaction) IdxTransactionTemplateId() int { //nolint:dupl false positive
	return 11
}

// SqlTransactionTemplateId return name of the column being indexed
func (b *BusinessTransaction) SqlTransactionTemplateId() string { //nolint:dupl false positive
	return `"transactionTemplateId"`
}

// ToArray receiver fields to slice
func (b *BusinessTransaction) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if b.Id != 0 {
		id = b.Id
	}
	return A.X{
		id,
		b.TenantCode,            // 1
		b.StartDate,             // 2
		b.EndDate,               // 3
		b.CreatedAt,             // 4
		b.CreatedBy,             // 5
		b.UpdatedAt,             // 6
		b.UpdatedBy,             // 7
		b.DeletedAt,             // 8
		b.DeletedBy,             // 9
		b.RestoredBy,            // 10
		b.TransactionTemplateId, // 11
	}
}

// FromArray convert slice to receiver fields
func (b *BusinessTransaction) FromArray(a A.X) *BusinessTransaction { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.TenantCode = X.ToS(a[1])
	b.StartDate = X.ToS(a[2])
	b.EndDate = X.ToS(a[3])
	b.CreatedAt = X.ToI(a[4])
	b.CreatedBy = X.ToU(a[5])
	b.UpdatedAt = X.ToI(a[6])
	b.UpdatedBy = X.ToU(a[7])
	b.DeletedAt = X.ToI(a[8])
	b.DeletedBy = X.ToU(a[9])
	b.RestoredBy = X.ToU(a[10])
	b.TransactionTemplateId = X.ToU(a[11])
	return b
}

// FromUncensoredArray convert slice to receiver fields
func (b *BusinessTransaction) FromUncensoredArray(a A.X) *BusinessTransaction { //nolint:dupl false positive
	b.Id = X.ToU(a[0])
	b.TenantCode = X.ToS(a[1])
	b.StartDate = X.ToS(a[2])
	b.EndDate = X.ToS(a[3])
	b.CreatedAt = X.ToI(a[4])
	b.CreatedBy = X.ToU(a[5])
	b.UpdatedAt = X.ToI(a[6])
	b.UpdatedBy = X.ToU(a[7])
	b.DeletedAt = X.ToI(a[8])
	b.DeletedBy = X.ToU(a[9])
	b.RestoredBy = X.ToU(a[10])
	b.TransactionTemplateId = X.ToU(a[11])
	return b
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (b *BusinessTransaction) FindOffsetLimit(offset, limit uint32, idx string) []BusinessTransaction { //nolint:dupl false positive
	var rows []BusinessTransaction
	res, err := b.Adapter.Select(b.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `BusinessTransaction.FindOffsetLimit failed: `+b.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := BusinessTransaction{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (b *BusinessTransaction) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := b.Adapter.Select(b.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `BusinessTransaction.FindOffsetLimit failed: `+b.SpaceName()) {
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
func (b *BusinessTransaction) Total() int64 { //nolint:dupl false positive
	rows := b.Adapter.CallBoxSpace(b.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// BusinessTransactionFieldTypeMap returns key value of field name and key
var BusinessTransactionFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:                    Tt.Unsigned,
	`tenantCode`:            Tt.String,
	`startDate`:             Tt.String,
	`endDate`:               Tt.String,
	`createdAt`:             Tt.Integer,
	`createdBy`:             Tt.Unsigned,
	`updatedAt`:             Tt.Integer,
	`updatedBy`:             Tt.Unsigned,
	`deletedAt`:             Tt.Integer,
	`deletedBy`:             Tt.Unsigned,
	`restoredBy`:            Tt.Unsigned,
	`transactionTemplateId`: Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Coa DAO reader/query struct
type Coa struct {
	Adapter    *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id         uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantCode string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	Name       string      `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	Label      string      `json:"label" form:"label" query:"label" long:"label" msg:"label"`
	ParentId   uint64      `json:"parentId,string" form:"parentId" query:"parentId" long:"parentId" msg:"parentId"`
	Children   []any       `json:"children" form:"children" query:"children" long:"children" msg:"children"`
	CreatedAt  int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy  uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt  int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy  uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt  int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy  uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	Editable   bool        `json:"editable" form:"editable" query:"editable" long:"editable" msg:"editable"`
}

// NewCoa create new ORM reader/query object
func NewCoa(adapter *Tt.Adapter) *Coa {
	return &Coa{Adapter: adapter}
}

// SpaceName returns full package and table name
func (c *Coa) SpaceName() string { //nolint:dupl false positive
	return string(mFinance.TableCoa) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (c *Coa) SqlTableName() string { //nolint:dupl false positive
	return `"coa"`
}

func (c *Coa) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (c *Coa) FindById() bool { //nolint:dupl false positive
	res, err := c.Adapter.Select(c.SpaceName(), c.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{c.Id})
	if L.IsError(err, `Coa.FindById failed: `+c.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		c.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (c *Coa) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "name"
	, "label"
	, "parentId"
	, "children"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "editable"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (c *Coa) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "name"
	, "label"
	, "parentId"
	, "children"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "editable"
	`
}

// ToUpdateArray generate slice of update command
func (c *Coa) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, c.Id},
		A.X{`=`, 1, c.TenantCode},
		A.X{`=`, 2, c.Name},
		A.X{`=`, 3, c.Label},
		A.X{`=`, 4, c.ParentId},
		A.X{`=`, 5, c.Children},
		A.X{`=`, 6, c.CreatedAt},
		A.X{`=`, 7, c.CreatedBy},
		A.X{`=`, 8, c.UpdatedAt},
		A.X{`=`, 9, c.UpdatedBy},
		A.X{`=`, 10, c.DeletedAt},
		A.X{`=`, 11, c.DeletedBy},
		A.X{`=`, 12, c.RestoredBy},
		A.X{`=`, 13, c.Editable},
	}
}

// IdxId return name of the index
func (c *Coa) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (c *Coa) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (c *Coa) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (c *Coa) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxName return name of the index
func (c *Coa) IdxName() int { //nolint:dupl false positive
	return 2
}

// SqlName return name of the column being indexed
func (c *Coa) SqlName() string { //nolint:dupl false positive
	return `"name"`
}

// IdxLabel return name of the index
func (c *Coa) IdxLabel() int { //nolint:dupl false positive
	return 3
}

// SqlLabel return name of the column being indexed
func (c *Coa) SqlLabel() string { //nolint:dupl false positive
	return `"label"`
}

// IdxParentId return name of the index
func (c *Coa) IdxParentId() int { //nolint:dupl false positive
	return 4
}

// SqlParentId return name of the column being indexed
func (c *Coa) SqlParentId() string { //nolint:dupl false positive
	return `"parentId"`
}

// IdxChildren return name of the index
func (c *Coa) IdxChildren() int { //nolint:dupl false positive
	return 5
}

// SqlChildren return name of the column being indexed
func (c *Coa) SqlChildren() string { //nolint:dupl false positive
	return `"children"`
}

// IdxCreatedAt return name of the index
func (c *Coa) IdxCreatedAt() int { //nolint:dupl false positive
	return 6
}

// SqlCreatedAt return name of the column being indexed
func (c *Coa) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (c *Coa) IdxCreatedBy() int { //nolint:dupl false positive
	return 7
}

// SqlCreatedBy return name of the column being indexed
func (c *Coa) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (c *Coa) IdxUpdatedAt() int { //nolint:dupl false positive
	return 8
}

// SqlUpdatedAt return name of the column being indexed
func (c *Coa) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (c *Coa) IdxUpdatedBy() int { //nolint:dupl false positive
	return 9
}

// SqlUpdatedBy return name of the column being indexed
func (c *Coa) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (c *Coa) IdxDeletedAt() int { //nolint:dupl false positive
	return 10
}

// SqlDeletedAt return name of the column being indexed
func (c *Coa) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (c *Coa) IdxDeletedBy() int { //nolint:dupl false positive
	return 11
}

// SqlDeletedBy return name of the column being indexed
func (c *Coa) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (c *Coa) IdxRestoredBy() int { //nolint:dupl false positive
	return 12
}

// SqlRestoredBy return name of the column being indexed
func (c *Coa) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxEditable return name of the index
func (c *Coa) IdxEditable() int { //nolint:dupl false positive
	return 13
}

// SqlEditable return name of the column being indexed
func (c *Coa) SqlEditable() string { //nolint:dupl false positive
	return `"editable"`
}

// ToArray receiver fields to slice
func (c *Coa) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if c.Id != 0 {
		id = c.Id
	}
	return A.X{
		id,
		c.TenantCode, // 1
		c.Name,       // 2
		c.Label,      // 3
		c.ParentId,   // 4
		c.Children,   // 5
		c.CreatedAt,  // 6
		c.CreatedBy,  // 7
		c.UpdatedAt,  // 8
		c.UpdatedBy,  // 9
		c.DeletedAt,  // 10
		c.DeletedBy,  // 11
		c.RestoredBy, // 12
		c.Editable,   // 13
	}
}

// FromArray convert slice to receiver fields
func (c *Coa) FromArray(a A.X) *Coa { //nolint:dupl false positive
	c.Id = X.ToU(a[0])
	c.TenantCode = X.ToS(a[1])
	c.Name = X.ToS(a[2])
	c.Label = X.ToS(a[3])
	c.ParentId = X.ToU(a[4])
	c.Children = X.ToArr(a[5])
	c.CreatedAt = X.ToI(a[6])
	c.CreatedBy = X.ToU(a[7])
	c.UpdatedAt = X.ToI(a[8])
	c.UpdatedBy = X.ToU(a[9])
	c.DeletedAt = X.ToI(a[10])
	c.DeletedBy = X.ToU(a[11])
	c.RestoredBy = X.ToU(a[12])
	c.Editable = X.ToBool(a[13])
	return c
}

// FromUncensoredArray convert slice to receiver fields
func (c *Coa) FromUncensoredArray(a A.X) *Coa { //nolint:dupl false positive
	c.Id = X.ToU(a[0])
	c.TenantCode = X.ToS(a[1])
	c.Name = X.ToS(a[2])
	c.Label = X.ToS(a[3])
	c.ParentId = X.ToU(a[4])
	c.Children = X.ToArr(a[5])
	c.CreatedAt = X.ToI(a[6])
	c.CreatedBy = X.ToU(a[7])
	c.UpdatedAt = X.ToI(a[8])
	c.UpdatedBy = X.ToU(a[9])
	c.DeletedAt = X.ToI(a[10])
	c.DeletedBy = X.ToU(a[11])
	c.RestoredBy = X.ToU(a[12])
	c.Editable = X.ToBool(a[13])
	return c
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (c *Coa) FindOffsetLimit(offset, limit uint32, idx string) []Coa { //nolint:dupl false positive
	var rows []Coa
	res, err := c.Adapter.Select(c.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Coa.FindOffsetLimit failed: `+c.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Coa{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (c *Coa) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := c.Adapter.Select(c.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Coa.FindOffsetLimit failed: `+c.SpaceName()) {
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
func (c *Coa) Total() int64 { //nolint:dupl false positive
	rows := c.Adapter.CallBoxSpace(c.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// CoaFieldTypeMap returns key value of field name and key
var CoaFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:         Tt.Unsigned,
	`tenantCode`: Tt.String,
	`name`:       Tt.String,
	`label`:      Tt.String,
	`parentId`:   Tt.Unsigned,
	`children`:   Tt.Array,
	`createdAt`:  Tt.Integer,
	`createdBy`:  Tt.Unsigned,
	`updatedAt`:  Tt.Integer,
	`updatedBy`:  Tt.Unsigned,
	`deletedAt`:  Tt.Integer,
	`deletedBy`:  Tt.Unsigned,
	`restoredBy`: Tt.Unsigned,
	`editable`:   Tt.Boolean,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// TransactionJournal DAO reader/query struct
type TransactionJournal struct {
	Adapter               *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id                    uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantCode            string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	CoaId                 uint64      `json:"coaId,string" form:"coaId" query:"coaId" long:"coaId" msg:"coaId"`
	DebitIDR              int64       `json:"debitIDR,string" form:"debitIDR" query:"debitIDR" long:"debitIDR" msg:"debitIDR"`
	CreditIDR             int64       `json:"creditIDR,string" form:"creditIDR" query:"creditIDR" long:"creditIDR" msg:"creditIDR"`
	Descriptions          string      `json:"descriptions" form:"descriptions" query:"descriptions" long:"descriptions" msg:"descriptions"`
	Date                  string      `json:"date" form:"date" query:"date" long:"date" msg:"date"`
	DetailObj             string      `json:"detailObj" form:"detailObj" query:"detailObj" long:"detailObj" msg:"detailObj"`
	CreatedAt             int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy             uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt             int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy             uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt             int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy             uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy            uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	TransactionTemplateId uint64      `json:"transactionTemplateId,string" form:"transactionTemplateId" query:"transactionTemplateId" long:"transactionTemplateId" msg:"transactionTemplateId"`
}

// NewTransactionJournal create new ORM reader/query object
func NewTransactionJournal(adapter *Tt.Adapter) *TransactionJournal {
	return &TransactionJournal{Adapter: adapter}
}

// SpaceName returns full package and table name
func (t *TransactionJournal) SpaceName() string { //nolint:dupl false positive
	return string(mFinance.TableTransactionJournal) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (t *TransactionJournal) SqlTableName() string { //nolint:dupl false positive
	return `"transactionJournal"`
}

func (t *TransactionJournal) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (t *TransactionJournal) FindById() bool { //nolint:dupl false positive
	res, err := t.Adapter.Select(t.SpaceName(), t.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{t.Id})
	if L.IsError(err, `TransactionJournal.FindById failed: `+t.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		t.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (t *TransactionJournal) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "coaId"
	, "debitIDR"
	, "creditIDR"
	, "descriptions"
	, "date"
	, "detailObj"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "transactionTemplateId"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (t *TransactionJournal) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "coaId"
	, "debitIDR"
	, "creditIDR"
	, "descriptions"
	, "date"
	, "detailObj"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "transactionTemplateId"
	`
}

// ToUpdateArray generate slice of update command
func (t *TransactionJournal) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, t.Id},
		A.X{`=`, 1, t.TenantCode},
		A.X{`=`, 2, t.CoaId},
		A.X{`=`, 3, t.DebitIDR},
		A.X{`=`, 4, t.CreditIDR},
		A.X{`=`, 5, t.Descriptions},
		A.X{`=`, 6, t.Date},
		A.X{`=`, 7, t.DetailObj},
		A.X{`=`, 8, t.CreatedAt},
		A.X{`=`, 9, t.CreatedBy},
		A.X{`=`, 10, t.UpdatedAt},
		A.X{`=`, 11, t.UpdatedBy},
		A.X{`=`, 12, t.DeletedAt},
		A.X{`=`, 13, t.DeletedBy},
		A.X{`=`, 14, t.RestoredBy},
		A.X{`=`, 15, t.TransactionTemplateId},
	}
}

// IdxId return name of the index
func (t *TransactionJournal) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (t *TransactionJournal) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (t *TransactionJournal) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (t *TransactionJournal) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxCoaId return name of the index
func (t *TransactionJournal) IdxCoaId() int { //nolint:dupl false positive
	return 2
}

// SqlCoaId return name of the column being indexed
func (t *TransactionJournal) SqlCoaId() string { //nolint:dupl false positive
	return `"coaId"`
}

// IdxDebitIDR return name of the index
func (t *TransactionJournal) IdxDebitIDR() int { //nolint:dupl false positive
	return 3
}

// SqlDebitIDR return name of the column being indexed
func (t *TransactionJournal) SqlDebitIDR() string { //nolint:dupl false positive
	return `"debitIDR"`
}

// IdxCreditIDR return name of the index
func (t *TransactionJournal) IdxCreditIDR() int { //nolint:dupl false positive
	return 4
}

// SqlCreditIDR return name of the column being indexed
func (t *TransactionJournal) SqlCreditIDR() string { //nolint:dupl false positive
	return `"creditIDR"`
}

// IdxDescriptions return name of the index
func (t *TransactionJournal) IdxDescriptions() int { //nolint:dupl false positive
	return 5
}

// SqlDescriptions return name of the column being indexed
func (t *TransactionJournal) SqlDescriptions() string { //nolint:dupl false positive
	return `"descriptions"`
}

// IdxDate return name of the index
func (t *TransactionJournal) IdxDate() int { //nolint:dupl false positive
	return 6
}

// SqlDate return name of the column being indexed
func (t *TransactionJournal) SqlDate() string { //nolint:dupl false positive
	return `"date"`
}

// IdxDetailObj return name of the index
func (t *TransactionJournal) IdxDetailObj() int { //nolint:dupl false positive
	return 7
}

// SqlDetailObj return name of the column being indexed
func (t *TransactionJournal) SqlDetailObj() string { //nolint:dupl false positive
	return `"detailObj"`
}

// IdxCreatedAt return name of the index
func (t *TransactionJournal) IdxCreatedAt() int { //nolint:dupl false positive
	return 8
}

// SqlCreatedAt return name of the column being indexed
func (t *TransactionJournal) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (t *TransactionJournal) IdxCreatedBy() int { //nolint:dupl false positive
	return 9
}

// SqlCreatedBy return name of the column being indexed
func (t *TransactionJournal) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (t *TransactionJournal) IdxUpdatedAt() int { //nolint:dupl false positive
	return 10
}

// SqlUpdatedAt return name of the column being indexed
func (t *TransactionJournal) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (t *TransactionJournal) IdxUpdatedBy() int { //nolint:dupl false positive
	return 11
}

// SqlUpdatedBy return name of the column being indexed
func (t *TransactionJournal) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (t *TransactionJournal) IdxDeletedAt() int { //nolint:dupl false positive
	return 12
}

// SqlDeletedAt return name of the column being indexed
func (t *TransactionJournal) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (t *TransactionJournal) IdxDeletedBy() int { //nolint:dupl false positive
	return 13
}

// SqlDeletedBy return name of the column being indexed
func (t *TransactionJournal) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (t *TransactionJournal) IdxRestoredBy() int { //nolint:dupl false positive
	return 14
}

// SqlRestoredBy return name of the column being indexed
func (t *TransactionJournal) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxTransactionTemplateId return name of the index
func (t *TransactionJournal) IdxTransactionTemplateId() int { //nolint:dupl false positive
	return 15
}

// SqlTransactionTemplateId return name of the column being indexed
func (t *TransactionJournal) SqlTransactionTemplateId() string { //nolint:dupl false positive
	return `"transactionTemplateId"`
}

// ToArray receiver fields to slice
func (t *TransactionJournal) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if t.Id != 0 {
		id = t.Id
	}
	return A.X{
		id,
		t.TenantCode,            // 1
		t.CoaId,                 // 2
		t.DebitIDR,              // 3
		t.CreditIDR,             // 4
		t.Descriptions,          // 5
		t.Date,                  // 6
		t.DetailObj,             // 7
		t.CreatedAt,             // 8
		t.CreatedBy,             // 9
		t.UpdatedAt,             // 10
		t.UpdatedBy,             // 11
		t.DeletedAt,             // 12
		t.DeletedBy,             // 13
		t.RestoredBy,            // 14
		t.TransactionTemplateId, // 15
	}
}

// FromArray convert slice to receiver fields
func (t *TransactionJournal) FromArray(a A.X) *TransactionJournal { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantCode = X.ToS(a[1])
	t.CoaId = X.ToU(a[2])
	t.DebitIDR = X.ToI(a[3])
	t.CreditIDR = X.ToI(a[4])
	t.Descriptions = X.ToS(a[5])
	t.Date = X.ToS(a[6])
	t.DetailObj = X.ToS(a[7])
	t.CreatedAt = X.ToI(a[8])
	t.CreatedBy = X.ToU(a[9])
	t.UpdatedAt = X.ToI(a[10])
	t.UpdatedBy = X.ToU(a[11])
	t.DeletedAt = X.ToI(a[12])
	t.DeletedBy = X.ToU(a[13])
	t.RestoredBy = X.ToU(a[14])
	t.TransactionTemplateId = X.ToU(a[15])
	return t
}

// FromUncensoredArray convert slice to receiver fields
func (t *TransactionJournal) FromUncensoredArray(a A.X) *TransactionJournal { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantCode = X.ToS(a[1])
	t.CoaId = X.ToU(a[2])
	t.DebitIDR = X.ToI(a[3])
	t.CreditIDR = X.ToI(a[4])
	t.Descriptions = X.ToS(a[5])
	t.Date = X.ToS(a[6])
	t.DetailObj = X.ToS(a[7])
	t.CreatedAt = X.ToI(a[8])
	t.CreatedBy = X.ToU(a[9])
	t.UpdatedAt = X.ToI(a[10])
	t.UpdatedBy = X.ToU(a[11])
	t.DeletedAt = X.ToI(a[12])
	t.DeletedBy = X.ToU(a[13])
	t.RestoredBy = X.ToU(a[14])
	t.TransactionTemplateId = X.ToU(a[15])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *TransactionJournal) FindOffsetLimit(offset, limit uint32, idx string) []TransactionJournal { //nolint:dupl false positive
	var rows []TransactionJournal
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `TransactionJournal.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := TransactionJournal{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *TransactionJournal) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `TransactionJournal.FindOffsetLimit failed: `+t.SpaceName()) {
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
func (t *TransactionJournal) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// TransactionJournalFieldTypeMap returns key value of field name and key
var TransactionJournalFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:                    Tt.Unsigned,
	`tenantCode`:            Tt.String,
	`coaId`:                 Tt.Unsigned,
	`debitIDR`:              Tt.Integer,
	`creditIDR`:             Tt.Integer,
	`descriptions`:          Tt.String,
	`date`:                  Tt.String,
	`detailObj`:             Tt.String,
	`createdAt`:             Tt.Integer,
	`createdBy`:             Tt.Unsigned,
	`updatedAt`:             Tt.Integer,
	`updatedBy`:             Tt.Unsigned,
	`deletedAt`:             Tt.Integer,
	`deletedBy`:             Tt.Unsigned,
	`restoredBy`:            Tt.Unsigned,
	`transactionTemplateId`: Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// TransactionTemplate DAO reader/query struct
type TransactionTemplate struct {
	Adapter    *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id         uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantCode string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	Name       string      `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	Color      string      `json:"color" form:"color" query:"color" long:"color" msg:"color"`
	ImageURL   string      `json:"imageURL" form:"imageURL" query:"imageURL" long:"imageURL" msg:"imageURL"`
	CreatedAt  int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy  uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt  int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy  uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt  int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy  uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
}

// NewTransactionTemplate create new ORM reader/query object
func NewTransactionTemplate(adapter *Tt.Adapter) *TransactionTemplate {
	return &TransactionTemplate{Adapter: adapter}
}

// SpaceName returns full package and table name
func (t *TransactionTemplate) SpaceName() string { //nolint:dupl false positive
	return string(mFinance.TableTransactionTemplate) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (t *TransactionTemplate) SqlTableName() string { //nolint:dupl false positive
	return `"transactionTemplate"`
}

func (t *TransactionTemplate) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (t *TransactionTemplate) FindById() bool { //nolint:dupl false positive
	res, err := t.Adapter.Select(t.SpaceName(), t.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{t.Id})
	if L.IsError(err, `TransactionTemplate.FindById failed: `+t.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		t.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (t *TransactionTemplate) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "name"
	, "color"
	, "imageURL"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (t *TransactionTemplate) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "name"
	, "color"
	, "imageURL"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// ToUpdateArray generate slice of update command
func (t *TransactionTemplate) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, t.Id},
		A.X{`=`, 1, t.TenantCode},
		A.X{`=`, 2, t.Name},
		A.X{`=`, 3, t.Color},
		A.X{`=`, 4, t.ImageURL},
		A.X{`=`, 5, t.CreatedAt},
		A.X{`=`, 6, t.CreatedBy},
		A.X{`=`, 7, t.UpdatedAt},
		A.X{`=`, 8, t.UpdatedBy},
		A.X{`=`, 9, t.DeletedAt},
		A.X{`=`, 10, t.DeletedBy},
		A.X{`=`, 11, t.RestoredBy},
	}
}

// IdxId return name of the index
func (t *TransactionTemplate) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (t *TransactionTemplate) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (t *TransactionTemplate) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (t *TransactionTemplate) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxName return name of the index
func (t *TransactionTemplate) IdxName() int { //nolint:dupl false positive
	return 2
}

// SqlName return name of the column being indexed
func (t *TransactionTemplate) SqlName() string { //nolint:dupl false positive
	return `"name"`
}

// IdxColor return name of the index
func (t *TransactionTemplate) IdxColor() int { //nolint:dupl false positive
	return 3
}

// SqlColor return name of the column being indexed
func (t *TransactionTemplate) SqlColor() string { //nolint:dupl false positive
	return `"color"`
}

// IdxImageURL return name of the index
func (t *TransactionTemplate) IdxImageURL() int { //nolint:dupl false positive
	return 4
}

// SqlImageURL return name of the column being indexed
func (t *TransactionTemplate) SqlImageURL() string { //nolint:dupl false positive
	return `"imageURL"`
}

// IdxCreatedAt return name of the index
func (t *TransactionTemplate) IdxCreatedAt() int { //nolint:dupl false positive
	return 5
}

// SqlCreatedAt return name of the column being indexed
func (t *TransactionTemplate) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (t *TransactionTemplate) IdxCreatedBy() int { //nolint:dupl false positive
	return 6
}

// SqlCreatedBy return name of the column being indexed
func (t *TransactionTemplate) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (t *TransactionTemplate) IdxUpdatedAt() int { //nolint:dupl false positive
	return 7
}

// SqlUpdatedAt return name of the column being indexed
func (t *TransactionTemplate) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (t *TransactionTemplate) IdxUpdatedBy() int { //nolint:dupl false positive
	return 8
}

// SqlUpdatedBy return name of the column being indexed
func (t *TransactionTemplate) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (t *TransactionTemplate) IdxDeletedAt() int { //nolint:dupl false positive
	return 9
}

// SqlDeletedAt return name of the column being indexed
func (t *TransactionTemplate) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (t *TransactionTemplate) IdxDeletedBy() int { //nolint:dupl false positive
	return 10
}

// SqlDeletedBy return name of the column being indexed
func (t *TransactionTemplate) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (t *TransactionTemplate) IdxRestoredBy() int { //nolint:dupl false positive
	return 11
}

// SqlRestoredBy return name of the column being indexed
func (t *TransactionTemplate) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// ToArray receiver fields to slice
func (t *TransactionTemplate) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if t.Id != 0 {
		id = t.Id
	}
	return A.X{
		id,
		t.TenantCode, // 1
		t.Name,       // 2
		t.Color,      // 3
		t.ImageURL,   // 4
		t.CreatedAt,  // 5
		t.CreatedBy,  // 6
		t.UpdatedAt,  // 7
		t.UpdatedBy,  // 8
		t.DeletedAt,  // 9
		t.DeletedBy,  // 10
		t.RestoredBy, // 11
	}
}

// FromArray convert slice to receiver fields
func (t *TransactionTemplate) FromArray(a A.X) *TransactionTemplate { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantCode = X.ToS(a[1])
	t.Name = X.ToS(a[2])
	t.Color = X.ToS(a[3])
	t.ImageURL = X.ToS(a[4])
	t.CreatedAt = X.ToI(a[5])
	t.CreatedBy = X.ToU(a[6])
	t.UpdatedAt = X.ToI(a[7])
	t.UpdatedBy = X.ToU(a[8])
	t.DeletedAt = X.ToI(a[9])
	t.DeletedBy = X.ToU(a[10])
	t.RestoredBy = X.ToU(a[11])
	return t
}

// FromUncensoredArray convert slice to receiver fields
func (t *TransactionTemplate) FromUncensoredArray(a A.X) *TransactionTemplate { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantCode = X.ToS(a[1])
	t.Name = X.ToS(a[2])
	t.Color = X.ToS(a[3])
	t.ImageURL = X.ToS(a[4])
	t.CreatedAt = X.ToI(a[5])
	t.CreatedBy = X.ToU(a[6])
	t.UpdatedAt = X.ToI(a[7])
	t.UpdatedBy = X.ToU(a[8])
	t.DeletedAt = X.ToI(a[9])
	t.DeletedBy = X.ToU(a[10])
	t.RestoredBy = X.ToU(a[11])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *TransactionTemplate) FindOffsetLimit(offset, limit uint32, idx string) []TransactionTemplate { //nolint:dupl false positive
	var rows []TransactionTemplate
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `TransactionTemplate.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := TransactionTemplate{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *TransactionTemplate) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `TransactionTemplate.FindOffsetLimit failed: `+t.SpaceName()) {
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
func (t *TransactionTemplate) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// TransactionTemplateFieldTypeMap returns key value of field name and key
var TransactionTemplateFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:         Tt.Unsigned,
	`tenantCode`: Tt.String,
	`name`:       Tt.String,
	`color`:      Tt.String,
	`imageURL`:   Tt.String,
	`createdAt`:  Tt.Integer,
	`createdBy`:  Tt.Unsigned,
	`updatedAt`:  Tt.Integer,
	`updatedBy`:  Tt.Unsigned,
	`deletedAt`:  Tt.Integer,
	`deletedBy`:  Tt.Unsigned,
	`restoredBy`: Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// TransactionTplDetail DAO reader/query struct
type TransactionTplDetail struct {
	Adapter    *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id         uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	ParentId   uint64      `json:"parentId,string" form:"parentId" query:"parentId" long:"parentId" msg:"parentId"`
	TenantCode string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	CoaId      uint64      `json:"coaId,string" form:"coaId" query:"coaId" long:"coaId" msg:"coaId"`
	IsDebit    bool        `json:"isDebit" form:"isDebit" query:"isDebit" long:"isDebit" msg:"isDebit"`
	CreatedAt  int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy  uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt  int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy  uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt  int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy  uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	Attributes []any       `json:"attributes" form:"attributes" query:"attributes" long:"attributes" msg:"attributes"`
}

// NewTransactionTplDetail create new ORM reader/query object
func NewTransactionTplDetail(adapter *Tt.Adapter) *TransactionTplDetail {
	return &TransactionTplDetail{Adapter: adapter}
}

// SpaceName returns full package and table name
func (t *TransactionTplDetail) SpaceName() string { //nolint:dupl false positive
	return string(mFinance.TableTransactionTplDetail) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (t *TransactionTplDetail) SqlTableName() string { //nolint:dupl false positive
	return `"transactionTplDetail"`
}

func (t *TransactionTplDetail) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (t *TransactionTplDetail) FindById() bool { //nolint:dupl false positive
	res, err := t.Adapter.Select(t.SpaceName(), t.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{t.Id})
	if L.IsError(err, `TransactionTplDetail.FindById failed: `+t.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		t.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (t *TransactionTplDetail) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "parentId"
	, "tenantCode"
	, "coaId"
	, "isDebit"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "attributes"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (t *TransactionTplDetail) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "parentId"
	, "tenantCode"
	, "coaId"
	, "isDebit"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "attributes"
	`
}

// ToUpdateArray generate slice of update command
func (t *TransactionTplDetail) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, t.Id},
		A.X{`=`, 1, t.ParentId},
		A.X{`=`, 2, t.TenantCode},
		A.X{`=`, 3, t.CoaId},
		A.X{`=`, 4, t.IsDebit},
		A.X{`=`, 5, t.CreatedAt},
		A.X{`=`, 6, t.CreatedBy},
		A.X{`=`, 7, t.UpdatedAt},
		A.X{`=`, 8, t.UpdatedBy},
		A.X{`=`, 9, t.DeletedAt},
		A.X{`=`, 10, t.DeletedBy},
		A.X{`=`, 11, t.RestoredBy},
		A.X{`=`, 12, t.Attributes},
	}
}

// IdxId return name of the index
func (t *TransactionTplDetail) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (t *TransactionTplDetail) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxParentId return name of the index
func (t *TransactionTplDetail) IdxParentId() int { //nolint:dupl false positive
	return 1
}

// SqlParentId return name of the column being indexed
func (t *TransactionTplDetail) SqlParentId() string { //nolint:dupl false positive
	return `"parentId"`
}

// IdxTenantCode return name of the index
func (t *TransactionTplDetail) IdxTenantCode() int { //nolint:dupl false positive
	return 2
}

// SqlTenantCode return name of the column being indexed
func (t *TransactionTplDetail) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxCoaId return name of the index
func (t *TransactionTplDetail) IdxCoaId() int { //nolint:dupl false positive
	return 3
}

// SqlCoaId return name of the column being indexed
func (t *TransactionTplDetail) SqlCoaId() string { //nolint:dupl false positive
	return `"coaId"`
}

// IdxIsDebit return name of the index
func (t *TransactionTplDetail) IdxIsDebit() int { //nolint:dupl false positive
	return 4
}

// SqlIsDebit return name of the column being indexed
func (t *TransactionTplDetail) SqlIsDebit() string { //nolint:dupl false positive
	return `"isDebit"`
}

// IdxCreatedAt return name of the index
func (t *TransactionTplDetail) IdxCreatedAt() int { //nolint:dupl false positive
	return 5
}

// SqlCreatedAt return name of the column being indexed
func (t *TransactionTplDetail) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (t *TransactionTplDetail) IdxCreatedBy() int { //nolint:dupl false positive
	return 6
}

// SqlCreatedBy return name of the column being indexed
func (t *TransactionTplDetail) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (t *TransactionTplDetail) IdxUpdatedAt() int { //nolint:dupl false positive
	return 7
}

// SqlUpdatedAt return name of the column being indexed
func (t *TransactionTplDetail) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (t *TransactionTplDetail) IdxUpdatedBy() int { //nolint:dupl false positive
	return 8
}

// SqlUpdatedBy return name of the column being indexed
func (t *TransactionTplDetail) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (t *TransactionTplDetail) IdxDeletedAt() int { //nolint:dupl false positive
	return 9
}

// SqlDeletedAt return name of the column being indexed
func (t *TransactionTplDetail) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (t *TransactionTplDetail) IdxDeletedBy() int { //nolint:dupl false positive
	return 10
}

// SqlDeletedBy return name of the column being indexed
func (t *TransactionTplDetail) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (t *TransactionTplDetail) IdxRestoredBy() int { //nolint:dupl false positive
	return 11
}

// SqlRestoredBy return name of the column being indexed
func (t *TransactionTplDetail) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxAttributes return name of the index
func (t *TransactionTplDetail) IdxAttributes() int { //nolint:dupl false positive
	return 12
}

// SqlAttributes return name of the column being indexed
func (t *TransactionTplDetail) SqlAttributes() string { //nolint:dupl false positive
	return `"attributes"`
}

// ToArray receiver fields to slice
func (t *TransactionTplDetail) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if t.Id != 0 {
		id = t.Id
	}
	return A.X{
		id,
		t.ParentId,   // 1
		t.TenantCode, // 2
		t.CoaId,      // 3
		t.IsDebit,    // 4
		t.CreatedAt,  // 5
		t.CreatedBy,  // 6
		t.UpdatedAt,  // 7
		t.UpdatedBy,  // 8
		t.DeletedAt,  // 9
		t.DeletedBy,  // 10
		t.RestoredBy, // 11
		t.Attributes, // 12
	}
}

// FromArray convert slice to receiver fields
func (t *TransactionTplDetail) FromArray(a A.X) *TransactionTplDetail { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.ParentId = X.ToU(a[1])
	t.TenantCode = X.ToS(a[2])
	t.CoaId = X.ToU(a[3])
	t.IsDebit = X.ToBool(a[4])
	t.CreatedAt = X.ToI(a[5])
	t.CreatedBy = X.ToU(a[6])
	t.UpdatedAt = X.ToI(a[7])
	t.UpdatedBy = X.ToU(a[8])
	t.DeletedAt = X.ToI(a[9])
	t.DeletedBy = X.ToU(a[10])
	t.RestoredBy = X.ToU(a[11])
	t.Attributes = X.ToArr(a[12])
	return t
}

// FromUncensoredArray convert slice to receiver fields
func (t *TransactionTplDetail) FromUncensoredArray(a A.X) *TransactionTplDetail { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.ParentId = X.ToU(a[1])
	t.TenantCode = X.ToS(a[2])
	t.CoaId = X.ToU(a[3])
	t.IsDebit = X.ToBool(a[4])
	t.CreatedAt = X.ToI(a[5])
	t.CreatedBy = X.ToU(a[6])
	t.UpdatedAt = X.ToI(a[7])
	t.UpdatedBy = X.ToU(a[8])
	t.DeletedAt = X.ToI(a[9])
	t.DeletedBy = X.ToU(a[10])
	t.RestoredBy = X.ToU(a[11])
	t.Attributes = X.ToArr(a[12])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *TransactionTplDetail) FindOffsetLimit(offset, limit uint32, idx string) []TransactionTplDetail { //nolint:dupl false positive
	var rows []TransactionTplDetail
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `TransactionTplDetail.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := TransactionTplDetail{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *TransactionTplDetail) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `TransactionTplDetail.FindOffsetLimit failed: `+t.SpaceName()) {
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
func (t *TransactionTplDetail) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// TransactionTplDetailFieldTypeMap returns key value of field name and key
var TransactionTplDetailFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:         Tt.Unsigned,
	`parentId`:   Tt.Unsigned,
	`tenantCode`: Tt.String,
	`coaId`:      Tt.Unsigned,
	`isDebit`:    Tt.Boolean,
	`createdAt`:  Tt.Integer,
	`createdBy`:  Tt.Unsigned,
	`updatedAt`:  Tt.Integer,
	`updatedBy`:  Tt.Unsigned,
	`deletedAt`:  Tt.Integer,
	`deletedBy`:  Tt.Unsigned,
	`restoredBy`: Tt.Unsigned,
	`attributes`: Tt.Array,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
