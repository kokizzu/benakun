package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mInternal"
	"benakun/model/mInternal/wcInternal"
	"os"
	"time"

	"github.com/goccy/go-json"
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/X"
)

/*
Payment Notification header (from DOKU)

{
  "X-TIMESTAMP": "2025-01-19T08:59:16Z",
  "X-SIGNATURE": "HMACSHA256=Z4HbaEBoWfRDmhmt27fNkKqIg+KQlXYLaJRaKAPhMNQ=",
  "Client-Id": "BRN-0248-1736342136769",
  "Request-Id": "422fb1f9-577a-4e06-a11c-691fe553bc3c",
  "Signature": "HMACSHA256=Leck4V4j6LoGUT95ftNklJo9lt1uxNuzB1FxLghvC6I=",
  "Request-Timestamp": "2025-01-19T08:59:17Z"
}
*/

/*
Payment Notification body (from DOKU)

{
  "service": {
    "id": "EMONEY"
  },
  "acquirer": {
    "id": "DANA"
  },
  "channel": {
    "id": "EMONEY_DANA"
  },
  "order": {
    "invoice_number": "INV-2025-01-19T08:57:55Z",
    "amount": 120000
  },
  "transaction": {
    "status": "SUCCESS",
    "date": "2025-01-19T15:59:16Z",
    "original_request_id": "422fb1f9-577a-4e06-a11c-691fe553bc3c"
  },
  "additional_info": {
    "line_items": []
  }
}
*/

func (d *Domain) PaymentsNotifications(c *fiber.Ctx) error {
	respHeader := mInternal.DOKUPaymentNotificationHeader{
		XTimeStamp:       c.Get(`X-TIMESTAMP`),
		XSignature:       c.Get(`X-SIGNATURE`),
		ClientId:         c.Get(`Client-Id`),
		RequestId:        c.Get(`Request-Id`),
		Signature:        c.Get(`Signature`),
		RequestTimestamp: c.Get(`Request-Timestamp`),
	}

	// Validate Request
	if respHeader.ClientId != os.Getenv(`DOKU_CLIENT_ID`) {
		return c.SendStatus(400)
	}

	var respBody M.SX
	err := c.BodyParser(&respBody)
	L.IsError(err, `c.BodyParser: %#v`, respBody)

	order := X.ToMSX(respBody[`order`])
	invoiceNumber := X.ToS(order[`invoice_number`])

	channel := X.ToMSX(respBody[`channel`])
	channelId := X.ToS(channel[`id`])

	transaction := X.ToMSX(respBody[`transaction`])
	transactionStatus := X.ToS(transaction[`status`])

	invPayment := wcInternal.NewInvoicePaymentMutator(d.IntrOltp)
	invPayment.InvoiceNumber = invoiceNumber
	if !invPayment.FindByInvoiceNumber() {
		invPayment.SetInvoiceNumber(invoiceNumber)
	}

	var status = mInternal.InvoiceStatusPending
	switch transactionStatus {
	case mInternal.DokuTransactionStatusSuccess:
		status = mInternal.InvoiceStatusSuccess
	case mInternal.DokuTransactionStatusFailed:
		status = mInternal.InvoiceStatusFailed
	case mInternal.DokuTransactionStatusExpired:
		status = mInternal.InvoiceStatusExpired
	case mInternal.DokuTransactionStatusRefunded:
		status = mInternal.InvoiceStatusRefunded
	default:
		status = mInternal.InvoiceStatusPending
	}
	invPayment.SetStatus(status)
	invPayment.SetPaymentMethod(channelId)
	invPayment.SetResponseBody(string(c.Body()))

	respHeaderByte, err := json.Marshal(respHeader)
	L.IsError(err, `json.Marshal: %#v`, respHeader)

	invPayment.SetResponseHeader(string(respHeaderByte))

	if transactionStatus == mInternal.DokuTransactionStatusSuccess {
		user := wcAuth.NewUsersMutator(d.AuthOltp)
		user.Id = invPayment.UserId
		if !user.FindById() {
			return c.SendStatus(500)
		}

		now := time.Now()

		if user.SupportExpiredAt > now.Unix() {
			now = time.Unix(user.SupportExpiredAt, 0)
		}

		expiredAt := mInternal.GetSupportExpiredAtByAmount(invPayment.Amount, now)
		user.SetSupportExpiredAt(expiredAt)

		invPayment.SetSupportStartAt(now.Unix())
		invPayment.SetSupportEndAt(expiredAt)

		if !user.DoUpdateById() {
			return c.SendStatus(500)
		}
	}

	invPayment.DoOverwriteByInvoiceNumber()

	return c.SendStatus(200)
}
