package mInternal

import (
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kpango/fastime"
)

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

	InvoiceStatusPending   = `pending`
	InvoiceStatusSuccess   = `success`
	InvoiceStatusFailed    = `failed`
	InvoiceStatusCancelled = `cancelled`
	InvoiceStatusExpired   = `expired`
	InvoiceStatusRefunded  = `refunded`
)

const (
	DokuTransactionStatusSuccess  = `SUCCESS`
	DokuTransactionStatusFailed   = `FAILED`
	DokuTransactionStatusExpired  = `EXPIRED`
	DokuTransactionStatusRefunded = `REFUNDED`
)

const (
	CurrencyIDR = `IDR`
	CurrencyUSD = `USD`
)

const (
	PriceSupportMonthly   uint32 = 50_000  // IDR 50.000
	PriceSupportQuarterly uint32 = 120_000 // IDR 120.000
	PriceSupportYearly    uint32 = 450_000 // IDR 450.000
)

const (
	SupportDurationMonthly   = `monthly`
	SupportDurationQuarterly = `quarterly`
	SupportDurationYearly    = `yearly`
)

func GetSupportExpiredAtByAmount(amount uint32) int64 {
	now := fastime.Now()

	switch amount {
	case PriceSupportMonthly:
		return now.AddDate(0, 1, 0).Unix()
	case PriceSupportQuarterly:
		return now.AddDate(0, 3, 0).Unix()
	case PriceSupportYearly:
		return now.AddDate(1, 0, 0).Unix()
	}

	return 0
}

func IsValidSupportDuration(duration string) bool {
	switch duration {
	case SupportDurationMonthly, SupportDurationQuarterly, SupportDurationYearly:
		return true
	default:
		return false
	}
}

type DOKUPaymentNotificationHeader struct {
	XTimeStamp       string `json:"X-TIMESTAMP"`       // "2025-01-19T08:59:16Z"
	XSignature       string `json:"X-SIGNATURE"`       // "HMACSHA256=Z4HbaEBoWfRDmhmt27fNkKqIg+KQlXYLaJRaKAPhMNQ="
	ClientId         string `json:"Client-Id"`         // "BRN-0248-1736342136769"
	RequestId        string `json:"Request-Id"`        // "422fb1f9-577a-4e06-a11c-691fe553bc3c"
	Signature        string `json:"Signature"`         // "HMACSHA256=Leck4V4j6LoGUT95ftNklJo9lt1uxNuzB1FxLghvC6I="
	RequestTimestamp string `json:"Request-Timestamp"` // "2025-01-19T08:59:17Z"
}

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
		Unique1:         InvoiceNumber,
	},
}
