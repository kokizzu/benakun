package domain

import (
	"errors"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/T"

	"benakun/model/mAuth"
	"benakun/model/mAuth/wcAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminInviteUser.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminInviteUser.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminInviteUser.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminInviteUser.go
//go:generate farify doublequote --file TenantAdminInviteUser.go

type (
	TenantAdminInviteUserIn struct {
		RequestCommon
		Email string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	}

	TenantAdminInviteUserOut struct {
		ResponseCommon
		Message string `json:"message" form:"message" query:"message" long:"message" msg:"message"`
	}
)

const (
	TenantAdminInviteUserAction = `tenantAdmin/inviteUser`

	ErrTenantAdminInviteUserUserNotFound       = `user not found`
	ErrTenantAdminInviteUserInvalidUserEmail   = `invalid user email`
	ErrTenantAdminInviteUserInvalidTenantAdmin = `invalid tenant admin`
	ErrTenantAdminInviteUserInviteFailed       = `invite user failed`
	ErrTenantAdminInviteUserInvalidInvitation  = `cannot invite this user`

	TenantAdminInviteUserMsg = `User invited success`
)

func (d *Domain) TenantAdminInviteUser(in *TenantAdminInviteUserIn) (out TenantAdminInviteUserOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.Email == `` {
		out.SetError(400, ErrTenantAdminInviteUserInvalidUserEmail)
		return
	}

	tenantUser := wcAuth.NewUsersMutator(d.AuthOltp)
	tenantUser.Id = sess.UserId
	if !tenantUser.FindById() {
		out.SetError(400, ErrTenantAdminInviteUserUserNotFound)
		return
	}

	if tenantUser.TenantCode == `` {
		out.SetError(400, ErrTenantAdminInviteUserInvalidTenantAdmin)
		return
	}

	userToInvite := wcAuth.NewUsersMutator(d.AuthOltp)
	userToInvite.Email = in.Email
	if !userToInvite.FindByEmail() {
		out.SetError(400, ErrTenantAdminInviteUserUserNotFound)
		return
	}

	if userToInvite.TenantCode != `` {
		out.SetError(400, ErrTenantAdminInviteUserInvalidInvitation)
		return
	}

	mapState, err := mAuth.ToInvitationStateMap(userToInvite.InvitationState)
	if errors.Is(err, mAuth.ErrInvitationStateEmpty) {
		invState := mAuth.InviteState{
			TenantCode: tenantUser.TenantCode,
			State:      mAuth.InvitationStateInvited,
			Date:       T.DateStr(),
		}
		userToInvite.SetInvitationState(invState.ToStateString())
	} else {
		err := mapState.ModifyState(tenantUser.TenantCode, mAuth.InvitationStateInvited)
		if err != nil {
			out.SetError(400, err.Error())
			return
		}
		userToInvite.SetInvitationState(mapState.ToStateString())
	}

	if !userToInvite.DoUpdateByEmail() {
		userToInvite.HaveMutation()
		out.SetError(500, ErrTenantAdminInviteUserInviteFailed)
		return
	}

	d.runSubtask(func() {
		inviteRespUrl := in.Host + `/` + UserResponseInvitationAction + `?tenantCode=` + tenantUser.TenantCode + `&response=`
		err := d.Mailer.SendInviteUserEmail(tenantUser.TenantCode, userToInvite.Email, inviteRespUrl)
		L.IsError(err, `SendInviteUserEmail`)
		// TODO: insert failed event to clickhouse
	})

	out.Message = TenantAdminInviteUserMsg
	return
}
