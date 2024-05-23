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
	, "level"
	, "parentId"
	, "children"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
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
	`
}

// ToUpdateArray generate slice of update command
func (c *Coa) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, c.Id},
		A.X{`=`, 1, c.TenantCode},
		A.X{`=`, 2, c.Name},
		A.X{`=`, 3, c.Level},
		A.X{`=`, 4, c.ParentId},
		A.X{`=`, 5, c.Children},
		A.X{`=`, 6, c.CreatedAt},
		A.X{`=`, 7, c.CreatedBy},
		A.X{`=`, 8, c.UpdatedAt},
		A.X{`=`, 9, c.UpdatedBy},
		A.X{`=`, 10, c.DeletedAt},
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
	`level`:      Tt.Double,
	`parentId`:   Tt.Unsigned,
	`children`:   Tt.Array,
	`createdAt`:  Tt.Integer,
	`createdBy`:  Tt.Unsigned,
	`updatedAt`:  Tt.Integer,
	`updatedBy`:  Tt.Unsigned,
	`deletedAt`:  Tt.Integer,
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
	CompletedAt  int64       `json:"completedAt" form:"completedAt" query:"completedAt" long:"completedAt" msg:"completedAt"`
	Price        int64       `json:"price" form:"price" query:"price" long:"price" msg:"price"`
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

// SqlSelectAllFields generate Sql select fields
func (t *Transactions) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
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
	, "completedAt"
	, "price"
	, "descriptions"
	, "qty"
	`
}

// ToUpdateArray generate slice of update command
func (t *Transactions) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, t.Id},
		A.X{`=`, 1, t.TenantCode},
		A.X{`=`, 2, t.CreatedAt},
		A.X{`=`, 3, t.CreatedBy},
		A.X{`=`, 4, t.UpdatedAt},
		A.X{`=`, 5, t.UpdatedBy},
		A.X{`=`, 6, t.DeletedAt},
		A.X{`=`, 7, t.CompletedAt},
		A.X{`=`, 8, t.Price},
		A.X{`=`, 9, t.Descriptions},
		A.X{`=`, 10, t.Qty},
	}
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

// IdxCompletedAt return name of the index
func (t *Transactions) IdxCompletedAt() int { //nolint:dupl false positive
	return 7
}

// SqlCompletedAt return name of the column being indexed
func (t *Transactions) SqlCompletedAt() string { //nolint:dupl false positive
	return `"completedAt"`
}

// IdxPrice return name of the index
func (t *Transactions) IdxPrice() int { //nolint:dupl false positive
	return 8
}

// SqlPrice return name of the column being indexed
func (t *Transactions) SqlPrice() string { //nolint:dupl false positive
	return `"price"`
}

// IdxDescriptions return name of the index
func (t *Transactions) IdxDescriptions() int { //nolint:dupl false positive
	return 9
}

// SqlDescriptions return name of the column being indexed
func (t *Transactions) SqlDescriptions() string { //nolint:dupl false positive
	return `"descriptions"`
}

// IdxQty return name of the index
func (t *Transactions) IdxQty() int { //nolint:dupl false positive
	return 10
}

// SqlQty return name of the column being indexed
func (t *Transactions) SqlQty() string { //nolint:dupl false positive
	return `"qty"`
}

// ToArray receiver fields to slice
func (t *Transactions) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		t.Id,           // 0
		t.TenantCode,   // 1
		t.CreatedAt,    // 2
		t.CreatedBy,    // 3
		t.UpdatedAt,    // 4
		t.UpdatedBy,    // 5
		t.DeletedAt,    // 6
		t.CompletedAt,  // 7
		t.Price,        // 8
		t.Descriptions, // 9
		t.Qty,          // 10
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
	t.CompletedAt = X.ToI(a[7])
	t.Price = X.ToI(a[8])
	t.Descriptions = X.ToS(a[9])
	t.Qty = X.ToI(a[10])
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
	t.CompletedAt = X.ToI(a[7])
	t.Price = X.ToI(a[8])
	t.Descriptions = X.ToS(a[9])
	t.Qty = X.ToI(a[10])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *Transactions) FindOffsetLimit(offset, limit uint32, idx string) []Transactions { //nolint:dupl false positive
	var rows []Transactions
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Transactions.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Transactions{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *Transactions) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Transactions.FindOffsetLimit failed: `+t.SpaceName()) {
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
	`completedAt`:  Tt.Integer,
	`price`:        Tt.Integer,
	`descriptions`: Tt.String,
	`qty`:          Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
