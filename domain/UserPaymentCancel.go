package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserPaymentCancel.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserPaymentCancel.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserPaymentCancel.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserPaymentCancel.go
//go:generate farify doublequote --file UserPaymentCancel.go

type (
	UserPaymentCancelIn struct {
		RequestCommon
	}
	UserPaymentCancelOut struct {
		ResponseCommon
		Request RequestCommon `json:"request" form:"request" query:"request" long:"request" msg:"request"`
	}
)

const (
	UserPaymentCancelAction = `user/paymentCancel`
)

func (d *Domain) UserPaymentCancel(in *UserPaymentCancelIn) (out UserPaymentCancelOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.Request = in.RequestCommon
	return
}
