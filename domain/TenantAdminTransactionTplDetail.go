package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
	"benakun/model/zCrud"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminTransactionTplDetail.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminTransactionTplDetail.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminTransactionTplDetail.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminTransactionTplDetail.go
//go:generate farify doublequote --file TenantAdminTransactionTplDetail.go

type (
	TenantAdminTransactionTplDetailIn struct {
		RequestCommon
		Cmd                 string                         `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		TransactionTplDetail	*rqFinance.TransactionTplDetail `json:"transactionTplDetail" form:"transactionTplDetail" query:"transactionTplDetail" long:"transactionTplDetail" msg:"transactionTplDetail"`
	}
	TenantAdminTransactionTplDetailOut struct {
		ResponseCommon
		TransactionTplDetail	*rqFinance.TransactionTplDetail `json:"transactionTplDetail" form:"transactionTplDetail" query:"transactionTplDetail" long:"transactionTplDetail" msg:"transactionTplDetail"`
		TransactionTplDetails	*[]rqFinance.TransactionTplDetail `json:"transactionTplDetails" form:"transactionTplDetails" query:"transactionTplDetails" long:"transactionTplDetails" msg:"transactionTplDetails"`
	}
)

const (
	TenantAdminTransactionTplDetailAction = `tenantAdmin/transactionTplDetail`

	ErrTenantAdminTransactionTplDetailUnauthorized   = `unauthorized user`
	ErrTenantAdminTransactionTplDetailTenantNotFound = `tenant admin not found`
	ErrTenantAdminTransactionTplDetailUserNotFound   = `user not found`
	ErrTenantAdminTransactionTplDetailNotFound       = `transaction template not found`
)

func (d *Domain) TenantAdminTransactionTplDetail(in *TenantAdminTransactionTplDetailIn) (out TenantAdminTransactionTplDetailOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminTransactionTplDetailUnauthorized)
		return
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		if in.TransactionTplDetail.Id > 0 {
			trxTplDetail := rqFinance.NewTransactionTplDetail(d.AuthOltp)
			trxTplDetail.Id = in.TransactionTplDetail.Id
			if !trxTplDetail.FindById() {
				out.SetError(400, ``)
				return
			}
		}
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.TenantCode = user.TenantCode
		if !tenant.FindByTenantCode() {
			out.SetError(400, ErrTenantAdminTransactionTplDetailTenantNotFound)
			return
		}

		trxTplDetail := wcFinance.NewTransactionTplDetailMutator(d.AuthOltp)
		trxTplDetail.Id = in.TransactionTplDetail.Id

		var parent *wcFinance.TransactionTemplateMutator

		if in.TransactionTplDetail.ParentId > 0 {
			parent = wcFinance.NewTransactionTemplateMutator(d.AuthOltp)
			parent.Id = in.TransactionTplDetail.ParentId
			if !parent.FindById() {
				out.SetError(400, ``)
				return
			}
		}

		if trxTplDetail.Id > 0 {
			if !trxTplDetail.FindById() {
				out.SetError(400, ``)
				return
			}

			switch in.Cmd {
			case zCrud.CmdUpsert:
				if in.TransactionTplDetail.CoaId != trxTplDetail.CoaId {
					coa := rqFinance.NewCoa(d.AuthOltp)
					coa.Id = in.TransactionTplDetail.CoaId
					if !coa.FindById() {
						out.SetError(400, ``)
						return
					}

					trxTplDetail.SetCoaId(in.TransactionTplDetail.CoaId)
				}

				if in.TransactionTplDetail.IsDebit != trxTplDetail.IsDebit {
					trxTplDetail.SetIsDebit(in.TransactionTplDetail.IsDebit)
				}
			case zCrud.CmdDelete:
				if trxTplDetail.DeletedAt == 0 {
					trxTplDetail.SetDeletedAt(in.UnixNow())
					trxTplDetail.SetDeletedBy(sess.UserId)
				}
			case zCrud.CmdRestore:
				if trxTplDetail.DeletedAt > 0 {
					trxTplDetail.SetDeletedAt(0)
					trxTplDetail.SetRestoredBy(sess.UserId)
				}
			}
		} else {
			if in.TransactionTplDetail.CoaId != 0 {
				coa := rqFinance.NewCoa(d.AuthOltp)
				coa.Id = in.TransactionTplDetail.CoaId
				if !coa.FindById() {
					out.SetError(400, ``)
					return
				}

				trxTplDetail.SetCoaId(in.TransactionTplDetail.CoaId)
			}
			trxTplDetail.SetParentId(parent.Id)
			trxTplDetail.SetIsDebit(in.TransactionTplDetail.IsDebit)
			trxTplDetail.SetTenantCode(tenant.TenantCode)
		}

		if trxTplDetail.HaveMutation() {
			trxTplDetail.SetUpdatedAt(in.UnixNow())
			trxTplDetail.SetUpdatedBy(sess.UserId)
			if trxTplDetail.Id == 0 {
				trxTplDetail.SetCreatedAt(in.UnixNow())
				trxTplDetail.SetCreatedBy(sess.UserId)
			}
		}

		if !trxTplDetail.DoUpsertById() {
			out.SetError(400, ``)
			return
		}

		out.TransactionTplDetail = &trxTplDetail.TransactionTplDetail

		fallthrough
	case zCrud.CmdList:
		r := rqFinance.NewTransactionTplDetail(d.AuthOltp)
		r.ParentId = in.TransactionTplDetail.ParentId
		r.TenantCode = user.TenantCode
		r.ParentId = in.TransactionTplDetail.ParentId
		trxTplDetails := r.FindTrxTplDetailsByTenantByTrxTplId()

		out.TransactionTplDetails = &trxTplDetails
	}

	return
}
