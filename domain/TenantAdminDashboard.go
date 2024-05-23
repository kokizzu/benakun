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
		WithMeta		bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    		zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	TenantAdminDashboardOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Staffs [][]any `json:"staffs" form:"staffs" query:"staffs" long:"staffs" msg:"staffs"`
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
)

var TenantAdminDashboardMeta = zCrud.Meta{
	Fields: []zCrud.Field{
		{
			Name:      mAuth.Id,
			Label:     "ID",
			DataType:  zCrud.DataTypeInt,
			InputType: zCrud.InputTypeHidden,
		},
		{
			Name: mAuth.Email,
			Label: "Email",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeEmail,
		},
		{
			Name: mAuth.FullName,
			Label: "Full Name",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeText,
		},
		{
			Name: mAuth.Role,
			Label: "Role",
			DataType: zCrud.DataTypeString,
			InputType: zCrud.InputTypeCombobox,
			Ref:  []string{
				UserSegment, DataEntrySegment, TenantAdminSegment, ReportViewerSegment,
			},
		},
		{
			Name: mAuth.InvitationState,
			Label: "Status",
			ReadOnly: true,
		},
		{
			Name: mAuth.UpdatedAt,
			Label: "Updated At",
			InputType: zCrud.InputTypeDateTime,
			ReadOnly: true,
		},
	},
}

func (d *Domain) TenantAdminDashboard(in *TenantAdminDashboardIn) (out TenantAdminDashboardOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	L.Print(`Trigger 1`)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		L.Print(`Trigger 2`)
		return
	}

	L.Print(`Trigger 3`)
	user := wcAuth.NewUsersMutator(d.AuthOltp)
	L.Print(`Trigger 4`)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminDashboardUnauthorized)
		L.Print(`Trigger 5`)
		return
	}

	L.Print(`Trigger 6`)
	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	L.Print(`Trigger 7`)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
		out.SetError(400, ErrTenantAdminDashboardTenantNotFound)
		L.Print(`Trigger 8`)
		return
	}

	if in.WithMeta {
		out.Meta = &TenantAdminDashboardMeta
	}

	switch in.Cmd {
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore:
		if in.StaffEmail != `` {
			staff := wcAuth.NewUsersMutator(d.AuthOltp)
			staff.Email = in.StaffEmail
			if !staff.FindByEmail() {
				out.SetError(400, ErrTenantAdminDashboardUserNotFound)
				return
			}

			if in.Cmd == zCrud.CmdUpsert {
				if tenant.TenantCode == `` {
					out.SetError(400, ErrTenantAdminDashboardNotTenant)
					return
				}
				if staff.TenantCode != `` {
					out.SetError(400, ErrTenantAdminDashboardInvalidStaff)
					return
				}

				mapState, err := mAuth.ToInvitationStateMap(staff.InvitationState)
				if errors.Is(err, mAuth.ErrInvitationStateEmpty) {
					invState := mAuth.InviteState{
						TenantCode: tenant.TenantCode,
						State:      mAuth.InvitationStateInvited,
						Date:       T.DateStr(),
					}
					staff.SetInvitationState(invState.ToStateString())
				} else {
					err := mapState.ModifyState(tenant.TenantCode, mAuth.InvitationStateInvited)
					if err != nil {
						out.SetError(400, err.Error())
						return
					}
					staff.SetInvitationState(mapState.ToStateString())
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
