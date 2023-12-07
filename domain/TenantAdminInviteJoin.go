package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/wcAuth"

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

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminInviteJoinUserNotFound)
		return
	}

	if user.TenantCode == `` {
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
	invState := mAuth.InviteState{
		TenantCode: user.TenantCode,
		State:      mAuth.InvitationStateInvited,
		Date:       T.DateStr(),
	}
	if err != nil {
		userToInvite.SetInvitationState(invState.ToStateString())
	} else {
		err := mapState.ModifyState(user.TenantCode, mAuth.InvitationStateInvited)
		if err != nil {
			out.SetError(400, ErrTenantAdminInviteJoinInvalidInvitation)
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
		inviteRespUrl := in.Host + `/` + UserResponseInvitationAction + `?tenantCode=` + user.TenantCode + `&response=`
		err := d.Mailer.SendInviteUserEmail(user.TenantCode, userToInvite.Email, inviteRespUrl)
		L.IsError(err, `SendInviteUserEmail`)
		// TODO: insert failed event to clickhouse
	})

	out.Message = TenantAdminInviteJoinMsg
	return
}
