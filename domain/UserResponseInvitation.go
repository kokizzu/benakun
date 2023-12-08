package domain

import (
	"errors"

	"benakun/model/mAuth"
	"benakun/model/mAuth/wcAuth"
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

	if in.Response == mAuth.InvitationStateRespAccept {
		in.Response = mAuth.InvitationStateAccepted
	} else if in.Response == mAuth.InvitationStateRespReject {
		in.Response = mAuth.InvitationStateRejected
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

	mapState, err := mAuth.ToInvitationStateMap(user.InvitationState)
	if errors.Is(err, mAuth.ErrInvitationStateEmpty) {
		out.SetError(400, err.Error())
		out.Message = err.Error()
		return
	} else {
		err := mapState.ModifyState(in.TenantCode, in.Response)
		if err != nil {
			out.SetError(400, err.Error())
			out.Message = err.Error()
			return
		}
		user.SetInvitationState(mapState.ToStateString())
	}

	if in.Response == mAuth.InvitationStateAccepted {
		user.SetRole(DataEntrySegment)
		out.Message = user.Email + ` join to company ` + in.TenantCode
	} else if in.Response == mAuth.InvitationStateRejected {
		out.Message = user.Email + ` rejected the invitation`
	}

	if !user.DoUpdateById() {
		out.SetError(500, ErrUserResponseInvitationModificationFailed)
		out.Message = ErrUserResponseInvitationModificationFailed
		return
	}

	return
}
