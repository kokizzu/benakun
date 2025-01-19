package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"os"
	"strings"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/goccy/go-json"
	"github.com/google/uuid"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserPurchaseSupport.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserPurchaseSupport.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserPurchaseSupport.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserPurchaseSupport.go
//go:generate farify doublequote --file UserPurchaseSupport.go

type (
	UserPurchaseSupportIn struct {
		RequestCommon
		SuppportDuration string `json:"supportDuration" form:"supportDuration" query:"supportDuration" long:"supportDuration" msg:"supportDuration"`
		State            string `json:"state" form:"state" query:"state" long:"state" msg:"state"`
	}

	UserPurchaseSupportOut struct {
		ResponseCommon
		PaymentResponse any `json:"paymentResponse" form:"paymentResponse" query:"paymentResponse" long:"paymentResponse" msg:"paymentResponse"`
	}
)

const (
	UserPurchaseSupportAction = `user/purchaseSupport`

	ErrUserPurchaseSupportUserNotFound               = `user not found`
	ErrUserPurchaseSupportUserFailed                 = `failed to purchase support`
	ErrUserPurchaseSupportUserFailedPayment          = `failed to create payment, try again later`
	ErrUserPurchaseSupportUserConnection             = `failed to call DOKU API, try again later`
	ErrUserPurchaseSupportUserInvalidSupportDuration = `invalid support duration, must be monthly, quarterly, yearly`
)

const (
	AmountMonthly   uint32 = 50_000  // IDR 50.000
	AmountQuarterly uint32 = 120_000 // IDR 120.000
	AmountYearly    uint32 = 450_000 // IDR 450.000
)

const (
	SupportDurationMonthly   = `monthly`
	SupportDurationQuarterly = `quarterly`
	SupportDurationYearly    = `yearly`
)

func isValidSupportDuration(duration string) bool {
	switch duration {
	case SupportDurationMonthly, SupportDurationQuarterly, SupportDurationYearly:
		return true
	default:
		return false
	}
}

func (d *Domain) UserPurchaseSupport(in *UserPurchaseSupportIn) (out UserPurchaseSupportOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	switch in.State {
	case zCrud.StatePaymentRequest:
		if !isValidSupportDuration(in.SuppportDuration) {
			out.SetError(400, ErrUserPurchaseSupportUserInvalidSupportDuration)
			return
		}

		var amount uint32
		switch in.SuppportDuration {
		case SupportDurationMonthly:
			amount = AmountMonthly
		case SupportDurationQuarterly:
			amount = AmountQuarterly
		case SupportDurationYearly:
			amount = AmountYearly
		}

		timeNow := time.Now().UTC().Format(time.RFC3339)
		payload := M.SX{
			"order": M.SX{
				"amount":                amount,
				"invoice_number":        "INV-" + timeNow,
				"language":              "EN",
				"disable_retry_payment": true,
				"callback_url":          in.Host + `/` + GuestPaymentSuccessAction,
				"callback_url_cancel":   in.Host + `/` + GuestPaymentFailedAction,
				"callback_url_result":   in.Host + `/` + GuestPaymentSuccessAction,
				"auto_redirect":         true,
			},
			"payment": M.SX{
				"payment_due_date": 60, // in minutes
			},
			"additional_info": M.SX{
				"override_notification_url": os.Getenv(`DOKU_NOTIFICATION_URL`),
			},
		}

		jsonBody, err := json.Marshal(payload)
		if err != nil {
			out.SetError(500, ErrUserPurchaseSupportUserFailedPayment)
			return
		}

		requestId := uuid.NewString()
		digest := generateDOKUDigest(string(jsonBody))
		signature := generateDOKUSignature(
			os.Getenv(`DOKU_CLIENT_ID`),
			requestId,
			timeNow,
			`/checkout/v1/payment`,
			digest,
			os.Getenv(`DOKU_API_KEY`),
		)

		client := resty.New()
		resp, err := client.R().EnableTrace().
			SetHeader(`Client-Id`, os.Getenv(`DOKU_CLIENT_ID`)).
			SetHeader(`Request-Id`, requestId).
			SetHeader(`Request-Timestamp`, timeNow).
			SetHeader(`Request-Target`, `/checkout/v1/payment`).
			SetHeader(`Digest`, digest).
			SetHeader(`Signature`, signature).
			SetBody(payload).
			Post(os.Getenv(`DOKU_API_URL`))

		if err != nil {
			L.IsError(err, `resty.Post: %#v`, payload)
			out.SetError(500, ErrUserPurchaseSupportUserConnection)
			return
		}

		var respBody M.SX
		err = json.Unmarshal(resp.Body(), &respBody)
		if err != nil {
			L.IsError(err, `json.Unmarshal: %#v`, respBody)
			out.SetError(500, ErrUserPurchaseSupportUserConnection)
			return
		}

		out.PaymentResponse = respBody
	case zCrud.StatePaymentUpsert:
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
	}

	return
}
func generateDOKUDigest(jsonBody string) string {
	converted := []byte(jsonBody)
	hasher := sha256.New()
	hasher.Write(converted)

	hashedByte := hasher.Sum(nil)
	return (base64.StdEncoding.EncodeToString(hashedByte))
}

func generateDOKUSignature(clientId string, requestId string, requestTimestamp string, requestTarget string, digest string, secret string) string {
	// Prepare Signature Component
	var componentSignature strings.Builder
	componentSignature.WriteString("Client-Id:" + clientId)
	componentSignature.WriteString("\n")
	componentSignature.WriteString("Request-Id:" + requestId)
	componentSignature.WriteString("\n")
	componentSignature.WriteString("Request-Timestamp:" + requestTimestamp)
	componentSignature.WriteString("\n")
	componentSignature.WriteString("Request-Target:" + requestTarget)
	componentSignature.WriteString("\n")
	componentSignature.WriteString("Digest:" + digest)

	// Calculate HMAC-SHA256 base64 from all the components above
	key := []byte(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(componentSignature.String()))
	signature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	// Prepend encoded result with algorithm info HMACSHA256=
	return "HMACSHA256=" + signature
}
