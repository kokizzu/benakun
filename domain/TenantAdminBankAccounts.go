package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mBudget"
	"benakun/model/mBudget/rqBudget"
	"benakun/model/mBudget/wcBudget"
	"benakun/model/mFinance/wcFinance"
	"benakun/model/zCrud"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminBankAccounts.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminBankAccounts.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminBankAccounts.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminBankAccounts.go
//go:generate farify doublequote --file TenantAdminBankAccounts.go

type (
	TenantAdminBankAccountsIn struct {
		RequestCommon
		Cmd      string                `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		Account  *rqBudget.BankAccounts `json:"account" form:"account" query:"account" long:"account" msg:"account"`
		WithMeta bool                  `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn         `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	TenantAdminBankAccountsOut struct {
		ResponseCommon
		Pager    zCrud.PagerOut         `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta     *zCrud.Meta            `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Staffs   *[]rqAuth.Staff        `json:"staffs" form:"staffs" query:"staffs" long:"staffs" msg:"staffs"`
		Account  *rqBudget.BankAccounts `json:"account" form:"account" query:"account" long:"account" msg:"account"`
		Accounts [][]any                `json:"accounts" form:"accounts" query:"accounts" long:"accounts" msg:"accounts"`
	}
)

const (
	TenantAdminBankAccountsAction = `tenantAdmin/bankAccounts`

	ErrTenantAdminBankAccountsUnauthorized    			= `unauthorized user`
	ErrTenantAdminBankAccountsTenantNotFound  			= `tenant not found`
	ErrTenantAdminBankAccountsNotFound        			= `bank account not found`
	ErrTenantAdminBankAccountsSaveFailed      			= `bank account save failed`
	ErrTenantAdminBankAccountsParentNotFound  			= `parent bank account not found`
	ErrTenantAdminBankAccountsParentHaveChild 			= `parent bank account already have child`
	ErrTenantAdminBankAccountsStaffNotFound   			= `staff not found to choose account's owner`
	ErrTenantAdminBankAccountsNotTenant       			= `must be tenant admin to do this operation`
	ErrTenantAdminBankAccountsCustomerCoaNotFound 	= `customer coa not found`
	ErrTenantAdminBankAccountsCustomerCoaSaveFailed	= `failed to save account to customer coa`
	ErrTenantAdminBankAccountsCustomerCoaParentUpdateFailed = `failed to update parent of customer coa`
	ErrTenantAdminBankAccountsSupplierCoaNotFound 	= `supplier coa not found`
	ErrTenantAdminBankAccountsSupplierCoaSaveFailed	= `failed to save account to supplier coa`
	ErrTenantAdminBankAccountsSupplierCoaParentUpdateFailed = `failed to update parent of supplier coa`
	ErrTenantAdminBankAccountsBankCoaNotFound 			= `bank coa not found`
	ErrTenantAdminBankAccountsBankCoaSaveFailed			= `failed to save account to bank coa`
	ErrTenantAdminBankAccountsBankCoaParentUpdateFailed = `failed to update parent of bank coa`
)

var TenantAdminBankAccountsMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:     mBudget.Id,
			Label:    "ID",
			DataType: zCrud.DataTypeInt,
			ReadOnly: true,
		},
		{
			Name:      mBudget.Name,
			Label:     "Nama / Name",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBudget.AccountName,
			Label:     "Nama Akun / Account Name",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mBudget.AccountNumber,
			Label:     "Nomor Akun / Account Number",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeNumber,
		},
		{
			Name:     mBudget.BankName,
			Label:    "Nama Bank / Bank Name",
			DataType: zCrud.DataTypeString,
		},
		{
			Name:      mBudget.IsProfitCenter,
			Label:     "Pusat Profit / Profit Center ?",
			DataType:  zCrud.DataTypeBool,
			InputType: zCrud.InputTypeCheckbox,
		},
		{
			Name:        mBudget.StaffId,
			Label:       "Karyawan / Staff",
			DataType:    zCrud.DataTypeInt,
			InputType:   zCrud.InputTypeCombobox,
			Description: `Select staff`,
			ReadOnly: true,
		},
		{
			Name:      mBudget.IsCostCenter,
			Label:     "Pusat Biaya / Cost Center ?",
			DataType:  zCrud.DataTypeBool,
			InputType: zCrud.InputTypeCheckbox,
		},
		{
			Name:      mBudget.CreatedAt,
			Label:     "Dibuat pada / Created at",
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mBudget.UpdatedAt,
			Label:     "Diupdate pada / Updated at",
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
		{
			Name:      mBudget.DeletedAt,
			Label:     "Dihapus pada / Deleted at",
			InputType: zCrud.InputTypeDateTime,
			ReadOnly:  true,
		},
	},
}

func (d *Domain) TenantAdminBankAccounts(in *TenantAdminBankAccountsIn) (out TenantAdminBankAccountsOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminBankAccountsUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
		out.SetError(400, ErrTenantAdminBankAccountsTenantNotFound)
		return
	}

	if in.WithMeta {
		out.Meta = &TenantAdminBankAccountsMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		if in.Account.Id <= 0 {
			out.Meta = &SuperAdminUserManagementMeta

			if user.Role != TenantAdminSegment {
				out.SetError(400, ErrTenantAdminBankAccountsNotTenant)
				return
			}
			staff := rqAuth.NewUsers(d.AuthOltp)
			staffs := staff.FindStaffsByTenantCode(tenant.TenantCode)

			out.Staffs = &staffs
			return
		}

		account := rqBudget.NewBankAccounts(d.AuthOltp)
		account.Id = in.Account.Id

		if !account.FindById() {
			out.SetError(400, ErrTenantAdminBankAccountsNotFound)
			return
		}

		out.Account = account
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		if user.Role != TenantAdminSegment {
			out.SetError(400, ErrTenantAdminBankAccountsNotTenant)
			return
		}

		account := wcBudget.NewBankAccountsMutator(d.AuthOltp)
		account.Id = in.Account.Id

		if account.Id > 0 {
			if !account.FindById() {
				out.SetError(400, ErrTenantAdminBankAccountsNotFound)
				return
			}

			if in.Cmd == zCrud.CmdDelete {
				if account.DeletedAt == 0 {
					account.SetDeletedAt(in.UnixNow())
					account.SetDeletedBy(sess.UserId)
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if account.DeletedAt > 0 {
					account.SetDeletedAt(0)
					account.SetRestoredBy(sess.UserId)
				}
			}
		} else {
			if in.Account.ParentBankAccountId > 0 {
				parentAccount := rqBudget.NewBankAccounts(d.AuthOltp)
				parentAccount.Id = in.Account.ParentBankAccountId

				if !parentAccount.FindById() {
					out.SetError(400, ErrTenantAdminBankAccountsParentNotFound)
					return
				}

				if parentAccount.ChildBankAccountId > 0 {
					out.SetError(400, ErrTenantAdminBankAccountsParentHaveChild)
					return
				}

				account.SetParentBankAccountId(in.Account.ParentBankAccountId)
			}
			if in.Account.StaffId > 0 {
				staff := rqAuth.NewUsers(d.AuthOltp)
				if !staff.FindStaffByIdByTenantCode(in.Account.StaffId, tenant.TenantCode) {
					out.SetError(400, ErrTenantAdminBankAccountsStaffNotFound)
					return
				}

				account.SetStaffId(staff.Id)
			}

			if in.Account.IsProfitCenter {
				coaParent := wcFinance.NewCoaMutator(d.AuthOltp)
				coaParent.Id = tenant.CustomersCoaId
				if !coaParent.FindById() {
					out.SetError(400, ErrTenantAdminBankAccountsCustomerCoaNotFound)
					return
				}

				coa := wcFinance.NewCoaMutator(d.AuthOltp)
				coa.SetParentId(coaParent.Id)
				coa.SetName(in.Account.Name)
				coa.SetTenantCode(tenant.TenantCode)
				coa.SetCreatedAt(in.UnixNow())
				coa.SetCreatedBy(sess.UserId)
				coa.SetUpdatedAt(in.UnixNow())
				coa.SetUpdatedBy(sess.UserId)
				if !coa.DoInsert() {
					out.SetError(400, ErrTenantAdminBankAccountsCustomerCoaSaveFailed)
					return
				}

				children := coaParent.Children
				children = append(children, coa.Id)
				coaParent.SetChildren(children)
				coaParent.SetUpdatedAt(in.UnixNow())
				coaParent.SetUpdatedBy(sess.UserId)
				if !coaParent.DoUpdateById() {
					out.SetError(400, ErrTenantAdminBankAccountsCustomerCoaParentUpdateFailed)
					return
				}
			}
			account.SetIsProfitCenter(in.Account.IsProfitCenter)
			
			if in.Account.IsCostCenter {
				coaParent := wcFinance.NewCoaMutator(d.AuthOltp)
				coaParent.Id = tenant.SuppliersCoaId
				if !coaParent.FindById() {
					out.SetError(400, ErrTenantAdminBankAccountsSupplierCoaNotFound)
					return
				}

				coa := wcFinance.NewCoaMutator(d.AuthOltp)
				coa.SetParentId(coaParent.Id)
				coa.SetName(in.Account.Name)
				coa.SetTenantCode(tenant.TenantCode)
				coa.SetCreatedAt(in.UnixNow())
				coa.SetCreatedBy(sess.UserId)
				coa.SetUpdatedAt(in.UnixNow())
				coa.SetUpdatedBy(sess.UserId)
				if !coa.DoInsert() {
					out.SetError(400, ErrTenantAdminBankAccountsSupplierCoaSaveFailed)
					return
				}

				children := coaParent.Children
				children = append(children, coa.Id)
				coaParent.SetChildren(children)
				coaParent.SetUpdatedAt(in.UnixNow())
				coaParent.SetUpdatedBy(sess.UserId)
				if !coaParent.DoUpdateById() {
					out.SetError(400, ErrTenantAdminBankAccountsSupplierCoaParentUpdateFailed)
					return
				}
			} 
			account.SetIsCostCenter(in.Account.IsCostCenter)

			if !(in.Account.IsProfitCenter && in.Account.IsCostCenter) {
				coaParent := wcFinance.NewCoaMutator(d.AuthOltp)
				coaParent.Id = tenant.BanksCoaId
				if !coaParent.FindById() {
					out.SetError(400, ErrTenantAdminBankAccountsBankCoaNotFound)
					return
				}

				coa := wcFinance.NewCoaMutator(d.AuthOltp)
				coa.SetParentId(coaParent.Id)
				coa.SetName(in.Account.Name)
				coa.SetTenantCode(tenant.TenantCode)
				coa.SetCreatedAt(in.UnixNow())
				coa.SetCreatedBy(sess.UserId)
				coa.SetUpdatedAt(in.UnixNow())
				coa.SetUpdatedBy(sess.UserId)
				if !coa.DoInsert() {
					out.SetError(400, ErrTenantAdminBankAccountsBankCoaSaveFailed)
					return
				}

				children := coaParent.Children
				children = append(children, coa.Id)
				coaParent.SetChildren(children)
				coaParent.SetUpdatedAt(in.UnixNow())
				coaParent.SetUpdatedBy(sess.UserId)
				if !coaParent.DoUpdateById() {
					out.SetError(400, ErrTenantAdminBankAccountsBankCoaParentUpdateFailed)
					return
				}
			}
		}

		account.SetTenantCode(user.TenantCode)

		if in.Account.AccountName != `` {
			account.SetAccountName(in.Account.AccountName)
		}
		if in.Account.BankName != `` {
			account.SetBankName(in.Account.BankName)
		}
		if in.Account.Name != `` {
			account.SetName(in.Account.Name)
		}
		if in.Account.AccountNumber > 0 {
			account.SetAccountNumber(in.Account.AccountNumber)
		}

		if account.HaveMutation() {
			account.SetUpdatedAt(in.UnixNow())
			account.SetUpdatedBy(sess.UserId)
			if account.Id == 0 {
				account.SetCreatedAt(in.UnixNow())
				account.SetCreatedBy(sess.UserId)
			}
		}

		if !account.DoUpsertById() {
			out.SetError(500, ErrTenantAdminBankAccountsSaveFailed)
		}

		out.Account = &account.BankAccounts

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		accounts := rqBudget.NewBankAccounts(d.AuthOltp)
		accounts.TenantCode = user.TenantCode
		out.Accounts = accounts.FindByPagination(&TenantAdminBankAccountsMeta, &in.Pager, &out.Pager)
	}

	return
}
