package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"

	"benakun/model/zCrud"

	"benakun/domain"
	"benakun/model/mAuth/rqAuth"
)

func (w *WebServer) WebStatic(fw *fiber.App, d *domain.Domain) {

	fw.Get(`/privacy`, func(c *fiber.Ctx) error {
		return c.SendString(`TODO: replace with real privacy policy`)
	})

	fw.Get(`/tos`, func(c *fiber.Ctx) error {
		return c.SendString(`TODO: replace with real terms of service`)
	})
	fw.Get(`/`, func(c *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(c, d)
		google := d.GuestExternalAuth(&domain.GuestExternalAuthIn{
			RequestCommon: in.RequestCommon,
			Provider:      domain.OauthGoogle,
		})
		google.ResponseCommon.DecorateSession(c)
		return views.RenderIndex(c, M.SX{
			`title`:  `BenAkun`,
			`user`:   user,
			`google`: google.Link,

			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.SuperAdminUserManagementAction, func(ctx *fiber.Ctx) error {
		var in domain.SuperAdminUserManagementIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.SuperAdminUserManagementAction)
		if err != nil {
			return err
		}
		if notLogin(ctx, d, in.RequestCommon, true) {
			return ctx.Redirect(`/`, 302)
		}
		_, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.SuperAdminUserManagement(&in)
		return views.RenderSuperAdminUserManagement(ctx, M.SX{
			`title`:    `Users`,
			`segments`: segments,
			`users`:    out.Users,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})

	fw.Get(`/`+domain.GuestVerifyEmailAction, func(c *fiber.Ctx) error {
		var in domain.GuestVerifyEmailIn
		err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestVerifyEmailAction)
		var errStr, email string
		var verified bool
		if err != nil {
			errStr = err.Error()
		} else {
			out := d.GuestVerifyEmail(&in)
			verified = out.Ok
			errStr = out.Error
			email = out.Email
			out.DecorateSession(c)
		}
		return views.RenderGuestVerifyEmail(c, M.SX{
			`title`:    `Email Verification`,
			`verified`: verified,
			`email`:    email,
			`error`:    errStr,
		})
	})

	fw.Get(`/`+domain.GuestResetPasswordAction, func(c *fiber.Ctx) error {
		return views.RenderGuestResetPassword(c, M.SX{
			`title`: `Reset Password`,
		})
	})

	fw.Get(`/`+domain.SuperAdminDashboardAction, func(ctx *fiber.Ctx) error {
		var in domain.SuperAdminDashboardIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.SuperAdminDashboardAction)
		if err != nil {
			return err
		}
		if notLogin(ctx, d, in.RequestCommon, true) {
			return ctx.Redirect(`/`, 302)
		}
		_, segments := userInfoFromRequest(in.RequestCommon, d)
		// out := d.SuperAdminDashboard(&in)
		return views.RenderSuperAdminDashboard(ctx, M.SX{
			`title`:    `Users`,
			`segments`: segments,
		})
	})

}

func userInfoFromContext(c *fiber.Ctx, d *domain.Domain) (domain.UserProfileIn, *rqAuth.Users, M.SB) {
	var in domain.UserProfileIn
	err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserProfileAction)
	var user *rqAuth.Users
	segments := M.SB{}
	if err == nil {
		out := d.UserProfile(&in)
		user = out.User
		segments = out.Segments
	}
	return in, user, segments
}

func userInfoFromRequest(rc domain.RequestCommon, d *domain.Domain) (*rqAuth.Users, M.SB) {
	var user *rqAuth.Users
	segments := M.SB{}
	out := d.UserProfile(&domain.UserProfileIn{
		RequestCommon: rc,
	})
	user = out.User
	segments = out.Segments
	return user, segments
}

func notLogin(ctx *fiber.Ctx, d *domain.Domain, in domain.RequestCommon, superAdmin bool) bool {
	var check domain.ResponseCommon
	var sess *domain.Session

	if superAdmin {
		sess = d.MustSuperAdmin(in, &check)
	} else {
		sess = d.MustLogin(in, &check)
	}
	if sess == nil {
		// TODO: implement render error
		// _ = views.RenderError(ctx, M.SX{
		// 	`error`: check.Error,
		// })
		return true
	}
	return false
}
