package wcInternal

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mInternal/rqInternal"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// InvoicePaymentMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcInternal__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcInternal__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcInternal__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcInternal__ORM.GEN.go
type InvoicePaymentMutator struct {
	rqInternal.InvoicePayment
	mutations []A.X
	logs      []A.X
}

// NewInvoicePaymentMutator create new ORM writer/command object
func NewInvoicePaymentMutator(adapter *Tt.Adapter) (res *InvoicePaymentMutator) {
	res = &InvoicePaymentMutator{InvoicePayment: rqInternal.InvoicePayment{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (i *InvoicePaymentMutator) Logs() []A.X { //nolint:dupl false positive
	return i.logs
}

// HaveMutation check whether Set* methods ever called
func (i *InvoicePaymentMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(i.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (i *InvoicePaymentMutator) ClearMutations() { //nolint:dupl false positive
	i.mutations = []A.X{}
	i.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (i *InvoicePaymentMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := i.Adapter.Update(i.SpaceName(), i.UniqueIndexId(), A.X{i.Id}, i.ToUpdateArray())
	return !L.IsError(err, `InvoicePayment.DoOverwriteById failed: `+i.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (i *InvoicePaymentMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !i.HaveMutation() {
		return true
	}
	_, err := i.Adapter.Update(i.SpaceName(), i.UniqueIndexId(), A.X{i.Id}, i.mutations)
	return !L.IsError(err, `InvoicePayment.DoUpdateById failed: `+i.SpaceName())
}

// DoDeletePermanentById permanent delete
func (i *InvoicePaymentMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := i.Adapter.Delete(i.SpaceName(), i.UniqueIndexId(), A.X{i.Id})
	return !L.IsError(err, `InvoicePayment.DoDeletePermanentById failed: `+i.SpaceName())
}

// func (i *InvoicePaymentMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := i.ToArray()
//	_, err := i.Adapter.Upsert(i.SpaceName(), arr, A.X{
//		A.X{`=`, 0, i.Id},
//		A.X{`=`, 1, i.UserId},
//		A.X{`=`, 2, i.InvoiceNumber},
//		A.X{`=`, 3, i.Amount},
//		A.X{`=`, 4, i.Currency},
//		A.X{`=`, 5, i.PaymentMethod},
//		A.X{`=`, 6, i.ResponseHeader},
//		A.X{`=`, 7, i.ResponseBody},
//		A.X{`=`, 8, i.Status},
//		A.X{`=`, 9, i.CreatedAt},
//		A.X{`=`, 10, i.CreatedBy},
//		A.X{`=`, 11, i.UpdatedAt},
//		A.X{`=`, 12, i.UpdatedBy},
//		A.X{`=`, 13, i.DeletedAt},
//		A.X{`=`, 14, i.DeletedBy},
//		A.X{`=`, 15, i.RestoredBy},
//		A.X{`=`, 16, i.SupportStartAt},
//		A.X{`=`, 17, i.SupportEndAt},
//		A.X{`=`, 18, i.PaymentUrl},
//	})
//	return !L.IsError(err, `InvoicePayment.DoUpsert failed: `+i.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByInvoiceNumber update all columns, error if not exists, not using mutations/Set*
func (i *InvoicePaymentMutator) DoOverwriteByInvoiceNumber() bool { //nolint:dupl false positive
	_, err := i.Adapter.Update(i.SpaceName(), i.UniqueIndexInvoiceNumber(), A.X{i.InvoiceNumber}, i.ToUpdateArray())
	return !L.IsError(err, `InvoicePayment.DoOverwriteByInvoiceNumber failed: `+i.SpaceName())
}

// DoUpdateByInvoiceNumber update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (i *InvoicePaymentMutator) DoUpdateByInvoiceNumber() bool { //nolint:dupl false positive
	if !i.HaveMutation() {
		return true
	}
	_, err := i.Adapter.Update(i.SpaceName(), i.UniqueIndexInvoiceNumber(), A.X{i.InvoiceNumber}, i.mutations)
	return !L.IsError(err, `InvoicePayment.DoUpdateByInvoiceNumber failed: `+i.SpaceName())
}

// DoDeletePermanentByInvoiceNumber permanent delete
func (i *InvoicePaymentMutator) DoDeletePermanentByInvoiceNumber() bool { //nolint:dupl false positive
	_, err := i.Adapter.Delete(i.SpaceName(), i.UniqueIndexInvoiceNumber(), A.X{i.InvoiceNumber})
	return !L.IsError(err, `InvoicePayment.DoDeletePermanentByInvoiceNumber failed: `+i.SpaceName())
}

// DoInsert insert, error if already exists
func (i *InvoicePaymentMutator) DoInsert() bool { //nolint:dupl false positive
	arr := i.ToArray()
	row, err := i.Adapter.Insert(i.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			i.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `InvoicePayment.DoInsert failed: `+i.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (i *InvoicePaymentMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := i.ToArray()
	row, err := i.Adapter.Replace(i.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			i.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `InvoicePayment.DoUpsert failed: `+i.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != i.Id {
		i.mutations = append(i.mutations, A.X{`=`, 0, val})
		i.logs = append(i.logs, A.X{`id`, i.Id, val})
		i.Id = val
		return true
	}
	return false
}

// SetUserId create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetUserId(val uint64) bool { //nolint:dupl false positive
	if val != i.UserId {
		i.mutations = append(i.mutations, A.X{`=`, 1, val})
		i.logs = append(i.logs, A.X{`userId`, i.UserId, val})
		i.UserId = val
		return true
	}
	return false
}

// SetInvoiceNumber create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetInvoiceNumber(val string) bool { //nolint:dupl false positive
	if val != i.InvoiceNumber {
		i.mutations = append(i.mutations, A.X{`=`, 2, val})
		i.logs = append(i.logs, A.X{`invoiceNumber`, i.InvoiceNumber, val})
		i.InvoiceNumber = val
		return true
	}
	return false
}

// SetAmount create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetAmount(val int64) bool { //nolint:dupl false positive
	if val != i.Amount {
		i.mutations = append(i.mutations, A.X{`=`, 3, val})
		i.logs = append(i.logs, A.X{`amount`, i.Amount, val})
		i.Amount = val
		return true
	}
	return false
}

// SetCurrency create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetCurrency(val string) bool { //nolint:dupl false positive
	if val != i.Currency {
		i.mutations = append(i.mutations, A.X{`=`, 4, val})
		i.logs = append(i.logs, A.X{`currency`, i.Currency, val})
		i.Currency = val
		return true
	}
	return false
}

// SetPaymentMethod create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetPaymentMethod(val string) bool { //nolint:dupl false positive
	if val != i.PaymentMethod {
		i.mutations = append(i.mutations, A.X{`=`, 5, val})
		i.logs = append(i.logs, A.X{`paymentMethod`, i.PaymentMethod, val})
		i.PaymentMethod = val
		return true
	}
	return false
}

// SetResponseHeader create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetResponseHeader(val string) bool { //nolint:dupl false positive
	if val != i.ResponseHeader {
		i.mutations = append(i.mutations, A.X{`=`, 6, val})
		i.logs = append(i.logs, A.X{`responseHeader`, i.ResponseHeader, val})
		i.ResponseHeader = val
		return true
	}
	return false
}

// SetResponseBody create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetResponseBody(val string) bool { //nolint:dupl false positive
	if val != i.ResponseBody {
		i.mutations = append(i.mutations, A.X{`=`, 7, val})
		i.logs = append(i.logs, A.X{`responseBody`, i.ResponseBody, val})
		i.ResponseBody = val
		return true
	}
	return false
}

// SetStatus create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetStatus(val string) bool { //nolint:dupl false positive
	if val != i.Status {
		i.mutations = append(i.mutations, A.X{`=`, 8, val})
		i.logs = append(i.logs, A.X{`status`, i.Status, val})
		i.Status = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != i.CreatedAt {
		i.mutations = append(i.mutations, A.X{`=`, 9, val})
		i.logs = append(i.logs, A.X{`createdAt`, i.CreatedAt, val})
		i.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != i.CreatedBy {
		i.mutations = append(i.mutations, A.X{`=`, 10, val})
		i.logs = append(i.logs, A.X{`createdBy`, i.CreatedBy, val})
		i.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != i.UpdatedAt {
		i.mutations = append(i.mutations, A.X{`=`, 11, val})
		i.logs = append(i.logs, A.X{`updatedAt`, i.UpdatedAt, val})
		i.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != i.UpdatedBy {
		i.mutations = append(i.mutations, A.X{`=`, 12, val})
		i.logs = append(i.logs, A.X{`updatedBy`, i.UpdatedBy, val})
		i.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != i.DeletedAt {
		i.mutations = append(i.mutations, A.X{`=`, 13, val})
		i.logs = append(i.logs, A.X{`deletedAt`, i.DeletedAt, val})
		i.DeletedAt = val
		return true
	}
	return false
}

// SetDeletedBy create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != i.DeletedBy {
		i.mutations = append(i.mutations, A.X{`=`, 14, val})
		i.logs = append(i.logs, A.X{`deletedBy`, i.DeletedBy, val})
		i.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != i.RestoredBy {
		i.mutations = append(i.mutations, A.X{`=`, 15, val})
		i.logs = append(i.logs, A.X{`restoredBy`, i.RestoredBy, val})
		i.RestoredBy = val
		return true
	}
	return false
}

// SetSupportStartAt create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetSupportStartAt(val int64) bool { //nolint:dupl false positive
	if val != i.SupportStartAt {
		i.mutations = append(i.mutations, A.X{`=`, 16, val})
		i.logs = append(i.logs, A.X{`supportStartAt`, i.SupportStartAt, val})
		i.SupportStartAt = val
		return true
	}
	return false
}

// SetSupportEndAt create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetSupportEndAt(val int64) bool { //nolint:dupl false positive
	if val != i.SupportEndAt {
		i.mutations = append(i.mutations, A.X{`=`, 17, val})
		i.logs = append(i.logs, A.X{`supportEndAt`, i.SupportEndAt, val})
		i.SupportEndAt = val
		return true
	}
	return false
}

// SetPaymentUrl create mutations, should not duplicate
func (i *InvoicePaymentMutator) SetPaymentUrl(val string) bool { //nolint:dupl false positive
	if val != i.PaymentUrl {
		i.mutations = append(i.mutations, A.X{`=`, 18, val})
		i.logs = append(i.logs, A.X{`paymentUrl`, i.PaymentUrl, val})
		i.PaymentUrl = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (i *InvoicePaymentMutator) SetAll(from rqInternal.InvoicePayment, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
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
	if !excludeMap[`userId`] && (forceMap[`userId`] || from.UserId != 0) {
		i.UserId = from.UserId
		changed = true
	}
	if !excludeMap[`invoiceNumber`] && (forceMap[`invoiceNumber`] || from.InvoiceNumber != ``) {
		i.InvoiceNumber = S.Trim(from.InvoiceNumber)
		changed = true
	}
	if !excludeMap[`amount`] && (forceMap[`amount`] || from.Amount != 0) {
		i.Amount = from.Amount
		changed = true
	}
	if !excludeMap[`currency`] && (forceMap[`currency`] || from.Currency != ``) {
		i.Currency = S.Trim(from.Currency)
		changed = true
	}
	if !excludeMap[`paymentMethod`] && (forceMap[`paymentMethod`] || from.PaymentMethod != ``) {
		i.PaymentMethod = S.Trim(from.PaymentMethod)
		changed = true
	}
	if !excludeMap[`responseHeader`] && (forceMap[`responseHeader`] || from.ResponseHeader != ``) {
		i.ResponseHeader = S.Trim(from.ResponseHeader)
		changed = true
	}
	if !excludeMap[`responseBody`] && (forceMap[`responseBody`] || from.ResponseBody != ``) {
		i.ResponseBody = S.Trim(from.ResponseBody)
		changed = true
	}
	if !excludeMap[`status`] && (forceMap[`status`] || from.Status != ``) {
		i.Status = S.Trim(from.Status)
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
	if !excludeMap[`supportStartAt`] && (forceMap[`supportStartAt`] || from.SupportStartAt != 0) {
		i.SupportStartAt = from.SupportStartAt
		changed = true
	}
	if !excludeMap[`supportEndAt`] && (forceMap[`supportEndAt`] || from.SupportEndAt != 0) {
		i.SupportEndAt = from.SupportEndAt
		changed = true
	}
	if !excludeMap[`paymentUrl`] && (forceMap[`paymentUrl`] || from.PaymentUrl != ``) {
		i.PaymentUrl = S.Trim(from.PaymentUrl)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
