package wcInternal

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mInternal/rqInternal"

	"github.com/tarantool/go-tarantool/v2"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// PaymentMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcInternal__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcInternal__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcInternal__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcInternal__ORM.GEN.go
type PaymentMutator struct {
	rqInternal.Payment
	mutations *tarantool.Operations
	logs      []A.X
}

// NewPaymentMutator create new ORM writer/command object
func NewPaymentMutator(adapter *Tt.Adapter) (res *PaymentMutator) {
	res = &PaymentMutator{Payment: rqInternal.Payment{Adapter: adapter}}
	res.mutations = tarantool.NewOperations()
	return
}

// Logs get array of logs [field, old, new]
func (p *PaymentMutator) Logs() []A.X { //nolint:dupl false positive
	return p.logs
}

// HaveMutation check whether Set* methods ever called
func (p *PaymentMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(p.logs) > 0
}

// ClearMutations clear all previously called Set* methods
func (p *PaymentMutator) ClearMutations() { //nolint:dupl false positive
	p.mutations = tarantool.NewOperations()
	p.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (p *PaymentMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := p.Adapter.RetryDo(tarantool.NewUpdateRequest(p.SpaceName()).
		Index(p.UniqueIndexId()).
		Key(tarantool.UintKey{I: uint(p.Id)}).
		Operations(p.ToUpdateArray()),
	)
	return !L.IsError(err, `Payment.DoOverwriteById failed: `+p.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (p *PaymentMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !p.HaveMutation() {
		return true
	}
	_, err := p.Adapter.RetryDo(
		tarantool.NewUpdateRequest(p.SpaceName()).
			Index(p.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(p.Id)}).
			Operations(p.mutations),
	)
	return !L.IsError(err, `Payment.DoUpdateById failed: `+p.SpaceName())
}

// DoDeletePermanentById permanent delete
func (p *PaymentMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := p.Adapter.RetryDo(
		tarantool.NewDeleteRequest(p.SpaceName()).
			Index(p.UniqueIndexId()).
			Key(tarantool.UintKey{I: uint(p.Id)}),
	)
	return !L.IsError(err, `Payment.DoDeletePermanentById failed: `+p.SpaceName())
}

// DoInsert insert, error if already exists
func (p *PaymentMutator) DoInsert() bool { //nolint:dupl false positive
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
	return !L.IsError(err, `Payment.DoInsert failed: `+p.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// tarantool's replace/upsert can only match by primary key
// previous name: DoReplace
func (p *PaymentMutator) DoUpsertById() bool { //nolint:dupl false positive
	if p.Id > 0 {
		return p.DoUpdateById()
	}
	return p.DoInsert()
}

// SetId create mutations, should not duplicate
func (p *PaymentMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != p.Id {
		p.mutations.Assign(0, val)
		p.logs = append(p.logs, A.X{`id`, p.Id, val})
		p.Id = val
		return true
	}
	return false
}

// SetUserId create mutations, should not duplicate
func (p *PaymentMutator) SetUserId(val uint64) bool { //nolint:dupl false positive
	if val != p.UserId {
		p.mutations.Assign(1, val)
		p.logs = append(p.logs, A.X{`userId`, p.UserId, val})
		p.UserId = val
		return true
	}
	return false
}

// SetCode create mutations, should not duplicate
func (p *PaymentMutator) SetCode(val string) bool { //nolint:dupl false positive
	if val != p.Code {
		p.mutations.Assign(2, val)
		p.logs = append(p.logs, A.X{`code`, p.Code, val})
		p.Code = val
		return true
	}
	return false
}

// SetAmount create mutations, should not duplicate
func (p *PaymentMutator) SetAmount(val int64) bool { //nolint:dupl false positive
	if val != p.Amount {
		p.mutations.Assign(3, val)
		p.logs = append(p.logs, A.X{`amount`, p.Amount, val})
		p.Amount = val
		return true
	}
	return false
}

// SetCurrency create mutations, should not duplicate
func (p *PaymentMutator) SetCurrency(val string) bool { //nolint:dupl false positive
	if val != p.Currency {
		p.mutations.Assign(4, val)
		p.logs = append(p.logs, A.X{`currency`, p.Currency, val})
		p.Currency = val
		return true
	}
	return false
}

// SetPaymentMethod create mutations, should not duplicate
func (p *PaymentMutator) SetPaymentMethod(val string) bool { //nolint:dupl false positive
	if val != p.PaymentMethod {
		p.mutations.Assign(5, val)
		p.logs = append(p.logs, A.X{`paymentMethod`, p.PaymentMethod, val})
		p.PaymentMethod = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (p *PaymentMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.CreatedAt {
		p.mutations.Assign(6, val)
		p.logs = append(p.logs, A.X{`createdAt`, p.CreatedAt, val})
		p.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (p *PaymentMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.CreatedBy {
		p.mutations.Assign(7, val)
		p.logs = append(p.logs, A.X{`createdBy`, p.CreatedBy, val})
		p.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (p *PaymentMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != p.UpdatedAt {
		p.mutations.Assign(8, val)
		p.logs = append(p.logs, A.X{`updatedAt`, p.UpdatedAt, val})
		p.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (p *PaymentMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.UpdatedBy {
		p.mutations.Assign(9, val)
		p.logs = append(p.logs, A.X{`updatedBy`, p.UpdatedBy, val})
		p.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (p *PaymentMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != p.DeletedAt {
		p.mutations.Assign(10, val)
		p.logs = append(p.logs, A.X{`deletedAt`, p.DeletedAt, val})
		p.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (p *PaymentMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != p.DeletedBy {
		p.mutations.Assign(11, val)
		p.logs = append(p.logs, A.X{`deletedBy`, p.DeletedBy, val})
		p.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (p *PaymentMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != p.RestoredBy {
		p.mutations.Assign(12, val)
		p.logs = append(p.logs, A.X{`restoredBy`, p.RestoredBy, val})
		p.RestoredBy = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (p *PaymentMutator) SetAll(from rqInternal.Payment, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`userId`] && (forceMap[`userId`] || from.UserId != 0) {
		p.UserId = from.UserId
		changed = true
	}
	if !excludeMap[`code`] && (forceMap[`code`] || from.Code != ``) {
		p.Code = S.Trim(from.Code)
		changed = true
	}
	if !excludeMap[`amount`] && (forceMap[`amount`] || from.Amount != 0) {
		p.Amount = from.Amount
		changed = true
	}
	if !excludeMap[`currency`] && (forceMap[`currency`] || from.Currency != ``) {
		p.Currency = S.Trim(from.Currency)
		changed = true
	}
	if !excludeMap[`paymentMethod`] && (forceMap[`paymentMethod`] || from.PaymentMethod != ``) {
		p.PaymentMethod = S.Trim(from.PaymentMethod)
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
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
