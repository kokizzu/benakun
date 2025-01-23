package wcFinance

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mFinance/rqFinance"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// BusinessTransactionMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcFinance__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcFinance__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcFinance__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcFinance__ORM.GEN.go
type BusinessTransactionMutator struct {
	rqFinance.BusinessTransaction
	mutations []A.X
	logs      []A.X
}

// NewBusinessTransactionMutator create new ORM writer/command object
func NewBusinessTransactionMutator(adapter *Tt.Adapter) (res *BusinessTransactionMutator) {
	res = &BusinessTransactionMutator{BusinessTransaction: rqFinance.BusinessTransaction{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (b *BusinessTransactionMutator) Logs() []A.X { //nolint:dupl false positive
	return b.logs
}

// HaveMutation check whether Set* methods ever called
func (b *BusinessTransactionMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(b.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (b *BusinessTransactionMutator) ClearMutations() { //nolint:dupl false positive
	b.mutations = []A.X{}
	b.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (b *BusinessTransactionMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := b.Adapter.Update(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id}, b.ToUpdateArray())
	return !L.IsError(err, `BusinessTransaction.DoOverwriteById failed: `+b.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (b *BusinessTransactionMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !b.HaveMutation() {
		return true
	}
	_, err := b.Adapter.Update(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id}, b.mutations)
	return !L.IsError(err, `BusinessTransaction.DoUpdateById failed: `+b.SpaceName())
}

// DoDeletePermanentById permanent delete
func (b *BusinessTransactionMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := b.Adapter.Delete(b.SpaceName(), b.UniqueIndexId(), A.X{b.Id})
	return !L.IsError(err, `BusinessTransaction.DoDeletePermanentById failed: `+b.SpaceName())
}

// func (b *BusinessTransactionMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := b.ToArray()
//	_, err := b.Adapter.Upsert(b.SpaceName(), arr, A.X{
//		A.X{`=`, 0, b.Id},
//		A.X{`=`, 1, b.TenantCode},
//		A.X{`=`, 2, b.StartDate},
//		A.X{`=`, 3, b.EndDate},
//		A.X{`=`, 4, b.CreatedAt},
//		A.X{`=`, 5, b.CreatedBy},
//		A.X{`=`, 6, b.UpdatedAt},
//		A.X{`=`, 7, b.UpdatedBy},
//		A.X{`=`, 8, b.DeletedAt},
//		A.X{`=`, 9, b.DeletedBy},
//		A.X{`=`, 10, b.RestoredBy},
//		A.X{`=`, 11, b.TransactionTemplateId},
//	})
//	return !L.IsError(err, `BusinessTransaction.DoUpsert failed: `+b.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (b *BusinessTransactionMutator) DoInsert() bool { //nolint:dupl false positive
	arr := b.ToArray()
	row, err := b.Adapter.Insert(b.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			b.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `BusinessTransaction.DoInsert failed: `+b.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (b *BusinessTransactionMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := b.ToArray()
	row, err := b.Adapter.Replace(b.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			b.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `BusinessTransaction.DoUpsert failed: `+b.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != b.Id {
		b.mutations = append(b.mutations, A.X{`=`, 0, val})
		b.logs = append(b.logs, A.X{`id`, b.Id, val})
		b.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != b.TenantCode {
		b.mutations = append(b.mutations, A.X{`=`, 1, val})
		b.logs = append(b.logs, A.X{`tenantCode`, b.TenantCode, val})
		b.TenantCode = val
		return true
	}
	return false
}

// SetStartDate create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetStartDate(val string) bool { //nolint:dupl false positive
	if val != b.StartDate {
		b.mutations = append(b.mutations, A.X{`=`, 2, val})
		b.logs = append(b.logs, A.X{`startDate`, b.StartDate, val})
		b.StartDate = val
		return true
	}
	return false
}

// SetEndDate create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetEndDate(val string) bool { //nolint:dupl false positive
	if val != b.EndDate {
		b.mutations = append(b.mutations, A.X{`=`, 3, val})
		b.logs = append(b.logs, A.X{`endDate`, b.EndDate, val})
		b.EndDate = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != b.CreatedAt {
		b.mutations = append(b.mutations, A.X{`=`, 4, val})
		b.logs = append(b.logs, A.X{`createdAt`, b.CreatedAt, val})
		b.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.CreatedBy {
		b.mutations = append(b.mutations, A.X{`=`, 5, val})
		b.logs = append(b.logs, A.X{`createdBy`, b.CreatedBy, val})
		b.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != b.UpdatedAt {
		b.mutations = append(b.mutations, A.X{`=`, 6, val})
		b.logs = append(b.logs, A.X{`updatedAt`, b.UpdatedAt, val})
		b.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.UpdatedBy {
		b.mutations = append(b.mutations, A.X{`=`, 7, val})
		b.logs = append(b.logs, A.X{`updatedBy`, b.UpdatedBy, val})
		b.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != b.DeletedAt {
		b.mutations = append(b.mutations, A.X{`=`, 8, val})
		b.logs = append(b.logs, A.X{`deletedAt`, b.DeletedAt, val})
		b.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.DeletedBy {
		b.mutations = append(b.mutations, A.X{`=`, 9, val})
		b.logs = append(b.logs, A.X{`deletedBy`, b.DeletedBy, val})
		b.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != b.RestoredBy {
		b.mutations = append(b.mutations, A.X{`=`, 10, val})
		b.logs = append(b.logs, A.X{`restoredBy`, b.RestoredBy, val})
		b.RestoredBy = val
		return true
	}
	return false
}

// SetTransactionTemplateId create mutations, should not duplicate
func (b *BusinessTransactionMutator) SetTransactionTemplateId(val uint64) bool { //nolint:dupl false positive
	if val != b.TransactionTemplateId {
		b.mutations = append(b.mutations, A.X{`=`, 11, val})
		b.logs = append(b.logs, A.X{`transactionTemplateId`, b.TransactionTemplateId, val})
		b.TransactionTemplateId = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (b *BusinessTransactionMutator) SetAll(from rqFinance.BusinessTransaction, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		b.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		b.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`startDate`] && (forceMap[`startDate`] || from.StartDate != ``) {
		b.StartDate = S.Trim(from.StartDate)
		changed = true
	}
	if !excludeMap[`endDate`] && (forceMap[`endDate`] || from.EndDate != ``) {
		b.EndDate = S.Trim(from.EndDate)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		b.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		b.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		b.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		b.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		b.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		b.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		b.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`transactionTemplateId`] && (forceMap[`transactionTemplateId`] || from.TransactionTemplateId != 0) {
		b.TransactionTemplateId = from.TransactionTemplateId
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// CoaMutator DAO writer/command struct
type CoaMutator struct {
	rqFinance.Coa
	mutations []A.X
	logs      []A.X
}

// NewCoaMutator create new ORM writer/command object
func NewCoaMutator(adapter *Tt.Adapter) (res *CoaMutator) {
	res = &CoaMutator{Coa: rqFinance.Coa{Adapter: adapter}}
	res.Children = []any{}
	return
}

// Logs get array of logs [field, old, new]
func (c *CoaMutator) Logs() []A.X { //nolint:dupl false positive
	return c.logs
}

// HaveMutation check whether Set* methods ever called
func (c *CoaMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(c.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (c *CoaMutator) ClearMutations() { //nolint:dupl false positive
	c.mutations = []A.X{}
	c.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (c *CoaMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := c.Adapter.Update(c.SpaceName(), c.UniqueIndexId(), A.X{c.Id}, c.ToUpdateArray())
	return !L.IsError(err, `Coa.DoOverwriteById failed: `+c.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (c *CoaMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !c.HaveMutation() {
		return true
	}
	_, err := c.Adapter.Update(c.SpaceName(), c.UniqueIndexId(), A.X{c.Id}, c.mutations)
	return !L.IsError(err, `Coa.DoUpdateById failed: `+c.SpaceName())
}

// DoDeletePermanentById permanent delete
func (c *CoaMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := c.Adapter.Delete(c.SpaceName(), c.UniqueIndexId(), A.X{c.Id})
	return !L.IsError(err, `Coa.DoDeletePermanentById failed: `+c.SpaceName())
}

// func (c *CoaMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := c.ToArray()
//	_, err := c.Adapter.Upsert(c.SpaceName(), arr, A.X{
//		A.X{`=`, 0, c.Id},
//		A.X{`=`, 1, c.TenantCode},
//		A.X{`=`, 2, c.Name},
//		A.X{`=`, 3, c.Label},
//		A.X{`=`, 4, c.ParentId},
//		A.X{`=`, 5, c.Children},
//		A.X{`=`, 6, c.CreatedAt},
//		A.X{`=`, 7, c.CreatedBy},
//		A.X{`=`, 8, c.UpdatedAt},
//		A.X{`=`, 9, c.UpdatedBy},
//		A.X{`=`, 10, c.DeletedAt},
//		A.X{`=`, 11, c.DeletedBy},
//		A.X{`=`, 12, c.RestoredBy},
//		A.X{`=`, 13, c.Editable},
//	})
//	return !L.IsError(err, `Coa.DoUpsert failed: `+c.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (c *CoaMutator) DoInsert() bool { //nolint:dupl false positive
	arr := c.ToArray()
	row, err := c.Adapter.Insert(c.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			c.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Coa.DoInsert failed: `+c.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (c *CoaMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := c.ToArray()
	row, err := c.Adapter.Replace(c.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			c.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Coa.DoUpsert failed: `+c.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (c *CoaMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != c.Id {
		c.mutations = append(c.mutations, A.X{`=`, 0, val})
		c.logs = append(c.logs, A.X{`id`, c.Id, val})
		c.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (c *CoaMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != c.TenantCode {
		c.mutations = append(c.mutations, A.X{`=`, 1, val})
		c.logs = append(c.logs, A.X{`tenantCode`, c.TenantCode, val})
		c.TenantCode = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (c *CoaMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != c.Name {
		c.mutations = append(c.mutations, A.X{`=`, 2, val})
		c.logs = append(c.logs, A.X{`name`, c.Name, val})
		c.Name = val
		return true
	}
	return false
}

// SetLabel create mutations, should not duplicate
func (c *CoaMutator) SetLabel(val string) bool { //nolint:dupl false positive
	if val != c.Label {
		c.mutations = append(c.mutations, A.X{`=`, 3, val})
		c.logs = append(c.logs, A.X{`label`, c.Label, val})
		c.Label = val
		return true
	}
	return false
}

// SetParentId create mutations, should not duplicate
func (c *CoaMutator) SetParentId(val uint64) bool { //nolint:dupl false positive
	if val != c.ParentId {
		c.mutations = append(c.mutations, A.X{`=`, 4, val})
		c.logs = append(c.logs, A.X{`parentId`, c.ParentId, val})
		c.ParentId = val
		return true
	}
	return false
}

// SetChildren create mutations, should not duplicate
func (c *CoaMutator) SetChildren(val []any) bool { //nolint:dupl false positive
	c.mutations = append(c.mutations, A.X{`=`, 5, val})
	c.logs = append(c.logs, A.X{`children`, c.Children, val})
	c.Children = val
	return true
}

// SetCreatedAt create mutations, should not duplicate
func (c *CoaMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != c.CreatedAt {
		c.mutations = append(c.mutations, A.X{`=`, 6, val})
		c.logs = append(c.logs, A.X{`createdAt`, c.CreatedAt, val})
		c.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (c *CoaMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != c.CreatedBy {
		c.mutations = append(c.mutations, A.X{`=`, 7, val})
		c.logs = append(c.logs, A.X{`createdBy`, c.CreatedBy, val})
		c.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (c *CoaMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != c.UpdatedAt {
		c.mutations = append(c.mutations, A.X{`=`, 8, val})
		c.logs = append(c.logs, A.X{`updatedAt`, c.UpdatedAt, val})
		c.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (c *CoaMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != c.UpdatedBy {
		c.mutations = append(c.mutations, A.X{`=`, 9, val})
		c.logs = append(c.logs, A.X{`updatedBy`, c.UpdatedBy, val})
		c.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (c *CoaMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != c.DeletedAt {
		c.mutations = append(c.mutations, A.X{`=`, 10, val})
		c.logs = append(c.logs, A.X{`deletedAt`, c.DeletedAt, val})
		c.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (c *CoaMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != c.DeletedBy {
		c.mutations = append(c.mutations, A.X{`=`, 11, val})
		c.logs = append(c.logs, A.X{`deletedBy`, c.DeletedBy, val})
		c.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (c *CoaMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != c.RestoredBy {
		c.mutations = append(c.mutations, A.X{`=`, 12, val})
		c.logs = append(c.logs, A.X{`restoredBy`, c.RestoredBy, val})
		c.RestoredBy = val
		return true
	}
	return false
}

// SetEditable create mutations, should not duplicate
func (c *CoaMutator) SetEditable(val bool) bool { //nolint:dupl false positive
	if val != c.Editable {
		c.mutations = append(c.mutations, A.X{`=`, 13, val})
		c.logs = append(c.logs, A.X{`editable`, c.Editable, val})
		c.Editable = val
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
	if !excludeMap[`label`] && (forceMap[`label`] || from.Label != ``) {
		c.Label = S.Trim(from.Label)
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
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		c.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		c.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`editable`] && (forceMap[`editable`] || from.Editable != false) {
		c.Editable = from.Editable
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// TransactionJournalMutator DAO writer/command struct
type TransactionJournalMutator struct {
	rqFinance.TransactionJournal
	mutations []A.X
	logs      []A.X
}

// NewTransactionJournalMutator create new ORM writer/command object
func NewTransactionJournalMutator(adapter *Tt.Adapter) (res *TransactionJournalMutator) {
	res = &TransactionJournalMutator{TransactionJournal: rqFinance.TransactionJournal{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (t *TransactionJournalMutator) Logs() []A.X { //nolint:dupl false positive
	return t.logs
}

// HaveMutation check whether Set* methods ever called
func (t *TransactionJournalMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(t.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (t *TransactionJournalMutator) ClearMutations() { //nolint:dupl false positive
	t.mutations = []A.X{}
	t.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (t *TransactionJournalMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.ToUpdateArray())
	return !L.IsError(err, `TransactionJournal.DoOverwriteById failed: `+t.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TransactionJournalMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.mutations)
	return !L.IsError(err, `TransactionJournal.DoUpdateById failed: `+t.SpaceName())
}

// DoDeletePermanentById permanent delete
func (t *TransactionJournalMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Delete(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id})
	return !L.IsError(err, `TransactionJournal.DoDeletePermanentById failed: `+t.SpaceName())
}

// func (t *TransactionJournalMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := t.ToArray()
//	_, err := t.Adapter.Upsert(t.SpaceName(), arr, A.X{
//		A.X{`=`, 0, t.Id},
//		A.X{`=`, 1, t.TenantCode},
//		A.X{`=`, 2, t.CoaId},
//		A.X{`=`, 3, t.DebitIDR},
//		A.X{`=`, 4, t.CreditIDR},
//		A.X{`=`, 5, t.Descriptions},
//		A.X{`=`, 6, t.Date},
//		A.X{`=`, 7, t.DetailObj},
//		A.X{`=`, 8, t.CreatedAt},
//		A.X{`=`, 9, t.CreatedBy},
//		A.X{`=`, 10, t.UpdatedAt},
//		A.X{`=`, 11, t.UpdatedBy},
//		A.X{`=`, 12, t.DeletedAt},
//		A.X{`=`, 13, t.DeletedBy},
//		A.X{`=`, 14, t.RestoredBy},
//		A.X{`=`, 15, t.TransactionTemplateId},
//	})
//	return !L.IsError(err, `TransactionJournal.DoUpsert failed: `+t.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (t *TransactionJournalMutator) DoInsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.Insert(t.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `TransactionJournal.DoInsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (t *TransactionJournalMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.Replace(t.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `TransactionJournal.DoUpsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (t *TransactionJournalMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != t.Id {
		t.mutations = append(t.mutations, A.X{`=`, 0, val})
		t.logs = append(t.logs, A.X{`id`, t.Id, val})
		t.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (t *TransactionJournalMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != t.TenantCode {
		t.mutations = append(t.mutations, A.X{`=`, 1, val})
		t.logs = append(t.logs, A.X{`tenantCode`, t.TenantCode, val})
		t.TenantCode = val
		return true
	}
	return false
}

// SetCoaId create mutations, should not duplicate
func (t *TransactionJournalMutator) SetCoaId(val uint64) bool { //nolint:dupl false positive
	if val != t.CoaId {
		t.mutations = append(t.mutations, A.X{`=`, 2, val})
		t.logs = append(t.logs, A.X{`coaId`, t.CoaId, val})
		t.CoaId = val
		return true
	}
	return false
}

// SetDebitIDR create mutations, should not duplicate
func (t *TransactionJournalMutator) SetDebitIDR(val int64) bool { //nolint:dupl false positive
	if val != t.DebitIDR {
		t.mutations = append(t.mutations, A.X{`=`, 3, val})
		t.logs = append(t.logs, A.X{`debitIDR`, t.DebitIDR, val})
		t.DebitIDR = val
		return true
	}
	return false
}

// SetCreditIDR create mutations, should not duplicate
func (t *TransactionJournalMutator) SetCreditIDR(val int64) bool { //nolint:dupl false positive
	if val != t.CreditIDR {
		t.mutations = append(t.mutations, A.X{`=`, 4, val})
		t.logs = append(t.logs, A.X{`creditIDR`, t.CreditIDR, val})
		t.CreditIDR = val
		return true
	}
	return false
}

// SetDescriptions create mutations, should not duplicate
func (t *TransactionJournalMutator) SetDescriptions(val string) bool { //nolint:dupl false positive
	if val != t.Descriptions {
		t.mutations = append(t.mutations, A.X{`=`, 5, val})
		t.logs = append(t.logs, A.X{`descriptions`, t.Descriptions, val})
		t.Descriptions = val
		return true
	}
	return false
}

// SetDate create mutations, should not duplicate
func (t *TransactionJournalMutator) SetDate(val string) bool { //nolint:dupl false positive
	if val != t.Date {
		t.mutations = append(t.mutations, A.X{`=`, 6, val})
		t.logs = append(t.logs, A.X{`date`, t.Date, val})
		t.Date = val
		return true
	}
	return false
}

// SetDetailObj create mutations, should not duplicate
func (t *TransactionJournalMutator) SetDetailObj(val string) bool { //nolint:dupl false positive
	if val != t.DetailObj {
		t.mutations = append(t.mutations, A.X{`=`, 7, val})
		t.logs = append(t.logs, A.X{`detailObj`, t.DetailObj, val})
		t.DetailObj = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (t *TransactionJournalMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.CreatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 8, val})
		t.logs = append(t.logs, A.X{`createdAt`, t.CreatedAt, val})
		t.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (t *TransactionJournalMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.CreatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 9, val})
		t.logs = append(t.logs, A.X{`createdBy`, t.CreatedBy, val})
		t.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (t *TransactionJournalMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.UpdatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 10, val})
		t.logs = append(t.logs, A.X{`updatedAt`, t.UpdatedAt, val})
		t.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (t *TransactionJournalMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.UpdatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 11, val})
		t.logs = append(t.logs, A.X{`updatedBy`, t.UpdatedBy, val})
		t.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (t *TransactionJournalMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != t.DeletedAt {
		t.mutations = append(t.mutations, A.X{`=`, 12, val})
		t.logs = append(t.logs, A.X{`deletedAt`, t.DeletedAt, val})
		t.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (t *TransactionJournalMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.DeletedBy {
		t.mutations = append(t.mutations, A.X{`=`, 13, val})
		t.logs = append(t.logs, A.X{`deletedBy`, t.DeletedBy, val})
		t.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (t *TransactionJournalMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != t.RestoredBy {
		t.mutations = append(t.mutations, A.X{`=`, 14, val})
		t.logs = append(t.logs, A.X{`restoredBy`, t.RestoredBy, val})
		t.RestoredBy = val
		return true
	}
	return false
}

// SetTransactionTemplateId create mutations, should not duplicate
func (t *TransactionJournalMutator) SetTransactionTemplateId(val uint64) bool { //nolint:dupl false positive
	if val != t.TransactionTemplateId {
		t.mutations = append(t.mutations, A.X{`=`, 15, val})
		t.logs = append(t.logs, A.X{`transactionTemplateId`, t.TransactionTemplateId, val})
		t.TransactionTemplateId = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (t *TransactionJournalMutator) SetAll(from rqFinance.TransactionJournal, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`coaId`] && (forceMap[`coaId`] || from.CoaId != 0) {
		t.CoaId = from.CoaId
		changed = true
	}
	if !excludeMap[`debitIDR`] && (forceMap[`debitIDR`] || from.DebitIDR != 0) {
		t.DebitIDR = from.DebitIDR
		changed = true
	}
	if !excludeMap[`creditIDR`] && (forceMap[`creditIDR`] || from.CreditIDR != 0) {
		t.CreditIDR = from.CreditIDR
		changed = true
	}
	if !excludeMap[`descriptions`] && (forceMap[`descriptions`] || from.Descriptions != ``) {
		t.Descriptions = S.Trim(from.Descriptions)
		changed = true
	}
	if !excludeMap[`date`] && (forceMap[`date`] || from.Date != ``) {
		t.Date = S.Trim(from.Date)
		changed = true
	}
	if !excludeMap[`detailObj`] && (forceMap[`detailObj`] || from.DetailObj != ``) {
		t.DetailObj = S.Trim(from.DetailObj)
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
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		t.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		t.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`transactionTemplateId`] && (forceMap[`transactionTemplateId`] || from.TransactionTemplateId != 0) {
		t.TransactionTemplateId = from.TransactionTemplateId
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// TransactionTemplateMutator DAO writer/command struct
type TransactionTemplateMutator struct {
	rqFinance.TransactionTemplate
	mutations []A.X
	logs      []A.X
}

// NewTransactionTemplateMutator create new ORM writer/command object
func NewTransactionTemplateMutator(adapter *Tt.Adapter) (res *TransactionTemplateMutator) {
	res = &TransactionTemplateMutator{TransactionTemplate: rqFinance.TransactionTemplate{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (t *TransactionTemplateMutator) Logs() []A.X { //nolint:dupl false positive
	return t.logs
}

// HaveMutation check whether Set* methods ever called
func (t *TransactionTemplateMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(t.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (t *TransactionTemplateMutator) ClearMutations() { //nolint:dupl false positive
	t.mutations = []A.X{}
	t.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (t *TransactionTemplateMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.ToUpdateArray())
	return !L.IsError(err, `TransactionTemplate.DoOverwriteById failed: `+t.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TransactionTemplateMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.mutations)
	return !L.IsError(err, `TransactionTemplate.DoUpdateById failed: `+t.SpaceName())
}

// DoDeletePermanentById permanent delete
func (t *TransactionTemplateMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Delete(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id})
	return !L.IsError(err, `TransactionTemplate.DoDeletePermanentById failed: `+t.SpaceName())
}

// func (t *TransactionTemplateMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := t.ToArray()
//	_, err := t.Adapter.Upsert(t.SpaceName(), arr, A.X{
//		A.X{`=`, 0, t.Id},
//		A.X{`=`, 1, t.TenantCode},
//		A.X{`=`, 2, t.Name},
//		A.X{`=`, 3, t.Color},
//		A.X{`=`, 4, t.ImageURL},
//		A.X{`=`, 5, t.CreatedAt},
//		A.X{`=`, 6, t.CreatedBy},
//		A.X{`=`, 7, t.UpdatedAt},
//		A.X{`=`, 8, t.UpdatedBy},
//		A.X{`=`, 9, t.DeletedAt},
//		A.X{`=`, 10, t.DeletedBy},
//		A.X{`=`, 11, t.RestoredBy},
//	})
//	return !L.IsError(err, `TransactionTemplate.DoUpsert failed: `+t.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (t *TransactionTemplateMutator) DoInsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.Insert(t.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `TransactionTemplate.DoInsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (t *TransactionTemplateMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.Replace(t.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `TransactionTemplate.DoUpsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != t.Id {
		t.mutations = append(t.mutations, A.X{`=`, 0, val})
		t.logs = append(t.logs, A.X{`id`, t.Id, val})
		t.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != t.TenantCode {
		t.mutations = append(t.mutations, A.X{`=`, 1, val})
		t.logs = append(t.logs, A.X{`tenantCode`, t.TenantCode, val})
		t.TenantCode = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != t.Name {
		t.mutations = append(t.mutations, A.X{`=`, 2, val})
		t.logs = append(t.logs, A.X{`name`, t.Name, val})
		t.Name = val
		return true
	}
	return false
}

// SetColor create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetColor(val string) bool { //nolint:dupl false positive
	if val != t.Color {
		t.mutations = append(t.mutations, A.X{`=`, 3, val})
		t.logs = append(t.logs, A.X{`color`, t.Color, val})
		t.Color = val
		return true
	}
	return false
}

// SetImageURL create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetImageURL(val string) bool { //nolint:dupl false positive
	if val != t.ImageURL {
		t.mutations = append(t.mutations, A.X{`=`, 4, val})
		t.logs = append(t.logs, A.X{`imageURL`, t.ImageURL, val})
		t.ImageURL = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.CreatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 5, val})
		t.logs = append(t.logs, A.X{`createdAt`, t.CreatedAt, val})
		t.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.CreatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 6, val})
		t.logs = append(t.logs, A.X{`createdBy`, t.CreatedBy, val})
		t.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.UpdatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 7, val})
		t.logs = append(t.logs, A.X{`updatedAt`, t.UpdatedAt, val})
		t.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.UpdatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 8, val})
		t.logs = append(t.logs, A.X{`updatedBy`, t.UpdatedBy, val})
		t.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != t.DeletedAt {
		t.mutations = append(t.mutations, A.X{`=`, 9, val})
		t.logs = append(t.logs, A.X{`deletedAt`, t.DeletedAt, val})
		t.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.DeletedBy {
		t.mutations = append(t.mutations, A.X{`=`, 10, val})
		t.logs = append(t.logs, A.X{`deletedBy`, t.DeletedBy, val})
		t.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (t *TransactionTemplateMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != t.RestoredBy {
		t.mutations = append(t.mutations, A.X{`=`, 11, val})
		t.logs = append(t.logs, A.X{`restoredBy`, t.RestoredBy, val})
		t.RestoredBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (t *TransactionTemplateMutator) SetAll(from rqFinance.TransactionTemplate, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`name`] && (forceMap[`name`] || from.Name != ``) {
		t.Name = S.Trim(from.Name)
		changed = true
	}
	if !excludeMap[`color`] && (forceMap[`color`] || from.Color != ``) {
		t.Color = S.Trim(from.Color)
		changed = true
	}
	if !excludeMap[`imageURL`] && (forceMap[`imageURL`] || from.ImageURL != ``) {
		t.ImageURL = S.Trim(from.ImageURL)
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
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		t.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		t.RestoredBy = from.RestoredBy
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// TransactionTplDetailMutator DAO writer/command struct
type TransactionTplDetailMutator struct {
	rqFinance.TransactionTplDetail
	mutations []A.X
	logs      []A.X
}

// NewTransactionTplDetailMutator create new ORM writer/command object
func NewTransactionTplDetailMutator(adapter *Tt.Adapter) (res *TransactionTplDetailMutator) {
	res = &TransactionTplDetailMutator{TransactionTplDetail: rqFinance.TransactionTplDetail{Adapter: adapter}}
	res.Attributes = []any{}
	return
}

// Logs get array of logs [field, old, new]
func (t *TransactionTplDetailMutator) Logs() []A.X { //nolint:dupl false positive
	return t.logs
}

// HaveMutation check whether Set* methods ever called
func (t *TransactionTplDetailMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(t.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (t *TransactionTplDetailMutator) ClearMutations() { //nolint:dupl false positive
	t.mutations = []A.X{}
	t.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (t *TransactionTplDetailMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.ToUpdateArray())
	return !L.IsError(err, `TransactionTplDetail.DoOverwriteById failed: `+t.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TransactionTplDetailMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.mutations)
	return !L.IsError(err, `TransactionTplDetail.DoUpdateById failed: `+t.SpaceName())
}

// DoDeletePermanentById permanent delete
func (t *TransactionTplDetailMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Delete(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id})
	return !L.IsError(err, `TransactionTplDetail.DoDeletePermanentById failed: `+t.SpaceName())
}

// func (t *TransactionTplDetailMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := t.ToArray()
//	_, err := t.Adapter.Upsert(t.SpaceName(), arr, A.X{
//		A.X{`=`, 0, t.Id},
//		A.X{`=`, 1, t.ParentId},
//		A.X{`=`, 2, t.TenantCode},
//		A.X{`=`, 3, t.CoaId},
//		A.X{`=`, 4, t.IsDebit},
//		A.X{`=`, 5, t.CreatedAt},
//		A.X{`=`, 6, t.CreatedBy},
//		A.X{`=`, 7, t.UpdatedAt},
//		A.X{`=`, 8, t.UpdatedBy},
//		A.X{`=`, 9, t.DeletedAt},
//		A.X{`=`, 10, t.DeletedBy},
//		A.X{`=`, 11, t.RestoredBy},
//		A.X{`=`, 12, t.Attributes},
//	})
//	return !L.IsError(err, `TransactionTplDetail.DoUpsert failed: `+t.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (t *TransactionTplDetailMutator) DoInsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.Insert(t.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `TransactionTplDetail.DoInsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (t *TransactionTplDetailMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.Replace(t.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `TransactionTplDetail.DoUpsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != t.Id {
		t.mutations = append(t.mutations, A.X{`=`, 0, val})
		t.logs = append(t.logs, A.X{`id`, t.Id, val})
		t.Id = val
		return true
	}
	return false
}

// SetParentId create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetParentId(val uint64) bool { //nolint:dupl false positive
	if val != t.ParentId {
		t.mutations = append(t.mutations, A.X{`=`, 1, val})
		t.logs = append(t.logs, A.X{`parentId`, t.ParentId, val})
		t.ParentId = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != t.TenantCode {
		t.mutations = append(t.mutations, A.X{`=`, 2, val})
		t.logs = append(t.logs, A.X{`tenantCode`, t.TenantCode, val})
		t.TenantCode = val
		return true
	}
	return false
}

// SetCoaId create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetCoaId(val uint64) bool { //nolint:dupl false positive
	if val != t.CoaId {
		t.mutations = append(t.mutations, A.X{`=`, 3, val})
		t.logs = append(t.logs, A.X{`coaId`, t.CoaId, val})
		t.CoaId = val
		return true
	}
	return false
}

// SetIsDebit create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetIsDebit(val bool) bool { //nolint:dupl false positive
	if val != t.IsDebit {
		t.mutations = append(t.mutations, A.X{`=`, 4, val})
		t.logs = append(t.logs, A.X{`isDebit`, t.IsDebit, val})
		t.IsDebit = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.CreatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 5, val})
		t.logs = append(t.logs, A.X{`createdAt`, t.CreatedAt, val})
		t.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.CreatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 6, val})
		t.logs = append(t.logs, A.X{`createdBy`, t.CreatedBy, val})
		t.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.UpdatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 7, val})
		t.logs = append(t.logs, A.X{`updatedAt`, t.UpdatedAt, val})
		t.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.UpdatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 8, val})
		t.logs = append(t.logs, A.X{`updatedBy`, t.UpdatedBy, val})
		t.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != t.DeletedAt {
		t.mutations = append(t.mutations, A.X{`=`, 9, val})
		t.logs = append(t.logs, A.X{`deletedAt`, t.DeletedAt, val})
		t.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.DeletedBy {
		t.mutations = append(t.mutations, A.X{`=`, 10, val})
		t.logs = append(t.logs, A.X{`deletedBy`, t.DeletedBy, val})
		t.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != t.RestoredBy {
		t.mutations = append(t.mutations, A.X{`=`, 11, val})
		t.logs = append(t.logs, A.X{`restoredBy`, t.RestoredBy, val})
		t.RestoredBy = val
		return true
	}
	return false
}

// SetAttributes create mutations, should not duplicate
func (t *TransactionTplDetailMutator) SetAttributes(val []any) bool { //nolint:dupl false positive
	t.mutations = append(t.mutations, A.X{`=`, 12, val})
	t.logs = append(t.logs, A.X{`attributes`, t.Attributes, val})
	t.Attributes = val
	return true
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (t *TransactionTplDetailMutator) SetAll(from rqFinance.TransactionTplDetail, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`parentId`] && (forceMap[`parentId`] || from.ParentId != 0) {
		t.ParentId = from.ParentId
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		t.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`coaId`] && (forceMap[`coaId`] || from.CoaId != 0) {
		t.CoaId = from.CoaId
		changed = true
	}
	if !excludeMap[`isDebit`] && (forceMap[`isDebit`] || from.IsDebit != false) {
		t.IsDebit = from.IsDebit
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
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		t.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		t.RestoredBy = from.RestoredBy
		changed = true
	}
	if !excludeMap[`attributes`] && (forceMap[`attributes`] || from.Attributes != nil) {
		t.Attributes = from.Attributes
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
