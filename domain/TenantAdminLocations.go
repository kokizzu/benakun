package domain

import (
	"benakun/model/mAuth/wcAuth"
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
		Cmd      		string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta		bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    		zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	TenantAdminLocationsOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
	}
)

const (
	TenantAdminLocationsAction = `tenantAdmin/locations`

	ErrTenantAdminLocationsUnauthorized   = `unauthorized user`
	ErrTenantAdminLocationsTenantNotFound = `tenant admin not found`
)
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

	return
}
