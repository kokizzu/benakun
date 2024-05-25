package rqBusiness

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mBusiness"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Products DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqBusiness__ORM.GEN.go
type Products struct {
	Adapter    *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id         uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	TenantCode string      `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
	CreatedAt  int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy  uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt  int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy  uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt  int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy  uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	Name       string      `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	Detail     string      `json:"detail" form:"detail" query:"detail" long:"detail" msg:"detail"`
	Rule       string      `json:"rule" form:"rule" query:"rule" long:"rule" msg:"rule"`
	Kind       string      `json:"kind" form:"kind" query:"kind" long:"kind" msg:"kind"`
	CogsIDR    int64       `json:"cogsIDR" form:"cogsIDR" query:"cogsIDR" long:"cogsIDR" msg:"cogsIDR"`
}

// NewProducts create new ORM reader/query object
func NewProducts(adapter *Tt.Adapter) *Products {
	return &Products{Adapter: adapter}
}

// SpaceName returns full package and table name
func (p *Products) SpaceName() string { //nolint:dupl false positive
	return string(mBusiness.TableProducts) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (p *Products) SqlTableName() string { //nolint:dupl false positive
	return `"products"`
}

func (p *Products) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (p *Products) FindById() bool { //nolint:dupl false positive
	res, err := p.Adapter.Select(p.SpaceName(), p.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{p.Id})
	if L.IsError(err, `Products.FindById failed: `+p.SpaceName()) {
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
func (p *Products) SqlSelectAllFields() string { //nolint:dupl false positive
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
	, "detail"
	, "rule"
	, "kind"
	, "cogsIDR"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *Products) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
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
	, "detail"
	, "rule"
	, "kind"
	, "cogsIDR"
	`
}

// ToUpdateArray generate slice of update command
func (p *Products) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, p.Id},
		A.X{`=`, 1, p.TenantCode},
		A.X{`=`, 2, p.CreatedAt},
		A.X{`=`, 3, p.CreatedBy},
		A.X{`=`, 4, p.UpdatedAt},
		A.X{`=`, 5, p.UpdatedBy},
		A.X{`=`, 6, p.DeletedAt},
		A.X{`=`, 7, p.DeletedBy},
		A.X{`=`, 8, p.RestoredBy},
		A.X{`=`, 9, p.Name},
		A.X{`=`, 10, p.Detail},
		A.X{`=`, 11, p.Rule},
		A.X{`=`, 12, p.Kind},
		A.X{`=`, 13, p.CogsIDR},
	}
}

// IdxId return name of the index
func (p *Products) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (p *Products) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (p *Products) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (p *Products) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxCreatedAt return name of the index
func (p *Products) IdxCreatedAt() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedAt return name of the column being indexed
func (p *Products) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (p *Products) IdxCreatedBy() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedBy return name of the column being indexed
func (p *Products) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (p *Products) IdxUpdatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlUpdatedAt return name of the column being indexed
func (p *Products) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (p *Products) IdxUpdatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedBy return name of the column being indexed
func (p *Products) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (p *Products) IdxDeletedAt() int { //nolint:dupl false positive
	return 6
}

// SqlDeletedAt return name of the column being indexed
func (p *Products) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (p *Products) IdxDeletedBy() int { //nolint:dupl false positive
	return 7
}

// SqlDeletedBy return name of the column being indexed
func (p *Products) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (p *Products) IdxRestoredBy() int { //nolint:dupl false positive
	return 8
}

// SqlRestoredBy return name of the column being indexed
func (p *Products) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxName return name of the index
func (p *Products) IdxName() int { //nolint:dupl false positive
	return 9
}

// SqlName return name of the column being indexed
func (p *Products) SqlName() string { //nolint:dupl false positive
	return `"name"`
}

// IdxDetail return name of the index
func (p *Products) IdxDetail() int { //nolint:dupl false positive
	return 10
}

// SqlDetail return name of the column being indexed
func (p *Products) SqlDetail() string { //nolint:dupl false positive
	return `"detail"`
}

// IdxRule return name of the index
func (p *Products) IdxRule() int { //nolint:dupl false positive
	return 11
}

// SqlRule return name of the column being indexed
func (p *Products) SqlRule() string { //nolint:dupl false positive
	return `"rule"`
}

// IdxKind return name of the index
func (p *Products) IdxKind() int { //nolint:dupl false positive
	return 12
}

// SqlKind return name of the column being indexed
func (p *Products) SqlKind() string { //nolint:dupl false positive
	return `"kind"`
}

// IdxCogsIDR return name of the index
func (p *Products) IdxCogsIDR() int { //nolint:dupl false positive
	return 13
}

// SqlCogsIDR return name of the column being indexed
func (p *Products) SqlCogsIDR() string { //nolint:dupl false positive
	return `"cogsIDR"`
}

// ToArray receiver fields to slice
func (p *Products) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.TenantCode, // 1
		p.CreatedAt,  // 2
		p.CreatedBy,  // 3
		p.UpdatedAt,  // 4
		p.UpdatedBy,  // 5
		p.DeletedAt,  // 6
		p.DeletedBy,  // 7
		p.RestoredBy, // 8
		p.Name,       // 9
		p.Detail,     // 10
		p.Rule,       // 11
		p.Kind,       // 12
		p.CogsIDR,    // 13
	}
}

// FromArray convert slice to receiver fields
func (p *Products) FromArray(a A.X) *Products { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.TenantCode = X.ToS(a[1])
	p.CreatedAt = X.ToI(a[2])
	p.CreatedBy = X.ToU(a[3])
	p.UpdatedAt = X.ToI(a[4])
	p.UpdatedBy = X.ToU(a[5])
	p.DeletedAt = X.ToI(a[6])
	p.DeletedBy = X.ToU(a[7])
	p.RestoredBy = X.ToU(a[8])
	p.Name = X.ToS(a[9])
	p.Detail = X.ToS(a[10])
	p.Rule = X.ToS(a[11])
	p.Kind = X.ToS(a[12])
	p.CogsIDR = X.ToI(a[13])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *Products) FromUncensoredArray(a A.X) *Products { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.TenantCode = X.ToS(a[1])
	p.CreatedAt = X.ToI(a[2])
	p.CreatedBy = X.ToU(a[3])
	p.UpdatedAt = X.ToI(a[4])
	p.UpdatedBy = X.ToU(a[5])
	p.DeletedAt = X.ToI(a[6])
	p.DeletedBy = X.ToU(a[7])
	p.RestoredBy = X.ToU(a[8])
	p.Name = X.ToS(a[9])
	p.Detail = X.ToS(a[10])
	p.Rule = X.ToS(a[11])
	p.Kind = X.ToS(a[12])
	p.CogsIDR = X.ToI(a[13])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *Products) FindOffsetLimit(offset, limit uint32, idx string) []Products { //nolint:dupl false positive
	var rows []Products
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Products.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := Products{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *Products) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := p.Adapter.Select(p.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `Products.FindOffsetLimit failed: `+p.SpaceName()) {
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
func (p *Products) Total() int64 { //nolint:dupl false positive
	rows := p.Adapter.CallBoxSpace(p.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// ProductsFieldTypeMap returns key value of field name and key
var ProductsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:         Tt.Unsigned,
	`tenantCode`: Tt.String,
	`createdAt`:  Tt.Integer,
	`createdBy`:  Tt.Unsigned,
	`updatedAt`:  Tt.Integer,
	`updatedBy`:  Tt.Unsigned,
	`deletedAt`:  Tt.Integer,
	`deletedBy`:  Tt.Unsigned,
	`restoredBy`: Tt.Unsigned,
	`name`:       Tt.String,
	`detail`:     Tt.String,
	`rule`:       Tt.String,
	`kind`:       Tt.String,
	`cogsIDR`:    Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
