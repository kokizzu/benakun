package rqAuth

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mAuth"

	"github.com/tarantool/go-tarantool/v2"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Orgs DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqAuth__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqAuth__ORM.GEN.go
type Orgs struct {
	Adapter    *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id         uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantCode string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	Name       string      `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	HeadTitle  string      `json:"headTitle" form:"headTitle" query:"headTitle" long:"headTitle" msg:"headTitle"`
	ParentId   uint64      `json:"parentId,string" form:"parentId" query:"parentId" long:"parentId" msg:"parentId"`
	Children   []any       `json:"children" form:"children" query:"children" long:"children" msg:"children"`
	OrgType    uint64      `json:"orgType" form:"orgType" query:"orgType" long:"orgType" msg:"orgType"`
	CreatedAt  int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy  uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt  int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy  uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt  int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy  uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
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
	res, err := o.Adapter.RetryDo(
		tarantool.NewSelectRequest(o.SpaceName()).
			Index(o.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(o.Id)}),
	)
	if L.IsError(err, `Orgs.FindById failed: `+o.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			o.FromArray(row)
			return true
		}
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
	, "deletedBy"
	, "restoredBy"
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
	, "deletedBy"
	, "restoredBy"
	`
}

// ToUpdateArray generate slice of update command
func (o *Orgs) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, o.Id).
		Assign(1, o.TenantCode).
		Assign(2, o.Name).
		Assign(3, o.HeadTitle).
		Assign(4, o.ParentId).
		Assign(5, o.Children).
		Assign(6, o.OrgType).
		Assign(7, o.CreatedAt).
		Assign(8, o.CreatedBy).
		Assign(9, o.UpdatedAt).
		Assign(10, o.UpdatedBy).
		Assign(11, o.DeletedAt).
		Assign(12, o.DeletedBy).
		Assign(13, o.RestoredBy)
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

// IdxDeletedBy return name of the index
func (o *Orgs) IdxDeletedBy() int { //nolint:dupl false positive
	return 12
}

// SqlDeletedBy return name of the column being indexed
func (o *Orgs) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (o *Orgs) IdxRestoredBy() int { //nolint:dupl false positive
	return 13
}

// SqlRestoredBy return name of the column being indexed
func (o *Orgs) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
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
		o.DeletedBy,  // 12
		o.RestoredBy, // 13
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
	o.DeletedBy = X.ToU(a[12])
	o.RestoredBy = X.ToU(a[13])
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
	o.DeletedBy = X.ToU(a[12])
	o.RestoredBy = X.ToU(a[13])
	return o
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (o *Orgs) FindOffsetLimit(offset, limit uint32, idx string) []Orgs { //nolint:dupl false positive
	var rows []Orgs
	res, err := o.Adapter.RetryDo(
		tarantool.NewSelectRequest(o.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Orgs.FindOffsetLimit failed: `+o.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := Orgs{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (o *Orgs) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := o.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(o.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Orgs.FindOffsetLimit failed: `+o.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `Orgs.FindOffsetLimit failed: `+o.SpaceName()) {
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
func (o *Orgs) Total() int64 { //nolint:dupl false positive
	rows := o.Adapter.CallBoxSpace(o.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// OrgsFieldTypeMap returns key value of field name and key
var OrgsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
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
	`deletedBy`:  Tt.Unsigned,
	`restoredBy`: Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Sessions DAO reader/query struct
type Sessions struct {
	Adapter      *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	SessionToken string      `json:"sessionToken" form:"sessionToken" query:"sessionToken" long:"sessionToken" msg:"sessionToken"`
	UserId       uint64      `json:"userId,string" form:"userId" query:"userId" long:"userId" msg:"userId"`
	ExpiredAt    int64       `json:"expiredAt" form:"expiredAt" query:"expiredAt" long:"expiredAt" msg:"expiredAt"`
	Device       string      `json:"device" form:"device" query:"device" long:"device" msg:"device"`
	LoginAt      int64       `json:"loginAt" form:"loginAt" query:"loginAt" long:"loginAt" msg:"loginAt"`
	LoginIPs     string      `json:"loginIPs" form:"loginIPs" query:"loginIPs" long:"loginIPs" msg:"loginIPs"`
	TenantCode   string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
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
	res, err := s.Adapter.RetryDo(
		tarantool.NewSelectRequest(s.SpaceName()).
			Index(s.UniqueIndexSessionToken()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.StringKey{S: s.SessionToken}),
	)
	if L.IsError(err, `Sessions.FindBySessionToken failed: `+s.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			s.FromArray(row)
			return true
		}
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
func (s *Sessions) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, s.SessionToken).
		Assign(1, s.UserId).
		Assign(2, s.ExpiredAt).
		Assign(3, s.Device).
		Assign(4, s.LoginAt).
		Assign(5, s.LoginIPs).
		Assign(6, s.TenantCode)
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
	res, err := s.Adapter.RetryDo(
		tarantool.NewSelectRequest(s.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Sessions.FindOffsetLimit failed: `+s.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := Sessions{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (s *Sessions) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := s.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(s.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Sessions.FindOffsetLimit failed: `+s.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `Sessions.FindOffsetLimit failed: `+s.SpaceName()) {
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
func (s *Sessions) Total() int64 { //nolint:dupl false positive
	rows := s.Adapter.CallBoxSpace(s.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// SessionsFieldTypeMap returns key value of field name and key
var SessionsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
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
	Adapter        *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id             uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantCode     string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	CreatedAt      int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy      uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt      int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy      uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt      int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	ProductsCoaId  uint64      `json:"productsCoaId,string" form:"productsCoaId" query:"productsCoaId" long:"productsCoaId" msg:"productsCoaId"`
	SuppliersCoaId uint64      `json:"suppliersCoaId,string" form:"suppliersCoaId" query:"suppliersCoaId" long:"suppliersCoaId" msg:"suppliersCoaId"`
	CustomersCoaId uint64      `json:"customersCoaId,string" form:"customersCoaId" query:"customersCoaId" long:"customersCoaId" msg:"customersCoaId"`
	StaffsCoaId    uint64      `json:"staffsCoaId,string" form:"staffsCoaId" query:"staffsCoaId" long:"staffsCoaId" msg:"staffsCoaId"`
	BanksCoaId     uint64      `json:"banksCoaId,string" form:"banksCoaId" query:"banksCoaId" long:"banksCoaId" msg:"banksCoaId"`
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
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(t.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(t.Id)}),
	)
	if L.IsError(err, `Tenants.FindById failed: `+t.SpaceName()) {
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

// UniqueIndexTenantCode return unique index name
func (t *Tenants) UniqueIndexTenantCode() string { //nolint:dupl false positive
	return `tenantCode`
}

// FindByTenantCode Find one by TenantCode
func (t *Tenants) FindByTenantCode() bool { //nolint:dupl false positive
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(t.UniqueIndexTenantCode()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.StringKey{S: t.TenantCode}),
	)
	if L.IsError(err, `Tenants.FindByTenantCode failed: `+t.SpaceName()) {
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
func (t *Tenants) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "productsCoaId"
	, "suppliersCoaId"
	, "customersCoaId"
	, "staffsCoaId"
	, "banksCoaId"
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
	, "productsCoaId"
	, "suppliersCoaId"
	, "customersCoaId"
	, "staffsCoaId"
	, "banksCoaId"
	`
}

// ToUpdateArray generate slice of update command
func (t *Tenants) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, t.Id).
		Assign(1, t.TenantCode).
		Assign(2, t.CreatedAt).
		Assign(3, t.CreatedBy).
		Assign(4, t.UpdatedAt).
		Assign(5, t.UpdatedBy).
		Assign(6, t.DeletedAt).
		Assign(7, t.ProductsCoaId).
		Assign(8, t.SuppliersCoaId).
		Assign(9, t.CustomersCoaId).
		Assign(10, t.StaffsCoaId).
		Assign(11, t.BanksCoaId)
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

// IdxProductsCoaId return name of the index
func (t *Tenants) IdxProductsCoaId() int { //nolint:dupl false positive
	return 7
}

// SqlProductsCoaId return name of the column being indexed
func (t *Tenants) SqlProductsCoaId() string { //nolint:dupl false positive
	return `"productsCoaId"`
}

// IdxSuppliersCoaId return name of the index
func (t *Tenants) IdxSuppliersCoaId() int { //nolint:dupl false positive
	return 8
}

// SqlSuppliersCoaId return name of the column being indexed
func (t *Tenants) SqlSuppliersCoaId() string { //nolint:dupl false positive
	return `"suppliersCoaId"`
}

// IdxCustomersCoaId return name of the index
func (t *Tenants) IdxCustomersCoaId() int { //nolint:dupl false positive
	return 9
}

// SqlCustomersCoaId return name of the column being indexed
func (t *Tenants) SqlCustomersCoaId() string { //nolint:dupl false positive
	return `"customersCoaId"`
}

// IdxStaffsCoaId return name of the index
func (t *Tenants) IdxStaffsCoaId() int { //nolint:dupl false positive
	return 10
}

// SqlStaffsCoaId return name of the column being indexed
func (t *Tenants) SqlStaffsCoaId() string { //nolint:dupl false positive
	return `"staffsCoaId"`
}

// IdxBanksCoaId return name of the index
func (t *Tenants) IdxBanksCoaId() int { //nolint:dupl false positive
	return 11
}

// SqlBanksCoaId return name of the column being indexed
func (t *Tenants) SqlBanksCoaId() string { //nolint:dupl false positive
	return `"banksCoaId"`
}

// ToArray receiver fields to slice
func (t *Tenants) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if t.Id != 0 {
		id = t.Id
	}
	return A.X{
		id,
		t.TenantCode,     // 1
		t.CreatedAt,      // 2
		t.CreatedBy,      // 3
		t.UpdatedAt,      // 4
		t.UpdatedBy,      // 5
		t.DeletedAt,      // 6
		t.ProductsCoaId,  // 7
		t.SuppliersCoaId, // 8
		t.CustomersCoaId, // 9
		t.StaffsCoaId,    // 10
		t.BanksCoaId,     // 11
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
	t.ProductsCoaId = X.ToU(a[7])
	t.SuppliersCoaId = X.ToU(a[8])
	t.CustomersCoaId = X.ToU(a[9])
	t.StaffsCoaId = X.ToU(a[10])
	t.BanksCoaId = X.ToU(a[11])
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
	t.ProductsCoaId = X.ToU(a[7])
	t.SuppliersCoaId = X.ToU(a[8])
	t.CustomersCoaId = X.ToU(a[9])
	t.StaffsCoaId = X.ToU(a[10])
	t.BanksCoaId = X.ToU(a[11])
	return t
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (t *Tenants) FindOffsetLimit(offset, limit uint32, idx string) []Tenants { //nolint:dupl false positive
	var rows []Tenants
	res, err := t.Adapter.RetryDo(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Tenants.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := Tenants{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (t *Tenants) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := t.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(t.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Tenants.FindOffsetLimit failed: `+t.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `Tenants.FindOffsetLimit failed: `+t.SpaceName()) {
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
func (t *Tenants) Total() int64 { //nolint:dupl false positive
	rows := t.Adapter.CallBoxSpace(t.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// TenantsFieldTypeMap returns key value of field name and key
var TenantsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:             Tt.Unsigned,
	`tenantCode`:     Tt.String,
	`createdAt`:      Tt.Integer,
	`createdBy`:      Tt.Unsigned,
	`updatedAt`:      Tt.Integer,
	`updatedBy`:      Tt.Unsigned,
	`deletedAt`:      Tt.Integer,
	`productsCoaId`:  Tt.Unsigned,
	`suppliersCoaId`: Tt.Unsigned,
	`customersCoaId`: Tt.Unsigned,
	`staffsCoaId`:    Tt.Unsigned,
	`banksCoaId`:     Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Users DAO reader/query struct
type Users struct {
	Adapter            *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id                 uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	Email              string      `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	Password           string      `json:"password" form:"password" query:"password" long:"password" msg:"password"`
	CreatedAt          int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy          uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt          int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy          uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt          int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	PasswordSetAt      int64       `json:"passwordSetAt" form:"passwordSetAt" query:"passwordSetAt" long:"passwordSetAt" msg:"passwordSetAt"`
	SecretCode         string      `json:"secretCode" form:"secretCode" query:"secretCode" long:"secretCode" msg:"secretCode"`
	SecretCodeAt       int64       `json:"secretCodeAt" form:"secretCodeAt" query:"secretCodeAt" long:"secretCodeAt" msg:"secretCodeAt"`
	VerificationSentAt int64       `json:"verificationSentAt" form:"verificationSentAt" query:"verificationSentAt" long:"verificationSentAt" msg:"verificationSentAt"`
	VerifiedAt         int64       `json:"verifiedAt" form:"verifiedAt" query:"verifiedAt" long:"verifiedAt" msg:"verifiedAt"`
	LastLoginAt        int64       `json:"lastLoginAt" form:"lastLoginAt" query:"lastLoginAt" long:"lastLoginAt" msg:"lastLoginAt"`
	FullName           string      `json:"fullName" form:"fullName" query:"fullName" long:"fullName" msg:"fullName"`
	TenantCode         string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	Role               string      `json:"role" form:"role" query:"role" long:"role" msg:"role"`
	InvitationState    string      `json:"invitationState" form:"invitationState" query:"invitationState" long:"invitationState" msg:"invitationState"`
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
	res, err := u.Adapter.RetryDo(
		tarantool.NewSelectRequest(u.SpaceName()).
			Index(u.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(u.Id)}),
	)
	if L.IsError(err, `Users.FindById failed: `+u.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			u.FromArray(row)
			return true
		}
	}
	return false
}

// UniqueIndexEmail return unique index name
func (u *Users) UniqueIndexEmail() string { //nolint:dupl false positive
	return `email`
}

// FindByEmail Find one by Email
func (u *Users) FindByEmail() bool { //nolint:dupl false positive
	res, err := u.Adapter.RetryDo(
		tarantool.NewSelectRequest(u.SpaceName()).
			Index(u.UniqueIndexEmail()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.StringKey{S: u.Email}),
	)
	if L.IsError(err, `Users.FindByEmail failed: `+u.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			u.FromArray(row)
			return true
		}
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
func (u *Users) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, u.Id).
		Assign(1, u.Email).
		Assign(2, u.Password).
		Assign(3, u.CreatedAt).
		Assign(4, u.CreatedBy).
		Assign(5, u.UpdatedAt).
		Assign(6, u.UpdatedBy).
		Assign(7, u.DeletedAt).
		Assign(8, u.PasswordSetAt).
		Assign(9, u.SecretCode).
		Assign(10, u.SecretCodeAt).
		Assign(11, u.VerificationSentAt).
		Assign(12, u.VerifiedAt).
		Assign(13, u.LastLoginAt).
		Assign(14, u.FullName).
		Assign(15, u.TenantCode).
		Assign(16, u.Role).
		Assign(17, u.InvitationState)
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
	res, err := u.Adapter.RetryDo(
		tarantool.NewSelectRequest(u.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Users.FindOffsetLimit failed: `+u.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := Users{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (u *Users) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := u.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(u.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Users.FindOffsetLimit failed: `+u.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `Users.FindOffsetLimit failed: `+u.SpaceName()) {
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
func (u *Users) Total() int64 { //nolint:dupl false positive
	rows := u.Adapter.CallBoxSpace(u.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// UsersFieldTypeMap returns key value of field name and key
var UsersFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
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
