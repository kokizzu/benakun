package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"

	"benakun/domain"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mBusiness/rqBusiness"
	"benakun/model/mFinance/rqFinance"
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

		var myCompany *rqAuth.Orgs = &rqAuth.Orgs{}
		if user != nil && user.TenantCode != `` {
			myCompany = rqAuth.NewOrgs(d.AuthOltp)
			myCompany.FindCompanyByTenantCode(user.TenantCode)
		}

		return views.RenderIndex(c, M.SX{
			`title`:     `BenAkun`,
			`user`:      user,
			`google`:    google.Link,
			`segments`:  segments,
			`myCompany`: myCompany,
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
		if err != nil {
			msg = err.Error()
			L.Print(`Error response: `, msg)
		} else {
			out := d.UserResponseInvitation(&in)
			msg = out.Message
		}
		return views.RenderUserResponsejoin(c, M.SX{
			`title`:   `Respond invitation`,
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

	fw.Get(`/`+domain.DataEntryDashboardAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notDataEntryLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		tenantCode, err := domain.GetTenantCodeByHost(in.Host)
		if err != nil {
			return ctx.Redirect(`/`, 302)
		}

		r := rqFinance.NewTransactionTemplate(d.AuthOltp)
		r.TenantCode = tenantCode
		trxTemplates := r.FindTransactionTemplatesByTenant()

		return views.RenderDataEntryDashboard(ctx, M.SX{
			`title`:                `Data Entry Dashboard`,
			`user`:                 user,
			`segments`:             segments,
			`transactionTemplates`: trxTemplates,
		})
	})

	fw.Get(`/`+domain.DataEntryTemplateAction+`/:templateId`, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notDataEntryLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		tenantCode, err := domain.GetTenantCodeByHost(in.Host)
		if err != nil {
			L.Print(`TENANT CODE:`, tenantCode)
			return ctx.Redirect(`/`, 302)
		}

		templateId := ctx.Params("templateId")

		trxTemplate := rqFinance.NewTransactionTemplate(d.AuthOltp)
		trxTemplate.Id = S.ToU(templateId)
		trxTemplate.TenantCode = tenantCode
		if !trxTemplate.FindByIdByTenantCode() {
			return ctx.Redirect(`/`+domain.DataEntryDashboardAction, 302)
		}

		return views.RenderDataEntryTemplatesTemplate(ctx, M.SX{
			`title`:      `Data Entry Template`,
			`user`:       user,
			`segments`:   segments,
			`transactiontemplate`: trxTemplate,
		})
	})

	fw.Get(`/`+domain.DataEntryTransactionEntryAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notDataEntryLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		return views.RenderDataEntryTransactionEntry(ctx, M.SX{
			`title`:    `Transaction Entry`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.TenantAdminDashboardAction, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminDashboardIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminDashboardAction)
		if err != nil {
			return err
		}
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.TenantAdminDashboard(&in)
		return views.RenderTenantAdminDashboard(ctx, M.SX{
			`title`:    `Tenant Admin Dashboard`,
			`user`:     user,
			`segments`: segments,
			`staffs`:   out.Staffs,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
		})
	})

	fw.Get(`/`+domain.TenantAdminLocationsAction, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminLocationsIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminLocationsAction)
		if err != nil {
			return err
		}
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.TenantAdminLocations(&in)
		return views.RenderTenantAdminLocations(ctx, M.SX{
			`title`:     `Tenant Admin Locations`,
			`user`:      user,
			`segments`:  segments,
			`fields`:    out.Meta.Fields,
			`pager`:     out.Pager,
			`locations`: out.Locations,
		})
	})

	fw.Get(`/`+domain.TenantAdminProductsAction, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminProductsIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminProductsAction)
		if err != nil {
			return err
		}
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.TenantAdminProducts(&in)
		return views.RenderTenantAdminProducts(ctx, M.SX{
			`title`:    `Tenant Admin Products`,
			`user`:     user,
			`segments`: segments,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
			`product`:  out.Product,
			`products`: out.Products,
		})
	})

	fw.Get(`/`+domain.TenantAdminBudgetingAction, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminBudgetingIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminBudgetingAction)
		if err != nil {
			return err
		}
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.Cmd = zCrud.CmdList

		out := d.TenantAdminBudgeting(&in)
		return views.RenderTenantAdminBudgeting(ctx, M.SX{
			`title`:    `Tenant Admin Budgeting`,
			`user`:     user,
			`segments`: segments,
			`orgs`:     out.Orgs,
		})
	})

	fw.Get(`/`+domain.TenantAdminOrganizationAction, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminOrganizationIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminOrganizationAction)
		if err != nil {
			return err
		}
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.Cmd = zCrud.CmdList

		out := d.TenantAdminOrganization(&in)
		return views.RenderTenantAdminOrganization(ctx, M.SX{
			`title`:    `Tenant Admin Organization`,
			`user`:     user,
			`segments`: segments,
			`orgs`:     out.Orgs,
		})
	})

	fw.Get(`/`+domain.TenantAdminCoaAction, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminCoaIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminCoaAction)
		if err != nil {
			return err
		}
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.Cmd = zCrud.CmdList

		out := d.TenantAdminCoa(&in)
		return views.RenderTenantAdminCoa(ctx, M.SX{
			`title`:    `Tenant Admin Coa`,
			`user`:     user,
			`segments`: segments,
			`coas`:     out.Coas,
		})
	})

	fw.Get(`/`+domain.TenantAdminTransactionAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderTenantAdminTransaction(ctx, M.SX{
			`title`:    `Tenant Admin Transaction`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.TenantAdminBankAccountsAction, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminBankAccountsIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminBankAccountsAction)
		if err != nil {
			return err
		}
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)

		r := rqAuth.NewUsers(d.AuthOltp)
		staffs := r.FindStaffsChoicesByTenantCode(user.TenantCode)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.TenantAdminBankAccounts(&in)
		return views.RenderTenantAdminBankAccounts(ctx, M.SX{
			`title`:    `Tenant Admin Bank Accounts`,
			`user`:     user,
			`segments`: segments,
			`accounts`: out.Accounts,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
			`staffs`:   staffs,
		})
	})

	fw.Get(`/`+domain.TenantAdminInventoryChangesAction, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminInventoryChangesIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminInventoryChangesAction)
		if err != nil {
			return err
		}
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)
		r := rqBusiness.NewProducts(d.AuthOltp)
		products := r.FindProductsChoicesByTenantCode(user.TenantCode)

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.TenantAdminInventoryChanges(&in)
		return views.RenderTenantAdminInventoryChanges(ctx, M.SX{
			`title`:            `Tenant Admin Inventory Changes`,
			`user`:             user,
			`segments`:         segments,
			`fields`:           out.Meta.Fields,
			`pager`:            out.Pager,
			`inventoryChanges`: out.InventoryChanges,
			`products`:         products,
		})
	})

	fw.Get(`/`+domain.TenantAdminInventoryChangesAction+`/:productId`, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		in.RequestCommon.Action = domain.TenantAdminCoaAction

		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		productId := ctx.Params(`productId`)
		product := rqBusiness.NewProducts(d.AuthOltp)
		product.Id = S.ToU(productId)
		if !product.FindById() {
			return ctx.Redirect(`/`+domain.TenantAdminInventoryChangesAction, 302)
		}

		r := rqBusiness.NewProducts(d.AuthOltp)
		products := r.FindProductsChoicesByTenantCode(user.TenantCode)

		out := d.TenantAdminInventoryChanges(&domain.TenantAdminInventoryChangesIn{
			RequestCommon: in.RequestCommon,
			ProductId:     product.Id,
			Cmd:           zCrud.CmdList,
			WithMeta:      true,
		})

		return views.RenderTenantAdminInventoryChangesProduct(ctx, M.SX{
			`title`:            `Tenant Admin Products's Inventory Changes`,
			`user`:             user,
			`segments`:         segments,
			`product`:          product,
			`products`: 				products,
			`fields`:           out.Meta.Fields,
			`pager`:            out.Pager,
			`inventoryChanges`: out.InventoryChanges,
		})
	})

	fw.Get(`/`+domain.TenantAdminTransactionTemplateAction, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminTransactionTemplateIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminTransactionTemplateAction)
		if err != nil {
			return err
		}
		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)

		in.Cmd = zCrud.CmdList

		out := d.TenantAdminTransactionTemplate(&in)
		return views.RenderTenantAdminTransactionTemplate(ctx, M.SX{
			`title`:                `Tenant Admin Transaction Template`,
			`user`:                 user,
			`segments`:             segments,
			`transactionTemplates`: out.TransactionTemplates,
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
		out := d.SuperAdminDashboard(&in)

		return views.RenderSuperAdminDashboard(ctx, M.SX{
			`title`:                  `Super Admin Dashboard`,
			`segments`:               segments,
			`user`:                   user,
			`registeredUserTotal`:    out.RegisteredUserTotal,
			`registeredUserToday`:    out.RegisteredUserToday,
			`requestsPerDate`:        out.RequestsPerDate,
			`uniqueUserPerDate`:      out.UniqueUserPerDate,
			`uniqueIpPerDate`:        out.UniqueIpPerDate,
			`countPerActionsPerDate`: out.CountPerActionsPerDate,
		})
	})

	fw.Get(`/`+domain.SuperAdminAccessLogAction, func(ctx *fiber.Ctx) error {
		var in domain.SuperAdminAccessLogIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.SuperAdminAccessLogAction)
		if err != nil {
			return err
		}
		if notLogin(d, in.RequestCommon, true) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		out := d.SuperAdminAccessLog(&in)

		return views.RenderSuperAdminAccessLog(ctx, M.SX{
			`title`:    `Super Admin Access Log`,
			`segments`: segments,
			`user`:     user,
			`logs`: out.Logs,
			`fields`: out.Meta.Fields,
			`pager`: out.Pager,	
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
			`title`:    `Super Admin User Management`,
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
			`title`:    `Super Admin Tenant Management`,
			`segments`: segments,
			`tenants`:  out.Tenants,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
			`user`:     user,
			`tenant`:   out.Tenant,
		})
	})

	fw.Get(`/apidocs`, func(ctx *fiber.Ctx) error {
		return views.RenderApidocsIndex(ctx, M.SX{
			`title`:       `Benakun API Docs`,
			`description`: `Restful API Documentation of Benakun`,
		})
	})

	fw.Get(`/debug`, func(ctx *fiber.Ctx) error {
		return views.RenderDebug(ctx, M.SX{})
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
	var segments = M.SB{}
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

func notTenantLogin(d *domain.Domain, in domain.RequestCommon) bool {
	var check domain.ResponseCommon

	if sess := d.MustTenantAdmin(in, &check); sess == nil {
		return true
	}

	return false
}

func notDataEntryLogin(d *domain.Domain, in domain.RequestCommon) bool {
	var check domain.ResponseCommon

	if sess := d.MustDataEntry(in, &check); sess == nil {
		return true
	}

	return false
}
