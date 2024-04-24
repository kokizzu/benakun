package rqAuth

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	`benakun/model/mAuth`

	`github.com/tarantool/go-tarantool`

	`github.com/kokizzu/gotro/A`
	`github.com/kokizzu/gotro/D/Tt`
	`github.com/kokizzu/gotro/L`
	`github.com/kokizzu/gotro/X`
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqAuth__ORM.GEN.go
// Coa DAO reader/query struct
type Coa struct {
	Adapter *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	Id         uint64
	TenantCode string
	Name       string
	Level      float64
	ParentId   uint64
	Children   []any
	CreatedAt  int64
	CreatedBy  uint64
	UpdatedAt  int64
	UpdatedBy  uint64
	DeletedAt  int64
}

// NewCoa create new ORM reader/query object
func NewCoa(adapter *Tt.Adapter) *Coa {
	return &Coa{Adapter: adapter}
}

// SpaceName returns full package and table name
func (c *Coa) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableCoa) // casting required to string from Tt.TableName
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
	rows := c.Adapter.CallBoxSpace(c.SpaceName() + `:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// CoaFieldTypeMap returns key value of field name and key
var CoaFieldTypeMap = map[string]Tt.DataType { //nolint:dupl false positive
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

// Orgs DAO reader/query struct
type Orgs struct {
	Adapter *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	Id         uint64
	TenantCode string
	Name       string
	HeadTitle  string
	ParentId   uint64
	Children   []any
	OrgType    uint64
	CreatedAt  int64
	CreatedBy  uint64
	UpdatedAt  int64
	UpdatedBy  uint64
	DeletedAt  int64
}

// NewOrgs create new ORM reader/query object
func NewOrgs(adapter *Tt.Adapter) *Orgs {
	return &Orgs{Adapter: adapter}
}

// SpaceName returns full package and table name
func (o *Orgs) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableOrgs) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (o *Orgs) SqlTableName() string { //nolint:dupl false positive
	return `"orgs"`
}

func (o *Orgs) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (o *Orgs) FindById() bool { //nolint:dupl false positive
	res, err := o.Adapter.Select(o.SpaceName(), o.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{o.Id})
	if L.IsError(err, `Orgs.FindById failed: `+o.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		o.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (o *Orgs) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "name"
	, "headTitle"
	, "parentId"
	, "children"
	, "orgType"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (o *Orgs) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "name"
	, "headTitle"
	, "parentId"
	, "children"
	, "orgType"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	`
}

// ToUpdateArray generate slice of update command
func (o *Orgs) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, o.Id},
		A.X{`=`, 1, o.TenantCode},
		A.X{`=`, 2, o.Name},
		A.X{`=`, 3, o.HeadTitle},
		A.X{`=`, 4, o.ParentId},
		A.X{`=`, 5, o.Children},
		A.X{`=`, 6, o.OrgType},
		A.X{`=`, 7, o.CreatedAt},
		A.X{`=`, 8, o.CreatedBy},
		A.X{`=`, 9, o.UpdatedAt},
		A.X{`=`, 10, o.UpdatedBy},
		A.X{`=`, 11, o.DeletedAt},
	}
}

// IdxId return name of the index
func (o *Orgs) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (o *Orgs) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (o *Orgs) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (o *Orgs) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxName return name of the index
func (o *Orgs) IdxName() int { //nolint:dupl false positive
	return 2
}

// SqlName return name of the column being indexed
func (o *Orgs) SqlName() string { //nolint:dupl false positive
	return `"name"`
}

// IdxHeadTitle return name of the index
func (o *Orgs) IdxHeadTitle() int { //nolint:dupl false positive
	return 3
}

// SqlHeadTitle return name of the column being indexed
func (o *Orgs) SqlHeadTitle() string { //nolint:dupl false positive
	return `"headTitle"`
}

// IdxParentId return name of the index
func (o *Orgs) IdxParentId() int { //nolint:dupl false positive
	return 4
}

// SqlParentId return name of the column being indexed
func (o *Orgs) SqlParentId() string { //nolint:dupl false positive
	return `"parentId"`
}

// IdxChildren return name of the index
func (o *Orgs) IdxChildren() int { //nolint:dupl false positive
	return 5
}

// SqlChildren return name of the column being indexed
func (o *Orgs) SqlChildren() string { //nolint:dupl false positive
	return `"children"`
}

// IdxOrgType return name of the index
func (o *Orgs) IdxOrgType() int { //nolint:dupl false positive
	return 6
}

// SqlOrgType return name of the column being indexed
func (o *Orgs) SqlOrgType() string { //nolint:dupl false positive
	return `"orgType"`
}

// IdxCreatedAt return name of the index
func (o *Orgs) IdxCreatedAt() int { //nolint:dupl false positive
	return 7
}

// SqlCreatedAt return name of the column being indexed
func (o *Orgs) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (o *Orgs) IdxCreatedBy() int { //nolint:dupl false positive
	return 8
}

// SqlCreatedBy return name of the column being indexed
func (o *Orgs) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (o *Orgs) IdxUpdatedAt() int { //nolint:dupl false positive
	return 9
}

// SqlUpdatedAt return name of the column being indexed
func (o *Orgs) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (o *Orgs) IdxUpdatedBy() int { //nolint:dupl false positive
	return 10
}

// SqlUpdatedBy return name of the column being indexed
func (o *Orgs) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (o *Orgs) IdxDeletedAt() int { //nolint:dupl false positive
	return 11
}

// SqlDeletedAt return name of the column being indexed
func (o *Orgs) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// ToArray receiver fields to slice
func (o *Orgs) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if o.Id != 0 {
		id = o.Id
	}
	return A.X{
		id,
		o.TenantCode, // 1
		o.Name,       // 2
		o.HeadTitle,  // 3
		o.ParentId,   // 4
		o.Children,   // 5
		o.OrgType,    // 6
		o.CreatedAt,  // 7
		o.CreatedBy,  // 8
		o.UpdatedAt,  // 9
		o.UpdatedBy,  // 10
		o.DeletedAt,  // 11
	}
}

// FromArray convert slice to receiver fields
func (o *Orgs) FromArray(a A.X) *Orgs { //nolint:dupl false positive
	o.Id = X.ToU(a[0])
	o.TenantCode = X.ToS(a[1])
	o.Name = X.ToS(a[2])
	o.HeadTitle = X.ToS(a[3])
	o.ParentId = X.ToU(a[4])
	o.Children = X.ToArr(a[5])
	o.OrgType = X.ToU(a[6])
	o.CreatedAt = X.ToI(a[7])
	o.CreatedBy = X.ToU(a[8])
	o.UpdatedAt = X.ToI(a[9])
	o.UpdatedBy = X.ToU(a[10])
	o.DeletedAt = X.ToI(a[11])
	return o
}

// FromUncensoredArray convert slice to receiver fields
func (o *Orgs) FromUncensoredArray(a A.X) *Orgs { //nolint:dupl false positive
	o.Id = X.ToU(a[0])
	o.TenantCode = X.ToS(a[1])
	o.Name = X.ToS(a[2])
	o.HeadTitle = X.ToS(a[3])
	o.ParentId = X.ToU(a[4])
	o.Children = X.ToArr(a[5])
	o.OrgType = X.ToU(a[6])
	o.CreatedAt = X.ToI(a[7])
	o.CreatedBy = X.ToU(a[8])
	o.UpdatedAt = X.ToI(a[9])
	o.UpdatedBy = X.ToU(a[10])
	o.DeletedAt = X.ToI(a[11])
	return o
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (o *Orgs) FindOffsetLimit(offset, limit uint32, idx string) []Orgs { //nolint:dupl false positive
	var rows []Orgs
	res, err := o.Adapter.Select(o.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Orgs.FindOffsetLimit failed: `+o.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Orgs{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (o *Orgs) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := o.Adapter.Select(o.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Orgs.FindOffsetLimit failed: `+o.SpaceName()) {
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
func (o *Orgs) Total() int64 { //nolint:dupl false positive
	rows := o.Adapter.CallBoxSpace(o.SpaceName() + `:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// OrgsFieldTypeMap returns key value of field name and key
var OrgsFieldTypeMap = map[string]Tt.DataType { //nolint:dupl false positive
	`id`:         Tt.Unsigned,
	`tenantCode`: Tt.String,
	`name`:       Tt.String,
	`headTitle`:  Tt.String,
	`parentId`:   Tt.Unsigned,
	`children`:   Tt.Array,
	`orgType`:    Tt.Unsigned,
	`createdAt`:  Tt.Integer,
	`createdBy`:  Tt.Unsigned,
	`updatedAt`:  Tt.Integer,
	`updatedBy`:  Tt.Unsigned,
	`deletedAt`:  Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Sessions DAO reader/query struct
type Sessions struct {
	Adapter *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	SessionToken string
	UserId       uint64
	ExpiredAt    int64
	Device       string
	LoginAt      int64
	LoginIPs     string
	TenantCode   string
}

// NewSessions create new ORM reader/query object
func NewSessions(adapter *Tt.Adapter) *Sessions {
	return &Sessions{Adapter: adapter}
}

// SpaceName returns full package and table name
func (s *Sessions) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableSessions) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (s *Sessions) SqlTableName() string { //nolint:dupl false positive
	return `"sessions"`
}

// UniqueIndexSessionToken return unique index name
func (s *Sessions) UniqueIndexSessionToken() string { //nolint:dupl false positive
	return `sessionToken`
}

// FindBySessionToken Find one by SessionToken
func (s *Sessions) FindBySessionToken() bool { //nolint:dupl false positive
	res, err := s.Adapter.Select(s.SpaceName(), s.UniqueIndexSessionToken(), 0, 1, tarantool.IterEq, A.X{s.SessionToken})
	if L.IsError(err, `Sessions.FindBySessionToken failed: `+s.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		s.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (s *Sessions) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "sessionToken"
	, "userId"
	, "expiredAt"
	, "device"
	, "loginAt"
	, "loginIPs"
	, "tenantCode"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (s *Sessions) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "sessionToken"
	, "userId"
	, "expiredAt"
	, "device"
	, "loginAt"
	, "loginIPs"
	, "tenantCode"
	`
}

// ToUpdateArray generate slice of update command
func (s *Sessions) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, s.SessionToken},
		A.X{`=`, 1, s.UserId},
		A.X{`=`, 2, s.ExpiredAt},
		A.X{`=`, 3, s.Device},
		A.X{`=`, 4, s.LoginAt},
		A.X{`=`, 5, s.LoginIPs},
		A.X{`=`, 6, s.TenantCode},
	}
}

// IdxSessionToken return name of the index
func (s *Sessions) IdxSessionToken() int { //nolint:dupl false positive
	return 0
}

// SqlSessionToken return name of the column being indexed
func (s *Sessions) SqlSessionToken() string { //nolint:dupl false positive
	return `"sessionToken"`
}

// IdxUserId return name of the index
func (s *Sessions) IdxUserId() int { //nolint:dupl false positive
	return 1
}

// SqlUserId return name of the column being indexed
func (s *Sessions) SqlUserId() string { //nolint:dupl false positive
	return `"userId"`
}

// IdxExpiredAt return name of the index
func (s *Sessions) IdxExpiredAt() int { //nolint:dupl false positive
	return 2
}

// SqlExpiredAt return name of the column being indexed
func (s *Sessions) SqlExpiredAt() string { //nolint:dupl false positive
	return `"expiredAt"`
}

// IdxDevice return name of the index
func (s *Sessions) IdxDevice() int { //nolint:dupl false positive
	return 3
}

// SqlDevice return name of the column being indexed
func (s *Sessions) SqlDevice() string { //nolint:dupl false positive
	return `"device"`
}

// IdxLoginAt return name of the index
func (s *Sessions) IdxLoginAt() int { //nolint:dupl false positive
	return 4
}

// SqlLoginAt return name of the column being indexed
func (s *Sessions) SqlLoginAt() string { //nolint:dupl false positive
	return `"loginAt"`
}

// IdxLoginIPs return name of the index
func (s *Sessions) IdxLoginIPs() int { //nolint:dupl false positive
	return 5
}

// SqlLoginIPs return name of the column being indexed
func (s *Sessions) SqlLoginIPs() string { //nolint:dupl false positive
	return `"loginIPs"`
}

// IdxTenantCode return name of the index
func (s *Sessions) IdxTenantCode() int { //nolint:dupl false positive
	return 6
}

// SqlTenantCode return name of the column being indexed
func (s *Sessions) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// ToArray receiver fields to slice
func (s *Sessions) ToArray() A.X { //nolint:dupl false positive
	return A.X{
		s.SessionToken, // 0
		s.UserId,       // 1
		s.ExpiredAt,    // 2
		s.Device,       // 3
		s.LoginAt,      // 4
		s.LoginIPs,     // 5
		s.TenantCode,   // 6
	}
}

// FromArray convert slice to receiver fields
func (s *Sessions) FromArray(a A.X) *Sessions { //nolint:dupl false positive
	s.SessionToken = X.ToS(a[0])
	s.UserId = X.ToU(a[1])
	s.ExpiredAt = X.ToI(a[2])
	s.Device = X.ToS(a[3])
	s.LoginAt = X.ToI(a[4])
	s.LoginIPs = X.ToS(a[5])
	s.TenantCode = X.ToS(a[6])
	return s
}

// FromUncensoredArray convert slice to receiver fields
func (s *Sessions) FromUncensoredArray(a A.X) *Sessions { //nolint:dupl false positive
	s.SessionToken = X.ToS(a[0])
	s.UserId = X.ToU(a[1])
	s.ExpiredAt = X.ToI(a[2])
	s.Device = X.ToS(a[3])
	s.LoginAt = X.ToI(a[4])
	s.LoginIPs = X.ToS(a[5])
	s.TenantCode = X.ToS(a[6])
	return s
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (s *Sessions) FindOffsetLimit(offset, limit uint32, idx string) []Sessions { //nolint:dupl false positive
	var rows []Sessions
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Sessions.FindOffsetLimit failed: `+s.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Sessions{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (s *Sessions) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := s.Adapter.Select(s.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Sessions.FindOffsetLimit failed: `+s.SpaceName()) {
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
func (s *Sessions) Total() int64 { //nolint:dupl false positive
	rows := s.Adapter.CallBoxSpace(s.SpaceName() + `:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// SessionsFieldTypeMap returns key value of field name and key
var SessionsFieldTypeMap = map[string]Tt.DataType { //nolint:dupl false positive
	`sessionToken`: Tt.String,
	`userId`:       Tt.Unsigned,
	`expiredAt`:    Tt.Integer,
	`device`:       Tt.String,
	`loginAt`:      Tt.Integer,
	`loginIPs`:     Tt.String,
	`tenantCode`:   Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Tenants DAO reader/query struct
type Tenants struct {
	Adapter *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	Id         uint64
	TenantCode string
	CreatedAt  int64
	CreatedBy  uint64
	UpdatedAt  int64
	UpdatedBy  uint64
	DeletedAt  int64
}

// NewTenants create new ORM reader/query object
func NewTenants(adapter *Tt.Adapter) *Tenants {
	return &Tenants{Adapter: adapter}
}

// SpaceName returns full package and table name
func (t *Tenants) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableTenants) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (t *Tenants) SqlTableName() string { //nolint:dupl false positive
	return `"tenants"`
}

func (t *Tenants) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (t *Tenants) FindById() bool { //nolint:dupl false positive
	res, err := t.Adapter.Select(t.SpaceName(), t.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{t.Id})
	if L.IsError(err, `Tenants.FindById failed: `+t.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		t.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexTenantCode return unique index name
func (t *Tenants) UniqueIndexTenantCode() string { //nolint:dupl false positive
	return `tenantCode`
}

// FindByTenantCode Find one by TenantCode
func (t *Tenants) FindByTenantCode() bool { //nolint:dupl false positive
	res, err := t.Adapter.Select(t.SpaceName(), t.UniqueIndexTenantCode(), 0, 1, tarantool.IterEq, A.X{t.TenantCode})
	if L.IsError(err, `Tenants.FindByTenantCode failed: `+t.SpaceName()) {
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
func (t *Tenants) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (t *Tenants) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	`
}

// ToUpdateArray generate slice of update command
func (t *Tenants) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, t.Id},
		A.X{`=`, 1, t.TenantCode},
		A.X{`=`, 2, t.CreatedAt},
		A.X{`=`, 3, t.CreatedBy},
		A.X{`=`, 4, t.UpdatedAt},
		A.X{`=`, 5, t.UpdatedBy},
		A.X{`=`, 6, t.DeletedAt},
	}
}

// IdxId return name of the index
func (t *Tenants) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (t *Tenants) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (t *Tenants) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (t *Tenants) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxCreatedAt return name of the index
func (t *Tenants) IdxCreatedAt() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedAt return name of the column being indexed
func (t *Tenants) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (t *Tenants) IdxCreatedBy() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedBy return name of the column being indexed
func (t *Tenants) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (t *Tenants) IdxUpdatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlUpdatedAt return name of the column being indexed
func (t *Tenants) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (t *Tenants) IdxUpdatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedBy return name of the column being indexed
func (t *Tenants) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (t *Tenants) IdxDeletedAt() int { //nolint:dupl false positive
	return 6
}

// SqlDeletedAt return name of the column being indexed
func (t *Tenants) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// ToArray receiver fields to slice
func (t *Tenants) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if t.Id != 0 {
		id = t.Id
	}
	return A.X{
		id,
		t.TenantCode, // 1
		t.CreatedAt,  // 2
		t.CreatedBy,  // 3
		t.UpdatedAt,  // 4
		t.UpdatedBy,  // 5
		t.DeletedAt,  // 6
	}
}

// FromArray convert slice to receiver fields
func (t *Tenants) FromArray(a A.X) *Tenants { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantCode = X.ToS(a[1])
	t.CreatedAt = X.ToI(a[2])
	t.CreatedBy = X.ToU(a[3])
	t.UpdatedAt = X.ToI(a[4])
	t.UpdatedBy = X.ToU(a[5])
	t.DeletedAt = X.ToI(a[6])
	return t
}

// FromUncensoredArray convert slice to receiver fields
func (t *Tenants) FromUncensoredArray(a A.X) *Tenants { //nolint:dupl false positive
	t.Id = X.ToU(a[0])
	t.TenantCode = X.ToS(a[1])
	t.CreatedAt = X.ToI(a[2])
	t.CreatedBy = X.ToU(a[3])
	t.UpdatedAt = X.ToI(a[4])
	t.UpdatedBy = X.ToU(a[5])
	t.DeletedAt = X.ToI(a[6])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *Tenants) FindOffsetLimit(offset, limit uint32, idx string) []Tenants { //nolint:dupl false positive
	var rows []Tenants
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Tenants.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Tenants{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *Tenants) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := t.Adapter.Select(t.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Tenants.FindOffsetLimit failed: `+t.SpaceName()) {
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
func (t *Tenants) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName() + `:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// TenantsFieldTypeMap returns key value of field name and key
var TenantsFieldTypeMap = map[string]Tt.DataType { //nolint:dupl false positive
	`id`:         Tt.Unsigned,
	`tenantCode`: Tt.String,
	`createdAt`:  Tt.Integer,
	`createdBy`:  Tt.Unsigned,
	`updatedAt`:  Tt.Integer,
	`updatedBy`:  Tt.Unsigned,
	`deletedAt`:  Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Transactions DAO reader/query struct
type Transactions struct {
	Adapter *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	Id           uint64
	TenantCode   string
	CreatedAt    int64
	CreatedBy    uint64
	UpdatedAt    int64
	UpdatedBy    uint64
	DeletedAt    int64
	CompletedAt  int64
	Price        int64
	Descriptions string
	Qty          int64
}

// NewTransactions create new ORM reader/query object
func NewTransactions(adapter *Tt.Adapter) *Transactions {
	return &Transactions{Adapter: adapter}
}

// SpaceName returns full package and table name
func (t *Transactions) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableTransactions) // casting required to string from Tt.TableName
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
	rows := t.Adapter.CallBoxSpace(t.SpaceName() + `:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// TransactionsFieldTypeMap returns key value of field name and key
var TransactionsFieldTypeMap = map[string]Tt.DataType { //nolint:dupl false positive
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

// Users DAO reader/query struct
type Users struct {
	Adapter *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-"`
	Id                 uint64
	Email              string
	Password           string
	CreatedAt          int64
	CreatedBy          uint64
	UpdatedAt          int64
	UpdatedBy          uint64
	DeletedAt          int64
	PasswordSetAt      int64
	SecretCode         string
	SecretCodeAt       int64
	VerificationSentAt int64
	VerifiedAt         int64
	LastLoginAt        int64
	FullName           string
	TenantCode         string
	Role               string
	InvitationState    string
}

// NewUsers create new ORM reader/query object
func NewUsers(adapter *Tt.Adapter) *Users {
	return &Users{Adapter: adapter}
}

// SpaceName returns full package and table name
func (u *Users) SpaceName() string { //nolint:dupl false positive
	return string(mAuth.TableUsers) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (u *Users) SqlTableName() string { //nolint:dupl false positive
	return `"users"`
}

func (u *Users) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (u *Users) FindById() bool { //nolint:dupl false positive
	res, err := u.Adapter.Select(u.SpaceName(), u.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{u.Id})
	if L.IsError(err, `Users.FindById failed: `+u.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		u.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexEmail return unique index name
func (u *Users) UniqueIndexEmail() string { //nolint:dupl false positive
	return `email`
}

// FindByEmail Find one by Email
func (u *Users) FindByEmail() bool { //nolint:dupl false positive
	res, err := u.Adapter.Select(u.SpaceName(), u.UniqueIndexEmail(), 0, 1, tarantool.IterEq, A.X{u.Email})
	if L.IsError(err, `Users.FindByEmail failed: `+u.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		u.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (u *Users) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "email"
	, "password"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "passwordSetAt"
	, "secretCode"
	, "secretCodeAt"
	, "verificationSentAt"
	, "verifiedAt"
	, "lastLoginAt"
	, "fullName"
	, "tenantCode"
	, "role"
	, "invitationState"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (u *Users) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "email"
	, "password"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "passwordSetAt"
	, "secretCode"
	, "secretCodeAt"
	, "verificationSentAt"
	, "verifiedAt"
	, "lastLoginAt"
	, "fullName"
	, "tenantCode"
	, "role"
	, "invitationState"
	`
}

// ToUpdateArray generate slice of update command
func (u *Users) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, u.Id},
		A.X{`=`, 1, u.Email},
		A.X{`=`, 2, u.Password},
		A.X{`=`, 3, u.CreatedAt},
		A.X{`=`, 4, u.CreatedBy},
		A.X{`=`, 5, u.UpdatedAt},
		A.X{`=`, 6, u.UpdatedBy},
		A.X{`=`, 7, u.DeletedAt},
		A.X{`=`, 8, u.PasswordSetAt},
		A.X{`=`, 9, u.SecretCode},
		A.X{`=`, 10, u.SecretCodeAt},
		A.X{`=`, 11, u.VerificationSentAt},
		A.X{`=`, 12, u.VerifiedAt},
		A.X{`=`, 13, u.LastLoginAt},
		A.X{`=`, 14, u.FullName},
		A.X{`=`, 15, u.TenantCode},
		A.X{`=`, 16, u.Role},
		A.X{`=`, 17, u.InvitationState},
	}
}

// IdxId return name of the index
func (u *Users) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (u *Users) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxEmail return name of the index
func (u *Users) IdxEmail() int { //nolint:dupl false positive
	return 1
}

// SqlEmail return name of the column being indexed
func (u *Users) SqlEmail() string { //nolint:dupl false positive
	return `"email"`
}

// IdxPassword return name of the index
func (u *Users) IdxPassword() int { //nolint:dupl false positive
	return 2
}

// SqlPassword return name of the column being indexed
func (u *Users) SqlPassword() string { //nolint:dupl false positive
	return `"password"`
}

// IdxCreatedAt return name of the index
func (u *Users) IdxCreatedAt() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedAt return name of the column being indexed
func (u *Users) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (u *Users) IdxCreatedBy() int { //nolint:dupl false positive
	return 4
}

// SqlCreatedBy return name of the column being indexed
func (u *Users) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (u *Users) IdxUpdatedAt() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedAt return name of the column being indexed
func (u *Users) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (u *Users) IdxUpdatedBy() int { //nolint:dupl false positive
	return 6
}

// SqlUpdatedBy return name of the column being indexed
func (u *Users) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (u *Users) IdxDeletedAt() int { //nolint:dupl false positive
	return 7
}

// SqlDeletedAt return name of the column being indexed
func (u *Users) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxPasswordSetAt return name of the index
func (u *Users) IdxPasswordSetAt() int { //nolint:dupl false positive
	return 8
}

// SqlPasswordSetAt return name of the column being indexed
func (u *Users) SqlPasswordSetAt() string { //nolint:dupl false positive
	return `"passwordSetAt"`
}

// IdxSecretCode return name of the index
func (u *Users) IdxSecretCode() int { //nolint:dupl false positive
	return 9
}

// SqlSecretCode return name of the column being indexed
func (u *Users) SqlSecretCode() string { //nolint:dupl false positive
	return `"secretCode"`
}

// IdxSecretCodeAt return name of the index
func (u *Users) IdxSecretCodeAt() int { //nolint:dupl false positive
	return 10
}

// SqlSecretCodeAt return name of the column being indexed
func (u *Users) SqlSecretCodeAt() string { //nolint:dupl false positive
	return `"secretCodeAt"`
}

// IdxVerificationSentAt return name of the index
func (u *Users) IdxVerificationSentAt() int { //nolint:dupl false positive
	return 11
}

// SqlVerificationSentAt return name of the column being indexed
func (u *Users) SqlVerificationSentAt() string { //nolint:dupl false positive
	return `"verificationSentAt"`
}

// IdxVerifiedAt return name of the index
func (u *Users) IdxVerifiedAt() int { //nolint:dupl false positive
	return 12
}

// SqlVerifiedAt return name of the column being indexed
func (u *Users) SqlVerifiedAt() string { //nolint:dupl false positive
	return `"verifiedAt"`
}

// IdxLastLoginAt return name of the index
func (u *Users) IdxLastLoginAt() int { //nolint:dupl false positive
	return 13
}

// SqlLastLoginAt return name of the column being indexed
func (u *Users) SqlLastLoginAt() string { //nolint:dupl false positive
	return `"lastLoginAt"`
}

// IdxFullName return name of the index
func (u *Users) IdxFullName() int { //nolint:dupl false positive
	return 14
}

// SqlFullName return name of the column being indexed
func (u *Users) SqlFullName() string { //nolint:dupl false positive
	return `"fullName"`
}

// IdxTenantCode return name of the index
func (u *Users) IdxTenantCode() int { //nolint:dupl false positive
	return 15
}

// SqlTenantCode return name of the column being indexed
func (u *Users) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxRole return name of the index
func (u *Users) IdxRole() int { //nolint:dupl false positive
	return 16
}

// SqlRole return name of the column being indexed
func (u *Users) SqlRole() string { //nolint:dupl false positive
	return `"role"`
}

// IdxInvitationState return name of the index
func (u *Users) IdxInvitationState() int { //nolint:dupl false positive
	return 17
}

// SqlInvitationState return name of the column being indexed
func (u *Users) SqlInvitationState() string { //nolint:dupl false positive
	return `"invitationState"`
}

// CensorFields remove sensitive fields for output
func (u *Users) CensorFields() { //nolint:dupl false positive
	u.Password = ``
	u.SecretCode = ``
	u.SecretCodeAt = 0
	}
// ToArray receiver fields to slice
func (u *Users) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if u.Id != 0 {
		id = u.Id
	}
	return A.X{
		id,
		u.Email,              // 1
		u.Password,           // 2
		u.CreatedAt,          // 3
		u.CreatedBy,          // 4
		u.UpdatedAt,          // 5
		u.UpdatedBy,          // 6
		u.DeletedAt,          // 7
		u.PasswordSetAt,      // 8
		u.SecretCode,         // 9
		u.SecretCodeAt,       // 10
		u.VerificationSentAt, // 11
		u.VerifiedAt,         // 12
		u.LastLoginAt,        // 13
		u.FullName,           // 14
		u.TenantCode,         // 15
		u.Role,               // 16
		u.InvitationState,    // 17
	}
}

// FromArray convert slice to receiver fields
func (u *Users) FromArray(a A.X) *Users { //nolint:dupl false positive
	u.Id = X.ToU(a[0])
	u.Email = X.ToS(a[1])
	u.Password = X.ToS(a[2])
	u.CreatedAt = X.ToI(a[3])
	u.CreatedBy = X.ToU(a[4])
	u.UpdatedAt = X.ToI(a[5])
	u.UpdatedBy = X.ToU(a[6])
	u.DeletedAt = X.ToI(a[7])
	u.PasswordSetAt = X.ToI(a[8])
	u.SecretCode = X.ToS(a[9])
	u.SecretCodeAt = X.ToI(a[10])
	u.VerificationSentAt = X.ToI(a[11])
	u.VerifiedAt = X.ToI(a[12])
	u.LastLoginAt = X.ToI(a[13])
	u.FullName = X.ToS(a[14])
	u.TenantCode = X.ToS(a[15])
	u.Role = X.ToS(a[16])
	u.InvitationState = X.ToS(a[17])
	return u
}

// FromUncensoredArray convert slice to receiver fields
func (u *Users) FromUncensoredArray(a A.X) *Users { //nolint:dupl false positive
	u.Id = X.ToU(a[0])
	u.Email = X.ToS(a[1])
	u.CreatedAt = X.ToI(a[3])
	u.CreatedBy = X.ToU(a[4])
	u.UpdatedAt = X.ToI(a[5])
	u.UpdatedBy = X.ToU(a[6])
	u.DeletedAt = X.ToI(a[7])
	u.PasswordSetAt = X.ToI(a[8])
	u.VerificationSentAt = X.ToI(a[11])
	u.VerifiedAt = X.ToI(a[12])
	u.LastLoginAt = X.ToI(a[13])
	u.FullName = X.ToS(a[14])
	u.TenantCode = X.ToS(a[15])
	u.Role = X.ToS(a[16])
	u.InvitationState = X.ToS(a[17])
	return u
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (u *Users) FindOffsetLimit(offset, limit uint32, idx string) []Users { //nolint:dupl false positive
	var rows []Users
	res, err := u.Adapter.Select(u.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Users.FindOffsetLimit failed: `+u.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Users{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (u *Users) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := u.Adapter.Select(u.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Users.FindOffsetLimit failed: `+u.SpaceName()) {
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
func (u *Users) Total() int64 { //nolint:dupl false positive
	rows := u.Adapter.CallBoxSpace(u.SpaceName() + `:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// UsersFieldTypeMap returns key value of field name and key
var UsersFieldTypeMap = map[string]Tt.DataType { //nolint:dupl false positive
	`id`:                 Tt.Unsigned,
	`email`:              Tt.String,
	`password`:           Tt.String,
	`createdAt`:          Tt.Integer,
	`createdBy`:          Tt.Unsigned,
	`updatedAt`:          Tt.Integer,
	`updatedBy`:          Tt.Unsigned,
	`deletedAt`:          Tt.Integer,
	`passwordSetAt`:      Tt.Integer,
	`secretCode`:         Tt.String,
	`secretCodeAt`:       Tt.Integer,
	`verificationSentAt`: Tt.Integer,
	`verifiedAt`:         Tt.Integer,
	`lastLoginAt`:        Tt.Integer,
	`fullName`:           Tt.String,
	`tenantCode`:         Tt.String,
	`role`:               Tt.String,
	`invitationState`:    Tt.String,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

