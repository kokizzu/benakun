package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"

	"benakun/domain"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/zCrud"
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
			`title`:    `BenAkun`,
			`user`:     user,
			`google`:   google.Link,
			`segments`: segments,
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

	fw.Get(`/`+domain.UserResponseInvitationAction, func(c *fiber.Ctx) error {
		var in domain.UserResponseInvitationIn
		err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserResponseInvitationAction)
		var msg string
		var ok bool
		if err != nil {
			ok = false
			msg = err.Error()
		} else {
			out := d.UserResponseInvitation(&in)
			ok = out.Ok
			msg = out.Message
		}
		return views.RenderUserResponsejoin(c, M.SX{
			`title`:   `Respond invitation`,
			`ok`:      ok,
			`message`: msg,
		})
	})

	fw.Get(`/`+domain.UserProfileAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(d, in.RequestCommon, false) {
			return ctx.Redirect(`/`, 302)
		}
		in.RequestCommon.Action = domain.UserSessionsActiveAction
		out := d.UserSessionsActive(&domain.UserSessionsActiveIn{
			RequestCommon: in.RequestCommon,
		})
		return views.RenderUserProfile(ctx, M.SX{
			`title`:          `Profile`,
			`user`:           user,
			`segments`:       segments,
			`activeSessions`: out.SessionsActive,
		})
	})

	fw.Get(`/`+domain.DataEntrySegment, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(d, in.RequestCommon, false) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderDataEntryDashboard(ctx, M.SX{
			`title`:    `Data Entry Dashboard`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.ReportViewerDashboardAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(d, in.RequestCommon, false) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderReportViewerDashboard(ctx, M.SX{
			`title`:    `Report Viewer Dashboard`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.TenantAdminDashboardAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(d, in.RequestCommon, false) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderTenantAdminDashboard(ctx, M.SX{
			`title`:    `Tenant Admin Dashboard`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.TenantAdminBudgetingAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(d, in.RequestCommon, false) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderTenantAdminBudgeting(ctx, M.SX{
			`title`:    `Tenant Admin Budgeting`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.TenantAdminOrganizationAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(d, in.RequestCommon, false) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderTenantAdminBudgeting(ctx, M.SX{
			`title`:    `Tenant Admin Organization`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.SuperAdminDashboardAction, func(ctx *fiber.Ctx) error {
		var in domain.SuperAdminDashboardIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.SuperAdminDashboardAction)
		if err != nil {
			return err
		}
		if notLogin(d, in.RequestCommon, true) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		// out := d.SuperAdminDashboard(&in)
		return views.RenderSuperAdminDashboard(ctx, M.SX{
			`title`:    `Users`,
			`segments`: segments,
			`user`:     user,
		})
	})

	fw.Get(`/`+domain.SuperAdminUserManagementAction, func(ctx *fiber.Ctx) error {
		var in domain.SuperAdminUserManagementIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.SuperAdminUserManagementAction)
		if err != nil {
			return err
		}
		if notLogin(d, in.RequestCommon, true) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.SuperAdminUserManagement(&in)
		return views.RenderSuperAdminUserManagement(ctx, M.SX{
			`title`:    `Users`,
			`segments`: segments,
			`users`:    out.Users,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
			`user`:     user,
		})
	})

	fw.Get(`/`+domain.SuperAdminTenantManagementAction, func(ctx *fiber.Ctx) error {
		var in domain.SuperAdminTenantManagementIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.SuperAdminTenantManagementAction)
		if err != nil {
			return err
		}
		if notLogin(d, in.RequestCommon, true) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.SuperAdminTenantManagement(&in)
		return views.RenderSuperAdminTenantManagement(ctx, M.SX{
			`title`:    `Users`,
			`segments`: segments,
			`tenants`:  out.Tenants,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
			`user`:     user,
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

func notLogin(d *domain.Domain, in domain.RequestCommon, superAdmin bool) bool {
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
