package domain

import (
	"benakun/model/mAuth/wcAuth"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminTransaction.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminTransaction.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminTransaction.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminTransaction.go
//go:generate farify doublequote --file TenantAdminTransaction.go

type (
	TenantAdminTransactionIn struct {
		RequestCommon
	}

	TenantAdminTransactionOut struct {
		ResponseCommon
	}
)

const (
	TenantAdminTransactionAction = `tenantAdmin/transaction`

	ErrTenantAdminTransactionUnauthorized                 = `unauthorized user`
)

func (d *Domain) TenantAdminTransaction(in *TenantAdminTransactionIn) (out TenantAdminTransactionOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminTransactionUnauthorized)
		return
	}

	return
}