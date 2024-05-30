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

// ProductsMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcBusiness__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcBusiness__ORM.GEN.go
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
