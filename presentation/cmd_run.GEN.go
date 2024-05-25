package presentation

import (
	"os"

	"benakun/domain"
)


// Code generated by 1_codegen_test.go DO NOT EDIT.


func cmdRun(b *domain.Domain, action string, payload []byte) {
	switch action {
	case domain.DataEntryDashboardAction:
		in := domain.DataEntryDashboardIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.DataEntryDashboard(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestAutoLoginAction:
		in := domain.GuestAutoLoginIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestAutoLogin(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestDebugAction:
		in := domain.GuestDebugIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestDebug(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestExternalAuthAction:
		in := domain.GuestExternalAuthIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestExternalAuth(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestForgotPasswordAction:
		in := domain.GuestForgotPasswordIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestForgotPassword(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestLoginAction:
		in := domain.GuestLoginIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestLogin(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestOauthCallbackAction:
		in := domain.GuestOauthCallbackIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestOauthCallback(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestRegisterAction:
		in := domain.GuestRegisterIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestRegister(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestResendVerificationEmailAction:
		in := domain.GuestResendVerificationEmailIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestResendVerificationEmail(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestResetPasswordAction:
		in := domain.GuestResetPasswordIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestResetPassword(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.GuestVerifyEmailAction:
		in := domain.GuestVerifyEmailIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.GuestVerifyEmail(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.ReportViewerDashboardAction:
		in := domain.ReportViewerDashboardIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.ReportViewerDashboard(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.SuperAdminDashboardAction:
		in := domain.SuperAdminDashboardIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.SuperAdminDashboard(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.SuperAdminTenantManagementAction:
		in := domain.SuperAdminTenantManagementIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.SuperAdminTenantManagement(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.SuperAdminUserManagementAction:
		in := domain.SuperAdminUserManagementIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.SuperAdminUserManagement(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminBankAccountsAction:
		in := domain.TenantAdminBankAccountsIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminBankAccounts(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminBudgetingAction:
		in := domain.TenantAdminBudgetingIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminBudgeting(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminCoaAction:
		in := domain.TenantAdminCoaIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminCoa(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminCreateBudgetPlanAction:
		in := domain.TenantAdminCreateBudgetPlanIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminCreateBudgetPlan(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminCreateCoaChildAction:
		in := domain.TenantAdminCreateCoaChildIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminCreateCoaChild(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminCreateOrganizationChildAction:
		in := domain.TenantAdminCreateOrganizationChildIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminCreateOrganizationChild(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminDashboardAction:
		in := domain.TenantAdminDashboardIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminDashboard(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminDeleteCoaChildAction:
		in := domain.TenantAdminDeleteCoaChildIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminDeleteCoaChild(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminDeleteOrganizationChildAction:
		in := domain.TenantAdminDeleteOrganizationChildIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminDeleteOrganizationChild(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminGetBudgetPlansAction:
		in := domain.TenantAdminGetBudgetPlansIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminGetBudgetPlans(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminMoveOrganizationChildAction:
		in := domain.TenantAdminMoveOrganizationChildIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminMoveOrganizationChild(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminOrganizationAction:
		in := domain.TenantAdminOrganizationIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminOrganization(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminProductsAction:
		in := domain.TenantAdminProductsIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminProducts(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminRestoreCoaChildAction:
		in := domain.TenantAdminRestoreCoaChildIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminRestoreCoaChild(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminRestoreOrganizationChildAction:
		in := domain.TenantAdminRestoreOrganizationChildIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminRestoreOrganizationChild(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminTransactionAction:
		in := domain.TenantAdminTransactionIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminTransaction(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminUpdateBudgetPlanAction:
		in := domain.TenantAdminUpdateBudgetPlanIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminUpdateBudgetPlan(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminUpdateCoaChildAction:
		in := domain.TenantAdminUpdateCoaChildIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminUpdateCoaChild(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.TenantAdminUpdateOrganizationChildAction:
		in := domain.TenantAdminUpdateOrganizationChildIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.TenantAdminUpdateOrganizationChild(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserAutoLoginLinkAction:
		in := domain.UserAutoLoginLinkIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserAutoLoginLink(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserChangePasswordAction:
		in := domain.UserChangePasswordIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserChangePassword(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserCreateCompanyAction:
		in := domain.UserCreateCompanyIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserCreateCompany(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserLogoutAction:
		in := domain.UserLogoutIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserLogout(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserProfileAction:
		in := domain.UserProfileIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserProfile(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserResponseInvitationAction:
		in := domain.UserResponseInvitationIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserResponseInvitation(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserSessionKillAction:
		in := domain.UserSessionKillIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserSessionKill(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserSessionsActiveAction:
		in := domain.UserSessionsActiveIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserSessionsActive(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	case domain.UserUpdateProfileAction:
		in := domain.UserUpdateProfileIn{}
		if !in.RequestCommon.FromCli(action, payload, &in) {
			return
		}
		out := b.UserUpdateProfile(&in)
		in.RequestCommon.ToCli(os.Stdout, out, out.ResponseCommon)

	}
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
