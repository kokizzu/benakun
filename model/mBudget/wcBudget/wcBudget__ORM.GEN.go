package wcBudget

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mBudget/rqBudget"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// PlansMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcBudget__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcBudget__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcBudget__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcBudget__ORM.GEN.go
type PlansMutator struct {
	rqBudget.Plans
	mutations []A.X
	logs      []A.X
}

// NewPlansMutator create new ORM writer/command object
func NewPlansMutator(adapter *Tt.Adapter) (res *PlansMutator) {
	res = &PlansMutator{Plans: rqBudget.Plans{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (p *PlansMutator) Logs() []A.X { //nolint:dupl false positive
	return p.logs
}

// HaveMutation check whether Set* methods ever called
func (p *PlansMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *PlansMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = []A.X{}
	p.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (p *PlansMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.ToUpdateArray())
	return !L.IsError(err, `Plans.DoOverwriteById failed: `+p.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PlansMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.Update(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id}, p.mutations)
	return !L.IsError(err, `Plans.DoUpdateById failed: `+p.SpaceName())
}

// DoDeletePermanentById permanent delete
func (p *PlansMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := p.Adapter.Delete(p.SpaceName(), p.UniqueIndexId(), A.X{p.Id})
	return !L.IsError(err, `Plans.DoDeletePermanentById failed: `+p.SpaceName())
}

// func (p *PlansMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := p.ToArray()
//	_, err := p.Adapter.Upsert(p.SpaceName(), arr, A.X{
//		A.X{`=`, 0, p.Id},
//		A.X{`=`, 1, p.PlanType},
//		A.X{`=`, 2, p.ParentId},
//		A.X{`=`, 3, p.CreatedAt},
//		A.X{`=`, 4, p.CreatedBy},
//		A.X{`=`, 5, p.UpdatedAt},
//		A.X{`=`, 6, p.UpdatedBy},
//		A.X{`=`, 7, p.DeletedAt},
//		A.X{`=`, 8, p.Title},
//		A.X{`=`, 9, p.Description},
//		A.X{`=`, 10, p.OrgId},
//		A.X{`=`, 11, p.PerYear},
//		A.X{`=`, 12, p.BudgetIDR},
//		A.X{`=`, 13, p.BudgetUSD},
//		A.X{`=`, 14, p.BudgetEUR},
//	})
//	return !L.IsError(err, `Plans.DoUpsert failed: `+p.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (p *PlansMutator) DoInsert() bool { //nolint:dupl false positive
	arr := p.ToArray()
	row, err := p.Adapter.Insert(p.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Plans.DoInsert failed: `+p.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (p *PlansMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := p.ToArray()
	row, err := p.Adapter.Replace(p.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			p.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Plans.DoUpsert failed: `+p.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (p *PlansMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != p.Id {
		p.mutations = append(p.mutations, A.X{`=`, 0, val})
		p.logs = append(p.logs, A.X{`id`, p.Id, val})
		p.Id = val
		return true
	}
	return false
}

// SetPlanType create mutations, should not duplicate
func (p *PlansMutator) SetPlanType(val string) bool { //nolint:dupl false positive
	if val != p.PlanType {
		p.mutations = append(p.mutations, A.X{`=`, 1, val})
		p.logs = append(p.logs, A.X{`planType`, p.PlanType, val})
		p.PlanType = val
		return true
	}
	return false
}

// SetParentId create mutations, should not duplicate
func (p *PlansMutator) SetParentId(val uint64) bool { //nolint:dupl false positive
	if val != p.ParentId {
		p.mutations = append(p.mutations, A.X{`=`, 2, val})
		p.logs = append(p.logs, A.X{`parentId`, p.ParentId, val})
		p.ParentId = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (p *PlansMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.CreatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 3, val})
		p.logs = append(p.logs, A.X{`createdAt`, p.CreatedAt, val})
		p.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (p *PlansMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.CreatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 4, val})
		p.logs = append(p.logs, A.X{`createdBy`, p.CreatedBy, val})
		p.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (p *PlansMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.UpdatedAt {
		p.mutations = append(p.mutations, A.X{`=`, 5, val})
		p.logs = append(p.logs, A.X{`updatedAt`, p.UpdatedAt, val})
		p.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (p *PlansMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.UpdatedBy {
		p.mutations = append(p.mutations, A.X{`=`, 6, val})
		p.logs = append(p.logs, A.X{`updatedBy`, p.UpdatedBy, val})
		p.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (p *PlansMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != p.DeletedAt {
		p.mutations = append(p.mutations, A.X{`=`, 7, val})
		p.logs = append(p.logs, A.X{`deletedAt`, p.DeletedAt, val})
		p.DeletedAt = val
		return true
	}
	return false
}

// SetTitle create mutations, should not duplicate
func (p *PlansMutator) SetTitle(val string) bool { //nolint:dupl false positive
	if val != p.Title {
		p.mutations = append(p.mutations, A.X{`=`, 8, val})
		p.logs = append(p.logs, A.X{`title`, p.Title, val})
		p.Title = val
		return true
	}
	return false
}

// SetDescription create mutations, should not duplicate
func (p *PlansMutator) SetDescription(val string) bool { //nolint:dupl false positive
	if val != p.Description {
		p.mutations = append(p.mutations, A.X{`=`, 9, val})
		p.logs = append(p.logs, A.X{`description`, p.Description, val})
		p.Description = val
		return true
	}
	return false
}

// SetOrgId create mutations, should not duplicate
func (p *PlansMutator) SetOrgId(val uint64) bool { //nolint:dupl false positive
	if val != p.OrgId {
		p.mutations = append(p.mutations, A.X{`=`, 10, val})
		p.logs = append(p.logs, A.X{`orgId`, p.OrgId, val})
		p.OrgId = val
		return true
	}
	return false
}

// SetPerYear create mutations, should not duplicate
func (p *PlansMutator) SetPerYear(val int64) bool { //nolint:dupl false positive
	if val != p.PerYear {
		p.mutations = append(p.mutations, A.X{`=`, 11, val})
		p.logs = append(p.logs, A.X{`perYear`, p.PerYear, val})
		p.PerYear = val
		return true
	}
	return false
}

// SetBudgetIDR create mutations, should not duplicate
func (p *PlansMutator) SetBudgetIDR(val int64) bool { //nolint:dupl false positive
	if val != p.BudgetIDR {
		p.mutations = append(p.mutations, A.X{`=`, 12, val})
		p.logs = append(p.logs, A.X{`budgetIDR`, p.BudgetIDR, val})
		p.BudgetIDR = val
		return true
	}
	return false
}

// SetBudgetUSD create mutations, should not duplicate
func (p *PlansMutator) SetBudgetUSD(val int64) bool { //nolint:dupl false positive
	if val != p.BudgetUSD {
		p.mutations = append(p.mutations, A.X{`=`, 13, val})
		p.logs = append(p.logs, A.X{`budgetUSD`, p.BudgetUSD, val})
		p.BudgetUSD = val
		return true
	}
	return false
}

// SetBudgetEUR create mutations, should not duplicate
func (p *PlansMutator) SetBudgetEUR(val int64) bool { //nolint:dupl false positive
	if val != p.BudgetEUR {
		p.mutations = append(p.mutations, A.X{`=`, 14, val})
		p.logs = append(p.logs, A.X{`budgetEUR`, p.BudgetEUR, val})
		p.BudgetEUR = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (p *PlansMutator) SetAll(from rqBudget.Plans, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`planType`] && (forceMap[`planType`] || from.PlanType != ``) {
		p.PlanType = S.Trim(from.PlanType)
		changed = true
	}
	if !excludeMap[`parentId`] && (forceMap[`parentId`] || from.ParentId != 0) {
		p.ParentId = from.ParentId
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
	if !excludeMap[`title`] && (forceMap[`title`] || from.Title != ``) {
		p.Title = S.Trim(from.Title)
		changed = true
	}
	if !excludeMap[`description`] && (forceMap[`description`] || from.Description != ``) {
		p.Description = S.Trim(from.Description)
		changed = true
	}
	if !excludeMap[`orgId`] && (forceMap[`orgId`] || from.OrgId != 0) {
		p.OrgId = from.OrgId
		changed = true
	}
	if !excludeMap[`perYear`] && (forceMap[`perYear`] || from.PerYear != 0) {
		p.PerYear = from.PerYear
		changed = true
	}
	if !excludeMap[`budgetIDR`] && (forceMap[`budgetIDR`] || from.BudgetIDR != 0) {
		p.BudgetIDR = from.BudgetIDR
		changed = true
	}
	if !excludeMap[`budgetUSD`] && (forceMap[`budgetUSD`] || from.BudgetUSD != 0) {
		p.BudgetUSD = from.BudgetUSD
		changed = true
	}
	if !excludeMap[`budgetEUR`] && (forceMap[`budgetEUR`] || from.BudgetEUR != 0) {
		p.BudgetEUR = from.BudgetEUR
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
