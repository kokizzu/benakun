package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/saAuth"
	"benakun/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file SuperAdminAccessLog.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type SuperAdminAccessLog.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type SuperAdminAccessLog.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type SuperAdminAccessLog.go
//go:generate farify doublequote --file SuperAdminAccessLog.go

type (
	SuperAdminAccessLogIn struct {
		RequestCommon

		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
	}

	SuperAdminAccessLogOut struct {
		ResponseCommon

		Pager zCrud.PagerOut      `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Logs  []saAuth.ActionLogs `json:"logs" form:"logs" query:"logs" long:"logs" msg:"logs"`
		Meta  *zCrud.Meta         `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
	}
)

const (
	SuperAdminAccessLogAction = `superAdmin/accessLog`

	ErrSuperAdminAccessLogUnauthorized = `you are unauthorized to do this operation`
)

var (
	SuperAdminAccessLogsMeta = zCrud.Meta{
		Fields: []zCrud.Field{
			{
				Name:      mAuth.CreatedAt,
				Label:     `Created At`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeDateTime,
			},
			{
				Name:      mAuth.RequestId,
				Label:     `Request ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.ActorId,
				Label:     `Actor ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeNumber,
			},
			{
				Name:      mAuth.Action,
				Label:     `Action`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.StatusCode,
				Label:     `Status Code`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeNumber,
			},
			{
				Name:      mAuth.Traces,
				Label:     `Traces`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.Error,
				Label:     `Error`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.IpAddr4,
				Label:     `IP Address 4`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.IpAddr6,
				Label:     `IP Address 6`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.UserAgent,
				Label:     `User Agent`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeString,
				InputType: zCrud.InputTypeText,
			},
			{
				Name:      mAuth.Latency,
				Label:     `Latency`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeFloat,
				InputType: zCrud.InputTypeNumber,
			},
			{
				Name:      mAuth.RefId,
				Label:     `Ref ID`,
				ReadOnly:  true,
				DataType:  zCrud.DataTypeInt,
				InputType: zCrud.InputTypeNumber,
			},
		},
	}
)

func (d *Domain) SuperAdminAccessLog(in *SuperAdminAccessLogIn) (out SuperAdminAccessLogOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustSuperAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrSuperAdminAccessLogUnauthorized)
		return
	}

	if in.WithMeta {
		out.Meta = &SuperAdminAccessLogsMeta
	}

	if len(in.Pager.Order) == 0 {
		in.Pager.Order = []string{`-createdAt`}
	}

	r := saAuth.NewActionLogs(d.AuthOlap)
	out.Logs = r.FindByPagination(&in.Pager, &out.Pager)

	return
}
