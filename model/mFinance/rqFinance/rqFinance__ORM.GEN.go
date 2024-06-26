package rqFinance

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mFinance"

	"github.com/tarantool/go-tarantool/v2"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Coa DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqFinance__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqFinance__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqFinance__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqFinance__ORM.GEN.go
type Coa struct {
	Adapter    *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id         uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantCode string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	Name       string      `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	Level      float64     `json:"level" form:"level" query:"level" long:"level" msg:"level"`
	ParentId   uint64      `json:"parentId,string" form:"parentId" query:"parentId" long:"parentId" msg:"parentId"`
	Children   []any       `json:"children" form:"children" query:"children" long:"children" msg:"children"`
	CreatedAt  int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy  uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt  int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy  uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt  int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy  uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
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
	res, err := c.Adapter.RetryDo(
		tarantool.NewSelectRequest(c.SpaceName()).
			Index(c.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(c.Id)}),
	)
	if L.IsError(err, `Coa.FindById failed: `+c.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			c.FromArray(row)
			return true
		}
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (c *Coa) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "name"
	, "level"
	, "parentId"
	, "children"
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
func (c *Coa) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "name"
	, "level"
	, "parentId"
	, "children"
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
func (c *Coa) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, c.Id).
		Assign(1, c.TenantCode).
		Assign(2, c.Name).
		Assign(3, c.Level).
		Assign(4, c.ParentId).
		Assign(5, c.Children).
		Assign(6, c.CreatedAt).
		Assign(7, c.CreatedBy).
		Assign(8, c.UpdatedAt).
		Assign(9, c.UpdatedBy).
		Assign(10, c.DeletedAt).
		Assign(11, c.DeletedBy).
		Assign(12, c.RestoredBy)
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

// IdxLevel return name of the index
func (c *Coa) IdxLevel() int { //nolint:dupl false positive
	return 3
}

// SqlLevel return name of the column being indexed
func (c *Coa) SqlLevel() string { //nolint:dupl false positive
	return `"level"`
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
		c.Level,      // 3
		c.ParentId,   // 4
		c.Children,   // 5
		c.CreatedAt,  // 6
		c.CreatedBy,  // 7
		c.UpdatedAt,  // 8
		c.UpdatedBy,  // 9
		c.DeletedAt,  // 10
		c.DeletedBy,  // 11
		c.RestoredBy, // 12
	}
}

// FromArray convert slice to receiver fields
func (c *Coa) FromArray(a A.X) *Coa { //nolint:dupl false positive
	c.Id = X.ToU(a[0])
	c.TenantCode = X.ToS(a[1])
	c.Name = X.ToS(a[2])
	c.Level = X.ToF(a[3])
	c.ParentId = X.ToU(a[4])
	c.Children = X.ToArr(a[5])
	c.CreatedAt = X.ToI(a[6])
	c.CreatedBy = X.ToU(a[7])
	c.UpdatedAt = X.ToI(a[8])
	c.UpdatedBy = X.ToU(a[9])
	c.DeletedAt = X.ToI(a[10])
	c.DeletedBy = X.ToU(a[11])
	c.RestoredBy = X.ToU(a[12])
	return c
}

// FromUncensoredArray convert slice to receiver fields
func (c *Coa) FromUncensoredArray(a A.X) *Coa { //nolint:dupl false positive
	c.Id = X.ToU(a[0])
	c.TenantCode = X.ToS(a[1])
	c.Name = X.ToS(a[2])
	c.Level = X.ToF(a[3])
	c.ParentId = X.ToU(a[4])
	c.Children = X.ToArr(a[5])
	c.CreatedAt = X.ToI(a[6])
	c.CreatedBy = X.ToU(a[7])
	c.UpdatedAt = X.ToI(a[8])
	c.UpdatedBy = X.ToU(a[9])
	c.DeletedAt = X.ToI(a[10])
	c.DeletedBy = X.ToU(a[11])
	c.RestoredBy = X.ToU(a[12])
	return c
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (c *Coa) FindOffsetLimit(offset, limit uint32, idx string) []Coa { //nolint:dupl false positive
	var rows []Coa
	res, err := c.Adapter.RetryDo(
		tarantool.NewSelectRequest(c.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Coa.FindOffsetLimit failed: `+c.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := Coa{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (c *Coa) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := c.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(c.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Coa.FindOffsetLimit failed: `+c.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `Coa.FindOffsetLimit failed: `+c.SpaceName()) {
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
	`level`:      Tt.Double,
	`parentId`:   Tt.Unsigned,
	`children`:   Tt.Array,
	`createdAt`:  Tt.Integer,
	`createdBy`:  Tt.Unsigned,
	`updatedAt`:  Tt.Integer,
	`updatedBy`:  Tt.Unsigned,
	`deletedAt`:  Tt.Integer,
	`deletedBy`:  Tt.Unsigned,
	`restoredBy`: Tt.Unsigned,
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
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(t.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(t.Id)}),
	)
	if L.IsError(err, `TransactionTemplate.FindById failed: `+t.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			t.FromArray(row)
			return true
		}
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
func (t *TransactionTemplate) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, t.Id).
		Assign(1, t.TenantCode).
		Assign(2, t.Name).
		Assign(3, t.Color).
		Assign(4, t.ImageURL).
		Assign(5, t.CreatedAt).
		Assign(6, t.CreatedBy).
		Assign(7, t.UpdatedAt).
		Assign(8, t.UpdatedBy).
		Assign(9, t.DeletedAt).
		Assign(10, t.DeletedBy).
		Assign(11, t.RestoredBy)
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
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `TransactionTemplate.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := TransactionTemplate{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *TransactionTemplate) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := t.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `TransactionTemplate.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `TransactionTemplate.FindOffsetLimit failed: `+t.SpaceName()) {
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

// TransactionTemplateDetail DAO reader/query struct
type TransactionTemplateDetail struct {
	Adapter    *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id         uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
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
}

// NewTransactionTemplateDetail create new ORM reader/query object
func NewTransactionTemplateDetail(adapter *Tt.Adapter) *TransactionTemplateDetail {
	return &TransactionTemplateDetail{Adapter: adapter}
}

// SpaceName returns full package and table name
func (t *TransactionTemplateDetail) SpaceName() string { //nolint:dupl false positive
	return string(mFinance.TableTransactionTemplateDetail) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (t *TransactionTemplateDetail) SqlTableName() string { //nolint:dupl false positive
	return `"transactionTemplateDetail"`
}

func (t *TransactionTemplateDetail) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (t *TransactionTemplateDetail) FindById() bool { //nolint:dupl false positive
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(t.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(t.Id)}),
	)
	if L.IsError(err, `TransactionTemplateDetail.FindById failed: `+t.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			t.FromArray(row)
			return true
		}
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (t *TransactionTemplateDetail) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
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
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (t *TransactionTemplateDetail) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
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
	`
}

// ToUpdateArray generate slice of update command
func (t *TransactionTemplateDetail) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, t.Id).
		Assign(1, t.TenantCode).
		Assign(2, t.CoaId).
		Assign(3, t.IsDebit).
		Assign(4, t.CreatedAt).
		Assign(5, t.CreatedBy).
		Assign(6, t.UpdatedAt).
		Assign(7, t.UpdatedBy).
		Assign(8, t.DeletedAt).
		Assign(9, t.DeletedBy).
		Assign(10, t.RestoredBy)
}

// IdxId return name of the index
func (t *TransactionTemplateDetail) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (t *TransactionTemplateDetail) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (t *TransactionTemplateDetail) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (t *TransactionTemplateDetail) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxCoaId return name of the index
func (t *TransactionTemplateDetail) IdxCoaId() int { //nolint:dupl false positive
	return 2
}

// SqlCoaId return name of the column being indexed
func (t *TransactionTemplateDetail) SqlCoaId() string { //nolint:dupl false positive
	return `"coaId"`
}

// IdxIsDebit return name of the index
func (t *TransactionTemplateDetail) IdxIsDebit() int { //nolint:dupl false positive
	return 3
}

// SqlIsDebit return name of the column being indexed
func (t *TransactionTemplateDetail) SqlIsDebit() string { //nolint:dupl false positive
	return `"isDebit"`
}

// IdxCreatedAt return name of the index
func (t *TransactionTemplateDetail) IdxCreatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlCreatedAt return name of the column being indexed
func (t *TransactionTemplateDetail) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (t *TransactionTemplateDetail) IdxCreatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlCreatedBy return name of the column being indexed
func (t *TransactionTemplateDetail) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (t *TransactionTemplateDetail) IdxUpdatedAt() int { //nolint:dupl false positive
	return 6
}

// SqlUpdatedAt return name of the column being indexed
func (t *TransactionTemplateDetail) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (t *TransactionTemplateDetail) IdxUpdatedBy() int { //nolint:dupl false positive
	return 7
}

// SqlUpdatedBy return name of the column being indexed
func (t *TransactionTemplateDetail) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (t *TransactionTemplateDetail) IdxDeletedAt() int { //nolint:dupl false positive
	return 8
}

// SqlDeletedAt return name of the column being indexed
func (t *TransactionTemplateDetail) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (t *TransactionTemplateDetail) IdxDeletedBy() int { //nolint:dupl false positive
	return 9
}

// SqlDeletedBy return name of the column being indexed
func (t *TransactionTemplateDetail) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (t *TransactionTemplateDetail) IdxRestoredBy() int { //nolint:dupl false positive
	return 10
}

// SqlRestoredBy return name of the column being indexed
func (t *TransactionTemplateDetail) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// ToArray receiver fields to slice
func (t *TransactionTemplateDetail) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if t.Id != 0 {
		id = t.Id
	}
	return A.X{
		id,
		t.TenantCode, // 1
		t.CoaId,      // 2
		t.IsDebit,    // 3
		t.CreatedAt,  // 4
		t.CreatedBy,  // 5
		t.UpdatedAt,  // 6
		t.UpdatedBy,  // 7
		t.DeletedAt,  // 8
		t.DeletedBy,  // 9
		t.RestoredBy, // 10
	}
}

// FromArray convert slice to receiver fields
func (t *TransactionTemplateDetail) FromArray(a A.X) *TransactionTemplateDetail { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantCode = X.ToS(a[1])
	t.CoaId = X.ToU(a[2])
	t.IsDebit = X.ToBool(a[3])
	t.CreatedAt = X.ToI(a[4])
	t.CreatedBy = X.ToU(a[5])
	t.UpdatedAt = X.ToI(a[6])
	t.UpdatedBy = X.ToU(a[7])
	t.DeletedAt = X.ToI(a[8])
	t.DeletedBy = X.ToU(a[9])
	t.RestoredBy = X.ToU(a[10])
	return t
}

// FromUncensoredArray convert slice to receiver fields
func (t *TransactionTemplateDetail) FromUncensoredArray(a A.X) *TransactionTemplateDetail { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantCode = X.ToS(a[1])
	t.CoaId = X.ToU(a[2])
	t.IsDebit = X.ToBool(a[3])
	t.CreatedAt = X.ToI(a[4])
	t.CreatedBy = X.ToU(a[5])
	t.UpdatedAt = X.ToI(a[6])
	t.UpdatedBy = X.ToU(a[7])
	t.DeletedAt = X.ToI(a[8])
	t.DeletedBy = X.ToU(a[9])
	t.RestoredBy = X.ToU(a[10])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *TransactionTemplateDetail) FindOffsetLimit(offset, limit uint32, idx string) []TransactionTemplateDetail { //nolint:dupl false positive
	var rows []TransactionTemplateDetail
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `TransactionTemplateDetail.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := TransactionTemplateDetail{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *TransactionTemplateDetail) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := t.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `TransactionTemplateDetail.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `TransactionTemplateDetail.FindOffsetLimit failed: `+t.SpaceName()) {
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
func (t *TransactionTemplateDetail) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// TransactionTemplateDetailFieldTypeMap returns key value of field name and key
var TransactionTemplateDetailFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:         Tt.Unsigned,
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
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Transactions DAO reader/query struct
type Transactions struct {
	Adapter      *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id           uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantCode   string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	CreatedAt    int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy    uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt    int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy    uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt    int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy    uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy   uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	CompletedAt  int64       `json:"completedAt" form:"completedAt" query:"completedAt" long:"completedAt" msg:"completedAt"`
	Price        uint64      `json:"price" form:"price" query:"price" long:"price" msg:"price"`
	Descriptions string      `json:"descriptions" form:"descriptions" query:"descriptions" long:"descriptions" msg:"descriptions"`
	Qty          int64       `json:"qty" form:"qty" query:"qty" long:"qty" msg:"qty"`
}

// NewTransactions create new ORM reader/query object
func NewTransactions(adapter *Tt.Adapter) *Transactions {
	return &Transactions{Adapter: adapter}
}

// SpaceName returns full package and table name
func (t *Transactions) SpaceName() string { //nolint:dupl false positive
	return string(mFinance.TableTransactions) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (t *Transactions) SqlTableName() string { //nolint:dupl false positive
	return `"transactions"`
}

func (t *Transactions) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (t *Transactions) FindById() bool { //nolint:dupl false positive
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(t.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(t.Id)}),
	)
	if L.IsError(err, `Transactions.FindById failed: `+t.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			t.FromArray(row)
			return true
		}
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (t *Transactions) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "completedAt"
	, "price"
	, "descriptions"
	, "qty"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (t *Transactions) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "completedAt"
	, "price"
	, "descriptions"
	, "qty"
	`
}

// ToUpdateArray generate slice of update command
func (t *Transactions) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, t.Id).
		Assign(1, t.TenantCode).
		Assign(2, t.CreatedAt).
		Assign(3, t.CreatedBy).
		Assign(4, t.UpdatedAt).
		Assign(5, t.UpdatedBy).
		Assign(6, t.DeletedAt).
		Assign(7, t.DeletedBy).
		Assign(8, t.RestoredBy).
		Assign(9, t.CompletedAt).
		Assign(10, t.Price).
		Assign(11, t.Descriptions).
		Assign(12, t.Qty)
}

// IdxId return name of the index
func (t *Transactions) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (t *Transactions) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (t *Transactions) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (t *Transactions) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxCreatedAt return name of the index
func (t *Transactions) IdxCreatedAt() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedAt return name of the column being indexed
func (t *Transactions) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (t *Transactions) IdxCreatedBy() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedBy return name of the column being indexed
func (t *Transactions) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (t *Transactions) IdxUpdatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlUpdatedAt return name of the column being indexed
func (t *Transactions) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (t *Transactions) IdxUpdatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedBy return name of the column being indexed
func (t *Transactions) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (t *Transactions) IdxDeletedAt() int { //nolint:dupl false positive
	return 6
}

// SqlDeletedAt return name of the column being indexed
func (t *Transactions) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (t *Transactions) IdxDeletedBy() int { //nolint:dupl false positive
	return 7
}

// SqlDeletedBy return name of the column being indexed
func (t *Transactions) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (t *Transactions) IdxRestoredBy() int { //nolint:dupl false positive
	return 8
}

// SqlRestoredBy return name of the column being indexed
func (t *Transactions) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxCompletedAt return name of the index
func (t *Transactions) IdxCompletedAt() int { //nolint:dupl false positive
	return 9
}

// SqlCompletedAt return name of the column being indexed
func (t *Transactions) SqlCompletedAt() string { //nolint:dupl false positive
	return `"completedAt"`
}

// IdxPrice return name of the index
func (t *Transactions) IdxPrice() int { //nolint:dupl false positive
	return 10
}

// SqlPrice return name of the column being indexed
func (t *Transactions) SqlPrice() string { //nolint:dupl false positive
	return `"price"`
}

// IdxDescriptions return name of the index
func (t *Transactions) IdxDescriptions() int { //nolint:dupl false positive
	return 11
}

// SqlDescriptions return name of the column being indexed
func (t *Transactions) SqlDescriptions() string { //nolint:dupl false positive
	return `"descriptions"`
}

// IdxQty return name of the index
func (t *Transactions) IdxQty() int { //nolint:dupl false positive
	return 12
}

// SqlQty return name of the column being indexed
func (t *Transactions) SqlQty() string { //nolint:dupl false positive
	return `"qty"`
}

// ToArray receiver fields to slice
func (t *Transactions) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if t.Id != 0 {
		id = t.Id
	}
	return A.X{
		id,
		t.TenantCode,   // 1
		t.CreatedAt,    // 2
		t.CreatedBy,    // 3
		t.UpdatedAt,    // 4
		t.UpdatedBy,    // 5
		t.DeletedAt,    // 6
		t.DeletedBy,    // 7
		t.RestoredBy,   // 8
		t.CompletedAt,  // 9
		t.Price,        // 10
		t.Descriptions, // 11
		t.Qty,          // 12
	}
}

// FromArray convert slice to receiver fields
func (t *Transactions) FromArray(a A.X) *Transactions { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantCode = X.ToS(a[1])
	t.CreatedAt = X.ToI(a[2])
	t.CreatedBy = X.ToU(a[3])
	t.UpdatedAt = X.ToI(a[4])
	t.UpdatedBy = X.ToU(a[5])
	t.DeletedAt = X.ToI(a[6])
	t.DeletedBy = X.ToU(a[7])
	t.RestoredBy = X.ToU(a[8])
	t.CompletedAt = X.ToI(a[9])
	t.Price = X.ToU(a[10])
	t.Descriptions = X.ToS(a[11])
	t.Qty = X.ToI(a[12])
	return t
}

// FromUncensoredArray convert slice to receiver fields
func (t *Transactions) FromUncensoredArray(a A.X) *Transactions { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantCode = X.ToS(a[1])
	t.CreatedAt = X.ToI(a[2])
	t.CreatedBy = X.ToU(a[3])
	t.UpdatedAt = X.ToI(a[4])
	t.UpdatedBy = X.ToU(a[5])
	t.DeletedAt = X.ToI(a[6])
	t.DeletedBy = X.ToU(a[7])
	t.RestoredBy = X.ToU(a[8])
	t.CompletedAt = X.ToI(a[9])
	t.Price = X.ToU(a[10])
	t.Descriptions = X.ToS(a[11])
	t.Qty = X.ToI(a[12])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *Transactions) FindOffsetLimit(offset, limit uint32, idx string) []Transactions { //nolint:dupl false positive
	var rows []Transactions
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Transactions.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := Transactions{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *Transactions) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := t.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Transactions.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `Transactions.FindOffsetLimit failed: `+t.SpaceName()) {
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
func (t *Transactions) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// TransactionsFieldTypeMap returns key value of field name and key
var TransactionsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:           Tt.Unsigned,
	`tenantCode`:   Tt.String,
	`createdAt`:    Tt.Integer,
	`createdBy`:    Tt.Unsigned,
	`updatedAt`:    Tt.Integer,
	`updatedBy`:    Tt.Unsigned,
	`deletedAt`:    Tt.Integer,
	`deletedBy`:    Tt.Unsigned,
	`restoredBy`:   Tt.Unsigned,
	`completedAt`:  Tt.Integer,
	`price`:        Tt.Unsigned,
	`descriptions`: Tt.String,
	`qty`:          Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
