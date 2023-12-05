package domain

import (
	"benakun/model/mAuth/wcAuth"

	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file UserResponseInvitation.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type UserResponseInvitation.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type UserResponseInvitation.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type UserResponseInvitation.go
//go:generate farify doublequote --file UserResponseInvitation.go

type (
	UserResponseInvitationIn struct {
		RequestCommon
		TenantCode string `json:"tenantCode" form:"tenantCode" query:"tenantCode" long:"tenantCode" msg:"tenantCode"`
		Response   string `json:"response" form:"response" query:"response" long:"response" msg:"response"`
	}
	UserResponseInvitationOut struct {
		ResponseCommon
		Message string `json:"message" form:"message" query:"message" long:"message" msg:"message"`
	}
)

const (
	UserResponseInvitationAction = `user/responseInvitation`

	ErrUserResponseInvitationInvalidResponse    = `invalid response for invitation`
	ErrUserResponseInvitationUserNotFound       = `user not found on invitation`
	ErrUserResponseInvitationTenantNotFound     = `tenant admin not found on invitation`
	ErrUserResponseInvitationModificationFailed = `failed verify invitation user`
)

func (d *Domain) UserResponseInvitation(in *UserResponseInvitationIn) (out UserResponseInvitationOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	if in.Response == InvitationStateRespAccept {
		in.Response = InvitationStateAccepted
	} else if in.Response == InvitationStateRespReject {
		in.Response = InvitationStateRejected
	} else {
		out.SetError(400, ErrUserResponseInvitationInvalidResponse)
		out.Message = ErrUserResponseInvitationInvalidResponse
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrUserResponseInvitationUserNotFound)
		out.Message = ErrUserResponseInvitationUserNotFound
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = in.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrUserResponseInvitationTenantNotFound)
		out.Message = ErrUserResponseInvitationTenantNotFound
		return
	}

	tenantUser := wcAuth.NewUsersMutator(d.AuthOltp)
	tenantUser.TenantCode = in.TenantCode
	if !tenantUser.FindByTenantCode() {
		out.SetError(400, ErrUserResponseInvitationTenantNotFound)
		out.Message = ErrUserResponseInvitationTenantNotFound
		return
	}

	mapState, err := ToInvitationStateMap(user.InvitationState)
	if err != nil {
		out.SetError(400, ErrUserResponseInvitationInvalidResponse)
		out.Message = ErrUserResponseInvitationInvalidResponse
	} else {
		err := mapState.ModifyState(in.TenantCode, in.Response)
		if err != nil {
			out.SetError(400, ErrUserResponseInvitationInvalidResponse)
			out.Message = ErrUserResponseInvitationModificationFailed
			return
		}
		user.SetInvitationState(mapState.ToStateString())
	}

	if !user.DoUpdateById() {
		out.SetError(500, ErrUserResponseInvitationModificationFailed)
		out.Message = ErrUserResponseInvitationModificationFailed
		return
	}

	if in.Response == InvitationStateAccepted {
		out.Message = user.Email + ` to join accepted the invitation`
	} else if in.Response == InvitationStateRejected {
		out.Message = user.Email + ` rejected the invitation`
	}

	if TenantState == InvitationStateRevoked {
		d.runSubtask(func() {
			err := d.Mailer.SendResponseStateTenantEmail(tenantUser.Email, user.Email, in.Response)
			L.IsError(err, `SendResponseStateTenantEmail`)
		})
	}

	return
}
