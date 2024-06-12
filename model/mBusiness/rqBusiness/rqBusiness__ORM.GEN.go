package rqBusiness

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mBusiness"

	"github.com/tarantool/go-tarantool/v2"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// InventoryChanges DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqBusiness__ORM.GEN.go
type InventoryChanges struct {
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
	StockDelta uint64      `json:"stockDelta" form:"stockDelta" query:"stockDelta" long:"stockDelta" msg:"stockDelta"`
	ProductId  uint64      `json:"productId,string" form:"productId" query:"productId" long:"productId" msg:"productId"`
	LocationId uint64      `json:"locationId,string" form:"locationId" query:"locationId" long:"locationId" msg:"locationId"`
	SpendingId uint64      `json:"spendingId,string" form:"spendingId" query:"spendingId" long:"spendingId" msg:"spendingId"`
	ExpenseId  uint64      `json:"expenseId,string" form:"expenseId" query:"expenseId" long:"expenseId" msg:"expenseId"`
}

// NewInventoryChanges create new ORM reader/query object
func NewInventoryChanges(adapter *Tt.Adapter) *InventoryChanges {
	return &InventoryChanges{Adapter: adapter}
}

// SpaceName returns full package and table name
func (i *InventoryChanges) SpaceName() string { //nolint:dupl false positive
	return string(mBusiness.TableInventoryChanges) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (i *InventoryChanges) SqlTableName() string { //nolint:dupl false positive
	return `"inventoryChanges"`
}

func (i *InventoryChanges) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (i *InventoryChanges) FindById() bool { //nolint:dupl false positive
	res, err := i.Adapter.RetryDo(
		tarantool.NewSelectRequest(i.SpaceName()).
			Index(i.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(i.Id)}),
	)
	if L.IsError(err, `InventoryChanges.FindById failed: `+i.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			i.FromArray(row)
			return true
		}
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (i *InventoryChanges) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "stockDelta"
	, "productId"
	, "locationId"
	, "spendingId"
	, "expenseId"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (i *InventoryChanges) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "tenantCode"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "stockDelta"
	, "productId"
	, "locationId"
	, "spendingId"
	, "expenseId"
	`
}

// ToUpdateArray generate slice of update command
func (i *InventoryChanges) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, i.Id).
		Assign(1, i.TenantCode).
		Assign(2, i.CreatedAt).
		Assign(3, i.CreatedBy).
		Assign(4, i.UpdatedAt).
		Assign(5, i.UpdatedBy).
		Assign(6, i.DeletedAt).
		Assign(7, i.DeletedBy).
		Assign(8, i.RestoredBy).
		Assign(9, i.StockDelta).
		Assign(10, i.ProductId).
		Assign(11, i.LocationId).
		Assign(12, i.SpendingId).
		Assign(13, i.ExpenseId)
}

// IdxId return name of the index
func (i *InventoryChanges) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (i *InventoryChanges) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (i *InventoryChanges) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (i *InventoryChanges) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxCreatedAt return name of the index
func (i *InventoryChanges) IdxCreatedAt() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedAt return name of the column being indexed
func (i *InventoryChanges) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (i *InventoryChanges) IdxCreatedBy() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedBy return name of the column being indexed
func (i *InventoryChanges) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (i *InventoryChanges) IdxUpdatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlUpdatedAt return name of the column being indexed
func (i *InventoryChanges) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (i *InventoryChanges) IdxUpdatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedBy return name of the column being indexed
func (i *InventoryChanges) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (i *InventoryChanges) IdxDeletedAt() int { //nolint:dupl false positive
	return 6
}

// SqlDeletedAt return name of the column being indexed
func (i *InventoryChanges) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (i *InventoryChanges) IdxDeletedBy() int { //nolint:dupl false positive
	return 7
}

// SqlDeletedBy return name of the column being indexed
func (i *InventoryChanges) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (i *InventoryChanges) IdxRestoredBy() int { //nolint:dupl false positive
	return 8
}

// SqlRestoredBy return name of the column being indexed
func (i *InventoryChanges) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxStockDelta return name of the index
func (i *InventoryChanges) IdxStockDelta() int { //nolint:dupl false positive
	return 9
}

// SqlStockDelta return name of the column being indexed
func (i *InventoryChanges) SqlStockDelta() string { //nolint:dupl false positive
	return `"stockDelta"`
}

// IdxProductId return name of the index
func (i *InventoryChanges) IdxProductId() int { //nolint:dupl false positive
	return 10
}

// SqlProductId return name of the column being indexed
func (i *InventoryChanges) SqlProductId() string { //nolint:dupl false positive
	return `"productId"`
}

// IdxLocationId return name of the index
func (i *InventoryChanges) IdxLocationId() int { //nolint:dupl false positive
	return 11
}

// SqlLocationId return name of the column being indexed
func (i *InventoryChanges) SqlLocationId() string { //nolint:dupl false positive
	return `"locationId"`
}

// IdxSpendingId return name of the index
func (i *InventoryChanges) IdxSpendingId() int { //nolint:dupl false positive
	return 12
}

// SqlSpendingId return name of the column being indexed
func (i *InventoryChanges) SqlSpendingId() string { //nolint:dupl false positive
	return `"spendingId"`
}

// IdxExpenseId return name of the index
func (i *InventoryChanges) IdxExpenseId() int { //nolint:dupl false positive
	return 13
}

// SqlExpenseId return name of the column being indexed
func (i *InventoryChanges) SqlExpenseId() string { //nolint:dupl false positive
	return `"expenseId"`
}

// ToArray receiver fields to slice
func (i *InventoryChanges) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if i.Id != 0 {
		id = i.Id
	}
	return A.X{
		id,
		i.TenantCode, // 1
		i.CreatedAt,  // 2
		i.CreatedBy,  // 3
		i.UpdatedAt,  // 4
		i.UpdatedBy,  // 5
		i.DeletedAt,  // 6
		i.DeletedBy,  // 7
		i.RestoredBy, // 8
		i.StockDelta, // 9
		i.ProductId,  // 10
		i.LocationId, // 11
		i.SpendingId, // 12
		i.ExpenseId,  // 13
	}
}

// FromArray convert slice to receiver fields
func (i *InventoryChanges) FromArray(a A.X) *InventoryChanges { //nolint:dupl false positive
	i.Id = X.ToU(a[0])
	i.TenantCode = X.ToS(a[1])
	i.CreatedAt = X.ToI(a[2])
	i.CreatedBy = X.ToU(a[3])
	i.UpdatedAt = X.ToI(a[4])
	i.UpdatedBy = X.ToU(a[5])
	i.DeletedAt = X.ToI(a[6])
	i.DeletedBy = X.ToU(a[7])
	i.RestoredBy = X.ToU(a[8])
	i.StockDelta = X.ToU(a[9])
	i.ProductId = X.ToU(a[10])
	i.LocationId = X.ToU(a[11])
	i.SpendingId = X.ToU(a[12])
	i.ExpenseId = X.ToU(a[13])
	return i
}

// FromUncensoredArray convert slice to receiver fields
func (i *InventoryChanges) FromUncensoredArray(a A.X) *InventoryChanges { //nolint:dupl false positive
	i.Id = X.ToU(a[0])
	i.TenantCode = X.ToS(a[1])
	i.CreatedAt = X.ToI(a[2])
	i.CreatedBy = X.ToU(a[3])
	i.UpdatedAt = X.ToI(a[4])
	i.UpdatedBy = X.ToU(a[5])
	i.DeletedAt = X.ToI(a[6])
	i.DeletedBy = X.ToU(a[7])
	i.RestoredBy = X.ToU(a[8])
	i.StockDelta = X.ToU(a[9])
	i.ProductId = X.ToU(a[10])
	i.LocationId = X.ToU(a[11])
	i.SpendingId = X.ToU(a[12])
	i.ExpenseId = X.ToU(a[13])
	return i
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (i *InventoryChanges) FindOffsetLimit(offset, limit uint32, idx string) []InventoryChanges { //nolint:dupl false positive
	var rows []InventoryChanges
	res, err := i.Adapter.RetryDo(
		tarantool.NewSelectRequest(i.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `InventoryChanges.FindOffsetLimit failed: `+i.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := InventoryChanges{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (i *InventoryChanges) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := i.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(i.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `InventoryChanges.FindOffsetLimit failed: `+i.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `InventoryChanges.FindOffsetLimit failed: `+i.SpaceName()) {
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
func (i *InventoryChanges) Total() int64 { //nolint:dupl false positive
	rows := i.Adapter.CallBoxSpace(i.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// InventoryChangesFieldTypeMap returns key value of field name and key
var InventoryChangesFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:         Tt.Unsigned,
	`tenantCode`: Tt.String,
	`createdAt`:  Tt.Integer,
	`createdBy`:  Tt.Unsigned,
	`updatedAt`:  Tt.Integer,
	`updatedBy`:  Tt.Unsigned,
	`deletedAt`:  Tt.Integer,
	`deletedBy`:  Tt.Unsigned,
	`restoredBy`: Tt.Unsigned,
	`stockDelta`: Tt.Unsigned,
	`productId`:  Tt.Unsigned,
	`locationId`: Tt.Unsigned,
	`spendingId`: Tt.Unsigned,
	`expenseId`:  Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Locations DAO reader/query struct
type Locations struct {
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
	Name         string      `json:"name" form:"name" query:"name" long:"name" msg:"name"`
	Country      string      `json:"country" form:"country" query:"country" long:"country" msg:"country"`
	StateProvice string      `json:"stateProvice" form:"stateProvice" query:"stateProvice" long:"stateProvice" msg:"stateProvice"`
	CityRegency  string      `json:"cityRegency" form:"cityRegency" query:"cityRegency" long:"cityRegency" msg:"cityRegency"`
	Subdistrict  string      `json:"subdistrict" form:"subdistrict" query:"subdistrict" long:"subdistrict" msg:"subdistrict"`
	Village      string      `json:"village" form:"village" query:"village" long:"village" msg:"village"`
	RwBanjar     string      `json:"rwBanjar" form:"rwBanjar" query:"rwBanjar" long:"rwBanjar" msg:"rwBanjar"`
	RtNeigb      string      `json:"rtNeigb" form:"rtNeigb" query:"rtNeigb" long:"rtNeigb" msg:"rtNeigb"`
	Address      string      `json:"address" form:"address" query:"address" long:"address" msg:"address"`
	Lat          float64     `json:"lat" form:"lat" query:"lat" long:"lat" msg:"lat"`
	Lng          float64     `json:"lng" form:"lng" query:"lng" long:"lng" msg:"lng"`
}

// NewLocations create new ORM reader/query object
func NewLocations(adapter *Tt.Adapter) *Locations {
	return &Locations{Adapter: adapter}
}

// SpaceName returns full package and table name
func (l *Locations) SpaceName() string { //nolint:dupl false positive
	return string(mBusiness.TableLocations) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (l *Locations) SqlTableName() string { //nolint:dupl false positive
	return `"locations"`
}

func (l *Locations) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (l *Locations) FindById() bool { //nolint:dupl false positive
	res, err := l.Adapter.RetryDo(
		tarantool.NewSelectRequest(l.SpaceName()).
			Index(l.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(l.Id)}),
	)
	if L.IsError(err, `Locations.FindById failed: `+l.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			l.FromArray(row)
			return true
		}
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (l *Locations) SqlSelectAllFields() string { //nolint:dupl false positive
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
	, "country"
	, "stateProvice"
	, "cityRegency"
	, "subdistrict"
	, "village"
	, "rwBanjar"
	, "rtNeigb"
	, "address"
	, "lat"
	, "lng"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (l *Locations) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
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
	, "country"
	, "stateProvice"
	, "cityRegency"
	, "subdistrict"
	, "village"
	, "rwBanjar"
	, "rtNeigb"
	, "address"
	, "lat"
	, "lng"
	`
}

// ToUpdateArray generate slice of update command
func (l *Locations) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, l.Id).
		Assign(1, l.TenantCode).
		Assign(2, l.CreatedAt).
		Assign(3, l.CreatedBy).
		Assign(4, l.UpdatedAt).
		Assign(5, l.UpdatedBy).
		Assign(6, l.DeletedAt).
		Assign(7, l.DeletedBy).
		Assign(8, l.RestoredBy).
		Assign(9, l.Name).
		Assign(10, l.Country).
		Assign(11, l.StateProvice).
		Assign(12, l.CityRegency).
		Assign(13, l.Subdistrict).
		Assign(14, l.Village).
		Assign(15, l.RwBanjar).
		Assign(16, l.RtNeigb).
		Assign(17, l.Address).
		Assign(18, l.Lat).
		Assign(19, l.Lng)
}

// IdxId return name of the index
func (l *Locations) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (l *Locations) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxTenantCode return name of the index
func (l *Locations) IdxTenantCode() int { //nolint:dupl false positive
	return 1
}

// SqlTenantCode return name of the column being indexed
func (l *Locations) SqlTenantCode() string { //nolint:dupl false positive
	return `"tenantCode"`
}

// IdxCreatedAt return name of the index
func (l *Locations) IdxCreatedAt() int { //nolint:dupl false positive
	return 2
}

// SqlCreatedAt return name of the column being indexed
func (l *Locations) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (l *Locations) IdxCreatedBy() int { //nolint:dupl false positive
	return 3
}

// SqlCreatedBy return name of the column being indexed
func (l *Locations) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (l *Locations) IdxUpdatedAt() int { //nolint:dupl false positive
	return 4
}

// SqlUpdatedAt return name of the column being indexed
func (l *Locations) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (l *Locations) IdxUpdatedBy() int { //nolint:dupl false positive
	return 5
}

// SqlUpdatedBy return name of the column being indexed
func (l *Locations) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (l *Locations) IdxDeletedAt() int { //nolint:dupl false positive
	return 6
}

// SqlDeletedAt return name of the column being indexed
func (l *Locations) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (l *Locations) IdxDeletedBy() int { //nolint:dupl false positive
	return 7
}

// SqlDeletedBy return name of the column being indexed
func (l *Locations) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (l *Locations) IdxRestoredBy() int { //nolint:dupl false positive
	return 8
}

// SqlRestoredBy return name of the column being indexed
func (l *Locations) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxName return name of the index
func (l *Locations) IdxName() int { //nolint:dupl false positive
	return 9
}

// SqlName return name of the column being indexed
func (l *Locations) SqlName() string { //nolint:dupl false positive
	return `"name"`
}

// IdxCountry return name of the index
func (l *Locations) IdxCountry() int { //nolint:dupl false positive
	return 10
}

// SqlCountry return name of the column being indexed
func (l *Locations) SqlCountry() string { //nolint:dupl false positive
	return `"country"`
}

// IdxStateProvice return name of the index
func (l *Locations) IdxStateProvice() int { //nolint:dupl false positive
	return 11
}

// SqlStateProvice return name of the column being indexed
func (l *Locations) SqlStateProvice() string { //nolint:dupl false positive
	return `"stateProvice"`
}

// IdxCityRegency return name of the index
func (l *Locations) IdxCityRegency() int { //nolint:dupl false positive
	return 12
}

// SqlCityRegency return name of the column being indexed
func (l *Locations) SqlCityRegency() string { //nolint:dupl false positive
	return `"cityRegency"`
}

// IdxSubdistrict return name of the index
func (l *Locations) IdxSubdistrict() int { //nolint:dupl false positive
	return 13
}

// SqlSubdistrict return name of the column being indexed
func (l *Locations) SqlSubdistrict() string { //nolint:dupl false positive
	return `"subdistrict"`
}

// IdxVillage return name of the index
func (l *Locations) IdxVillage() int { //nolint:dupl false positive
	return 14
}

// SqlVillage return name of the column being indexed
func (l *Locations) SqlVillage() string { //nolint:dupl false positive
	return `"village"`
}

// IdxRwBanjar return name of the index
func (l *Locations) IdxRwBanjar() int { //nolint:dupl false positive
	return 15
}

// SqlRwBanjar return name of the column being indexed
func (l *Locations) SqlRwBanjar() string { //nolint:dupl false positive
	return `"rwBanjar"`
}

// IdxRtNeigb return name of the index
func (l *Locations) IdxRtNeigb() int { //nolint:dupl false positive
	return 16
}

// SqlRtNeigb return name of the column being indexed
func (l *Locations) SqlRtNeigb() string { //nolint:dupl false positive
	return `"rtNeigb"`
}

// IdxAddress return name of the index
func (l *Locations) IdxAddress() int { //nolint:dupl false positive
	return 17
}

// SqlAddress return name of the column being indexed
func (l *Locations) SqlAddress() string { //nolint:dupl false positive
	return `"address"`
}

// IdxLat return name of the index
func (l *Locations) IdxLat() int { //nolint:dupl false positive
	return 18
}

// SqlLat return name of the column being indexed
func (l *Locations) SqlLat() string { //nolint:dupl false positive
	return `"lat"`
}

// IdxLng return name of the index
func (l *Locations) IdxLng() int { //nolint:dupl false positive
	return 19
}

// SqlLng return name of the column being indexed
func (l *Locations) SqlLng() string { //nolint:dupl false positive
	return `"lng"`
}

// ToArray receiver fields to slice
func (l *Locations) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if l.Id != 0 {
		id = l.Id
	}
	return A.X{
		id,
		l.TenantCode,   // 1
		l.CreatedAt,    // 2
		l.CreatedBy,    // 3
		l.UpdatedAt,    // 4
		l.UpdatedBy,    // 5
		l.DeletedAt,    // 6
		l.DeletedBy,    // 7
		l.RestoredBy,   // 8
		l.Name,         // 9
		l.Country,      // 10
		l.StateProvice, // 11
		l.CityRegency,  // 12
		l.Subdistrict,  // 13
		l.Village,      // 14
		l.RwBanjar,     // 15
		l.RtNeigb,      // 16
		l.Address,      // 17
		l.Lat,          // 18
		l.Lng,          // 19
	}
}

// FromArray convert slice to receiver fields
func (l *Locations) FromArray(a A.X) *Locations { //nolint:dupl false positive
	l.Id = X.ToU(a[0])
	l.TenantCode = X.ToS(a[1])
	l.CreatedAt = X.ToI(a[2])
	l.CreatedBy = X.ToU(a[3])
	l.UpdatedAt = X.ToI(a[4])
	l.UpdatedBy = X.ToU(a[5])
	l.DeletedAt = X.ToI(a[6])
	l.DeletedBy = X.ToU(a[7])
	l.RestoredBy = X.ToU(a[8])
	l.Name = X.ToS(a[9])
	l.Country = X.ToS(a[10])
	l.StateProvice = X.ToS(a[11])
	l.CityRegency = X.ToS(a[12])
	l.Subdistrict = X.ToS(a[13])
	l.Village = X.ToS(a[14])
	l.RwBanjar = X.ToS(a[15])
	l.RtNeigb = X.ToS(a[16])
	l.Address = X.ToS(a[17])
	l.Lat = X.ToF(a[18])
	l.Lng = X.ToF(a[19])
	return l
}

// FromUncensoredArray convert slice to receiver fields
func (l *Locations) FromUncensoredArray(a A.X) *Locations { //nolint:dupl false positive
	l.Id = X.ToU(a[0])
	l.TenantCode = X.ToS(a[1])
	l.CreatedAt = X.ToI(a[2])
	l.CreatedBy = X.ToU(a[3])
	l.UpdatedAt = X.ToI(a[4])
	l.UpdatedBy = X.ToU(a[5])
	l.DeletedAt = X.ToI(a[6])
	l.DeletedBy = X.ToU(a[7])
	l.RestoredBy = X.ToU(a[8])
	l.Name = X.ToS(a[9])
	l.Country = X.ToS(a[10])
	l.StateProvice = X.ToS(a[11])
	l.CityRegency = X.ToS(a[12])
	l.Subdistrict = X.ToS(a[13])
	l.Village = X.ToS(a[14])
	l.RwBanjar = X.ToS(a[15])
	l.RtNeigb = X.ToS(a[16])
	l.Address = X.ToS(a[17])
	l.Lat = X.ToF(a[18])
	l.Lng = X.ToF(a[19])
	return l
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (l *Locations) FindOffsetLimit(offset, limit uint32, idx string) []Locations { //nolint:dupl false positive
	var rows []Locations
	res, err := l.Adapter.RetryDo(
		tarantool.NewSelectRequest(l.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Locations.FindOffsetLimit failed: `+l.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := Locations{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (l *Locations) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := l.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(l.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Locations.FindOffsetLimit failed: `+l.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `Locations.FindOffsetLimit failed: `+l.SpaceName()) {
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
func (l *Locations) Total() int64 { //nolint:dupl false positive
	rows := l.Adapter.CallBoxSpace(l.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// LocationsFieldTypeMap returns key value of field name and key
var LocationsFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:           Tt.Unsigned,
	`tenantCode`:   Tt.String,
	`createdAt`:    Tt.Integer,
	`createdBy`:    Tt.Unsigned,
	`updatedAt`:    Tt.Integer,
	`updatedBy`:    Tt.Unsigned,
	`deletedAt`:    Tt.Integer,
	`deletedBy`:    Tt.Unsigned,
	`restoredBy`:   Tt.Unsigned,
	`name`:         Tt.String,
	`country`:      Tt.String,
	`stateProvice`: Tt.String,
	`cityRegency`:  Tt.String,
	`subdistrict`:  Tt.String,
	`village`:      Tt.String,
	`rwBanjar`:     Tt.String,
	`rtNeigb`:      Tt.String,
	`address`:      Tt.String,
	`lat`:          Tt.Double,
	`lng`:          Tt.Double,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// Products DAO reader/query struct
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
	CogsIDR    uint64      `json:"cogsIDR,string" form:"cogsIDR" query:"cogsIDR" long:"cogsIDR" msg:"cogsIDR"`
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
	res, err := p.Adapter.RetryDo(
		tarantool.NewSelectRequest(p.SpaceName()).
			Index(p.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(p.Id)}),
	)
	if L.IsError(err, `Products.FindById failed: `+p.SpaceName()) {
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
func (p *Products) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, p.Id).
		Assign(1, p.TenantCode).
		Assign(2, p.CreatedAt).
		Assign(3, p.CreatedBy).
		Assign(4, p.UpdatedAt).
		Assign(5, p.UpdatedBy).
		Assign(6, p.DeletedAt).
		Assign(7, p.DeletedBy).
		Assign(8, p.RestoredBy).
		Assign(9, p.Name).
		Assign(10, p.Detail).
		Assign(11, p.Rule).
		Assign(12, p.Kind).
		Assign(13, p.CogsIDR)
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
	p.CogsIDR = X.ToU(a[13])
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
	p.CogsIDR = X.ToU(a[13])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *Products) FindOffsetLimit(offset, limit uint32, idx string) []Products { //nolint:dupl false positive
	var rows []Products
	res, err := p.Adapter.RetryDo(
		tarantool.NewSelectRequest(p.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Products.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := Products{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *Products) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := p.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(p.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Products.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `Products.FindOffsetLimit failed: `+p.SpaceName()) {
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
	`cogsIDR`:    Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
