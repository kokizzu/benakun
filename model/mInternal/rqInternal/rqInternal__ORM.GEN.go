package rqInternal

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mInternal"

	"github.com/tarantool/go-tarantool"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// InvoicePayment DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqInternal__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqInternal__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqInternal__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqInternal__ORM.GEN.go
type InvoicePayment struct {
	Adapter        *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id             uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	UserId         uint64      `json:"userId,string" form:"userId" query:"userId" long:"userId" msg:"userId"`
	InvoiceNumber  string      `json:"invoiceNumber" form:"invoiceNumber" query:"invoiceNumber" long:"invoiceNumber" msg:"invoiceNumber"`
	Amount         int64       `json:"amount" form:"amount" query:"amount" long:"amount" msg:"amount"`
	Currency       string      `json:"currency" form:"currency" query:"currency" long:"currency" msg:"currency"`
	PaymentMethod  string      `json:"paymentMethod" form:"paymentMethod" query:"paymentMethod" long:"paymentMethod" msg:"paymentMethod"`
	ResponseHeader string      `json:"responseHeader" form:"responseHeader" query:"responseHeader" long:"responseHeader" msg:"responseHeader"`
	ResponseBody   string      `json:"responseBody" form:"responseBody" query:"responseBody" long:"responseBody" msg:"responseBody"`
	Status         string      `json:"status" form:"status" query:"status" long:"status" msg:"status"`
	CreatedAt      int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy      uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt      int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy      uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt      int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy      uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy     uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
	SupportStartAt int64       `json:"supportStartAt" form:"supportStartAt" query:"supportStartAt" long:"supportStartAt" msg:"supportStartAt"`
	SupportEndAt   int64       `json:"supportEndAt" form:"supportEndAt" query:"supportEndAt" long:"supportEndAt" msg:"supportEndAt"`
}

// NewInvoicePayment create new ORM reader/query object
func NewInvoicePayment(adapter *Tt.Adapter) *InvoicePayment {
	return &InvoicePayment{Adapter: adapter}
}

// SpaceName returns full package and table name
func (i *InvoicePayment) SpaceName() string { //nolint:dupl false positive
	return string(mInternal.TableInvoicePayment) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (i *InvoicePayment) SqlTableName() string { //nolint:dupl false positive
	return `"invoicePayment"`
}

func (i *InvoicePayment) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (i *InvoicePayment) FindById() bool { //nolint:dupl false positive
	res, err := i.Adapter.Select(i.SpaceName(), i.UniqueIndexId(), 0, 1, tarantool.IterEq, A.X{i.Id})
	if L.IsError(err, `InvoicePayment.FindById failed: `+i.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		i.FromArray(rows[0])
		return true
	}
	return false
}

// UniqueIndexInvoiceNumber return unique index name
func (i *InvoicePayment) UniqueIndexInvoiceNumber() string { //nolint:dupl false positive
	return `invoiceNumber`
}

// FindByInvoiceNumber Find one by InvoiceNumber
func (i *InvoicePayment) FindByInvoiceNumber() bool { //nolint:dupl false positive
	res, err := i.Adapter.Select(i.SpaceName(), i.UniqueIndexInvoiceNumber(), 0, 1, tarantool.IterEq, A.X{i.InvoiceNumber})
	if L.IsError(err, `InvoicePayment.FindByInvoiceNumber failed: `+i.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		i.FromArray(rows[0])
		return true
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (i *InvoicePayment) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "userId"
	, "invoiceNumber"
	, "amount"
	, "currency"
	, "paymentMethod"
	, "responseHeader"
	, "responseBody"
	, "status"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "supportStartAt"
	, "supportEndAt"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (i *InvoicePayment) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "userId"
	, "invoiceNumber"
	, "amount"
	, "currency"
	, "paymentMethod"
	, "responseHeader"
	, "responseBody"
	, "status"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	, "supportStartAt"
	, "supportEndAt"
	`
}

// ToUpdateArray generate slice of update command
func (i *InvoicePayment) ToUpdateArray() A.X { //nolint:dupl false positive
	return A.X{
		A.X{`=`, 0, i.Id},
		A.X{`=`, 1, i.UserId},
		A.X{`=`, 2, i.InvoiceNumber},
		A.X{`=`, 3, i.Amount},
		A.X{`=`, 4, i.Currency},
		A.X{`=`, 5, i.PaymentMethod},
		A.X{`=`, 6, i.ResponseHeader},
		A.X{`=`, 7, i.ResponseBody},
		A.X{`=`, 8, i.Status},
		A.X{`=`, 9, i.CreatedAt},
		A.X{`=`, 10, i.CreatedBy},
		A.X{`=`, 11, i.UpdatedAt},
		A.X{`=`, 12, i.UpdatedBy},
		A.X{`=`, 13, i.DeletedAt},
		A.X{`=`, 14, i.DeletedBy},
		A.X{`=`, 15, i.RestoredBy},
		A.X{`=`, 16, i.SupportStartAt},
		A.X{`=`, 17, i.SupportEndAt},
	}
}

// IdxId return name of the index
func (i *InvoicePayment) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (i *InvoicePayment) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxUserId return name of the index
func (i *InvoicePayment) IdxUserId() int { //nolint:dupl false positive
	return 1
}

// SqlUserId return name of the column being indexed
func (i *InvoicePayment) SqlUserId() string { //nolint:dupl false positive
	return `"userId"`
}

// IdxInvoiceNumber return name of the index
func (i *InvoicePayment) IdxInvoiceNumber() int { //nolint:dupl false positive
	return 2
}

// SqlInvoiceNumber return name of the column being indexed
func (i *InvoicePayment) SqlInvoiceNumber() string { //nolint:dupl false positive
	return `"invoiceNumber"`
}

// IdxAmount return name of the index
func (i *InvoicePayment) IdxAmount() int { //nolint:dupl false positive
	return 3
}

// SqlAmount return name of the column being indexed
func (i *InvoicePayment) SqlAmount() string { //nolint:dupl false positive
	return `"amount"`
}

// IdxCurrency return name of the index
func (i *InvoicePayment) IdxCurrency() int { //nolint:dupl false positive
	return 4
}

// SqlCurrency return name of the column being indexed
func (i *InvoicePayment) SqlCurrency() string { //nolint:dupl false positive
	return `"currency"`
}

// IdxPaymentMethod return name of the index
func (i *InvoicePayment) IdxPaymentMethod() int { //nolint:dupl false positive
	return 5
}

// SqlPaymentMethod return name of the column being indexed
func (i *InvoicePayment) SqlPaymentMethod() string { //nolint:dupl false positive
	return `"paymentMethod"`
}

// IdxResponseHeader return name of the index
func (i *InvoicePayment) IdxResponseHeader() int { //nolint:dupl false positive
	return 6
}

// SqlResponseHeader return name of the column being indexed
func (i *InvoicePayment) SqlResponseHeader() string { //nolint:dupl false positive
	return `"responseHeader"`
}

// IdxResponseBody return name of the index
func (i *InvoicePayment) IdxResponseBody() int { //nolint:dupl false positive
	return 7
}

// SqlResponseBody return name of the column being indexed
func (i *InvoicePayment) SqlResponseBody() string { //nolint:dupl false positive
	return `"responseBody"`
}

// IdxStatus return name of the index
func (i *InvoicePayment) IdxStatus() int { //nolint:dupl false positive
	return 8
}

// SqlStatus return name of the column being indexed
func (i *InvoicePayment) SqlStatus() string { //nolint:dupl false positive
	return `"status"`
}

// IdxCreatedAt return name of the index
func (i *InvoicePayment) IdxCreatedAt() int { //nolint:dupl false positive
	return 9
}

// SqlCreatedAt return name of the column being indexed
func (i *InvoicePayment) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (i *InvoicePayment) IdxCreatedBy() int { //nolint:dupl false positive
	return 10
}

// SqlCreatedBy return name of the column being indexed
func (i *InvoicePayment) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (i *InvoicePayment) IdxUpdatedAt() int { //nolint:dupl false positive
	return 11
}

// SqlUpdatedAt return name of the column being indexed
func (i *InvoicePayment) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (i *InvoicePayment) IdxUpdatedBy() int { //nolint:dupl false positive
	return 12
}

// SqlUpdatedBy return name of the column being indexed
func (i *InvoicePayment) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (i *InvoicePayment) IdxDeletedAt() int { //nolint:dupl false positive
	return 13
}

// SqlDeletedAt return name of the column being indexed
func (i *InvoicePayment) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (i *InvoicePayment) IdxDeletedBy() int { //nolint:dupl false positive
	return 14
}

// SqlDeletedBy return name of the column being indexed
func (i *InvoicePayment) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (i *InvoicePayment) IdxRestoredBy() int { //nolint:dupl false positive
	return 15
}

// SqlRestoredBy return name of the column being indexed
func (i *InvoicePayment) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// IdxSupportStartAt return name of the index
func (i *InvoicePayment) IdxSupportStartAt() int { //nolint:dupl false positive
	return 16
}

// SqlSupportStartAt return name of the column being indexed
func (i *InvoicePayment) SqlSupportStartAt() string { //nolint:dupl false positive
	return `"supportStartAt"`
}

// IdxSupportEndAt return name of the index
func (i *InvoicePayment) IdxSupportEndAt() int { //nolint:dupl false positive
	return 17
}

// SqlSupportEndAt return name of the column being indexed
func (i *InvoicePayment) SqlSupportEndAt() string { //nolint:dupl false positive
	return `"supportEndAt"`
}

// ToArray receiver fields to slice
func (i *InvoicePayment) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if i.Id != 0 {
		id = i.Id
	}
	return A.X{
		id,
		i.UserId,         // 1
		i.InvoiceNumber,  // 2
		i.Amount,         // 3
		i.Currency,       // 4
		i.PaymentMethod,  // 5
		i.ResponseHeader, // 6
		i.ResponseBody,   // 7
		i.Status,         // 8
		i.CreatedAt,      // 9
		i.CreatedBy,      // 10
		i.UpdatedAt,      // 11
		i.UpdatedBy,      // 12
		i.DeletedAt,      // 13
		i.DeletedBy,      // 14
		i.RestoredBy,     // 15
		i.SupportStartAt, // 16
		i.SupportEndAt,   // 17
	}
}

// FromArray convert slice to receiver fields
func (i *InvoicePayment) FromArray(a A.X) *InvoicePayment { //nolint:dupl false positive
	i.Id = X.ToU(a[0])
	i.UserId = X.ToU(a[1])
	i.InvoiceNumber = X.ToS(a[2])
	i.Amount = X.ToI(a[3])
	i.Currency = X.ToS(a[4])
	i.PaymentMethod = X.ToS(a[5])
	i.ResponseHeader = X.ToS(a[6])
	i.ResponseBody = X.ToS(a[7])
	i.Status = X.ToS(a[8])
	i.CreatedAt = X.ToI(a[9])
	i.CreatedBy = X.ToU(a[10])
	i.UpdatedAt = X.ToI(a[11])
	i.UpdatedBy = X.ToU(a[12])
	i.DeletedAt = X.ToI(a[13])
	i.DeletedBy = X.ToU(a[14])
	i.RestoredBy = X.ToU(a[15])
	i.SupportStartAt = X.ToI(a[16])
	i.SupportEndAt = X.ToI(a[17])
	return i
}

// FromUncensoredArray convert slice to receiver fields
func (i *InvoicePayment) FromUncensoredArray(a A.X) *InvoicePayment { //nolint:dupl false positive
	i.Id = X.ToU(a[0])
	i.UserId = X.ToU(a[1])
	i.InvoiceNumber = X.ToS(a[2])
	i.Amount = X.ToI(a[3])
	i.Currency = X.ToS(a[4])
	i.PaymentMethod = X.ToS(a[5])
	i.ResponseHeader = X.ToS(a[6])
	i.ResponseBody = X.ToS(a[7])
	i.Status = X.ToS(a[8])
	i.CreatedAt = X.ToI(a[9])
	i.CreatedBy = X.ToU(a[10])
	i.UpdatedAt = X.ToI(a[11])
	i.UpdatedBy = X.ToU(a[12])
	i.DeletedAt = X.ToI(a[13])
	i.DeletedBy = X.ToU(a[14])
	i.RestoredBy = X.ToU(a[15])
	i.SupportStartAt = X.ToI(a[16])
	i.SupportEndAt = X.ToI(a[17])
	return i
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (i *InvoicePayment) FindOffsetLimit(offset, limit uint32, idx string) []InvoicePayment { //nolint:dupl false positive
	var rows []InvoicePayment
	res, err := i.Adapter.Select(i.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `InvoicePayment.FindOffsetLimit failed: `+i.SpaceName()) {
		return rows
	}
	for _, row := range res.Tuples() {
		item := InvoicePayment{}
		rows = append(rows, *item.FromArray(row))
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (i *InvoicePayment) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	res, err := i.Adapter.Select(i.SpaceName(), idx, offset, limit, tarantool.IterAll, A.X{})
	if L.IsError(err, `InvoicePayment.FindOffsetLimit failed: `+i.SpaceName()) {
		return rows, Tt.QueryMetaFrom(res, err)
	}
	tuples := res.Tuples()
	rows = make([]A.X, len(tuples))
	for z, row := range tuples {
		rows[z] = row
	}
	return rows, Tt.QueryMetaFrom(res, nil)
}

// Total count number of rows
func (i *InvoicePayment) Total() int64 { //nolint:dupl false positive
	rows := i.Adapter.CallBoxSpace(i.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// InvoicePaymentFieldTypeMap returns key value of field name and key
var InvoicePaymentFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:             Tt.Unsigned,
	`userId`:         Tt.Unsigned,
	`invoiceNumber`:  Tt.String,
	`amount`:         Tt.Integer,
	`currency`:       Tt.String,
	`paymentMethod`:  Tt.String,
	`responseHeader`: Tt.String,
	`responseBody`:   Tt.String,
	`status`:         Tt.String,
	`createdAt`:      Tt.Integer,
	`createdBy`:      Tt.Unsigned,
	`updatedAt`:      Tt.Integer,
	`updatedBy`:      Tt.Unsigned,
	`deletedAt`:      Tt.Integer,
	`deletedBy`:      Tt.Unsigned,
	`restoredBy`:     Tt.Unsigned,
	`supportStartAt`: Tt.Integer,
	`supportEndAt`:   Tt.Integer,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
