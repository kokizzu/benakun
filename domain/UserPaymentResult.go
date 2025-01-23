package domain

import (
	"benakun/model/mInternal/rqInternal"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserPaymentResult.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserPaymentResult.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserPaymentResult.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserPaymentResult.go
//go:generate farify doublequote --file UserPaymentResult.go

type (
	UserPaymentResultIn struct {
		RequestCommon
		InvoiceNumber string `json:"invoiceNumber" form:"invoiceNumber" query:"invoiceNumber" long:"invoiceNumber" msg:"invoiceNumber"`
	}
	UserPaymentResultOut struct {
		ResponseCommon
		InvoicePayment *rqInternal.InvoicePayment `json:"invoicePayment" form:"invoicePayment" query:"invoicePayment" long:"invoicePayment" msg:"invoicePayment"`
	}
)

const (
	UserPaymentResultAction = `user/paymentResult`

	ErrUserPaymentResultUnauthorized = `unauthorized to view payment result`
	ErrUserPaymentResultNotFound     = `invoice payment not found`
)

func (d *Domain) UserPaymentResult(in *UserPaymentResultIn) (out UserPaymentResultOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(401, ErrUserPaymentResultUnauthorized)
		return
	}

	invPayment := rqInternal.NewInvoicePayment(d.IntrOltp)
	invPayment.InvoiceNumber = in.InvoiceNumber
	if !invPayment.FindByInvoiceNumber() {
		out.SetError(404, ErrUserPaymentResultNotFound)
		return
	}

	out.InvoicePayment = invPayment

	return
}
