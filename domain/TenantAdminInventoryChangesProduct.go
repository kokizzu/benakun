package domain

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

func (d *Domain) TenantAdminInventoryChangesProduct(in *TenantAdminInventoryChangesProductIn) (out TenantAdminInventoryChangesProductOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	return
}
