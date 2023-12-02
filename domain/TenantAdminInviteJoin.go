package domain

import (
	"benakun/model/mAuth/wcAuth"

	"github.com/kokizzu/gotro/L"
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

	// TODO: define string as array of string to store in db
	// [
	//  	'tenant:[john_0937]:accepted:2024/01/23',
	// 		'tenant:[budi_1882]:rejected:2023/12/15'
	// ]
	// TODO: convert it back to array when update

	userToInvite.SetInvitationState(InviteJoinInvited(user.TenantCode))
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
