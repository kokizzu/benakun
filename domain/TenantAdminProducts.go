package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBusiness"
	"benakun/model/mBusiness/rqBusiness"
	"benakun/model/zCrud"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminProducts.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminProducts.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminProducts.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminProducts.go
//go:generate farify doublequote --file TenantAdminProducts.go

type (
	TenantAdminProductsIn struct {
		RequestCommon
		Cmd      		string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta		bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    		zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Product *rqBusiness.Products `json:"product" form:"product" query:"product" long:"product" msg:"product"`
	}
	TenantAdminProductsOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Product *rqBusiness.Products `json:"product" form:"product" query:"product" long:"product" msg:"product"`
		Products [][]any `json:"staffs" form:"staffs" query:"staffs" long:"staffs" msg:"staffs"`
	}
)

const (
	TenantAdminProductsAction = `tenantAdmin/products`

	ErrTenantAdminProductsUnauthorized   = `unauthorized user`
	ErrTenantAdminProductsTenantNotFound = `tenant admin not found`
)

var TenantAdminProductsMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name: mBusiness.Id,
			Label: "ID",
			DataType: zCrud.DataTypeInt,
			ReadOnly: true,
		},
	},
}

func (d *Domain) TenantAdminProducts(in *TenantAdminProductsIn) (out TenantAdminProductsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrTenantAdminProductsUnauthorized)
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminProductsUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
		out.SetError(400, ErrTenantAdminProductsTenantNotFound)
		return
	}

	if in.WithMeta {
		out.Meta = &TenantAdminProductsMeta
	}

	return
}
