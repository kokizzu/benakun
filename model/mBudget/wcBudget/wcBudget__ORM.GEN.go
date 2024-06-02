package wcBudget

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mBudget/rqBudget"

	"github.com/tarantool/go-tarantool/v2"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// BankAccountsMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcBudget__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcBudget__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcBudget__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcBudget__ORM.GEN.go
type BankAccountsMutator struct {
	rqBudget.BankAccounts
	mutations *tarantool.Operations
	logs      []A.X
}

// NewBankAccountsMutator create new ORM writer/command object
func NewBankAccountsMutator(adapter *Tt.Adapter) (res *BankAccountsMutator) {
	res = &BankAccountsMutator{BankAccounts: rqBudget.BankAccounts{Adapter: adapter}}
	res.mutations = tarantool.NewOperations()
	return
}

// Logs get array of logs [field, old, new]
func (b *BankAccountsMutator) Logs() []A.X { //nolint:dupl false positive
	return b.logs
}

// HaveMutation check whether Set* methods ever called
func (b *BankAccountsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(b.logs) > 0
}

// ClearMutations clear all previously called Set* methods
func (b *BankAccountsMutator) ClearMutations() { //nolint:dupl false positive
	b.mutations = tarantool.NewOperations()
	b.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (b *BankAccountsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := b.Adapter.RetryDo(tarantool.NewUpdateRequest(b.SpaceName()).
		Index(b.UniqueIndexId()).
		Key(tarantool.UintKey{I: uint(b.Id)}).
		Operations(b.ToUpdateArray()),
	)
	return !L.IsError(err, `BankAccounts.DoOverwriteById failed: `+b.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (b *BankAccountsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !b.HaveMutation() {
		return true
	}
	_, err := b.Adapter.RetryDo(
		tarantool.NewUpdateRequest(b.SpaceName()).
			Index(b.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(b.Id)}).
			Operations(b.mutations),
	)
	return !L.IsError(err, `BankAccounts.DoUpdateById failed: `+b.SpaceName())
}

// DoDeletePermanentById permanent delete
func (b *BankAccountsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := b.Adapter.RetryDo(
		tarantool.NewDeleteRequest(b.SpaceName()).
			Index(b.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(b.Id)}),
	)
	return !L.IsError(err, `BankAccounts.DoDeletePermanentById failed: `+b.SpaceName())
}

// DoInsert insert, error if already exists
func (b *BankAccountsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := b.ToArray()
	row, err := b.Adapter.RetryDo(
		tarantool.NewInsertRequest(b.SpaceName()).
			Tuple(arr),
	)
	if err == nil {
		if len(row) > 0 {
			if cells, ok := row[0].([]any); ok && len(cells) > 0 {
				b.Id = X.ToU(cells[0])
			}
		}
	}
	return !L.IsError(err, `BankAccounts.DoInsert failed: `+b.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// tarantool's replace/upsert can only match by primary key
// previous name: DoReplace
func (b *BankAccountsMutator) DoUpsertById() bool { //nolint:dupl false positive
	if b.Id > 0 {
		return b.DoUpdateById()
	}
	return b.DoInsert()
}

// SetId create mutations, should not duplicate
func (b *BankAccountsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != b.Id {
		b.mutations.Assign(0, val)
		b.logs = append(b.logs, A.X{`id`, b.Id, val})
		b.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (b *BankAccountsMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != b.TenantCode {
		b.mutations.Assign(1, val)
		b.logs = append(b.logs, A.X{`tenantCode`, b.TenantCode, val})
		b.TenantCode = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (b *BankAccountsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != b.CreatedAt {
		b.mutations.Assign(2, val)
		b.logs = append(b.logs, A.X{`createdAt`, b.CreatedAt, val})
		b.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (b *BankAccountsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.CreatedBy {
		b.mutations.Assign(3, val)
		b.logs = append(b.logs, A.X{`createdBy`, b.CreatedBy, val})
		b.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (b *BankAccountsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != b.UpdatedAt {
		b.mutations.Assign(4, val)
		b.logs = append(b.logs, A.X{`updatedAt`, b.UpdatedAt, val})
		b.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (b *BankAccountsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.UpdatedBy {
		b.mutations.Assign(5, val)
		b.logs = append(b.logs, A.X{`updatedBy`, b.UpdatedBy, val})
		b.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (b *BankAccountsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != b.DeletedAt {
		b.mutations.Assign(6, val)
		b.logs = append(b.logs, A.X{`deletedAt`, b.DeletedAt, val})
		b.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (b *BankAccountsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != b.DeletedBy {
		b.mutations.Assign(7, val)
		b.logs = append(b.logs, A.X{`deletedBy`, b.DeletedBy, val})
		b.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (b *BankAccountsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != b.RestoredBy {
		b.mutations.Assign(8, val)
		b.logs = append(b.logs, A.X{`restoredBy`, b.RestoredBy, val})
		b.RestoredBy = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (b *BankAccountsMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != b.Name {
		b.mutations.Assign(9, val)
		b.logs = append(b.logs, A.X{`name`, b.Name, val})
		b.Name = val
		return true
	}
	return false
}

// SetParentBankAccountId create mutations, should not duplicate
func (b *BankAccountsMutator) SetParentBankAccountId(val uint64) bool { //nolint:dupl false positive
	if val != b.ParentBankAccountId {
		b.mutations.Assign(10, val)
		b.logs = append(b.logs, A.X{`parentBankAccountId`, b.ParentBankAccountId, val})
		b.ParentBankAccountId = val
		return true
	}
	return false
}

// SetChildBankAccountId create mutations, should not duplicate
func (b *BankAccountsMutator) SetChildBankAccountId(val uint64) bool { //nolint:dupl false positive
	if val != b.ChildBankAccountId {
		b.mutations.Assign(11, val)
		b.logs = append(b.logs, A.X{`childBankAccountId`, b.ChildBankAccountId, val})
		b.ChildBankAccountId = val
		return true
	}
	return false
}

// SetAccountNumber create mutations, should not duplicate
func (b *BankAccountsMutator) SetAccountNumber(val int64) bool { //nolint:dupl false positive
	if val != b.AccountNumber {
		b.mutations.Assign(12, val)
		b.logs = append(b.logs, A.X{`accountNumber`, b.AccountNumber, val})
		b.AccountNumber = val
		return true
	}
	return false
}

// SetBankName create mutations, should not duplicate
func (b *BankAccountsMutator) SetBankName(val string) bool { //nolint:dupl false positive
	if val != b.BankName {
		b.mutations.Assign(13, val)
		b.logs = append(b.logs, A.X{`bankName`, b.BankName, val})
		b.BankName = val
		return true
	}
	return false
}

// SetAccountName create mutations, should not duplicate
func (b *BankAccountsMutator) SetAccountName(val string) bool { //nolint:dupl false positive
	if val != b.AccountName {
		b.mutations.Assign(14, val)
		b.logs = append(b.logs, A.X{`accountName`, b.AccountName, val})
		b.AccountName = val
		return true
	}
	return false
}

// SetIsProfitCenter create mutations, should not duplicate
func (b *BankAccountsMutator) SetIsProfitCenter(val bool) bool { //nolint:dupl false positive
	if val != b.IsProfitCenter {
		b.mutations.Assign(15, val)
		b.logs = append(b.logs, A.X{`isProfitCenter `, b.IsProfitCenter, val})
		b.IsProfitCenter = val
		return true
	}
	return false
}

// SetIsCostCenter create mutations, should not duplicate
func (b *BankAccountsMutator) SetIsCostCenter(val bool) bool { //nolint:dupl false positive
	if val != b.IsCostCenter {
		b.mutations.Assign(16, val)
		b.logs = append(b.logs, A.X{`isCostCenter`, b.IsCostCenter, val})
		b.IsCostCenter = val
		return true
	}
	return false
}

// SetStaffId create mutations, should not duplicate
func (b *BankAccountsMutator) SetStaffId(val uint64) bool { //nolint:dupl false positive
	if val != b.StaffId {
		b.mutations.Assign(17, val)
		b.logs = append(b.logs, A.X{`staffId`, b.StaffId, val})
		b.StaffId = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (b *BankAccountsMutator) SetAll(from rqBudget.BankAccounts, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`name`] && (forceMap[`name`] || from.Name != ``) {
		b.Name = S.Trim(from.Name)
		changed = true
	}
	if !excludeMap[`parentBankAccountId`] && (forceMap[`parentBankAccountId`] || from.ParentBankAccountId != 0) {
		b.ParentBankAccountId = from.ParentBankAccountId
		changed = true
	}
	if !excludeMap[`childBankAccountId`] && (forceMap[`childBankAccountId`] || from.ChildBankAccountId != 0) {
		b.ChildBankAccountId = from.ChildBankAccountId
		changed = true
	}
	if !excludeMap[`accountNumber`] && (forceMap[`accountNumber`] || from.AccountNumber != 0) {
		b.AccountNumber = from.AccountNumber
		changed = true
	}
	if !excludeMap[`bankName`] && (forceMap[`bankName`] || from.BankName != ``) {
		b.BankName = S.Trim(from.BankName)
		changed = true
	}
	if !excludeMap[`accountName`] && (forceMap[`accountName`] || from.AccountName != ``) {
		b.AccountName = S.Trim(from.AccountName)
		changed = true
	}
	if !excludeMap[`isProfitCenter `] && (forceMap[`isProfitCenter `] || from.IsProfitCenter != false) {
		b.IsProfitCenter = from.IsProfitCenter
		changed = true
	}
	if !excludeMap[`isCostCenter`] && (forceMap[`isCostCenter`] || from.IsCostCenter != false) {
		b.IsCostCenter = from.IsCostCenter
		changed = true
	}
	if !excludeMap[`staffId`] && (forceMap[`staffId`] || from.StaffId != 0) {
		b.StaffId = from.StaffId
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// PlansMutator DAO writer/command struct
type PlansMutator struct {
	rqBudget.Plans
	mutations *tarantool.Operations
	logs      []A.X
}

// NewPlansMutator create new ORM writer/command object
func NewPlansMutator(adapter *Tt.Adapter) (res *PlansMutator) {
	res = &PlansMutator{Plans: rqBudget.Plans{Adapter: adapter}}
	res.mutations = tarantool.NewOperations()
	return
}

// Logs get array of logs [field, old, new]
func (p *PlansMutator) Logs() []A.X { //nolint:dupl false positive
	return p.logs
}

// HaveMutation check whether Set* methods ever called
func (p *PlansMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.logs) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *PlansMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = tarantool.NewOperations()
	p.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (p *PlansMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := p.Adapter.RetryDo(tarantool.NewUpdateRequest(p.SpaceName()).
		Index(p.UniqueIndexId()).
		Key(tarantool.UintKey{I: uint(p.Id)}).
		Operations(p.ToUpdateArray()),
	)
	return !L.IsError(err, `Plans.DoOverwriteById failed: `+p.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PlansMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.RetryDo(
		tarantool.NewUpdateRequest(p.SpaceName()).
			Index(p.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(p.Id)}).
			Operations(p.mutations),
	)
	return !L.IsError(err, `Plans.DoUpdateById failed: `+p.SpaceName())
}

// DoDeletePermanentById permanent delete
func (p *PlansMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := p.Adapter.RetryDo(
		tarantool.NewDeleteRequest(p.SpaceName()).
			Index(p.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(p.Id)}),
	)
	return !L.IsError(err, `Plans.DoDeletePermanentById failed: `+p.SpaceName())
}

// DoInsert insert, error if already exists
func (p *PlansMutator) DoInsert() bool { //nolint:dupl false positive
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
	return !L.IsError(err, `Plans.DoInsert failed: `+p.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// tarantool's replace/upsert can only match by primary key
// previous name: DoReplace
func (p *PlansMutator) DoUpsertById() bool { //nolint:dupl false positive
	if p.Id > 0 {
		return p.DoUpdateById()
	}
	return p.DoInsert()
}

// SetId create mutations, should not duplicate
func (p *PlansMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != p.Id {
		p.mutations.Assign(0, val)
		p.logs = append(p.logs, A.X{`id`, p.Id, val})
		p.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (p *PlansMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != p.TenantCode {
		p.mutations.Assign(1, val)
		p.logs = append(p.logs, A.X{`tenantCode`, p.TenantCode, val})
		p.TenantCode = val
		return true
	}
	return false
}

// SetPlanType create mutations, should not duplicate
func (p *PlansMutator) SetPlanType(val string) bool { //nolint:dupl false positive
	if val != p.PlanType {
		p.mutations.Assign(2, val)
		p.logs = append(p.logs, A.X{`planType`, p.PlanType, val})
		p.PlanType = val
		return true
	}
	return false
}

// SetParentId create mutations, should not duplicate
func (p *PlansMutator) SetParentId(val uint64) bool { //nolint:dupl false positive
	if val != p.ParentId {
		p.mutations.Assign(3, val)
		p.logs = append(p.logs, A.X{`parentId`, p.ParentId, val})
		p.ParentId = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (p *PlansMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.CreatedAt {
		p.mutations.Assign(4, val)
		p.logs = append(p.logs, A.X{`createdAt`, p.CreatedAt, val})
		p.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (p *PlansMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.CreatedBy {
		p.mutations.Assign(5, val)
		p.logs = append(p.logs, A.X{`createdBy`, p.CreatedBy, val})
		p.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (p *PlansMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.UpdatedAt {
		p.mutations.Assign(6, val)
		p.logs = append(p.logs, A.X{`updatedAt`, p.UpdatedAt, val})
		p.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (p *PlansMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.UpdatedBy {
		p.mutations.Assign(7, val)
		p.logs = append(p.logs, A.X{`updatedBy`, p.UpdatedBy, val})
		p.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (p *PlansMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != p.DeletedAt {
		p.mutations.Assign(8, val)
		p.logs = append(p.logs, A.X{`deletedAt`, p.DeletedAt, val})
		p.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (p *PlansMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.DeletedBy {
		p.mutations.Assign(9, val)
		p.logs = append(p.logs, A.X{`deletedBy`, p.DeletedBy, val})
		p.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (p *PlansMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != p.RestoredBy {
		p.mutations.Assign(10, val)
		p.logs = append(p.logs, A.X{`restoredBy`, p.RestoredBy, val})
		p.RestoredBy = val
		return true
	}
	return false
}

// SetTitle create mutations, should not duplicate
func (p *PlansMutator) SetTitle(val string) bool { //nolint:dupl false positive
	if val != p.Title {
		p.mutations.Assign(11, val)
		p.logs = append(p.logs, A.X{`title`, p.Title, val})
		p.Title = val
		return true
	}
	return false
}

// SetDescription create mutations, should not duplicate
func (p *PlansMutator) SetDescription(val string) bool { //nolint:dupl false positive
	if val != p.Description {
		p.mutations.Assign(12, val)
		p.logs = append(p.logs, A.X{`description`, p.Description, val})
		p.Description = val
		return true
	}
	return false
}

// SetOrgId create mutations, should not duplicate
func (p *PlansMutator) SetOrgId(val uint64) bool { //nolint:dupl false positive
	if val != p.OrgId {
		p.mutations.Assign(13, val)
		p.logs = append(p.logs, A.X{`orgId`, p.OrgId, val})
		p.OrgId = val
		return true
	}
	return false
}

// SetYearOf create mutations, should not duplicate
func (p *PlansMutator) SetYearOf(val uint64) bool { //nolint:dupl false positive
	if val != p.YearOf {
		p.mutations.Assign(14, val)
		p.logs = append(p.logs, A.X{`yearOf`, p.YearOf, val})
		p.YearOf = val
		return true
	}
	return false
}

// SetBudgetIDR create mutations, should not duplicate
func (p *PlansMutator) SetBudgetIDR(val int64) bool { //nolint:dupl false positive
	if val != p.BudgetIDR {
		p.mutations.Assign(15, val)
		p.logs = append(p.logs, A.X{`budgetIDR`, p.BudgetIDR, val})
		p.BudgetIDR = val
		return true
	}
	return false
}

// SetBudgetUSD create mutations, should not duplicate
func (p *PlansMutator) SetBudgetUSD(val int64) bool { //nolint:dupl false positive
	if val != p.BudgetUSD {
		p.mutations.Assign(16, val)
		p.logs = append(p.logs, A.X{`budgetUSD`, p.BudgetUSD, val})
		p.BudgetUSD = val
		return true
	}
	return false
}

// SetQuantity create mutations, should not duplicate
func (p *PlansMutator) SetQuantity(val int64) bool { //nolint:dupl false positive
	if val != p.Quantity {
		p.mutations.Assign(17, val)
		p.logs = append(p.logs, A.X{`quantity`, p.Quantity, val})
		p.Quantity = val
		return true
	}
	return false
}

// SetUnit create mutations, should not duplicate
func (p *PlansMutator) SetUnit(val string) bool { //nolint:dupl false positive
	if val != p.Unit {
		p.mutations.Assign(18, val)
		p.logs = append(p.logs, A.X{`unit`, p.Unit, val})
		p.Unit = val
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
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		p.TenantCode = S.Trim(from.TenantCode)
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
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		p.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		p.RestoredBy = from.RestoredBy
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
	if !excludeMap[`yearOf`] && (forceMap[`yearOf`] || from.YearOf != 0) {
		p.YearOf = from.YearOf
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
	if !excludeMap[`quantity`] && (forceMap[`quantity`] || from.Quantity != 0) {
		p.Quantity = from.Quantity
		changed = true
	}
	if !excludeMap[`unit`] && (forceMap[`unit`] || from.Unit != ``) {
		p.Unit = S.Trim(from.Unit)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
