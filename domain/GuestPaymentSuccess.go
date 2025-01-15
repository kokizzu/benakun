package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file GuestPaymentSuccess.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type GuestPaymentSuccess.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type GuestPaymentSuccess.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type GuestPaymentSuccess.go
//go:generate farify doublequote --file GuestPaymentSuccess.go

type (
	GuestPaymentSuccessIn struct {
		RequestCommon
	}
	GuestPaymentSuccessOut struct {
		ResponseCommon
		Request RequestCommon `json:"request" form:"request" query:"request" long:"request" msg:"request"`
	}
)

const (
	GuestPaymentSuccessAction = `guest/paymentSuccess`
)

func (d *Domain) GuestPaymentSuccess(in *GuestPaymentSuccessIn) (out GuestPaymentSuccessOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.Request = in.RequestCommon
	return
}
