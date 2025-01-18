package mInternal

import "github.com/kokizzu/gotro/D/Tt"

const (
	Id            = `id`
	UserId        = `userId`
	Code          = `code`
	Amount        = `amount`
	Currency      = `currency`
	PaymentMethod = `paymentMethod`
	CreatedAt     = `createdAt`
	CreatedBy     = `createdBy`
	UpdatedAt     = `updatedAt`
	UpdatedBy     = `updatedBy`
	DeletedAt     = `deletedAt`
	DeletedBy     = `deletedBy`
	RestoredBy    = `restoredBy`
)

const (
	TablePayment Tt.TableName = `payment`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TablePayment: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{UserId, Tt.Unsigned},
			{Code, Tt.String},
			{Amount, Tt.Integer},
			{Currency, Tt.String},
			{PaymentMethod, Tt.String},
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
