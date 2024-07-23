package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBusiness"
	"benakun/model/mBusiness/rqBusiness"
	"benakun/model/mBusiness/wcBusiness"
	"benakun/model/zCrud"

	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminInventoryChanges.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminInventoryChanges.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminInventoryChanges.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminInventoryChanges.go
//go:generate farify doublequote --file TenantAdminInventoryChanges.go

type (
	TenantAdminInventoryChangesIn struct {
		RequestCommon
		Cmd      				string        							`json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		ProductId  			uint64      								`json:"productId,string" form:"productId" query:"productId" long:"productId" msg:"productId"`
		InventoryChange	rqBusiness.InventoryChanges `json:"inventoryChange" form:"inventoryChange" query:"inventoryChange" long:"inventoryChange" msg:"inventoryChange"`
		WithMeta				bool          							`json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    				zCrud.PagerIn 							`json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	TenantAdminInventoryChangesOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		InventoryChanges	[][]any `json:"inventoryChanges" form:"inventoryChanges" query:"inventoryChanges" long:"inventoryChanges" msg:"inventoryChanges"`
	}
)

const (
	TenantAdminInventoryChangesAction = `tenantAdmin/inventoryChanges`

	ErrTenantAdminInventoryChangesUnauthorized   			= `unauthorized user`
	ErrTenantAdminInventoryChangesTenantNotFound 			= `tenant admin not found`
	ErrTenantAdminInventoryChangesNotFound 						= `inventory not found` 
	ErrTenantAdminInventoryChangesNotTenant						= `must be tenant admin to do this operation`
	ErrTenantAdminInventoryChangesProductNotFound			= `product not found`
	ErrTenantAdminInventoryChangesSaveFailed      		= `inventory changes save failed`
)

var TenantAdminInventoryChangesMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mBusiness.Id,
			Label:     "ID",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly: true,
		},
		{
			Name: mBusiness.StockDelta,
			Label: "Stock Delta",
			DataType: zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name: mBusiness.ProductId,
			Label: "Produk / Product",
			DataType: zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			Description: "Select product",
			ReadOnly: true,
		},
		{
			Name: mBusiness.CreatedAt,
			Label: "Dibuat pada / Created at",
			InputType: zCrud.InputTypeHidden,
			ReadOnly: true,
		},
		{
			Name: mBusiness.UpdatedAt,
			Label: "Diperbarui pada / Updated at",
			InputType: zCrud.InputTypeHidden,
			ReadOnly: true,
		},
		{
			Name: mBusiness.DeletedAt,
			Label: "Dihapus pada / Deleted at",
			InputType: zCrud.InputTypeHidden,
			ReadOnly: true,
		},
	},
}

func (d *Domain) TenantAdminInventoryChanges(in *TenantAdminInventoryChangesIn) (out TenantAdminInventoryChangesOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminInventoryChangesUnauthorized)
		return
	}

	if in.WithMeta {
		out.Meta = &TenantAdminInventoryChangesMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		// TODO:
		// - get products
		// - get locations
		// - get spendings
		// - get expenses
		L.Print(`cmd:`, in.Cmd)
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		if user.Role != TenantAdminSegment {
			out.SetError(400, ErrTenantAdminInventoryChangesNotTenant)
			return
		}

		invChange := wcBusiness.NewInventoryChangesMutator(d.AuthOltp)
		invChange.Id = in.InventoryChange.Id

		if invChange.Id > 0 {
			if !invChange.FindById() {
				out.SetError(400, ErrTenantAdminInventoryChangesNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if invChange.DeletedAt == 0 {
					invChange.SetDeletedAt(in.UnixNow())
					invChange.SetDeletedBy(sess.UserId)
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if invChange.DeletedAt > 0 {
					invChange.SetDeletedAt(0)
					invChange.SetRestoredBy(sess.UserId)
				}
			}
		}

		invChange.SetTenantCode(user.TenantCode)
		invChange.SetStockDelta(in.InventoryChange.StockDelta)

		if invChange.HaveMutation() {
			invChange.SetUpdatedAt(in.UnixNow())
			invChange.SetUpdatedBy(sess.UserId)
			if invChange.Id == 0 {
				invChange.SetCreatedAt(in.UnixNow())
				invChange.SetCreatedBy(sess.UserId)
			}
		}

		if !invChange.DoUpsertById() {
			out.SetError(500, ErrTenantAdminInventoryChangesSaveFailed)
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		invChange := rqBusiness.NewInventoryChanges(d.AuthOltp)
		invChange.TenantCode = user.TenantCode
		if in.ProductId != 0 {
			invChange.ProductId = in.ProductId
			out.InventoryChanges = invChange.FindByPaginationByProduct(&TenantAdminInventoryChangesMeta, &in.Pager, &out.Pager)
		} else {
			out.InventoryChanges = invChange.FindByPagination(&TenantAdminInventoryChangesMeta, &in.Pager, &out.Pager)
		}
	}

	return
}
