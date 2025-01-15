package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file GuestPaymentFailed.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type GuestPaymentFailed.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type GuestPaymentFailed.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type GuestPaymentFailed.go
//go:generate farify doublequote --file GuestPaymentFailed.go

type (
	GuestPaymentFailedIn struct {
		RequestCommon
	}
	GuestPaymentFailedOut struct {
		ResponseCommon
		Request RequestCommon `json:"request" form:"request" query:"request" long:"request" msg:"request"`
	}
)

const (
	GuestPaymentFailedAction = `guest/paymentFailed`
)

func (d *Domain) GuestPaymentFailed(in *GuestPaymentFailedIn) (out GuestPaymentFailedOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.Request = in.RequestCommon
	return
}
