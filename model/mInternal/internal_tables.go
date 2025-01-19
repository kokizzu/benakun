package mInternal

import "github.com/kokizzu/gotro/D/Tt"

const (
	Id             = `id`
	UserId         = `userId`
	InvoiceNumber  = `invoiceNumber`
	Amount         = `amount`
	Currency       = `currency`
	PaymentMethod  = `paymentMethod`
	ResponseHeader = `responseHeader`
	ResponseBody   = `responseBody`
	Status         = `status`
	CreatedAt      = `createdAt`
	CreatedBy      = `createdBy`
	UpdatedAt      = `updatedAt`
	UpdatedBy      = `updatedBy`
	DeletedAt      = `deletedAt`
	DeletedBy      = `deletedBy`
	RestoredBy     = `restoredBy`
)

const (
	TableInvoicePayment Tt.TableName = `invoicePayment`

	StatusPending  = `pending`
	StatusSuccess  = `success`
	StatusFailed   = `failed`
	StatusCanceled = `canceled`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableInvoicePayment: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{UserId, Tt.Unsigned},
			{InvoiceNumber, Tt.String},
			{Amount, Tt.Integer},
			{Currency, Tt.String},
			{PaymentMethod, Tt.String},
			{ResponseHeader, Tt.String},
			{ResponseBody, Tt.String},
			{Status, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
	},
}
