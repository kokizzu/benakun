package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"

	"github.com/kokizzu/gotro/T"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file SuperAdminUserManagement.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type SuperAdminUserManagement.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type SuperAdminUserManagement.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type SuperAdminUserManagement.go
//go:generate farify doublequote --file SuperAdminUserManagement.go

type (
	SuperAdminUserManagementIn struct {
		RequestCommon
		Cmd         string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		TenantAdmin string        `json:"tenantAdmin" form:"tenantAdmin" query:"tenantAdmin" long:"tenantAdmin" msg:"tenantAdmin"`
		User        rqAuth.Users  `json:"user" form:"user" query:"user" long:"user" msg:"user"`
		WithMeta    bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager       zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	SuperAdminUserManagementOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		User  *rqAuth.Users  `json:"user" form:"user" query:"user" long:"user" msg:"user"`
		// listing
		Users [][]any `json:"users" form:"users" query:"users" long:"users" msg:"users"`
	}
)

const (
	SuperAdminUserManagementAction              = `superAdmin/userManagement`
	ErrUserIdNotFound                           = `user id not found`
	ErrTenantCodeNotFound                       = `tenant code is not allready`
	ErrInvalidRole                              = `invalid role`
	ErrUserSaveFailed                           = `user save failed`
	ErrUsersEmailDuplicate                      = `email already by another user`
	ErrUsersEmailEmpty                          = `email cannot be empty`
	ErrSuperAdminUserManagementTenantAdminEmpty = `tenant admin cannot be empty`
	ErrSuperAdminUserManagementTenantCodeEmpty  = `tenant code cannot be empty`
)

var SuperAdminUserManagementMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mAuth.Id,
			Label:     "ID",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly:  true,
		},
		{
			Name:     mAuth.TenantCode,
			Label:    "Kode Tenant / Tenant Code",
			DataType: zCrud.DataTypeString,
			ReadOnly: true,
		},
		{
			Name:      mAuth.Email,
			Label:     "Surel / Email",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mAuth.FullName,
			Label:     "Nama Lengkap / Full Name",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name:      mAuth.Role,
			Label:     "Peran / Role",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeCombobox,
			Ref: []string{
				mAuth.RoleUser, mAuth.RoleDataEntry, mAuth.RoleReportViewer, mAuth.RoleTenantAdmin,
			},
		},
		{
			Name:      mAuth.InvitationState,
			Label:     "Status Udangan / Invitation State",
			DataType:  zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly:  true,
		},
		{
			Name:      mAuth.CreatedAt,
			Label:     `Dibuat pada / Created at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mAuth.UpdatedAt,
			Label:     `Diperbarui pada / Updated at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mAuth.DeletedAt,
			Label:     `Dihapus pada / Deleted at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
		{
			Name:      mAuth.LastLoginAt,
			Label:     `Login terakhir pada / Last login at`,
			ReadOnly:  true,
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeDateTime,
		},
	},
}

func (d *Domain) SuperAdminUserManagement(in *SuperAdminUserManagementIn) (out SuperAdminUserManagementOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	out.refId = in.User.Id

	sess := d.MustSuperAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.WithMeta {
		out.Meta = &SuperAdminUserManagementMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		if in.User.Id <= 0 {
			out.Meta = &SuperAdminUserManagementMeta
		}

		user := rqAuth.NewUsers(d.AuthOltp)
		user.Id = in.User.Id
		if !user.FindById() {
			out.SetError(400, ErrUserIdNotFound)
			return
		}
		user.CensorFields()
		out.User = user
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		user := wcAuth.NewUsersMutator(d.AuthOltp)
		user.Id = in.User.Id
		if user.Id > 0 {
			if !user.FindById() {
				out.SetError(400, ErrUserIdNotFound)
				return
			}

			if in.Cmd == zCrud.CmdUpsert {
				if in.User.Email != `` && in.User.Email != user.Email {
					dup := rqAuth.NewUsers(d.AuthOltp)
					dup.Email = in.User.Email
					if dup.FindByEmail() {
						out.SetError(400, ErrUsersEmailDuplicate)
						return
					}
					user.SetEmail(in.User.Email)
				}
			} else if in.Cmd == zCrud.CmdDelete {
				if user.DeletedAt == 0 {
					user.SetDeletedAt(in.UnixNow())
				}
			} else if in.Cmd == zCrud.CmdRestore {
				if user.DeletedAt > 0 {
					user.SetDeletedAt(0)
				}
			}
		} else {
			if in.User.Email != `` {
				dup := rqAuth.NewUsers(d.AuthOltp)
				dup.Email = in.User.Email
				if dup.FindByEmail() {
					out.SetError(400, ErrUsersEmailDuplicate)
					return
				}
				user.SetEmail(in.User.Email)
				user.SetVerifiedAt(0)
			} else {
				out.SetError(400, ErrUsersEmailEmpty)
				return
			}

			user.SetEncryptedPassword(mAuth.DefaultPassword, in.UnixNow())

			if in.User.Role != `` {
				if !mAuth.IsValidRole(in.User.Role) {
					out.SetError(400, ErrInvalidRole)
					return
				}

				switch in.User.Role {
				case mAuth.RoleDataEntry, mAuth.RoleReportViewer:
					if in.TenantAdmin != `` {
						tenantAdmin := rqAuth.NewTenants(d.AuthOltp)
						tenantAdmin.TenantCode = in.TenantAdmin
						if !tenantAdmin.FindByTenantCode() {
							out.SetError(400, ErrTenantCodeNotFound)
							return
						}

						invState := mAuth.InviteState{
							TenantCode: in.TenantAdmin,
							State:      mAuth.InvitationStateAccepted,
							Date:       T.DateStr(),
						}

						user.SetInvitationState(invState.ToStateString())
					} else {
						out.SetError(400, ErrSuperAdminUserManagementTenantAdminEmpty)
						return
					}
				case mAuth.RoleTenantAdmin:
					if in.User.TenantCode != `` {
						tenant := rqAuth.NewTenants(d.AuthOltp)
						tenant.TenantCode = in.User.TenantCode
						if !tenant.FindByTenantCode() {
							out.SetError(400, ErrTenantCodeNotFound)
							return
						}
						user.SetTenantCode(in.User.TenantCode)
					} else {
						out.SetError(400, ErrSuperAdminUserManagementTenantCodeEmpty)
						return
					}
				}
			}
		}

		if in.User.FullName != user.FullName {
			user.SetFullName(in.User.FullName)
		}

		if user.HaveMutation() {
			user.SetUpdatedAt(in.UnixNow())
			user.SetUpdatedBy(sess.UserId)
			if user.Id == 0 {
				user.SetCreatedAt(in.UnixNow())
				user.SetCreatedBy(sess.UserId)
			}
		}

		if !user.DoUpsert() {
			out.SetError(500, ErrUserSaveFailed)
		}

		user.CensorFields()
		out.User = &user.Users

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		r := rqAuth.NewUsers(d.AuthOltp)
		out.Users = r.FindByPagination(&SuperAdminUserManagementMeta, &in.Pager, &out.Pager)
	}
	return
}
