package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"os"
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
		State string `json:"state" form:"state" query:"state" long:"state" msg:"state"`
	}

	UserPurchaseSupportOut struct {
		ResponseCommon
	}
)

const (
	UserPurchaseSupportAction = `user/purchaseSupport`

	ErrUserPurchaseSupportUserNotFound      = `user not found`
	ErrUserPurchaseSupportUserFailed        = `failed to purchase support`
	ErrUserPurchaseSupportUserFailedPayment = `failed to create payment, try again later`
	ErrUserPurchaseSupportUserConnection    = `failed to call DOKU API, try again later`
)

func (d *Domain) UserPurchaseSupport(in *UserPurchaseSupportIn) (out UserPurchaseSupportOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	switch in.State {
	case zCrud.StatePaymentRequest:
		timeNow := time.Now().UTC().Format(time.RFC3339)
		payload := M.SX{
			"order": M.SX{
				"amount":         1000,
				"invoice_number": "INV-" + timeNow,
			},
			"payment": M.SX{
				"payment_due_date": 60,
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
			SetHeaders(M.SS{
				`Content-Type`:      `application/json`,
				`Client-Id`:         os.Getenv(`DOKU_CLIENT_ID`),
				`Request-Id`:        requestId,
				`Request-Timestamp`: timeNow,
				`Request-Target`:    `/checkout/v1/payment`,
				`Digest`:            digest,
				`Signature`:         signature,
			}).
			SetBody(payload).
			Get(os.Getenv(`DOKU_API_URL`))
		if err != nil {
			out.SetError(500, ErrUserPurchaseSupportUserConnection)
			return
		}

		fmt.Println(`Result:`, resp.Result())
		fmt.Println(`IsSuccess:`, resp.IsSuccess())
		fmt.Println(`StatusCode:`, resp.StatusCode())
		fmt.Println(`Status:`, resp.Status())
		fmt.Println(`Headers:`, resp.Header())
		fmt.Println(`Body:`, string(resp.Body()))

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
	var componentSignature = `Client-Id: ` + clientId + `
Request-Id: ` + requestId + `
Request-Timestamp: ` + requestTimestamp + `
Request-Target: ` + requestTarget + `
Digest: ` + digest

	L.Print(componentSignature)

	// Calculate HMAC-SHA256 base64 from all the components above
	h := hmac.New(sha256.New, []byte(secret))
	h.Write([]byte(componentSignature))
	hmacHash := h.Sum(nil)
	signature := base64.StdEncoding.EncodeToString(hmacHash)
	// Prepend encoded result with algorithm info HMACSHA256=
	return "HMACSHA256=" + signature
}
