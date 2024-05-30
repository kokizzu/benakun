package wcFinance

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mFinance/rqFinance"

	"github.com/tarantool/go-tarantool/v2"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// CoaMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcFinance__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcFinance__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcFinance__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcFinance__ORM.GEN.go
type CoaMutator struct {
	rqFinance.Coa
	mutations *tarantool.Operations
	logs      []A.X
}

// NewCoaMutator create new ORM writer/command object
func NewCoaMutator(adapter *Tt.Adapter) (res *CoaMutator) {
	res = &CoaMutator{Coa: rqFinance.Coa{Adapter: adapter}}
	res.mutations = tarantool.NewOperations()
	res.Children = []any{}
	return
}

// Logs get array of logs [field, old, new]
func (c *CoaMutator) Logs() []A.X { //nolint:dupl false positive
	return c.logs
}

// HaveMutation check whether Set* methods ever called
func (c *CoaMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(c.logs) > 0
}

// ClearMutations clear all previously called Set* methods
func (c *CoaMutator) ClearMutations() { //nolint:dupl false positive
	c.mutations = tarantool.NewOperations()
	c.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (c *CoaMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := c.Adapter.RetryDo(tarantool.NewUpdateRequest(c.SpaceName()).
		Index(c.UniqueIndexId()).
		Key(tarantool.UintKey{I: uint(c.Id)}).
		Operations(c.ToUpdateArray()),
	)
	return !L.IsError(err, `Coa.DoOverwriteById failed: `+c.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (c *CoaMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !c.HaveMutation() {
		return true
	}
	_, err := c.Adapter.RetryDo(
		tarantool.NewUpdateRequest(c.SpaceName()).
			Index(c.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(c.Id)}).
			Operations(c.mutations),
	)
	return !L.IsError(err, `Coa.DoUpdateById failed: `+c.SpaceName())
}

// DoDeletePermanentById permanent delete
func (c *CoaMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := c.Adapter.RetryDo(
		tarantool.NewDeleteRequest(c.SpaceName()).
			Index(c.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(c.Id)}),
	)
	return !L.IsError(err, `Coa.DoDeletePermanentById failed: `+c.SpaceName())
}

// DoInsert insert, error if already exists
func (c *CoaMutator) DoInsert() bool { //nolint:dupl false positive
	arr := c.ToArray()
	row, err := c.Adapter.RetryDo(
		tarantool.NewInsertRequest(c.SpaceName()).
			Tuple(arr),
	)
	if err == nil {
		if len(row) > 0 {
			if cells, ok := row[0].([]any); ok && len(cells) > 0 {
				c.Id = X.ToU(cells[0])
			}
		}
	}
	return !L.IsError(err, `Coa.DoInsert failed: `+c.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// tarantool's replace/upsert can only match by primary key
// previous name: DoReplace
func (c *CoaMutator) DoUpsertById() bool { //nolint:dupl false positive
	if c.Id > 0 {
		return c.DoUpdateById()
	}
	return c.DoInsert()
}

// SetId create mutations, should not duplicate
func (c *CoaMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != c.Id {
		c.mutations.Assign(0, val)
		c.logs = append(c.logs, A.X{`id`, c.Id, val})
		c.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (c *CoaMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != c.TenantCode {
		c.mutations.Assign(1, val)
		c.logs = append(c.logs, A.X{`tenantCode`, c.TenantCode, val})
		c.TenantCode = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (c *CoaMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != c.Name {
		c.mutations.Assign(2, val)
		c.logs = append(c.logs, A.X{`name`, c.Name, val})
		c.Name = val
		return true
	}
	return false
}

// SetLevel create mutations, should not duplicate
func (c *CoaMutator) SetLevel(val float64) bool { //nolint:dupl false positive
	if val != c.Level {
		c.mutations.Assign(3, val)
		c.logs = append(c.logs, A.X{`level`, c.Level, val})
		c.Level = val
		return true
	}
	return false
}

// SetParentId create mutations, should not duplicate
func (c *CoaMutator) SetParentId(val uint64) bool { //nolint:dupl false positive
	if val != c.ParentId {
		c.mutations.Assign(4, val)
		c.logs = append(c.logs, A.X{`parentId`, c.ParentId, val})
		c.ParentId = val
		return true
	}
	return false
}

// SetChildren create mutations, should not duplicate
func (c *CoaMutator) SetChildren(val []any) bool { //nolint:dupl false positive
	c.mutations.Assign(5, val)
	c.logs = append(c.logs, A.X{`children`, c.Children, val})
	c.Children = val
	return true
}

// SetCreatedAt create mutations, should not duplicate
func (c *CoaMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != c.CreatedAt {
		c.mutations.Assign(6, val)
		c.logs = append(c.logs, A.X{`createdAt`, c.CreatedAt, val})
		c.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (c *CoaMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != c.CreatedBy {
		c.mutations.Assign(7, val)
		c.logs = append(c.logs, A.X{`createdBy`, c.CreatedBy, val})
		c.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (c *CoaMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != c.UpdatedAt {
		c.mutations.Assign(8, val)
		c.logs = append(c.logs, A.X{`updatedAt`, c.UpdatedAt, val})
		c.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (c *CoaMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != c.UpdatedBy {
		c.mutations.Assign(9, val)
		c.logs = append(c.logs, A.X{`updatedBy`, c.UpdatedBy, val})
		c.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (c *CoaMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != c.DeletedAt {
		c.mutations.Assign(10, val)
		c.logs = append(c.logs, A.X{`deletedAt`, c.DeletedAt, val})
		c.DeletedAt = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (c *CoaMutator) SetAll(from rqFinance.Coa, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		c.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		c.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`name`] && (forceMap[`name`] || from.Name != ``) {
		c.Name = S.Trim(from.Name)
		changed = true
	}
	if !excludeMap[`level`] && (forceMap[`level`] || from.Level != 0) {
		c.Level = from.Level
		changed = true
	}
	if !excludeMap[`parentId`] && (forceMap[`parentId`] || from.ParentId != 0) {
		c.ParentId = from.ParentId
		changed = true
	}
	if !excludeMap[`children`] && (forceMap[`children`] || from.Children != nil) {
		c.Children = from.Children
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		c.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		c.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		c.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		c.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		c.DeletedAt = from.DeletedAt
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// TransactionsMutator DAO writer/command struct
type TransactionsMutator struct {
	rqFinance.Transactions
	mutations *tarantool.Operations
	logs      []A.X
}

// NewTransactionsMutator create new ORM writer/command object
func NewTransactionsMutator(adapter *Tt.Adapter) (res *TransactionsMutator) {
	res = &TransactionsMutator{Transactions: rqFinance.Transactions{Adapter: adapter}}
	res.mutations = tarantool.NewOperations()
	return
}

// Logs get array of logs [field, old, new]
func (t *TransactionsMutator) Logs() []A.X { //nolint:dupl false positive
	return t.logs
}

// HaveMutation check whether Set* methods ever called
func (t *TransactionsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(t.logs) > 0
}

// ClearMutations clear all previously called Set* methods
func (t *TransactionsMutator) ClearMutations() { //nolint:dupl false positive
	t.mutations = tarantool.NewOperations()
	t.logs = []A.X{}
}

// DoInsert insert, error if already exists
func (t *TransactionsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	_, err := t.Adapter.RetryDo(
		tarantool.NewInsertRequest(t.SpaceName()).
			Tuple(arr),
	)
	return !L.IsError(err, `Transactions.DoInsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (t *TransactionsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != t.Id {
		t.mutations.Assign(0, val)
		t.logs = append(t.logs, A.X{`id`, t.Id, val})
		t.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (t *TransactionsMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != t.TenantCode {
		t.mutations.Assign(1, val)
		t.logs = append(t.logs, A.X{`tenantCode`, t.TenantCode, val})
		t.TenantCode = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (t *TransactionsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.CreatedAt {
		t.mutations.Assign(2, val)
		t.logs = append(t.logs, A.X{`createdAt`, t.CreatedAt, val})
		t.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (t *TransactionsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.CreatedBy {
		t.mutations.Assign(3, val)
		t.logs = append(t.logs, A.X{`createdBy`, t.CreatedBy, val})
		t.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (t *TransactionsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.UpdatedAt {
		t.mutations.Assign(4, val)
		t.logs = append(t.logs, A.X{`updatedAt`, t.UpdatedAt, val})
		t.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (t *TransactionsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.UpdatedBy {
		t.mutations.Assign(5, val)
		t.logs = append(t.logs, A.X{`updatedBy`, t.UpdatedBy, val})
		t.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (t *TransactionsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != t.DeletedAt {
		t.mutations.Assign(6, val)
		t.logs = append(t.logs, A.X{`deletedAt`, t.DeletedAt, val})
		t.DeletedAt = val
		return true
	}
	return false
}

// SetCompletedAt create mutations, should not duplicate
func (t *TransactionsMutator) SetCompletedAt(val int64) bool { //nolint:dupl false positive
	if val != t.CompletedAt {
		t.mutations.Assign(7, val)
		t.logs = append(t.logs, A.X{`completedAt`, t.CompletedAt, val})
		t.CompletedAt = val
		return true
	}
	return false
}

// SetPrice create mutations, should not duplicate
func (t *TransactionsMutator) SetPrice(val int64) bool { //nolint:dupl false positive
	if val != t.Price {
		t.mutations.Assign(8, val)
		t.logs = append(t.logs, A.X{`price`, t.Price, val})
		t.Price = val
		return true
	}
	return false
}

// SetDescriptions create mutations, should not duplicate
func (t *TransactionsMutator) SetDescriptions(val string) bool { //nolint:dupl false positive
	if val != t.Descriptions {
		t.mutations.Assign(9, val)
		t.logs = append(t.logs, A.X{`descriptions`, t.Descriptions, val})
		t.Descriptions = val
		return true
	}
	return false
}

// SetQty create mutations, should not duplicate
func (t *TransactionsMutator) SetQty(val int64) bool { //nolint:dupl false positive
	if val != t.Qty {
		t.mutations.Assign(10, val)
		t.logs = append(t.logs, A.X{`qty`, t.Qty, val})
		t.Qty = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (t *TransactionsMutator) SetAll(from rqFinance.Transactions, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		t.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		t.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		t.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		t.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		t.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		t.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		t.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`completedAt`] && (forceMap[`completedAt`] || from.CompletedAt != 0) {
		t.CompletedAt = from.CompletedAt
		changed = true
	}
	if !excludeMap[`price`] && (forceMap[`price`] || from.Price != 0) {
		t.Price = from.Price
		changed = true
	}
	if !excludeMap[`descriptions`] && (forceMap[`descriptions`] || from.Descriptions != ``) {
		t.Descriptions = S.Trim(from.Descriptions)
		changed = true
	}
	if !excludeMap[`qty`] && (forceMap[`qty`] || from.Qty != 0) {
		t.Qty = from.Qty
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
