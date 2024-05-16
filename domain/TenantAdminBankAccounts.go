package domain

import (
	"benakun/model/mAuth/wcAuth"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminBankAccounts.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminBankAccounts.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminBankAccounts.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminBankAccounts.go
//go:generate farify doublequote --file TenantAdminBankAccounts.go

type (
	TenantAdminBankAccountsIn struct {
		RequestCommon
	}
	TenantAdminBankAccountsOut struct {
		ResponseCommon
	}
)

const (
	TenantAdminBankAccountsAction = `tenantAdmin/bankAccounts`
)

func (d *Domain) TenantAdminBankAccounts(in *TenantAdminBankAccountsIn) (out TenantAdminBankAccountsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminCoaUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminCoaTenantNotFound)
		return
	}

	return
}
