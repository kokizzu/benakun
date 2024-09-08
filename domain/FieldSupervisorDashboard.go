package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/zCrud"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file FieldSupervisorDashboard.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type FieldSupervisorDashboard.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type FieldSupervisorDashboard.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type FieldSupervisorDashboard.go
//go:generate farify doublequote --file FieldSupervisorDashboard.go

type (
	FieldSupervisorDashboardIn struct {
		RequestCommon
		Cmd      string               `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		WithMeta bool                 `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn        `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Transaction *rqFinance.Transactions `json:"transaction" form:"transaction" query:"transaction" long:"transaction" msg:"transaction"`
	}
	FieldSupervisorDashboardOut struct {
		ResponseCommon
		Pager    zCrud.PagerOut       `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta     *zCrud.Meta          `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Transaction *rqFinance.Transactions `json:"transaction" form:"transaction" query:"transaction" long:"transaction" msg:"transaction"`
		Transactions [][]any              `json:"transactions" form:"transactions" query:"transactions" long:"transactions" msg:"transactions"`
	}
)

const (
	FieldSupervisorDashboardAction = `fieldSupervisor/dashboard`
)

var FieldSupervisorDashboardMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name: mFinance.Id,
			Label: "ID",
			DataType: zCrud.DataTypeInt,
			ReadOnly: true,
		},
		{
			Name: mFinance.TransactionTplId,
			Label: "Templat Transaksi / Transaction Template",
			DataType: zCrud.DataTypeInt,
			InputType: zCrud.InputTypeCombobox,
			Description: "Select Transaction Template",
			ReadOnly: true,
		},
		{
			Name: mFinance.StartDate,
			Label: "Tanggal Mulai / Start Date",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name: mFinance.EndDate,
			Label: "Tanggal Selesai / End Date",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mFinance.CreatedAt,
			Label:     `Dibuat pada / Created at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mFinance.UpdatedAt,
			Label:     `Diperbarui pada / Updated at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mFinance.DeletedAt,
			Label:     `Dihapus pada / Deleted at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
	},
}

func (d *Domain) FieldSupervisorDashboard(in *FieldSupervisorDashboardIn) (out FieldSupervisorDashboardOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustFieldSupervisor(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrDataEntryTransactionEntryUnauthorized)
		return
	}

	if !sess.IsFieldSupervisor {
		out.SetError(400, ErrDataEntryTransactionEntryNotDataEntry)
		return
	}

	tenantCode, err := GetTenantCodeByHost(in.Host)
	if err != nil {
		out.SetError(400, ErrDataEntryTransactionEntryTenantNotFound)
		return
	}

	user := rqAuth.NewUsers(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrDataEntryTransactionEntryUserNotFound)
		return
	}

	tenant := rqAuth.NewTenants(d.AuthOltp)
	tenant.TenantCode = tenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrDataEntryTransactionEntryTenantNotFound)
		return
	}

	if in.WithMeta {
		out.Meta = &FieldSupervisorDashboardMeta
	}
	
	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert:
	case zCrud.CmdList:
		r := rqFinance.NewBusinessTransaction(d.AuthOltp)
		r.TenantCode = tenant.TenantCode
		out.Transactions = r.FindByPagination(&FieldSupervisorDashboardMeta, &in.Pager, &out.Pager)
	}

	return
}
