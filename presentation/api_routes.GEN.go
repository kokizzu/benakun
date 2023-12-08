package presentation

// Code generated by 1_codegen_test.go DO NOT EDIT.

import (
	"github.com/gofiber/fiber/v2"

	"benakun/domain"
)

func ApiRoutes(fw *fiber.App, d *domain.Domain) {

	// DataEntryDashboard
	fw.Post("/"+domain.DataEntryDashboardAction, func(c *fiber.Ctx) error {
		in := domain.DataEntryDashboardIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.DataEntryDashboardAction); err != nil {
			return nil
		}
		out := d.DataEntryDashboard(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestAutoLogin
	fw.Post("/"+domain.GuestAutoLoginAction, func(c *fiber.Ctx) error {
		in := domain.GuestAutoLoginIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestAutoLoginAction); err != nil {
			return nil
		}
		out := d.GuestAutoLogin(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestDebug
	fw.Post("/"+domain.GuestDebugAction, func(c *fiber.Ctx) error {
		in := domain.GuestDebugIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestDebugAction); err != nil {
			return nil
		}
		out := d.GuestDebug(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestExternalAuth
	fw.Post("/"+domain.GuestExternalAuthAction, func(c *fiber.Ctx) error {
		in := domain.GuestExternalAuthIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestExternalAuthAction); err != nil {
			return nil
		}
		out := d.GuestExternalAuth(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestForgotPassword
	fw.Post("/"+domain.GuestForgotPasswordAction, func(c *fiber.Ctx) error {
		in := domain.GuestForgotPasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestForgotPasswordAction); err != nil {
			return nil
		}
		out := d.GuestForgotPassword(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestLogin
	fw.Post("/"+domain.GuestLoginAction, func(c *fiber.Ctx) error {
		in := domain.GuestLoginIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestLoginAction); err != nil {
			return nil
		}
		out := d.GuestLogin(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestOauthCallback
	fw.Post("/"+domain.GuestOauthCallbackAction, func(c *fiber.Ctx) error {
		in := domain.GuestOauthCallbackIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestOauthCallbackAction); err != nil {
			return nil
		}
		out := d.GuestOauthCallback(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestRegister
	fw.Post("/"+domain.GuestRegisterAction, func(c *fiber.Ctx) error {
		in := domain.GuestRegisterIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestRegisterAction); err != nil {
			return nil
		}
		out := d.GuestRegister(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestResendVerificationEmail
	fw.Post("/"+domain.GuestResendVerificationEmailAction, func(c *fiber.Ctx) error {
		in := domain.GuestResendVerificationEmailIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestResendVerificationEmailAction); err != nil {
			return nil
		}
		out := d.GuestResendVerificationEmail(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestResetPassword
	fw.Post("/"+domain.GuestResetPasswordAction, func(c *fiber.Ctx) error {
		in := domain.GuestResetPasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestResetPasswordAction); err != nil {
			return nil
		}
		out := d.GuestResetPassword(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// GuestVerifyEmail
	fw.Post("/"+domain.GuestVerifyEmailAction, func(c *fiber.Ctx) error {
		in := domain.GuestVerifyEmailIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.GuestVerifyEmailAction); err != nil {
			return nil
		}
		out := d.GuestVerifyEmail(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// ReportViewerDashboard
	fw.Post("/"+domain.ReportViewerDashboardAction, func(c *fiber.Ctx) error {
		in := domain.ReportViewerDashboardIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.ReportViewerDashboardAction); err != nil {
			return nil
		}
		out := d.ReportViewerDashboard(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// SuperAdminDashboard
	fw.Post("/"+domain.SuperAdminDashboardAction, func(c *fiber.Ctx) error {
		in := domain.SuperAdminDashboardIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.SuperAdminDashboardAction); err != nil {
			return nil
		}
		out := d.SuperAdminDashboard(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// SuperAdminTenantManagement
	fw.Post("/"+domain.SuperAdminTenantManagementAction, func(c *fiber.Ctx) error {
		in := domain.SuperAdminTenantManagementIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.SuperAdminTenantManagementAction); err != nil {
			return nil
		}
		out := d.SuperAdminTenantManagement(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// SuperAdminUserManagement
	fw.Post("/"+domain.SuperAdminUserManagementAction, func(c *fiber.Ctx) error {
		in := domain.SuperAdminUserManagementIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.SuperAdminUserManagementAction); err != nil {
			return nil
		}
		out := d.SuperAdminUserManagement(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// TenantAdminBudgeting
	fw.Post("/"+domain.TenantAdminBudgetingAction, func(c *fiber.Ctx) error {
		in := domain.TenantAdminBudgetingIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.TenantAdminBudgetingAction); err != nil {
			return nil
		}
		out := d.TenantAdminBudgeting(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// TenantAdminDashboard
	fw.Post("/"+domain.TenantAdminDashboardAction, func(c *fiber.Ctx) error {
		in := domain.TenantAdminDashboardIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.TenantAdminDashboardAction); err != nil {
			return nil
		}
		out := d.TenantAdminDashboard(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// TenantAdminInviteJoin
	fw.Post("/"+domain.TenantAdminInviteJoinAction, func(c *fiber.Ctx) error {
		in := domain.TenantAdminInviteJoinIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.TenantAdminInviteJoinAction); err != nil {
			return nil
		}
		out := d.TenantAdminInviteJoin(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// TenantAdminOrganization
	fw.Post("/"+domain.TenantAdminOrganizationAction, func(c *fiber.Ctx) error {
		in := domain.TenantAdminOrganizationIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.TenantAdminOrganizationAction); err != nil {
			return nil
		}
		out := d.TenantAdminOrganization(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// TenantAdminTerminateStaff
	fw.Post("/"+domain.TenantAdminTerminateStaffAction, func(c *fiber.Ctx) error {
		in := domain.TenantAdminTerminateStaffIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.TenantAdminTerminateStaffAction); err != nil {
			return nil
		}
		out := d.TenantAdminTerminateStaff(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserAutoLoginLink
	fw.Post("/"+domain.UserAutoLoginLinkAction, func(c *fiber.Ctx) error {
		in := domain.UserAutoLoginLinkIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserAutoLoginLinkAction); err != nil {
			return nil
		}
		out := d.UserAutoLoginLink(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserChangePassword
	fw.Post("/"+domain.UserChangePasswordAction, func(c *fiber.Ctx) error {
		in := domain.UserChangePasswordIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserChangePasswordAction); err != nil {
			return nil
		}
		out := d.UserChangePassword(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserCreateCompany
	fw.Post("/"+domain.UserCreateCompanyAction, func(c *fiber.Ctx) error {
		in := domain.UserCreateCompanyIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserCreateCompanyAction); err != nil {
			return nil
		}
		out := d.UserCreateCompany(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserLogout
	fw.Post("/"+domain.UserLogoutAction, func(c *fiber.Ctx) error {
		in := domain.UserLogoutIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserLogoutAction); err != nil {
			return nil
		}
		out := d.UserLogout(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserProfile
	fw.Post("/"+domain.UserProfileAction, func(c *fiber.Ctx) error {
		in := domain.UserProfileIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserProfileAction); err != nil {
			return nil
		}
		out := d.UserProfile(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserResponseInvitation
	fw.Post("/"+domain.UserResponseInvitationAction, func(c *fiber.Ctx) error {
		in := domain.UserResponseInvitationIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserResponseInvitationAction); err != nil {
			return nil
		}
		out := d.UserResponseInvitation(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserSessionKill
	fw.Post("/"+domain.UserSessionKillAction, func(c *fiber.Ctx) error {
		in := domain.UserSessionKillIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserSessionKillAction); err != nil {
			return nil
		}
		out := d.UserSessionKill(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserSessionsActive
	fw.Post("/"+domain.UserSessionsActiveAction, func(c *fiber.Ctx) error {
		in := domain.UserSessionsActiveIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserSessionsActiveAction); err != nil {
			return nil
		}
		out := d.UserSessionsActive(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

	// UserUpdateProfile
	fw.Post("/"+domain.UserUpdateProfileAction, func(c *fiber.Ctx) error {
		in := domain.UserUpdateProfileIn{}
		if err := webApiParseInput(c, &in.RequestCommon, &in, domain.UserUpdateProfileAction); err != nil {
			return nil
		}
		out := d.UserUpdateProfile(&in)
		return in.ToFiberCtx(c, out, &out.ResponseCommon, in)
	})

}

// Code generated by 1_codegen_test.go DO NOT EDIT.
