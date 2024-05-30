package wcBusiness

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mBusiness/rqBusiness"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// LocationsMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcBusiness__ORM.GEN.go
type LocationsMutator struct {
	rqBusiness.Locations
	mutations []A.X
	logs      []A.X
}

// NewLocationsMutator create new ORM writer/command object
func NewLocationsMutator(adapter *Tt.Adapter) (res *LocationsMutator) {
	res = &LocationsMutator{Locations: rqBusiness.Locations{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (l *LocationsMutator) Logs() []A.X { //nolint:dupl false positive
	return l.logs
}

// HaveMutation check whether Set* methods ever called
func (l *LocationsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(l.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (l *LocationsMutator) ClearMutations() { //nolint:dupl false positive
	l.mutations = []A.X{}
	l.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (l *LocationsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := l.Adapter.Update(l.SpaceName(), l.UniqueIndexId(), A.X{l.Id}, l.ToUpdateArray())
	return !L.IsError(err, `Locations.DoOverwriteById failed: `+l.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (l *LocationsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !l.HaveMutation() {
		return true
	}
	_, err := l.Adapter.Update(l.SpaceName(), l.UniqueIndexId(), A.X{l.Id}, l.mutations)
	return !L.IsError(err, `Locations.DoUpdateById failed: `+l.SpaceName())
}

// DoDeletePermanentById permanent delete
func (l *LocationsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := l.Adapter.Delete(l.SpaceName(), l.UniqueIndexId(), A.X{l.Id})
	return !L.IsError(err, `Locations.DoDeletePermanentById failed: `+l.SpaceName())
}

// func (l *LocationsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := l.ToArray()
//	_, err := l.Adapter.Upsert(l.SpaceName(), arr, A.X{
//		A.X{`=`, 0, l.Id},
//		A.X{`=`, 1, l.TenantCode},
//		A.X{`=`, 2, l.CreatedAt},
//		A.X{`=`, 3, l.CreatedBy},
//		A.X{`=`, 4, l.UpdatedAt},
//		A.X{`=`, 5, l.UpdatedBy},
//		A.X{`=`, 6, l.DeletedAt},
//		A.X{`=`, 7, l.DeletedBy},
//		A.X{`=`, 8, l.RestoredBy},
//		A.X{`=`, 9, l.Name},
//		A.X{`=`, 10, l.Country},
//		A.X{`=`, 11, l.StateProvice},
//		A.X{`=`, 12, l.CityRegency},
//		A.X{`=`, 13, l.Subdistrict},
//		A.X{`=`, 14, l.Village},
//		A.X{`=`, 15, l.RwBanjar},
//		A.X{`=`, 16, l.RtNeigb},
//		A.X{`=`, 17, l.Address},
//		A.X{`=`, 18, l.Lat},
//		A.X{`=`, 19, l.Lng},
//	})
//	return !L.IsError(err, `Locations.DoUpsert failed: `+l.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (l *LocationsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := l.ToArray()
	row, err := l.Adapter.Insert(l.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			l.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Locations.DoInsert failed: `+l.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (l *LocationsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := l.ToArray()
	row, err := l.Adapter.Replace(l.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			l.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Locations.DoUpsert failed: `+l.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (l *LocationsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != l.Id {
		l.mutations = append(l.mutations, A.X{`=`, 0, val})
		l.logs = append(l.logs, A.X{`id`, l.Id, val})
		l.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (l *LocationsMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != l.TenantCode {
		l.mutations = append(l.mutations, A.X{`=`, 1, val})
		l.logs = append(l.logs, A.X{`tenantCode`, l.TenantCode, val})
		l.TenantCode = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (l *LocationsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != l.CreatedAt {
		l.mutations = append(l.mutations, A.X{`=`, 2, val})
		l.logs = append(l.logs, A.X{`createdAt`, l.CreatedAt, val})
		l.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (l *LocationsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.CreatedBy {
		l.mutations = append(l.mutations, A.X{`=`, 3, val})
		l.logs = append(l.logs, A.X{`createdBy`, l.CreatedBy, val})
		l.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (l *LocationsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != l.UpdatedAt {
		l.mutations = append(l.mutations, A.X{`=`, 4, val})
		l.logs = append(l.logs, A.X{`updatedAt`, l.UpdatedAt, val})
		l.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (l *LocationsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.UpdatedBy {
		l.mutations = append(l.mutations, A.X{`=`, 5, val})
		l.logs = append(l.logs, A.X{`updatedBy`, l.UpdatedBy, val})
		l.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (l *LocationsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != l.DeletedAt {
		l.mutations = append(l.mutations, A.X{`=`, 6, val})
		l.logs = append(l.logs, A.X{`deletedAt`, l.DeletedAt, val})
		l.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (l *LocationsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != l.DeletedBy {
		l.mutations = append(l.mutations, A.X{`=`, 7, val})
		l.logs = append(l.logs, A.X{`deletedBy`, l.DeletedBy, val})
		l.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (l *LocationsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != l.RestoredBy {
		l.mutations = append(l.mutations, A.X{`=`, 8, val})
		l.logs = append(l.logs, A.X{`restoredBy`, l.RestoredBy, val})
		l.RestoredBy = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (l *LocationsMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != l.Name {
		l.mutations = append(l.mutations, A.X{`=`, 9, val})
		l.logs = append(l.logs, A.X{`name`, l.Name, val})
		l.Name = val
		return true
	}
	return false
}

// SetCountry create mutations, should not duplicate
func (l *LocationsMutator) SetCountry(val string) bool { //nolint:dupl false positive
	if val != l.Country {
		l.mutations = append(l.mutations, A.X{`=`, 10, val})
		l.logs = append(l.logs, A.X{`country`, l.Country, val})
		l.Country = val
		return true
	}
	return false
}

// SetStateProvice create mutations, should not duplicate
func (l *LocationsMutator) SetStateProvice(val string) bool { //nolint:dupl false positive
	if val != l.StateProvice {
		l.mutations = append(l.mutations, A.X{`=`, 11, val})
		l.logs = append(l.logs, A.X{`stateProvice`, l.StateProvice, val})
		l.StateProvice = val
		return true
	}
	return false
}

// SetCityRegency create mutations, should not duplicate
func (l *LocationsMutator) SetCityRegency(val string) bool { //nolint:dupl false positive
	if val != l.CityRegency {
		l.mutations = append(l.mutations, A.X{`=`, 12, val})
		l.logs = append(l.logs, A.X{`cityRegency`, l.CityRegency, val})
		l.CityRegency = val
		return true
	}
	return false
}

// SetSubdistrict create mutations, should not duplicate
func (l *LocationsMutator) SetSubdistrict(val string) bool { //nolint:dupl false positive
	if val != l.Subdistrict {
		l.mutations = append(l.mutations, A.X{`=`, 13, val})
		l.logs = append(l.logs, A.X{`subdistrict`, l.Subdistrict, val})
		l.Subdistrict = val
		return true
	}
	return false
}

// SetVillage create mutations, should not duplicate
func (l *LocationsMutator) SetVillage(val string) bool { //nolint:dupl false positive
	if val != l.Village {
		l.mutations = append(l.mutations, A.X{`=`, 14, val})
		l.logs = append(l.logs, A.X{`village`, l.Village, val})
		l.Village = val
		return true
	}
	return false
}

// SetRwBanjar create mutations, should not duplicate
func (l *LocationsMutator) SetRwBanjar(val string) bool { //nolint:dupl false positive
	if val != l.RwBanjar {
		l.mutations = append(l.mutations, A.X{`=`, 15, val})
		l.logs = append(l.logs, A.X{`rwBanjar`, l.RwBanjar, val})
		l.RwBanjar = val
		return true
	}
	return false
}

// SetRtNeigb create mutations, should not duplicate
func (l *LocationsMutator) SetRtNeigb(val string) bool { //nolint:dupl false positive
	if val != l.RtNeigb {
		l.mutations = append(l.mutations, A.X{`=`, 16, val})
		l.logs = append(l.logs, A.X{`rtNeigb`, l.RtNeigb, val})
		l.RtNeigb = val
		return true
	}
	return false
}

// SetAddress create mutations, should not duplicate
func (l *LocationsMutator) SetAddress(val string) bool { //nolint:dupl false positive
	if val != l.Address {
		l.mutations = append(l.mutations, A.X{`=`, 17, val})
		l.logs = append(l.logs, A.X{`address`, l.Address, val})
		l.Address = val
		return true
	}
	return false
}

// SetLat create mutations, should not duplicate
func (l *LocationsMutator) SetLat(val float64) bool { //nolint:dupl false positive
	if val != l.Lat {
		l.mutations = append(l.mutations, A.X{`=`, 18, val})
		l.logs = append(l.logs, A.X{`lat`, l.Lat, val})
		l.Lat = val
		return true
	}
	return false
}

// SetLng create mutations, should not duplicate
func (l *LocationsMutator) SetLng(val float64) bool { //nolint:dupl false positive
	if val != l.Lng {
		l.mutations = append(l.mutations, A.X{`=`, 19, val})
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
	mutations []A.X
	logs      []A.X
}

// NewProductsMutator create new ORM writer/command object
func NewProductsMutator(adapter *Tt.Adapter) (res *ProductsMutator) {
	res = &ProductsMutator{Products: rqBusiness.Products{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (p *ProductsMutator) Logs() []A.X { //nolint:dupl false positive
	return p.logs
}

// HaveMutation check whether Set* methods ever called
func (p *ProductsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *ProductsMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = []A.X{}
	p.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (p *ProductsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.ToUpdateArray())
	return !L.IsError(err, `Products.DoOverwriteById failed: `+p.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *ProductsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.mutations)
	return !L.IsError(err, `Products.DoUpdateById failed: `+p.SpaceName())
}

// DoDeletePermanentById permanent delete
func (p *ProductsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Delete(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id})
	return !L.IsError(err, `Products.DoDeletePermanentById failed: `+p.SpaceName())
}

// func (p *ProductsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := p.ToArray()
//	_, err := p.Adapter.Upsert(p.SpaceName(), arr, A.X{
//		A.X{`=`, 0, p.Id},
//		A.X{`=`, 1, p.TenantCode},
//		A.X{`=`, 2, p.CreatedAt},
//		A.X{`=`, 3, p.CreatedBy},
//		A.X{`=`, 4, p.UpdatedAt},
//		A.X{`=`, 5, p.UpdatedBy},
//		A.X{`=`, 6, p.DeletedAt},
//		A.X{`=`, 7, p.DeletedBy},
//		A.X{`=`, 8, p.RestoredBy},
//		A.X{`=`, 9, p.Name},
//		A.X{`=`, 10, p.Detail},
//		A.X{`=`, 11, p.Rule},
//		A.X{`=`, 12, p.Kind},
//		A.X{`=`, 13, p.CogsIDR},
//	})
//	return !L.IsError(err, `Products.DoUpsert failed: `+p.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (p *ProductsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := p.ToArray()
	row, err := p.Adapter.Insert(p.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Products.DoInsert failed: `+p.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (p *ProductsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := p.ToArray()
	row, err := p.Adapter.Replace(p.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Products.DoUpsert failed: `+p.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (p *ProductsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != p.Id {
		p.mutations = append(p.mutations, A.X{`=`, 0, val})
		p.logs = append(p.logs, A.X{`id`, p.Id, val})
		p.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (p *ProductsMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != p.TenantCode {
		p.mutations = append(p.mutations, A.X{`=`, 1, val})
		p.logs = append(p.logs, A.X{`tenantCode`, p.TenantCode, val})
		p.TenantCode = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (p *ProductsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.CreatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 2, val})
		p.logs = append(p.logs, A.X{`createdAt`, p.CreatedAt, val})
		p.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (p *ProductsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.CreatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 3, val})
		p.logs = append(p.logs, A.X{`createdBy`, p.CreatedBy, val})
		p.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (p *ProductsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.UpdatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 4, val})
		p.logs = append(p.logs, A.X{`updatedAt`, p.UpdatedAt, val})
		p.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (p *ProductsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.UpdatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 5, val})
		p.logs = append(p.logs, A.X{`updatedBy`, p.UpdatedBy, val})
		p.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (p *ProductsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != p.DeletedAt {
		p.mutations = append(p.mutations, A.X{`=`, 6, val})
		p.logs = append(p.logs, A.X{`deletedAt`, p.DeletedAt, val})
		p.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (p *ProductsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.DeletedBy {
		p.mutations = append(p.mutations, A.X{`=`, 7, val})
		p.logs = append(p.logs, A.X{`deletedBy`, p.DeletedBy, val})
		p.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (p *ProductsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != p.RestoredBy {
		p.mutations = append(p.mutations, A.X{`=`, 8, val})
		p.logs = append(p.logs, A.X{`restoredBy`, p.RestoredBy, val})
		p.RestoredBy = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (p *ProductsMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != p.Name {
		p.mutations = append(p.mutations, A.X{`=`, 9, val})
		p.logs = append(p.logs, A.X{`name`, p.Name, val})
		p.Name = val
		return true
	}
	return false
}

// SetDetail create mutations, should not duplicate
func (p *ProductsMutator) SetDetail(val string) bool { //nolint:dupl false positive
	if val != p.Detail {
		p.mutations = append(p.mutations, A.X{`=`, 10, val})
		p.logs = append(p.logs, A.X{`detail`, p.Detail, val})
		p.Detail = val
		return true
	}
	return false
}

// SetRule create mutations, should not duplicate
func (p *ProductsMutator) SetRule(val string) bool { //nolint:dupl false positive
	if val != p.Rule {
		p.mutations = append(p.mutations, A.X{`=`, 11, val})
		p.logs = append(p.logs, A.X{`rule`, p.Rule, val})
		p.Rule = val
		return true
	}
	return false
}

// SetKind create mutations, should not duplicate
func (p *ProductsMutator) SetKind(val string) bool { //nolint:dupl false positive
	if val != p.Kind {
		p.mutations = append(p.mutations, A.X{`=`, 12, val})
		p.logs = append(p.logs, A.X{`kind`, p.Kind, val})
		p.Kind = val
		return true
	}
	return false
}

// SetCogsIDR create mutations, should not duplicate
func (p *ProductsMutator) SetCogsIDR(val int64) bool { //nolint:dupl false positive
	if val != p.CogsIDR {
		p.mutations = append(p.mutations, A.X{`=`, 13, val})
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
