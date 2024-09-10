package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mFinance"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
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
		Transaction rqFinance.BusinessTransaction `json:"transaction" form:"transaction" query:"transaction" long:"transaction" msg:"transaction"`
		TransactionJournals []rqFinance.TransactionJournal `json:"transactionJournals" form:"transactionJournals" query:"transactionJournals" long:"transactionJournals" msg:"transactionJournals"`
	}
	FieldSupervisorDashboardOut struct {
		ResponseCommon
		Pager    zCrud.PagerOut       `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta     *zCrud.Meta          `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Transaction *rqFinance.BusinessTransaction `json:"transaction" form:"transaction" query:"transaction" long:"transaction" msg:"transaction"`
		Transactions [][]any              `json:"transactions" form:"transactions" query:"transactions" long:"transactions" msg:"transactions"`
	}
)

const (
	FieldSupervisorDashboardAction = `fieldSupervisor/dashboard`

	ErrFieldSupervisorDashboardTrxNotFound = `transaction id not found`
	ErrFieldSupervisorDashboardUnauthorized = `you are unauthorized as field supervisor`
	ErrFieldSupervisorDashboardNotFieldSupervisor = `must be field supervisor to do this operation`
	ErrFieldSupervisorDashboardTenantHostNotFound = `tenant host not found`
	ErrFieldSupervisorDashboardTenantNotFound = `tenant not found`
	ErrFieldSupervisorDashboardUserNotFound = `user not found`
	ErrFieldSupervisorDashboardUpdateFailed = `failed to update transaction`
	ErrFieldSupervisorDashboardTrxJournalNotFound = `transaction journal not found`
	ErrFieldSupervisorDashboardCoaNotFound = `coa id not found`
	ErrFieldSupervisorDashboardTrxJournalUpdateFailed = `failed to update transaction journal`
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
		out.SetError(400, ErrFieldSupervisorDashboardUnauthorized)
		return
	}

	if !sess.IsFieldSupervisor {
		out.SetError(400, ErrFieldSupervisorDashboardNotFieldSupervisor)
		return
	}

	tenantCode, err := GetTenantCodeByHost(in.Host)
	if err != nil {
		out.SetError(400, ErrFieldSupervisorDashboardTenantHostNotFound)
		return
	}

	user := rqAuth.NewUsers(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrFieldSupervisorDashboardUserNotFound)
		return
	}

	tenant := rqAuth.NewTenants(d.AuthOltp)
	tenant.TenantCode = tenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrFieldSupervisorDashboardTenantNotFound)
		return
	}

	if in.WithMeta {
		out.Meta = &FieldSupervisorDashboardMeta
	}
	
	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert:
		businessTrx := wcFinance.NewBusinessTransactionMutator(d.AuthOltp)
		businessTrx.Id = in.Transaction.Id

		if !businessTrx.FindById() {
			out.SetError(400, ErrFieldSupervisorDashboardTrxNotFound)
			return
		}

		if in.Transaction.StartDate != `` && in.Transaction.StartDate != businessTrx.StartDate {
			businessTrx.SetStartDate(in.Transaction.StartDate)
		}

		if in.Transaction.EndDate != `` && in.Transaction.EndDate != businessTrx.EndDate {
			businessTrx.SetEndDate(in.Transaction.EndDate)
		}

		if businessTrx.HaveMutation() {
			businessTrx.SetUpdatedAt(in.UnixNow())
			businessTrx.SetUpdatedBy(sess.UserId)
			if !businessTrx.DoUpdateById() {
				out.SetError(400, ErrFieldSupervisorDashboardUpdateFailed)
				return
			}
		}

		for _, v := range in.TransactionJournals {
			trxJournal := wcFinance.NewTransactionJournalMutator(d.AuthOltp)
			trxJournal.Id = v.Id

			if !trxJournal.FindById() {
				out.SetError(400, ErrFieldSupervisorDashboardTrxJournalNotFound)
				return
			}

			if trxJournal.TransactionTemplateId == businessTrx.TransactionTemplateId {
				if v.CoaId != trxJournal.CoaId && v.CoaId != 0 {
					coa := rqFinance.NewCoa(d.AuthOltp)
					coa.Id = v.CoaId

					if !coa.FindById() {
						out.SetError(400, ErrFieldSupervisorDashboardCoaNotFound)
						return
					}

					trxJournal.SetCoaId(coa.Id)
				}

				if v.Descriptions != trxJournal.Descriptions && v.Descriptions != `` {
					trxJournal.SetDescriptions(v.Descriptions)
				}

				if v.Date != trxJournal.Date && v.Date != `` && mFinance.IsValidDate(v.Date){
					trxJournal.SetDescriptions(v.Date)
				}

				if v.Descriptions != trxJournal.Descriptions && v.Descriptions != `` {
					trxJournal.SetDescriptions(v.Descriptions)
				}

				trxJournal.SetUpdatedAt(in.UnixNow())
				trxJournal.SetUpdatedBy(sess.UserId)

				if !trxJournal.DoUpdateById() {
					out.SetError(400, ErrFieldSupervisorDashboardTrxJournalUpdateFailed)
					return
				}
			}
		}

		out.Transaction = &businessTrx.BusinessTransaction

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		r := rqFinance.NewBusinessTransaction(d.AuthOltp)
		r.TenantCode = tenant.TenantCode
		out.Transactions = r.FindByPagination(&FieldSupervisorDashboardMeta, &in.Pager, &out.Pager)
	}

	return
}
