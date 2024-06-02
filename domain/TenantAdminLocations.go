package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBusiness"
	"benakun/model/mBusiness/rqBusiness"
	"benakun/model/mBusiness/wcBusiness"
	"benakun/model/zCrud"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminLocations.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminLocations.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminLocations.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminLocations.go
//go:generate farify doublequote --file TenantAdminLocations.go

type (
	TenantAdminLocationsIn struct {
		RequestCommon
		Cmd      string               `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool                 `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn        `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Location rqBusiness.Locations `json:"location" form:"location" query:"location" long:"location" msg:"location"`
	}
	TenantAdminLocationsOut struct {
		ResponseCommon
		Pager     zCrud.PagerOut        `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta      *zCrud.Meta           `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Location  *rqBusiness.Locations `json:"location" form:"location" query:"location" long:"location" msg:"location"`
		Locations [][]any               `json:"locations" form:"locations" query:"locations" long:"locations" msg:"locations"`
	}
)

const (
	TenantAdminLocationsAction = `tenantAdmin/locations`

	ErrTenantAdminLocationsUnauthorized   = `unauthorized user`
	ErrTenantAdminLocationsTenantNotFound = `tenant admin not found`
	ErrTenantAdminLocationsNotTenant      = `must be tenant admin to do this operation`
	ErrTenantAdminLocationsNotFound       = `office location not found`
	ErrTenantAdminLocationsSaveFailed     = `office location save failed`
)

var TenantAdminLocationsMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mBusiness.Id,
			Label:     "ID",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:      mBusiness.Name,
			Label:     "Name",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBusiness.Country,
			Label:     "Negara",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBusiness.StateProvince,
			Label:     "Provinsi",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBusiness.CityRegency,
			Label:     "Kota/Kabupaten",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBusiness.Subdistrict,
			Label:     "Kecamatan",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBusiness.Village,
			Label:     "Desa/Keluarahan",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBusiness.RwBanjar,
			Label:     "RW/Banjar/Dusun",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBusiness.RtNeigb,
			Label:     "RT/Lingkungan",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBusiness.Address,
			Label:     "Alamat",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeTextArea,
		},
		{
			Name:      mBusiness.Lat,
			Label:     "Latitude",
			DataType:  zCrud.DataTypeFloat,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name:      mBusiness.Lng,
			Label:     "Longitude",
			DataType:  zCrud.DataTypeFloat,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name:      mBusiness.CreatedAt,
			Label:     "Created At",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mBusiness.UpdatedAt,
			Label:     "Updated At",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mBusiness.DeletedAt,
			Label:     "Deleted At",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
	},
}

func (d *Domain) TenantAdminLocations(in *TenantAdminLocationsIn) (out TenantAdminLocationsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminLocationsUnauthorized)
		return
	}

	if in.WithMeta {
		out.Meta = &TenantAdminLocationsMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.TenantCode = user.TenantCode
		if !tenant.FindByTenantCode() {
			out.SetError(400, ErrTenantAdminLocationsNotTenant)
			return
		}

		location := wcBusiness.NewLocationsMutator(d.AuthOltp)
		location.Id = in.Location.Id
		if location.Id > 0 {
			if !location.FindById() {
				out.SetError(400, ErrTenantAdminLocationsNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if location.DeletedAt == 0 {
					location.SetDeletedAt(in.UnixNow())
					location.SetDeletedBy(sess.UserId)
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if location.DeletedAt > 0 {
					location.SetDeletedAt(0)
					location.SetRestoredBy(sess.UserId)
				}
			}
		} else {
			location.SetTenantCode(user.TenantCode)
			location.SetCreatedAt(in.UnixNow())
			location.SetCreatedBy(sess.UserId)
		}

		location.SetAll(in.Location, nil, nil)
		location.SetUpdatedAt(in.UnixNow())
		location.SetUpdatedBy(sess.UserId)

		if !location.DoUpsertById() {
			out.SetError(500, ErrTenantAdminLocationsSaveFailed)
		}

		out.Location = &location.Locations

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		r := rqBusiness.NewLocations(d.AuthOltp)
		r.TenantCode = user.TenantCode
		out.Locations = r.FindByPagination(&TenantAdminLocationsMeta, &in.Pager, &out.Pager)
	}

	return
}
