package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/wcAuth"

	"github.com/kokizzu/gotro/T"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminTerminateStaff.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminTerminateStaff.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminTerminateStaff.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminTerminateStaff.go
//go:generate farify doublequote --file TenantAdminTerminateStaff.go

type (
	TenantAdminTerminateStaffIn struct {
		RequestCommon
		Email string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	}

	TenantAdminTerminateStaffOut struct {
		ResponseCommon
		Message string `json:"message" form:"message" query:"message" long:"message" msg:"message"`
	}
)

const (
	TenantAdminTerminateStaffAction = `tenantAdmin/terminateStaff`

	ErrTenantAdminTerminateStaffUserNotFound     = `user not found`
	ErrTenantAdminTerminateStaffTenantNotFound   = `tenant admin not found`
	ErrTenantAdminTerminateStaffInvalidUserEmail = `invalid user email`

	TenantAdminTerminateStaffMsg = `staff successfully terminated`
)

func (d *Domain) TenantAdminTerminateStaff(in *TenantAdminTerminateStaffIn) (out TenantAdminTerminateStaffOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.Email == `` {
		out.SetError(400, ErrTenantAdminTerminateStaffInvalidUserEmail)
		return
	}

	tenantUser := wcAuth.NewUsersMutator(d.AuthOltp)
	tenantUser.Id = sess.UserId
	if !tenantUser.FindById() {
		out.SetError(400, ErrTenantAdminTerminateStaffUserNotFound)
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = in.Email
	if !user.FindByEmail() {
		out.SetError(400, ErrTenantAdminTerminateStaffUserNotFound)
		return
	}

	mapState, err := mAuth.ToInvitationStateMap(user.InvitationState)
	if err != nil {
		invState := mAuth.InviteState{
			TenantCode: tenantUser.TenantCode,
			State:      mAuth.InvitationStateTerminated,
			Date:       T.DateStr(),
		}
		user.SetInvitationState(invState.ToStateString())
	} else {
		err := mapState.ModifyState(tenantUser.TenantCode, mAuth.InvitationStateTerminated)
		if err != nil {
			out.SetError(400, ErrTenantAdminInviteJoinInvalidInvitation)
			return
		}
		user.SetInvitationState(mapState.ToStateString())
	}

	out.Message = TenantAdminTerminateStaffMsg
	return
}
