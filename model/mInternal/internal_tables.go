package mInternal

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/kokizzu/gotro/D/Tt"
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
	SupportStartAt = `supportStartAt`
	SupportEndAt   = `supportEndAt`
	PaymentURL     = `paymentUrl`
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

// Transaction status obtained from DOKU
// https://dashboard.doku.com/docs/docs/technical-references/status-mapping/
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
	PriceSupportMonthly   int64 = 50_000  // IDR 50.000
	PriceSupportQuarterly int64 = 120_000 // IDR 120.000
	PriceSupportYearly    int64 = 450_000 // IDR 450.000
)

const (
	SupportDurationMonthly   = `monthly`
	SupportDurationQuarterly = `quarterly`
	SupportDurationYearly    = `yearly`
)

func GetSupportExpiredAtByAmount(amount int64, startAt time.Time) int64 {
	switch amount {
	case PriceSupportMonthly:
		return startAt.AddDate(0, 1, 0).Unix() // 1 month
	case PriceSupportQuarterly:
		return startAt.AddDate(0, 3, 0).Unix() // 3 months
	case PriceSupportYearly:
		return startAt.AddDate(1, 0, 0).Unix() // 1 year
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

// Generate invoice code with format INV-YYMMDD-HHMMSS-RANDOM. Last 6 characters are random string
func GenerateInvoiceCode() string {
	const letterBytes = `ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890`

	b := make([]byte, 6)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}

	nowFormatted := time.Now().Format("20060102-150405")
	randString := string(b)
	return fmt.Sprintf("INV-%s-%s", nowFormatted, randString)
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
			{SupportStartAt, Tt.Integer},
			{SupportEndAt, Tt.Integer},
			{PaymentURL, Tt.String},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
		Unique1:         InvoiceNumber,
	},
}
