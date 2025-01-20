package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserPaymentResult.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserPaymentResult.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserPaymentResult.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserPaymentResult.go
//go:generate farify doublequote --file UserPaymentResult.go

type (
	UserPaymentResultIn struct {
		RequestCommon
	}
	UserPaymentResultOut struct {
		ResponseCommon
	}
)

const (
	UserPaymentResultAction = `user/paymentResult`
)

func (d *Domain) UserPaymentResult(in *UserPaymentResultIn) (out UserPaymentResultOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	return
}
