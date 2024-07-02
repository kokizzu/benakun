package domain

import (
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminInventoryChangesProduct.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminInventoryChangesProduct.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminInventoryChangesProduct.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminInventoryChangesProduct.go
//go:generate farify doublequote --file TenantAdminInventoryChangesProduct.go

type (
	TenantAdminInventoryChangesProductIn struct {
		RequestCommon
	}
	TenantAdminInventoryChangesProductOut struct {
		ResponseCommon
	}
)

const (
	TenantAdminInventoryChangesProductAction = `tenantAdmin/inventoryChanges/`

	ErrTenantAdminInventoryChangesProductUnauthorized   			= `unauthorized user`
)

func (d *Domain) TenantAdminInventoryChangesProduct(in *TenantAdminInventoryChangesProductIn) (out TenantAdminInventoryChangesProductOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminInventoryChangesProductUnauthorized)
		return
	}

	return
}
