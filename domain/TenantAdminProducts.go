package domain

import (
	"benakun/model"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBusiness"
	"benakun/model/mBusiness/rqBusiness"
	"benakun/model/mBusiness/wcBusiness"
	"benakun/model/mFinance"
	"benakun/model/mFinance/wcFinance"
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
		Cmd      string               `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool                 `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn        `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Product  *rqBusiness.Products `json:"product" form:"product" query:"product" long:"product" msg:"product"`
	}
	TenantAdminProductsOut struct {
		ResponseCommon
		Pager    zCrud.PagerOut       `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta     *zCrud.Meta          `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Product  *rqBusiness.Products `json:"product" form:"product" query:"product" long:"product" msg:"product"`
		Products [][]any              `json:"products" form:"products" query:"products" long:"products" msg:"products"`
	}
)

const (
	TenantAdminProductsAction = `tenantAdmin/products`

	ErrTenantAdminProductsUnauthorized          = `unauthorized user`
	ErrTenantAdminProductsProductNotFound       = `product not found`
	ErrTenantAdminProductsRuleNotValid          = `invalid product rule (must be fifo, lifo, average)`
	ErrTenantAdminProductsKindNotValid          = `invalid product kind (must be goods, service)`
	ErrTenantAdminProductsSaveFailed            = `product save failed`
	ErrTenantAdminProductsNotTenant             = `must be tenant admin to do this operation`
	ErrTenantAdminProductsCoaParentNotFound     = `coa parent not found`
	ErrTenantAdminProductsFailedSave            = `failed to create a new product`
	ErrTenantAdminProductsFailedUpdateCoaParent = `failed to update coa parent of product`
)

var TenantAdminProductsMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:     mBusiness.Id,
			Label:    "ID",
			DataType: zCrud.DataTypeInt,
			ReadOnly: true,
		},
		{
			Name:      mBusiness.Name,
			Label:     "Nama / Name",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBusiness.Detail,
			Label:     "Detail",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeTextArea,
		},
		{
			Name:        mBusiness.Rule,
			Label:       "Aturan / Rule",
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeCombobox,
			Description: `Product rule`,
			Ref: []string{
				mBusiness.RuleTypeFIFO, mBusiness.RuleTypeLIFO, mBusiness.RuleTypeAVERAGE,
			},
		},
		{
			Name:        mBusiness.Kind,
			Label:       "Jenis / Kind",
			DataType:    zCrud.DataTypeString,
			InputType:   zCrud.InputTypeCombobox,
			Description: `Product kind`,
			Ref: []string{
				mBusiness.KindTypeGOODS, mBusiness.KindTypeService,
			},
		},
		{
			Name:      mBusiness.CogsIDR,
			Label:     "HPP / COGS (IDR)",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name:      mBusiness.ProfitPercentage,
			Label:     "Persentase Profit / Profit Percentage",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypePercentage,
		},
		{
			Name:      mBusiness.CreatedAt,
			Label:     `Dibuat pada / Created at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mBusiness.UpdatedAt,
			Label:     `Diperbarui pada / Updated at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mBusiness.DeletedAt,
			Label:     `Dihapus pada / Deleted at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
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

	if in.WithMeta {
		out.Meta = &TenantAdminProductsMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		if in.Product.Id <= 0 {
			out.Meta = &TenantAdminProductsMeta
			return
		}

		product := rqBusiness.NewProducts(d.AuthOltp)
		product.Id = in.Product.Id
		if !product.FindById() {
			out.SetError(400, ErrTenantAdminProductsProductNotFound)
			return
		}

		out.Product = product
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.TenantCode = user.TenantCode
		if !tenant.FindByTenantCode() {
			out.SetError(400, ErrTenantAdminProductsNotTenant)
			return
		}

		product := wcBusiness.NewProductsMutator(d.AuthOltp)
		product.Id = in.Product.Id
		if product.Id > 0 {
			if !product.FindById() {
				out.SetError(400, ErrTenantAdminProductsProductNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if product.DeletedAt == 0 {
					product.SetDeletedAt(in.UnixNow())
					product.SetDeletedBy(sess.UserId)
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if product.DeletedAt > 0 {
					product.SetDeletedAt(0)
					product.SetRestoredBy(sess.UserId)
				}
			}
		} else {
			coaParent := wcFinance.NewCoaMutator(d.AuthOltp)
			coaParent.Id = tenant.ProductsCoaId
			if !coaParent.FindById() {
				out.SetError(400, ErrTenantAdminProductsCoaParentNotFound)
				return
			}

			coaChild := wcFinance.NewCoaMutator(d.AuthOltp)
			coaChild.SetLabel(model.COA_LABEL_PRODUCT)
			coaChild.SetName(in.Product.Name)
			coaChild.SetTenantCode(user.TenantCode)
			coaChild.SetParentId(coaParent.Id)
			coaChild.SetLabel(mFinance.LabelProduct)
			coaChild.SetCreatedAt(in.UnixNow())
			coaChild.SetUpdatedAt(in.UnixNow())

			if !coaChild.DoInsert() {
				out.SetError(400, ErrTenantAdminProductsFailedSave)
				return
			}

			moveToIdx := len(coaParent.Children)
			children := insertChildToIndex(coaParent.Children, coaChild.Id, moveToIdx)
			coaParent.SetChildren(children)
			if !coaParent.DoUpsertById() {
				out.SetError(400, ErrTenantAdminProductsFailedUpdateCoaParent)
				return
			}

			product.SetTenantCode(user.TenantCode)
		}

		if in.Product.Name != `` {
			product.SetName(in.Product.Name)
		}
		if in.Product.Detail != `` {
			product.SetDetail(in.Product.Detail)
		}
		if in.Product.Rule != `` {
			if !mBusiness.IsValidProductRule(in.Product.Rule) {
				out.SetError(400, ErrTenantAdminProductsRuleNotValid)
				return
			}
			product.SetRule(in.Product.Rule)
		}
		if in.Product.Kind != `` {
			if !mBusiness.IsValidProductKind(in.Product.Kind) {
				out.SetError(400, ErrTenantAdminProductsKindNotValid)
				return
			}
			product.SetKind(in.Product.Kind)
		}
		if in.Product.CogsIDR > 0 {
			product.SetCogsIDR(in.Product.CogsIDR)
		}
		if in.Product.ProfitPercentage > 0 {
			product.SetProfitPercentage(in.Product.ProfitPercentage)
		}

		if product.HaveMutation() {
			product.SetUpdatedAt(in.UnixNow())
			product.SetUpdatedBy(sess.UserId)
			if product.Id == 0 {
				product.SetCreatedAt(in.UnixNow())
				product.SetCreatedBy(sess.UserId)
			}
		}

		if !product.DoUpsertById() {
			out.SetError(500, ErrTenantAdminProductsSaveFailed)
		}

		out.Product = &product.Products

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		r := rqBusiness.NewProducts(d.AuthOltp)
		r.TenantCode = user.TenantCode
		out.Products = r.FindByPagination(&TenantAdminProductsMeta, &in.Pager, &out.Pager)
	}

	return
}
