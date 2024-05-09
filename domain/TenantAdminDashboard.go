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
		Cmd      string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		Staff    rqAuth.StaffWithInvitation `json:"staff" form:"staff" query:"staff" long:"staff" msg:"staff"`
		WithMeta bool          `json:"withMeta" form:"withMeta" query:"withMeta" long:"withMeta" msg:"withMeta"`
		Pager    zCrud.PagerIn `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
	}
	TenantAdminDashboardOut struct {
		ResponseCommon
		Pager zCrud.PagerOut `json:"pager" form:"pager" query:"pager" long:"pager" msg:"pager"`
		Meta  *zCrud.Meta    `json:"meta" form:"meta" query:"meta" long:"meta" msg:"meta"`
		Staffs []rqAuth.StaffWithInvitation `json:"staffs" form:"staffs" query:"staffs" long:"staffs" msg:"staffs"`
	}
)

const (
	TenantAdminDashboardAction = `tenantAdmin/dashboard`

	ErrTenantAdminDashboardUnauthorized   = `unauthorized user`
	ErrTenantAdminDashboardTenantNotFound = `tenant admin not found`
	ErrTenantAdminDashboardStaffIdRequired = `staff id is required`
	ErrStaffIdNotFound = `staff id not found`
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

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminDashboardUnauthorized)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminDashboardTenantNotFound)
		return
	}

	if in.WithMeta {
		out.Meta = &TenantAdminDashboardMeta
	}

	switch in.Cmd {
	case zCrud.CmdUpsert:
		if in.Staff.Email != `` {
			user := wcAuth.NewUsersMutator(d.AuthOltp)
			user.Email = in.Staff.Email
			if !user.FindByEmail() {
				out.SetError(400, ErrTenantAdminInviteUserUserNotFound)
				return
			}

			if user.TenantCode != `` {
				out.SetError(400, ErrTenantAdminInviteUserInvalidInvitation)
				return
			}

			mapState, err := mAuth.ToInvitationStateMap(user.InvitationState)
			if errors.Is(err, mAuth.ErrInvitationStateEmpty) {
				invState := mAuth.InviteState{
					TenantCode: tenant.TenantCode,
					State:      mAuth.InvitationStateInvited,
					Date:       T.DateStr(),
				}
				user.SetInvitationState(invState.ToStateString())
			} else {
				err := mapState.ModifyState(tenant.TenantCode, mAuth.InvitationStateInvited)
				if err != nil {
					out.SetError(400, err.Error())
					return
				}
				user.SetInvitationState(mapState.ToStateString())
			}

			if !user.DoUpdateByEmail() {
				user.HaveMutation()
				out.SetError(500, ErrTenantAdminInviteUserInviteFailed)
				return
			}

			d.runSubtask(func() {
				inviteRespUrl := in.Host + `/` + UserResponseInvitationAction + `?tenantCode=` + tenant.TenantCode + `&response=`
				err := d.Mailer.SendInviteUserEmail(tenant.TenantCode, user.Email, inviteRespUrl)
				L.IsError(err, `SendInviteUserEmail`)
				// TODO: insert failed event to clickhouse
			})
		} else {
			out.SetError(400, ErrTenantAdminDashboardStaffIdRequired)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdDelete:
		user := wcAuth.NewUsersMutator(d.AuthOltp)
		user.Email = in.Staff.Email
		if !user.FindByEmail() {
			out.SetError(400, ErrTenantAdminTerminateStaffUserNotFound)
			return
		}

		mapState, err := mAuth.ToInvitationStateMap(user.InvitationState)
		if errors.Is(err, mAuth.ErrInvitationStateEmpty) {
			out.SetError(400, ErrTenantAdminTerminateStaffEmptyState)
			return
		} else {
			err := mapState.ModifyState(tenant.TenantCode, mAuth.InvitationStateTerminated)
			if err != nil {
				out.SetError(400, err.Error())
				return
			}
			user.SetInvitationState(mapState.ToStateString())
		}

		user.SetRole(UserSegment)

		if !user.DoUpdateByEmail() {
			user.HaveMutation()
			out.SetError(500, ErrTenantAdminTerminateStaffFailed)
			return
		}

		if in.Pager.Page == 0 {
			break
		}

		fallthrough
	case zCrud.CmdList:
		staffs := user.FindUsersByTenant(tenant.TenantCode)
		out.Staffs = staffs
	}

	return
}
