package domain

import (
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
		Ok    bool   `json:"ok" form:"ok" query:"ok" long:"ok" msg:"ok"`
		Email string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	}
)

const (
	UserResponseInvitationAction = `user/responseInvitation`

	ErrUserResponseInvitationInvalidResponse     = `invalid response`
	ErrUserResponseInvitationUserNotFound        = `user not found`
	ErrUserResponseInvitationTenantAdminNotFound = `tenant admin not found`
	ErrUserResponseInvitationModificationFailed  = `failed verify invitation user`
)

func (d *Domain) UserResponseInvitation(in *UserResponseInvitationIn) (out UserResponseInvitationOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrUserResponseInvitationUserNotFound)
		return
	}

	tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
	tenant.TenantCode = in.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrUserResponseInvitationTenantAdminNotFound)
		return
	}

	if in.Response == `accept` {
		user.SetInvitationState(InviteJoinAccepted(in.TenantCode))
	} else if in.Response == `reject` {
		user.SetInvitationState(InviteJoinRejected(in.TenantCode))
	} else {
		out.SetError(400, ErrUserResponseInvitationInvalidResponse)
		return
	}

	if !user.DoUpdateById() {
		out.SetError(500, ErrUserResponseInvitationModificationFailed)
		return
	}

	out.Ok = true
	out.Email = user.Email
	return
}
