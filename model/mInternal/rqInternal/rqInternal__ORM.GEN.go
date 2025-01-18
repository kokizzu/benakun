package rqInternal

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mInternal"

	"github.com/tarantool/go-tarantool/v2"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/X"
)

// Payment DAO reader/query struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file rqInternal__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type rqInternal__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type rqInternal__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type rqInternal__ORM.GEN.go
type Payment struct {
	Adapter       *Tt.Adapter `json:"-" msg:"-" query:"-" form:"-" long:"adapter"`
	Id            uint64      `json:"id,string" form:"id" query:"id" long:"id" msg:"id"`
	UserId        uint64      `json:"userId,string" form:"userId" query:"userId" long:"userId" msg:"userId"`
	Code          string      `json:"code" form:"code" query:"code" long:"code" msg:"code"`
	Amount        int64       `json:"amount" form:"amount" query:"amount" long:"amount" msg:"amount"`
	Currency      string      `json:"currency" form:"currency" query:"currency" long:"currency" msg:"currency"`
	PaymentMethod string      `json:"paymentMethod" form:"paymentMethod" query:"paymentMethod" long:"paymentMethod" msg:"paymentMethod"`
	CreatedAt     int64       `json:"createdAt" form:"createdAt" query:"createdAt" long:"createdAt" msg:"createdAt"`
	CreatedBy     uint64      `json:"createdBy,string" form:"createdBy" query:"createdBy" long:"createdBy" msg:"createdBy"`
	UpdatedAt     int64       `json:"updatedAt" form:"updatedAt" query:"updatedAt" long:"updatedAt" msg:"updatedAt"`
	UpdatedBy     uint64      `json:"updatedBy,string" form:"updatedBy" query:"updatedBy" long:"updatedBy" msg:"updatedBy"`
	DeletedAt     int64       `json:"deletedAt" form:"deletedAt" query:"deletedAt" long:"deletedAt" msg:"deletedAt"`
	DeletedBy     uint64      `json:"deletedBy,string" form:"deletedBy" query:"deletedBy" long:"deletedBy" msg:"deletedBy"`
	RestoredBy    uint64      `json:"restoredBy,string" form:"restoredBy" query:"restoredBy" long:"restoredBy" msg:"restoredBy"`
}

// NewPayment create new ORM reader/query object
func NewPayment(adapter *Tt.Adapter) *Payment {
	return &Payment{Adapter: adapter}
}

// SpaceName returns full package and table name
func (p *Payment) SpaceName() string { //nolint:dupl false positive
	return string(mInternal.TablePayment) // casting required to string from Tt.TableName
}

// SqlTableName returns quoted table name
func (p *Payment) SqlTableName() string { //nolint:dupl false positive
	return `"payment"`
}

func (p *Payment) UniqueIndexId() string { //nolint:dupl false positive
	return `id`
}

// FindById Find one by Id
func (p *Payment) FindById() bool { //nolint:dupl false positive
	res, err := p.Adapter.RetryDo(
		tarantool.NewSelectRequest(p.SpaceName()).
			Index(p.UniqueIndexId()).
			Limit(1).
			Iterator(tarantool.IterEq).
			Key(tarantool.UintKey{I: uint(p.Id)}),
	)
	if L.IsError(err, `Payment.FindById failed: `+p.SpaceName()) {
		return false
	}
	if len(res) == 1 {
		if row, ok := res[0].([]any); ok {
			p.FromArray(row)
			return true
		}
	}
	return false
}

// SqlSelectAllFields generate Sql select fields
func (p *Payment) SqlSelectAllFields() string { //nolint:dupl false positive
	return ` "id"
	, "userId"
	, "code"
	, "amount"
	, "currency"
	, "paymentMethod"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// SqlSelectAllUncensoredFields generate Sql select fields
func (p *Payment) SqlSelectAllUncensoredFields() string { //nolint:dupl false positive
	return ` "id"
	, "userId"
	, "code"
	, "amount"
	, "currency"
	, "paymentMethod"
	, "createdAt"
	, "createdBy"
	, "updatedAt"
	, "updatedBy"
	, "deletedAt"
	, "deletedBy"
	, "restoredBy"
	`
}

// ToUpdateArray generate slice of update command
func (p *Payment) ToUpdateArray() *tarantool.Operations { //nolint:dupl false positive
	return tarantool.NewOperations().
		Assign(0, p.Id).
		Assign(1, p.UserId).
		Assign(2, p.Code).
		Assign(3, p.Amount).
		Assign(4, p.Currency).
		Assign(5, p.PaymentMethod).
		Assign(6, p.CreatedAt).
		Assign(7, p.CreatedBy).
		Assign(8, p.UpdatedAt).
		Assign(9, p.UpdatedBy).
		Assign(10, p.DeletedAt).
		Assign(11, p.DeletedBy).
		Assign(12, p.RestoredBy)
}

// IdxId return name of the index
func (p *Payment) IdxId() int { //nolint:dupl false positive
	return 0
}

// SqlId return name of the column being indexed
func (p *Payment) SqlId() string { //nolint:dupl false positive
	return `"id"`
}

// IdxUserId return name of the index
func (p *Payment) IdxUserId() int { //nolint:dupl false positive
	return 1
}

// SqlUserId return name of the column being indexed
func (p *Payment) SqlUserId() string { //nolint:dupl false positive
	return `"userId"`
}

// IdxCode return name of the index
func (p *Payment) IdxCode() int { //nolint:dupl false positive
	return 2
}

// SqlCode return name of the column being indexed
func (p *Payment) SqlCode() string { //nolint:dupl false positive
	return `"code"`
}

// IdxAmount return name of the index
func (p *Payment) IdxAmount() int { //nolint:dupl false positive
	return 3
}

// SqlAmount return name of the column being indexed
func (p *Payment) SqlAmount() string { //nolint:dupl false positive
	return `"amount"`
}

// IdxCurrency return name of the index
func (p *Payment) IdxCurrency() int { //nolint:dupl false positive
	return 4
}

// SqlCurrency return name of the column being indexed
func (p *Payment) SqlCurrency() string { //nolint:dupl false positive
	return `"currency"`
}

// IdxPaymentMethod return name of the index
func (p *Payment) IdxPaymentMethod() int { //nolint:dupl false positive
	return 5
}

// SqlPaymentMethod return name of the column being indexed
func (p *Payment) SqlPaymentMethod() string { //nolint:dupl false positive
	return `"paymentMethod"`
}

// IdxCreatedAt return name of the index
func (p *Payment) IdxCreatedAt() int { //nolint:dupl false positive
	return 6
}

// SqlCreatedAt return name of the column being indexed
func (p *Payment) SqlCreatedAt() string { //nolint:dupl false positive
	return `"createdAt"`
}

// IdxCreatedBy return name of the index
func (p *Payment) IdxCreatedBy() int { //nolint:dupl false positive
	return 7
}

// SqlCreatedBy return name of the column being indexed
func (p *Payment) SqlCreatedBy() string { //nolint:dupl false positive
	return `"createdBy"`
}

// IdxUpdatedAt return name of the index
func (p *Payment) IdxUpdatedAt() int { //nolint:dupl false positive
	return 8
}

// SqlUpdatedAt return name of the column being indexed
func (p *Payment) SqlUpdatedAt() string { //nolint:dupl false positive
	return `"updatedAt"`
}

// IdxUpdatedBy return name of the index
func (p *Payment) IdxUpdatedBy() int { //nolint:dupl false positive
	return 9
}

// SqlUpdatedBy return name of the column being indexed
func (p *Payment) SqlUpdatedBy() string { //nolint:dupl false positive
	return `"updatedBy"`
}

// IdxDeletedAt return name of the index
func (p *Payment) IdxDeletedAt() int { //nolint:dupl false positive
	return 10
}

// SqlDeletedAt return name of the column being indexed
func (p *Payment) SqlDeletedAt() string { //nolint:dupl false positive
	return `"deletedAt"`
}

// IdxDeletedBy return name of the index
func (p *Payment) IdxDeletedBy() int { //nolint:dupl false positive
	return 11
}

// SqlDeletedBy return name of the column being indexed
func (p *Payment) SqlDeletedBy() string { //nolint:dupl false positive
	return `"deletedBy"`
}

// IdxRestoredBy return name of the index
func (p *Payment) IdxRestoredBy() int { //nolint:dupl false positive
	return 12
}

// SqlRestoredBy return name of the column being indexed
func (p *Payment) SqlRestoredBy() string { //nolint:dupl false positive
	return `"restoredBy"`
}

// ToArray receiver fields to slice
func (p *Payment) ToArray() A.X { //nolint:dupl false positive
	var id any = nil
	if p.Id != 0 {
		id = p.Id
	}
	return A.X{
		id,
		p.UserId,        // 1
		p.Code,          // 2
		p.Amount,        // 3
		p.Currency,      // 4
		p.PaymentMethod, // 5
		p.CreatedAt,     // 6
		p.CreatedBy,     // 7
		p.UpdatedAt,     // 8
		p.UpdatedBy,     // 9
		p.DeletedAt,     // 10
		p.DeletedBy,     // 11
		p.RestoredBy,    // 12
	}
}

// FromArray convert slice to receiver fields
func (p *Payment) FromArray(a A.X) *Payment { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.UserId = X.ToU(a[1])
	p.Code = X.ToS(a[2])
	p.Amount = X.ToI(a[3])
	p.Currency = X.ToS(a[4])
	p.PaymentMethod = X.ToS(a[5])
	p.CreatedAt = X.ToI(a[6])
	p.CreatedBy = X.ToU(a[7])
	p.UpdatedAt = X.ToI(a[8])
	p.UpdatedBy = X.ToU(a[9])
	p.DeletedAt = X.ToI(a[10])
	p.DeletedBy = X.ToU(a[11])
	p.RestoredBy = X.ToU(a[12])
	return p
}

// FromUncensoredArray convert slice to receiver fields
func (p *Payment) FromUncensoredArray(a A.X) *Payment { //nolint:dupl false positive
	p.Id = X.ToU(a[0])
	p.UserId = X.ToU(a[1])
	p.Code = X.ToS(a[2])
	p.Amount = X.ToI(a[3])
	p.Currency = X.ToS(a[4])
	p.PaymentMethod = X.ToS(a[5])
	p.CreatedAt = X.ToI(a[6])
	p.CreatedBy = X.ToU(a[7])
	p.UpdatedAt = X.ToI(a[8])
	p.UpdatedBy = X.ToU(a[9])
	p.DeletedAt = X.ToI(a[10])
	p.DeletedBy = X.ToU(a[11])
	p.RestoredBy = X.ToU(a[12])
	return p
}

// FindOffsetLimit returns slice of struct, order by idx, eg. .UniqueIndex*()
func (p *Payment) FindOffsetLimit(offset, limit uint32, idx string) []Payment { //nolint:dupl false positive
	var rows []Payment
	res, err := p.Adapter.RetryDo(
		tarantool.NewSelectRequest(p.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Payment.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows
	}
	for _, row := range res {
		item := Payment{}
		row, ok := row.([]any)
		if ok {
			rows = append(rows, *item.FromArray(row))
		}
	}
	return rows
}

// FindArrOffsetLimit returns as slice of slice order by idx eg. .UniqueIndex*()
func (p *Payment) FindArrOffsetLimit(offset, limit uint32, idx string) ([]A.X, Tt.QueryMeta) { //nolint:dupl false positive
	var rows []A.X
	resp, err := p.Adapter.RetryDoResp(
		tarantool.NewSelectRequest(p.SpaceName()).
			Index(idx).
			Offset(offset).
			Limit(limit).
			Iterator(tarantool.IterAll),
	)
	if L.IsError(err, `Payment.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	res, err := resp.Decode()
	if L.IsError(err, `Payment.FindOffsetLimit failed: `+p.SpaceName()) {
		return rows, Tt.QueryMetaFrom(resp, err)
	}
	rows = make([]A.X, len(res))
	for _, row := range res {
		row, ok := row.([]any)
		if ok {
			rows = append(rows, row)
		}
	}
	return rows, Tt.QueryMetaFrom(resp, nil)
}

// Total count number of rows
func (p *Payment) Total() int64 { //nolint:dupl false positive
	rows := p.Adapter.CallBoxSpace(p.SpaceName()+`:count`, A.X{})
	if len(rows) > 0 && len(rows[0]) > 0 {
		return X.ToI(rows[0][0])
	}
	return 0
}

// PaymentFieldTypeMap returns key value of field name and key
var PaymentFieldTypeMap = map[string]Tt.DataType{ //nolint:dupl false positive
	`id`:            Tt.Unsigned,
	`userId`:        Tt.Unsigned,
	`code`:          Tt.String,
	`amount`:        Tt.Integer,
	`currency`:      Tt.String,
	`paymentMethod`: Tt.String,
	`createdAt`:     Tt.Integer,
	`createdBy`:     Tt.Unsigned,
	`updatedAt`:     Tt.Integer,
	`updatedBy`:     Tt.Unsigned,
	`deletedAt`:     Tt.Integer,
	`deletedBy`:     Tt.Unsigned,
	`restoredBy`:    Tt.Unsigned,
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
