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

// ProductsMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcBusiness__ORM.GEN.go
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
