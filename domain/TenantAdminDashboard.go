package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"
	"errors"

	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/T"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminDashboard.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminDashboard.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminDashboard.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminDashboard.go
//go:generate farify doublequote --file TenantAdminDashboard.go

type (
	TenantAdminDashboardIn struct {
		RequestCommon
		Cmd      		string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		StaffEmail	string				`json:"staffEmail" form:"staffEmail" query:"staffEmail" long:"staffEmail" msg:"staffEmail"`
		StaffRole		string				`json:"staffRole" form:"staffRole" query:"staffRole" long:"staffRole" msg:"staffRole"`
		TenantCode  string 				`json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
		Role 				string 				`json:"role" form:"role" query:"role" long:"role" msg:"role"`
		IsEdit  		bool 					`json:"isEdit" form:"isEdit" query:"isEdit" long:"isEdit" msg:"isEdit"`
		WithMeta		bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    		zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	TenantAdminDashboardOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Staffs [][]any `json:"staffs" form:"staffs" query:"staffs" long:"staffs" msg:"staffs"`
		StaffsForm *[]rqAuth.Staff `json:"staffsForm" form:"staffsForm" query:"staffsForm" long:"staffsForm" msg:"staffsForm"`
	}
)

const (
	TenantAdminDashboardAction = `tenantAdmin/dashboard`

	ErrTenantAdminDashboardUnauthorized   = `unauthorized user`
	ErrTenantAdminDashboardTenantNotFound = `tenant admin not found`
	ErrTenantAdminDashboardStaffEmailRequired = `staff email is required`
	ErrTenantAdminDashboardUserNotFound = `user not found` 
	ErrTenantAdminDashboardInvalidStaff = `invalid staff`
	ErrTenantAdminDashboardEmptyState = `failed to modify staff, state is empty`
	ErrTenantAdminDashboardFailed = `failed to update staff`
	ErrTenantAdminDashboardNotTenant = `cannot invite user if not tenant`
	ErrTenantAdminDashboardInvalidRole = `invalid staff role to modify`
	ErrTenantAdminDashboardInvalidStaffToInvite = `invalid staff to invite, please use valid email/role`
)

var TenantAdminDashboardMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mAuth.Id,
			Label:     "ID",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
			ReadOnly: true,
		},
		{
			Name: mAuth.Email,
			Label: "Surel/Email",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeEmail,
			ReadOnly: true,
		},
		{
			Name: mAuth.FullName,
			Label: "Nama Lengkap/Full Name",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
			ReadOnly: true,
		},
		{
			Name: mAuth.Role,
			Label: "Peran/Role",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeCombobox,
			Ref:  []string{
				UserSegment, DataEntrySegment, ReportViewerSegment,
			},
		},
		{
			Name: mAuth.InvitationState,
			Label: "Status",
			ReadOnly: true,
		},
		{
			Name: mAuth.CreatedAt,
			Label: "Dibuat pada/Created at",
			InputType: zCrud.InputTypeHidden,
			ReadOnly: true,
		},
		{
			Name: mAuth.UpdatedAt,
			Label: "Diperbarui pada/Updated at",
			InputType: zCrud.InputTypeHidden,
			ReadOnly: true,
		},
	},
}

func (d *Domain) TenantAdminDashboard(in *TenantAdminDashboardIn) (out TenantAdminDashboardOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	L.Print(`HOST: `, in.Host)

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminDashboardUnauthorized)
		return
	}

	if in.WithMeta {
		out.Meta = &TenantAdminDashboardMeta
	}

	switch in.Cmd {
	case zCrud.CmdForm:
		r := rqAuth.NewUsers(d.AuthOltp)
		staffs := r.FindStaffsByTenantCode(in.TenantCode)

		out.StaffsForm = &staffs
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.TenantCode = user.TenantCode
		if !tenant.FindByTenantCode() {
			out.SetError(400, ErrTenantAdminDashboardNotTenant)
			return
		}

		if in.StaffEmail != `` {
			staff := wcAuth.NewUsersMutator(d.AuthOltp)
			staff.Email = in.StaffEmail
			if !staff.FindByEmail() {
				out.SetError(400, ErrTenantAdminDashboardUserNotFound)
				return
			}

			if in.Cmd == zCrud.CmdUpsert {
				if staff.Role != in.Role && in.IsEdit {
					switch in.Role {
					case mAuth.RoleUser, mAuth.RoleDataEntry, mAuth.RoleReportViewer:
						break
					default:
						out.SetError(400, ErrTenantAdminDashboardInvalidRole)
						return
					}

					staff.SetRole(in.Role)
				} else {
					if staff.TenantCode != `` {
						out.SetError(400, ErrTenantAdminDashboardInvalidStaff)
						return
					}

					if in.StaffEmail == `` || in.StaffRole == `` {
						out.SetError(400, ErrTenantAdminDashboardInvalidStaffToInvite)
						return
					} 

					mapState, err := mAuth.ToInvitationStateMap(staff.InvitationState)
					if errors.Is(err, mAuth.ErrInvitationStateEmpty) {
						invState := mAuth.InviteState{
							TenantCode: tenant.TenantCode,
							Role: 			in.StaffRole,
							State:      mAuth.InvitationStateInvited,
							Date:       T.DateStr(),
						}
						staff.SetInvitationState(invState.ToStateString())
					} else {
						if err := mapState.ModifyState(tenant.TenantCode, mAuth.InvitationStateInvited); err != nil {
							out.SetError(400, err.Error())
							return
						}

						if err := mapState.ModifyRole(tenant.TenantCode, in.StaffRole); err != nil {
							out.SetError(400, err.Error())
							return
						}

						staff.SetInvitationState(mapState.ToStateString())
					}
				}
			} else if in.Cmd == zCrud.CmdDelete {
				mapState, err := mAuth.ToInvitationStateMap(staff.InvitationState)
				if errors.Is(err, mAuth.ErrInvitationStateEmpty) {
					out.SetError(400, ErrTenantAdminDashboardEmptyState)
					return
				} else {
					err := mapState.ModifyState(tenant.TenantCode, mAuth.InvitationStateTerminated)
					if err != nil {
						out.SetError(400, err.Error())
						return
					}
					staff.SetInvitationState(mapState.ToStateString())
				}

				staff.SetRole(UserSegment)
			} else if in.Cmd == zCrud.CmdRestore {
				mapState, err := mAuth.ToInvitationStateMap(staff.InvitationState)
				if errors.Is(err, mAuth.ErrInvitationStateEmpty) {
					out.SetError(400, ErrTenantAdminDashboardEmptyState)
					return
				} else {
					err := mapState.ModifyState(tenant.TenantCode, mAuth.InvitationStateInvited)
					if err != nil {
						out.SetError(400, err.Error())
						return
					}
					staff.SetInvitationState(mapState.ToStateString())
				}
			}

			staff.SetUpdatedAt(in.UnixNow())
			staff.SetUpdatedBy(user.Id)
			if !staff.DoUpdateByEmail() {
				staff.HaveMutation()
				out.SetError(500, ErrTenantAdminDashboardFailed)
				return
			}

			if in.Cmd == zCrud.CmdUpsert || in.Cmd == zCrud.CmdRestore {
				d.runSubtask(func() {
					inviteRespUrl := in.Host + `/` + UserResponseInvitationAction + `?tenantCode=` + tenant.TenantCode + `&response=`
					err := d.Mailer.SendInviteUserEmail(tenant.TenantCode, staff.Email, inviteRespUrl)
					L.IsError(err, `SendInviteUserEmail`)
					// TODO: insert failed event to clickhouse
				})
			}
		} else {
			out.SetError(400, ErrTenantAdminDashboardStaffEmailRequired)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		staff := rqAuth.NewUsers(d.AuthOltp)
		staff.TenantCode = user.TenantCode
		out.Staffs = staff.FindStaffByPagination(&TenantAdminDashboardMeta, &in.Pager, &out.Pager)
	}

	return
}
