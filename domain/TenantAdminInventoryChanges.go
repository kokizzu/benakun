package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBusiness"
	"benakun/model/zCrud"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminInventoryChanges.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminInventoryChanges.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminInventoryChanges.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminInventoryChanges.go
//go:generate farify doublequote --file TenantAdminInventoryChanges.go

type (
	TenantAdminInventoryChangesIn struct {
		RequestCommon
		Cmd      		string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		StaffEmail	string				`json:"staffEmail" form:"staffEmail" query:"staffEmail" long:"staffEmail" msg:"staffEmail"`
		TenantCode  string 				`json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
		Role 			string 				`json:"role" form:"role" query:"role" long:"role" msg:"role"`
		WithMeta		bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    		zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	TenantAdminInventoryChangesOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Staffs [][]any `json:"staffs" form:"staffs" query:"staffs" long:"staffs" msg:"staffs"`
		StaffsForm *[]rqAuth.Staff `json:"staffsForm" form:"staffsForm" query:"staffsForm" long:"staffsForm" msg:"staffsForm"`
	}
)

const (
	TenantAdminInventoryChangesAction = `tenantAdmin/inventoryChanges`

	ErrTenantAdminInventoryChangesUnauthorized   = `unauthorized user`
	ErrTenantAdminInventoryChangesTenantNotFound = `tenant admin not found`
	ErrTenantAdminInventoryChangesStaffEmailRequired = `staff email is required`
	ErrTenantAdminInventoryChangesUserNotFound = `user not found` 
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

	return
}
