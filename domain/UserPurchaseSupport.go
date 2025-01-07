package domain

import (
	"benakun/model/mAuth/wcAuth"
	"time"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserPurchaseSupport.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserPurchaseSupport.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserPurchaseSupport.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserPurchaseSupport.go
//go:generate farify doublequote --file UserPurchaseSupport.go

type (
	UserPurchaseSupportIn struct {
		RequestCommon
	}

	UserPurchaseSupportOut struct {
		ResponseCommon
	}
)

const (
	UserPurchaseSupportAction = `user/purchaseSupport`

	ErrUserPurchaseSupportUserNotFound = `user not found`
	ErrUserPurchaseSupportUserFailed   = `failed to purchase support`
)

func (d *Domain) UserPurchaseSupport(in *UserPurchaseSupportIn) (out UserPurchaseSupportOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrUserPurchaseSupportUserNotFound)
		return
	}

	now := time.Now()
	expiredAd := time.Date(now.Year()+1, now.Month(), now.Day(), 0, 0, 0, 0, now.Location())
	user.SetSupportExpiredAt(expiredAd.Unix())
	if !user.DoUpdateById() {
		out.SetError(500, ErrUserPurchaseSupportUserFailed)
		return
	}

	return
}
