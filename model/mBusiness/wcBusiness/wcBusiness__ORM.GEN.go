package wcBusiness

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mBusiness/rqBusiness"

	"github.com/tarantool/go-tarantool/v2"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// InventoryChangesMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcBusiness__ORM.GEN.go
type InventoryChangesMutator struct {
	rqBusiness.InventoryChanges
	mutations *tarantool.Operations
	logs      []A.X
}

// NewInventoryChangesMutator create new ORM writer/command object
func NewInventoryChangesMutator(adapter *Tt.Adapter) (res *InventoryChangesMutator) {
	res = &InventoryChangesMutator{InventoryChanges: rqBusiness.InventoryChanges{Adapter: adapter}}
	res.mutations = tarantool.NewOperations()
	return
}

// Logs get array of logs [field, old, new]
func (i *InventoryChangesMutator) Logs() []A.X { //nolint:dupl false positive
	return i.logs
}

// HaveMutation check whether Set* methods ever called
func (i *InventoryChangesMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(i.logs) > 0
}

// ClearMutations clear all previously called Set* methods
func (i *InventoryChangesMutator) ClearMutations() { //nolint:dupl false positive
	i.mutations = tarantool.NewOperations()
	i.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (i *InventoryChangesMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := i.Adapter.RetryDo(tarantool.NewUpdateRequest(i.SpaceName()).
		Index(i.UniqueIndexId()).
		Key(tarantool.UintKey{I: uint(i.Id)}).
		Operations(i.ToUpdateArray()),
	)
	return !L.IsError(err, `InventoryChanges.DoOverwriteById failed: `+i.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (i *InventoryChangesMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !i.HaveMutation() {
		return true
	}
	_, err := i.Adapter.RetryDo(
		tarantool.NewUpdateRequest(i.SpaceName()).
			Index(i.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(i.Id)}).
			Operations(i.mutations),
	)
	return !L.IsError(err, `InventoryChanges.DoUpdateById failed: `+i.SpaceName())
}

// DoDeletePermanentById permanent delete
func (i *InventoryChangesMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := i.Adapter.RetryDo(
		tarantool.NewDeleteRequest(i.SpaceName()).
			Index(i.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(i.Id)}),
	)
	return !L.IsError(err, `InventoryChanges.DoDeletePermanentById failed: `+i.SpaceName())
}

// DoInsert insert, error if already exists
func (i *InventoryChangesMutator) DoInsert() bool { //nolint:dupl false positive
	arr := i.ToArray()
	row, err := i.Adapter.RetryDo(
		tarantool.NewInsertRequest(i.SpaceName()).
			Tuple(arr),
	)
	if err == nil {
		if len(row) > 0 {
			if cells, ok := row[0].([]any); ok && len(cells) > 0 {
				i.Id = X.ToU(cells[0])
			}
		}
	}
	return !L.IsError(err, `InventoryChanges.DoInsert failed: `+i.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// tarantool's replace/upsert can only match by primary key
// previous name: DoReplace
func (i *InventoryChangesMutator) DoUpsertById() bool { //nolint:dupl false positive
	if i.Id > 0 {
		return i.DoUpdateById()
	}
	return i.DoInsert()
}

// SetId create mutations, should not duplicate
func (i *InventoryChangesMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != i.Id {
		i.mutations.Assign(0, val)
		i.logs = append(i.logs, A.X{`id`, i.Id, val})
		i.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (i *InventoryChangesMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != i.TenantCode {
		i.mutations.Assign(1, val)
		i.logs = append(i.logs, A.X{`tenantCode`, i.TenantCode, val})
		i.TenantCode = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (i *InventoryChangesMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != i.CreatedAt {
		i.mutations.Assign(2, val)
		i.logs = append(i.logs, A.X{`createdAt`, i.CreatedAt, val})
		i.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (i *InventoryChangesMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != i.CreatedBy {
		i.mutations.Assign(3, val)
		i.logs = append(i.logs, A.X{`createdBy`, i.CreatedBy, val})
		i.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (i *InventoryChangesMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != i.UpdatedAt {
		i.mutations.Assign(4, val)
		i.logs = append(i.logs, A.X{`updatedAt`, i.UpdatedAt, val})
		i.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (i *InventoryChangesMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != i.UpdatedBy {
		i.mutations.Assign(5, val)
		i.logs = append(i.logs, A.X{`updatedBy`, i.UpdatedBy, val})
		i.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (i *InventoryChangesMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != i.DeletedAt {
		i.mutations.Assign(6, val)
		i.logs = append(i.logs, A.X{`deletedAt`, i.DeletedAt, val})
		i.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (i *InventoryChangesMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != i.DeletedBy {
		i.mutations.Assign(7, val)
		i.logs = append(i.logs, A.X{`deletedBy`, i.DeletedBy, val})
		i.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (i *InventoryChangesMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != i.RestoredBy {
		i.mutations.Assign(8, val)
		i.logs = append(i.logs, A.X{`restoredBy`, i.RestoredBy, val})
		i.RestoredBy = val
		return true
	}
	return false
}

// SetStockDelta create mutations, should not duplicate
func (i *InventoryChangesMutator) SetStockDelta(val uint64) bool { //nolint:dupl false positive
	if val != i.StockDelta {
		i.mutations.Assign(9, val)
		i.logs = append(i.logs, A.X{`stockDelta`, i.StockDelta, val})
		i.StockDelta = val
		return true
	}
	return false
}

// SetProductId create mutations, should not duplicate
func (i *InventoryChangesMutator) SetProductId(val uint64) bool { //nolint:dupl false positive
	if val != i.ProductId {
		i.mutations.Assign(10, val)
		i.logs = append(i.logs, A.X{`productId`, i.ProductId, val})
		i.ProductId = val
		return true
	}
	return false
}

// SetLocationId create mutations, should not duplicate
func (i *InventoryChangesMutator) SetLocationId(val uint64) bool { //nolint:dupl false positive
	if val != i.LocationId {
		i.mutations.Assign(11, val)
		i.logs = append(i.logs, A.X{`locationId`, i.LocationId, val})
		i.LocationId = val
		return true
	}
	return false
}

// SetSpendingId create mutations, should not duplicate
func (i *InventoryChangesMutator) SetSpendingId(val uint64) bool { //nolint:dupl false positive
	if val != i.SpendingId {
		i.mutations.Assign(12, val)
		i.logs = append(i.logs, A.X{`spendingId`, i.SpendingId, val})
		i.SpendingId = val
		return true
	}
	return false
}

// SetExpenseId create mutations, should not duplicate
func (i *InventoryChangesMutator) SetExpenseId(val uint64) bool { //nolint:dupl false positive
	if val != i.ExpenseId {
		i.mutations.Assign(13, val)
		i.logs = append(i.logs, A.X{`expenseId`, i.ExpenseId, val})
		i.ExpenseId = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (i *InventoryChangesMutator) SetAll(from rqBusiness.InventoryChanges, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		i.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		i.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		i.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		i.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		i.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		i.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		i.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		i.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		i.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`stockDelta`] && (forceMap[`stockDelta`] || from.StockDelta != 0) {
		i.StockDelta = from.StockDelta
		changed = true
	}
	if !excludeMap[`productId`] && (forceMap[`productId`] || from.ProductId != 0) {
		i.ProductId = from.ProductId
		changed = true
	}
	if !excludeMap[`locationId`] && (forceMap[`locationId`] || from.LocationId != 0) {
		i.LocationId = from.LocationId
		changed = true
	}
	if !excludeMap[`spendingId`] && (forceMap[`spendingId`] || from.SpendingId != 0) {
		i.SpendingId = from.SpendingId
		changed = true
	}
	if !excludeMap[`expenseId`] && (forceMap[`expenseId`] || from.ExpenseId != 0) {
		i.ExpenseId = from.ExpenseId
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// LocationsMutator DAO writer/command struct
type LocationsMutator struct {
	rqBusiness.Locations
	mutations *tarantool.Operations
	logs      []A.X
}

// NewLocationsMutator create new ORM writer/command object
func NewLocationsMutator(adapter *Tt.Adapter) (res *LocationsMutator) {
	res = &LocationsMutator{Locations: rqBusiness.Locations{Adapter: adapter}}
	res.mutations = tarantool.NewOperations()
	return
}

// Logs get array of logs [field, old, new]
func (l *LocationsMutator) Logs() []A.X { //nolint:dupl false positive
	return l.logs
}

// HaveMutation check whether Set* methods ever called
func (l *LocationsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(l.logs) > 0
}

// ClearMutations clear all previously called Set* methods
func (l *LocationsMutator) ClearMutations() { //nolint:dupl false positive
	l.mutations = tarantool.NewOperations()
	l.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (l *LocationsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := l.Adapter.RetryDo(tarantool.NewUpdateRequest(l.SpaceName()).
		Index(l.UniqueIndexId()).
		Key(tarantool.UintKey{I: uint(l.Id)}).
		Operations(l.ToUpdateArray()),
	)
	return !L.IsError(err, `Locations.DoOverwriteById failed: `+l.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (l *LocationsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !l.HaveMutation() {
		return true
	}
	_, err := l.Adapter.RetryDo(
		tarantool.NewUpdateRequest(l.SpaceName()).
			Index(l.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(l.Id)}).
			Operations(l.mutations),
	)
	return !L.IsError(err, `Locations.DoUpdateById failed: `+l.SpaceName())
}

// DoDeletePermanentById permanent delete
func (l *LocationsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := l.Adapter.RetryDo(
		tarantool.NewDeleteRequest(l.SpaceName()).
			Index(l.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(l.Id)}),
	)
	return !L.IsError(err, `Locations.DoDeletePermanentById failed: `+l.SpaceName())
}

// DoInsert insert, error if already exists
func (l *LocationsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := l.ToArray()
	row, err := l.Adapter.RetryDo(
		tarantool.NewInsertRequest(l.SpaceName()).
			Tuple(arr),
	)
	if err == nil {
		if len(row) > 0 {
			if cells, ok := row[0].([]any); ok && len(cells) > 0 {
				l.Id = X.ToU(cells[0])
			}
		}
	}
	return !L.IsError(err, `Locations.DoInsert failed: `+l.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// tarantool's replace/upsert can only match by primary key
// previous name: DoReplace
func (l *LocationsMutator) DoUpsertById() bool { //nolint:dupl false positive
	if l.Id > 0 {
		return l.DoUpdateById()
	}
	return l.DoInsert()
}

// SetId create mutations, should not duplicate
func (l *LocationsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != l.Id {
		l.mutations.Assign(0, val)
		l.logs = append(l.logs, A.X{`id`, l.Id, val})
		l.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (l *LocationsMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != l.TenantCode {
		l.mutations.Assign(1, val)
		l.logs = append(l.logs, A.X{`tenantCode`, l.TenantCode, val})
		l.TenantCode = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (l *LocationsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != l.CreatedAt {
		l.mutations.Assign(2, val)
		l.logs = append(l.logs, A.X{`createdAt`, l.CreatedAt, val})
		l.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (l *LocationsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.CreatedBy {
		l.mutations.Assign(3, val)
		l.logs = append(l.logs, A.X{`createdBy`, l.CreatedBy, val})
		l.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (l *LocationsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != l.UpdatedAt {
		l.mutations.Assign(4, val)
		l.logs = append(l.logs, A.X{`updatedAt`, l.UpdatedAt, val})
		l.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (l *LocationsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.UpdatedBy {
		l.mutations.Assign(5, val)
		l.logs = append(l.logs, A.X{`updatedBy`, l.UpdatedBy, val})
		l.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (l *LocationsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != l.DeletedAt {
		l.mutations.Assign(6, val)
		l.logs = append(l.logs, A.X{`deletedAt`, l.DeletedAt, val})
		l.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (l *LocationsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.DeletedBy {
		l.mutations.Assign(7, val)
		l.logs = append(l.logs, A.X{`deletedBy`, l.DeletedBy, val})
		l.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (l *LocationsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != l.RestoredBy {
		l.mutations.Assign(8, val)
		l.logs = append(l.logs, A.X{`restoredBy`, l.RestoredBy, val})
		l.RestoredBy = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (l *LocationsMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != l.Name {
		l.mutations.Assign(9, val)
		l.logs = append(l.logs, A.X{`name`, l.Name, val})
		l.Name = val
		return true
	}
	return false
}

// SetCountry create mutations, should not duplicate
func (l *LocationsMutator) SetCountry(val string) bool { //nolint:dupl false positive
	if val != l.Country {
		l.mutations.Assign(10, val)
		l.logs = append(l.logs, A.X{`country`, l.Country, val})
		l.Country = val
		return true
	}
	return false
}

// SetStateProvice create mutations, should not duplicate
func (l *LocationsMutator) SetStateProvice(val string) bool { //nolint:dupl false positive
	if val != l.StateProvice {
		l.mutations.Assign(11, val)
		l.logs = append(l.logs, A.X{`stateProvice`, l.StateProvice, val})
		l.StateProvice = val
		return true
	}
	return false
}

// SetCityRegency create mutations, should not duplicate
func (l *LocationsMutator) SetCityRegency(val string) bool { //nolint:dupl false positive
	if val != l.CityRegency {
		l.mutations.Assign(12, val)
		l.logs = append(l.logs, A.X{`cityRegency`, l.CityRegency, val})
		l.CityRegency = val
		return true
	}
	return false
}

// SetSubdistrict create mutations, should not duplicate
func (l *LocationsMutator) SetSubdistrict(val string) bool { //nolint:dupl false positive
	if val != l.Subdistrict {
		l.mutations.Assign(13, val)
		l.logs = append(l.logs, A.X{`subdistrict`, l.Subdistrict, val})
		l.Subdistrict = val
		return true
	}
	return false
}

// SetVillage create mutations, should not duplicate
func (l *LocationsMutator) SetVillage(val string) bool { //nolint:dupl false positive
	if val != l.Village {
		l.mutations.Assign(14, val)
		l.logs = append(l.logs, A.X{`village`, l.Village, val})
		l.Village = val
		return true
	}
	return false
}

// SetRwBanjar create mutations, should not duplicate
func (l *LocationsMutator) SetRwBanjar(val string) bool { //nolint:dupl false positive
	if val != l.RwBanjar {
		l.mutations.Assign(15, val)
		l.logs = append(l.logs, A.X{`rwBanjar`, l.RwBanjar, val})
		l.RwBanjar = val
		return true
	}
	return false
}

// SetRtNeigb create mutations, should not duplicate
func (l *LocationsMutator) SetRtNeigb(val string) bool { //nolint:dupl false positive
	if val != l.RtNeigb {
		l.mutations.Assign(16, val)
		l.logs = append(l.logs, A.X{`rtNeigb`, l.RtNeigb, val})
		l.RtNeigb = val
		return true
	}
	return false
}

// SetAddress create mutations, should not duplicate
func (l *LocationsMutator) SetAddress(val string) bool { //nolint:dupl false positive
	if val != l.Address {
		l.mutations.Assign(17, val)
		l.logs = append(l.logs, A.X{`address`, l.Address, val})
		l.Address = val
		return true
	}
	return false
}

// SetLat create mutations, should not duplicate
func (l *LocationsMutator) SetLat(val float64) bool { //nolint:dupl false positive
	if val != l.Lat {
		l.mutations.Assign(18, val)
		l.logs = append(l.logs, A.X{`lat`, l.Lat, val})
		l.Lat = val
		return true
	}
	return false
}

// SetLng create mutations, should not duplicate
func (l *LocationsMutator) SetLng(val float64) bool { //nolint:dupl false positive
	if val != l.Lng {
		l.mutations.Assign(19, val)
		l.logs = append(l.logs, A.X{`lng`, l.Lng, val})
		l.Lng = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (l *LocationsMutator) SetAll(from rqBusiness.Locations, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		l.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		l.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		l.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		l.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		l.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		l.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		l.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		l.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		l.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`name`] && (forceMap[`name`] || from.Name != ``) {
		l.Name = S.Trim(from.Name)
		changed = true
	}
	if !excludeMap[`country`] && (forceMap[`country`] || from.Country != ``) {
		l.Country = S.Trim(from.Country)
		changed = true
	}
	if !excludeMap[`stateProvice`] && (forceMap[`stateProvice`] || from.StateProvice != ``) {
		l.StateProvice = S.Trim(from.StateProvice)
		changed = true
	}
	if !excludeMap[`cityRegency`] && (forceMap[`cityRegency`] || from.CityRegency != ``) {
		l.CityRegency = S.Trim(from.CityRegency)
		changed = true
	}
	if !excludeMap[`subdistrict`] && (forceMap[`subdistrict`] || from.Subdistrict != ``) {
		l.Subdistrict = S.Trim(from.Subdistrict)
		changed = true
	}
	if !excludeMap[`village`] && (forceMap[`village`] || from.Village != ``) {
		l.Village = S.Trim(from.Village)
		changed = true
	}
	if !excludeMap[`rwBanjar`] && (forceMap[`rwBanjar`] || from.RwBanjar != ``) {
		l.RwBanjar = S.Trim(from.RwBanjar)
		changed = true
	}
	if !excludeMap[`rtNeigb`] && (forceMap[`rtNeigb`] || from.RtNeigb != ``) {
		l.RtNeigb = S.Trim(from.RtNeigb)
		changed = true
	}
	if !excludeMap[`address`] && (forceMap[`address`] || from.Address != ``) {
		l.Address = S.Trim(from.Address)
		changed = true
	}
	if !excludeMap[`lat`] && (forceMap[`lat`] || from.Lat != 0) {
		l.Lat = from.Lat
		changed = true
	}
	if !excludeMap[`lng`] && (forceMap[`lng`] || from.Lng != 0) {
		l.Lng = from.Lng
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// ProductsMutator DAO writer/command struct
type ProductsMutator struct {
	rqBusiness.Products
	mutations *tarantool.Operations
	logs      []A.X
}

// NewProductsMutator create new ORM writer/command object
func NewProductsMutator(adapter *Tt.Adapter) (res *ProductsMutator) {
	res = &ProductsMutator{Products: rqBusiness.Products{Adapter: adapter}}
	res.mutations = tarantool.NewOperations()
	return
}

// Logs get array of logs [field, old, new]
func (p *ProductsMutator) Logs() []A.X { //nolint:dupl false positive
	return p.logs
}

// HaveMutation check whether Set* methods ever called
func (p *ProductsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.logs) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *ProductsMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = tarantool.NewOperations()
	p.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (p *ProductsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := p.Adapter.RetryDo(tarantool.NewUpdateRequest(p.SpaceName()).
		Index(p.UniqueIndexId()).
		Key(tarantool.UintKey{I: uint(p.Id)}).
		Operations(p.ToUpdateArray()),
	)
	return !L.IsError(err, `Products.DoOverwriteById failed: `+p.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *ProductsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.RetryDo(
		tarantool.NewUpdateRequest(p.SpaceName()).
			Index(p.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(p.Id)}).
			Operations(p.mutations),
	)
	return !L.IsError(err, `Products.DoUpdateById failed: `+p.SpaceName())
}

// DoDeletePermanentById permanent delete
func (p *ProductsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := p.Adapter.RetryDo(
		tarantool.NewDeleteRequest(p.SpaceName()).
			Index(p.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(p.Id)}),
	)
	return !L.IsError(err, `Products.DoDeletePermanentById failed: `+p.SpaceName())
}

// DoInsert insert, error if already exists
func (p *ProductsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := p.ToArray()
	row, err := p.Adapter.RetryDo(
		tarantool.NewInsertRequest(p.SpaceName()).
			Tuple(arr),
	)
	if err == nil {
		if len(row) > 0 {
			if cells, ok := row[0].([]any); ok && len(cells) > 0 {
				p.Id = X.ToU(cells[0])
			}
		}
	}
	return !L.IsError(err, `Products.DoInsert failed: `+p.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// tarantool's replace/upsert can only match by primary key
// previous name: DoReplace
func (p *ProductsMutator) DoUpsertById() bool { //nolint:dupl false positive
	if p.Id > 0 {
		return p.DoUpdateById()
	}
	return p.DoInsert()
}

// SetId create mutations, should not duplicate
func (p *ProductsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != p.Id {
		p.mutations.Assign(0, val)
		p.logs = append(p.logs, A.X{`id`, p.Id, val})
		p.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (p *ProductsMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != p.TenantCode {
		p.mutations.Assign(1, val)
		p.logs = append(p.logs, A.X{`tenantCode`, p.TenantCode, val})
		p.TenantCode = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (p *ProductsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.CreatedAt {
		p.mutations.Assign(2, val)
		p.logs = append(p.logs, A.X{`createdAt`, p.CreatedAt, val})
		p.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (p *ProductsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.CreatedBy {
		p.mutations.Assign(3, val)
		p.logs = append(p.logs, A.X{`createdBy`, p.CreatedBy, val})
		p.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (p *ProductsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.UpdatedAt {
		p.mutations.Assign(4, val)
		p.logs = append(p.logs, A.X{`updatedAt`, p.UpdatedAt, val})
		p.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (p *ProductsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.UpdatedBy {
		p.mutations.Assign(5, val)
		p.logs = append(p.logs, A.X{`updatedBy`, p.UpdatedBy, val})
		p.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (p *ProductsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != p.DeletedAt {
		p.mutations.Assign(6, val)
		p.logs = append(p.logs, A.X{`deletedAt`, p.DeletedAt, val})
		p.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (p *ProductsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.DeletedBy {
		p.mutations.Assign(7, val)
		p.logs = append(p.logs, A.X{`deletedBy`, p.DeletedBy, val})
		p.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (p *ProductsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != p.RestoredBy {
		p.mutations.Assign(8, val)
		p.logs = append(p.logs, A.X{`restoredBy`, p.RestoredBy, val})
		p.RestoredBy = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (p *ProductsMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != p.Name {
		p.mutations.Assign(9, val)
		p.logs = append(p.logs, A.X{`name`, p.Name, val})
		p.Name = val
		return true
	}
	return false
}

// SetDetail create mutations, should not duplicate
func (p *ProductsMutator) SetDetail(val string) bool { //nolint:dupl false positive
	if val != p.Detail {
		p.mutations.Assign(10, val)
		p.logs = append(p.logs, A.X{`detail`, p.Detail, val})
		p.Detail = val
		return true
	}
	return false
}

// SetRule create mutations, should not duplicate
func (p *ProductsMutator) SetRule(val string) bool { //nolint:dupl false positive
	if val != p.Rule {
		p.mutations.Assign(11, val)
		p.logs = append(p.logs, A.X{`rule`, p.Rule, val})
		p.Rule = val
		return true
	}
	return false
}

// SetKind create mutations, should not duplicate
func (p *ProductsMutator) SetKind(val string) bool { //nolint:dupl false positive
	if val != p.Kind {
		p.mutations.Assign(12, val)
		p.logs = append(p.logs, A.X{`kind`, p.Kind, val})
		p.Kind = val
		return true
	}
	return false
}

// SetCogsIDR create mutations, should not duplicate
func (p *ProductsMutator) SetCogsIDR(val int64) bool { //nolint:dupl false positive
	if val != p.CogsIDR {
		p.mutations.Assign(13, val)
		p.logs = append(p.logs, A.X{`cogsIDR`, p.CogsIDR, val})
		p.CogsIDR = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (p *ProductsMutator) SetAll(from rqBusiness.Products, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		p.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		p.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		p.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		p.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		p.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		p.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		p.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		p.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		p.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`name`] && (forceMap[`name`] || from.Name != ``) {
		p.Name = S.Trim(from.Name)
		changed = true
	}
	if !excludeMap[`detail`] && (forceMap[`detail`] || from.Detail != ``) {
		p.Detail = S.Trim(from.Detail)
		changed = true
	}
	if !excludeMap[`rule`] && (forceMap[`rule`] || from.Rule != ``) {
		p.Rule = S.Trim(from.Rule)
		changed = true
	}
	if !excludeMap[`kind`] && (forceMap[`kind`] || from.Kind != ``) {
		p.Kind = S.Trim(from.Kind)
		changed = true
	}
	if !excludeMap[`cogsIDR`] && (forceMap[`cogsIDR`] || from.CogsIDR != 0) {
		p.CogsIDR = from.CogsIDR
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
