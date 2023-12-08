package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/wcAuth"
	"errors"
	"fmt"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/T"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminInviteJoin.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminInviteJoin.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminInviteJoin.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminInviteJoin.go
//go:generate farify doublequote --file TenantAdminInviteJoin.go

type (
	TenantAdminInviteJoinIn struct {
		RequestCommon
		Email string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	}

	TenantAdminInviteJoinOut struct {
		ResponseCommon
		Message string `json:"message" form:"message" query:"message" long:"message" msg:"message"`
	}
)

const (
	TenantAdminInviteJoinAction = `tenantAdmin/inviteUser`

	ErrTenantAdminInviteJoinUserNotFound       = `user not found`
	ErrTenantAdminInviteJoinInvalidUserEmail   = `invalid user email`
	ErrTenantAdminInviteJoinInvalidTenantAdmin = `invalid tenant admin`
	ErrTenantAdminInviteJoinInviteFailed       = `invite user failed`
	ErrTenantAdminInviteJoinInvalidInvitation  = `cannot invite this user`

	TenantAdminInviteJoinMsg = `User invited success`
)

func (d *Domain) TenantAdminInviteJoin(in *TenantAdminInviteJoinIn) (out TenantAdminInviteJoinOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.Email == `` {
		out.SetError(400, ErrTenantAdminInviteJoinInvalidUserEmail)
		return
	}

	tenantUser := wcAuth.NewUsersMutator(d.AuthOltp)
	tenantUser.Id = sess.UserId
	if !tenantUser.FindById() {
		out.SetError(400, ErrTenantAdminInviteJoinUserNotFound)
		return
	}

	if tenantUser.TenantCode == `` {
		out.SetError(400, ErrTenantAdminInviteJoinInvalidTenantAdmin)
		return
	}

	userToInvite := wcAuth.NewUsersMutator(d.AuthOltp)
	userToInvite.Email = in.Email
	if !userToInvite.FindByEmail() {
		out.SetError(400, ErrTenantAdminInviteJoinUserNotFound)
		return
	}

	if userToInvite.TenantCode != `` {
		out.SetError(400, ErrTenantAdminInviteJoinInvalidInvitation)
		return
	}

	mapState, err := mAuth.ToInvitationStateMap(userToInvite.InvitationState)
	if errors.Is(err, fmt.Errorf(mAuth.ErrInvitationStateEmpty)) {
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
		out.SetError(500, ErrTenantAdminInviteJoinInviteFailed)
		return
	}

	d.runSubtask(func() {
		inviteRespUrl := in.Host + `/` + UserResponseInvitationAction + `?tenantCode=` + tenantUser.TenantCode + `&response=`
		err := d.Mailer.SendInviteUserEmail(tenantUser.TenantCode, userToInvite.Email, inviteRespUrl)
		L.IsError(err, `SendInviteUserEmail`)
		// TODO: insert failed event to clickhouse
	})

	out.Message = TenantAdminInviteJoinMsg
	return
}
