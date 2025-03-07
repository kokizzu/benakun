package presentation

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"

	"benakun/domain"
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mBusiness/rqBusiness"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mInternal/rqInternal"
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

	fw.Get(`/`+domain.UserPurchaseSupportAction, func(c *fiber.Ctx) error {
		return views.RenderUserPurchaseSupport(c, M.SX{
			`title`: `Purchase Benakun Support+`,
		})
	})

	fw.All(`/payments/notifications`, d.PaymentsNotifications)

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

		return views.RenderUserProfile(ctx, M.SX{
			`title`:    `Profile`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/user/profileSessions`, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(d, in.RequestCommon, false) {
			return ctx.Redirect(`/`, 302)
		}
		in.RequestCommon.Action = domain.UserSessionsActiveAction
		out := d.UserSessionsActive(&domain.UserSessionsActiveIn{
			RequestCommon: in.RequestCommon,
		})

		return views.RenderUserProfileSessions(ctx, M.SX{
			`title`:          `Sessions`,
			`user`:           user,
			`segments`:       segments,
			`activeSessions`: out.SessionsActive,
		})
	})

	fw.Get(`/user/profileInvoices`, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notLogin(d, in.RequestCommon, false) {
			return ctx.Redirect(`/`, 302)
		}

		inv := rqInternal.NewInvoicePayment(d.IntrOltp)
		var invoices = []rqInternal.InvoicePayment{}
		invoices = inv.FindInvoicesByUserId(user.Id)

		return views.RenderUserProfileInvoices(ctx, M.SX{
			`title`:    `Invoices`,
			`user`:     user,
			`segments`: segments,
			`invoices`: invoices,
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
		if notReportViewer(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}
		return views.RenderReportViewerDashboard(ctx, M.SX{
			`title`:    `Report Viewer Dashboard`,
			`user`:     user,
			`segments`: segments,
		})
	})

	fw.Get(`/`+domain.ReportViewerGeneralLedgerAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notReportViewer(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		tenantCode, err := domain.GetTenantCodeByHost(in.Host)
		if err != nil {
			return ctx.Redirect(`/`, 302)
		}

		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.TenantCode = tenantCode
		coaChoices := coa.FindCoasChoicesByTenant()

		trxJournal := rqFinance.NewTransactionJournal(d.AuthOltp)
		trxJournal.TenantCode = tenantCode
		trxJournals := trxJournal.FindTrxJournalsByTenant()

		return views.RenderReportViewerGeneralLedger(ctx, M.SX{
			`title`:               `Report Viewer General Ledger`,
			`user`:                user,
			`segments`:            segments,
			`coaChoices`:          coaChoices,
			`transactionJournals`: trxJournals,
		})
	})

	fw.Get(`/`+domain.ReportViewerTrialBalanceAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notReportViewer(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		tenantCode, err := domain.GetTenantCodeByHost(in.Host)
		if err != nil {
			return ctx.Redirect(`/`, 302)
		}

		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.TenantCode = tenantCode
		coaChoices := coa.FindCoasChoicesByTenant()

		trxJournal := rqFinance.NewTransactionJournal(d.AuthOltp)
		trxJournal.TenantCode = tenantCode

		stTime := time.Now().AddDate(0, 0, -7)
		startDate := stTime.Format("2006-01-02")

		edTime := time.Now()
		endDate := edTime.Format("2006-01-02")

		trxJournals := trxJournal.FindTrxJournalsByDateByTenant(startDate, endDate)

		return views.RenderReportViewerTrialBalance(ctx, M.SX{
			`title`:               `Report Viewer Trial Balance`,
			`user`:                user,
			`segments`:            segments,
			`coaChoices`:          coaChoices,
			`transactionJournals`: trxJournals,
		})
	})

	fw.Get(`/`+domain.ReportViewerFinancialPositionAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notReportViewer(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		tenantCode, err := domain.GetTenantCodeByHost(in.Host)
		if err != nil {
			return ctx.Redirect(`/`, 302)
		}

		trxJournal := rqFinance.NewTransactionJournal(d.AuthOltp)
		trxJournal.TenantCode = tenantCode
		financialPosition := trxJournal.FindReportOfFinancialPositionByTenant()

		return views.RenderReportViewerFinancialPosition(ctx, M.SX{
			`title`:             `Report Viewer Financial Position`,
			`user`:              user,
			`segments`:          segments,
			`financialPosition`: financialPosition,
		})
	})

	fw.Get(`/`+domain.ReportViewerLossIncomeStatementsAction, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)
		if notReportViewer(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		tenantCode, err := domain.GetTenantCodeByHost(in.Host)
		if err != nil {
			return ctx.Redirect(`/`, 302)
		}

		trxJournal := rqFinance.NewTransactionJournal(d.AuthOltp)
		trxJournal.TenantCode = tenantCode
		trxJournals := trxJournal.FindTrxJournalsLossIncomeByTenant()

		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.TenantCode = tenantCode
		coaChoices := coa.FindCoasChoicesByTenant()

		return views.RenderReportViewerLossIncomeStatements(ctx, M.SX{
			`title`:               `Report Viewer Loss Income Statements`,
			`user`:                user,
			`segments`:            segments,
			`transactionJournals`: trxJournals,
			`coaChoices`:          coaChoices,
		})
	})

	fw.Get(`/`+domain.FieldSupervisorDashboardAction, func(ctx *fiber.Ctx) error {
		var in domain.FieldSupervisorDashboardIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminProductsAction)
		if err != nil {
			return err
		}

		if notFieldSupervisorLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)

		tenantCode, err := domain.GetTenantCodeByHost(in.Host)
		if err != nil {
			return ctx.Redirect(`/`, 302)
		}

		r := rqFinance.NewTransactionTemplate(d.AuthOltp)
		r.TenantCode = tenantCode
		trxTemplates := r.FindTransactionTamplatesChoicesByTenant()

		in.WithMeta = true
		in.Cmd = zCrud.CmdList
		out := d.FieldSupervisorDashboard(&in)

		if out.Error != `` {
			return ctx.Redirect(`/`, 302)
		}

		return views.RenderFieldSupervisorDashboard(ctx, M.SX{
			`title`:                `Field Supervisor Dashboard`,
			`user`:                 user,
			`segments`:             segments,
			`pager`:                out.Pager,
			`fields`:               out.Meta.Fields,
			`transaction`:          out.Transaction,
			`transactions`:         out.Transactions,
			`transactionTemplates`: trxTemplates,
		})
	})

	fw.Get(`/`+domain.FieldSupervisorBusinessTransactionEditAction+`:trxId`, func(ctx *fiber.Ctx) error {
		in, user, segments := userInfoFromContext(ctx, d)

		if notFieldSupervisorLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		tenantCode, err := domain.GetTenantCodeByHost(in.Host)
		if err != nil {
			return ctx.Redirect(`/`, 302)
		}

		businessTrxId := ctx.Params(`trxId`)
		transaction := rqFinance.NewBusinessTransaction(d.AuthOltp)
		transaction.Id = S.ToU(businessTrxId)
		transaction.TenantCode = tenantCode
		if !transaction.FindByIdByTenantCode() {
			return views.Render404(ctx, M.SX{
				`title`:       `Business Transaction Not Found`,
				`description`: `Make sure given id is valid`,
			})
		}

		trxJournal := rqFinance.NewTransactionJournal(d.AuthOltp)
		trxJournal.TransactionTemplateId = transaction.TransactionTemplateId
		trxJournal.TenantCode = tenantCode

		trxJournals := trxJournal.FindTrxJournalsByTrxTemplateByTenant()

		org := rqAuth.NewOrgs(d.AuthOltp)
		org.FindCompanyByTenantCode(tenantCode)

		return views.RenderFieldSupervisorBusinessTransactionEdit(ctx, M.SX{
			`title`:               `Field Supervisor Edit Business Transaction`,
			`user`:                user,
			`segments`:            segments,
			`transaction`:         transaction,
			`transactionJournals`: trxJournals,
			`org`:                 org,
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
			return ctx.Redirect(`/`, 302)
		}

		templateId := ctx.Params("templateId")

		trxTemplate := rqFinance.NewTransactionTemplate(d.AuthOltp)
		trxTemplate.Id = S.ToU(templateId)
		trxTemplate.TenantCode = tenantCode
		if !trxTemplate.FindByIdByTenantCode() {
			return ctx.Redirect(`/`+domain.DataEntryDashboardAction, 302)
		}

		trxTplDetail := rqFinance.NewTransactionTplDetail(d.AuthOltp)
		trxTplDetail.ParentId = trxTemplate.Id
		trxTplDetail.TenantCode = tenantCode
		trxTplDetails := trxTplDetail.FindTrxTplDetailsByTenantByTrxTplId()

		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.TenantCode = tenantCode
		coas := coa.FindCoasChoicesByTenant()

		org := rqAuth.NewOrgs(d.AuthOltp)
		org.FindCompanyByTenantCode(tenantCode)

		coasWithChildren := make(map[string]map[string]string)
		for _, v := range trxTplDetails {
			if v.ParseAttributes().IsChildOnly {
				coa.ParentId = v.CoaId
				coaChildren := coa.FindCoasChoicesChildByParentByTenant()
				coasWithChildren[I.ToS(int64(v.CoaId))] = coaChildren
			}
		}

		return views.RenderDataEntryTemplatesTemplate(ctx, M.SX{
			`title`:                 `Data Entry for ` + trxTemplate.Name,
			`user`:                  user,
			`segments`:              segments,
			`transactiontemplate`:   trxTemplate,
			`transactionTplDetails`: trxTplDetails,
			`coas`:                  coas,
			`org`:                   org,
			`coasWithChildren`:      coasWithChildren,
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

	fw.Get(`/`+domain.TenantAdminManualJournalAction, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminManualJournalIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminManualJournalAction)
		if err != nil {
			return err
		}

		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.WithMeta = true
		in.Cmd = zCrud.CmdList

		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.TenantCode = user.TenantCode
		coas := coa.FindCoasChoicesByTenant()

		org := rqAuth.NewOrgs(d.AuthOltp)
		org.FindCompanyByTenantCode(user.TenantCode)

		out := d.TenantAdminManualJournal(&in)
		return views.RenderTenantAdminManualJournal(ctx, M.SX{
			`title`:               `Tenant Admin Manual Journal`,
			`user`:                user,
			`segments`:            segments,
			`fields`:              out.Meta.Fields,
			`pager`:               out.Pager,
			`transactionJournals`: out.TransactionJournals,
			`coas`:                coas,
			`org`:                 org,
		})
	})

	fw.Get(`/`+domain.TenantAdminManualJournalAction+`/edit/:id`, func(ctx *fiber.Ctx) error {
		var in domain.TenantAdminManualJournalIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.TenantAdminManualJournalAction)
		if err != nil {
			return err
		}

		if notTenantLogin(d, in.RequestCommon) {
			return ctx.Redirect(`/`, 302)
		}

		journalId, err := ctx.ParamsInt("id", 0)
		if err != nil || journalId == 0 {
			return views.Render404(ctx, M.SX{
				`title`:       `Invalid Journal ID`,
				`description`: `Transaction journal ID must be number`,
			})
		}

		user, segments := userInfoFromRequest(in.RequestCommon, d)
		in.Cmd = zCrud.CmdForm
		in.TransactionJournal.Id = uint64(journalId)
		out := d.TenantAdminManualJournal(&in)

		if out.TransactionJournal == nil {
			return views.Render404(ctx, M.SX{
				`title`:       `Journal Not Found`,
				`description`: `Cannot found transaction journal for ID: "` + I.ToS(int64(journalId)),
			})
		}

		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.TenantCode = user.TenantCode
		coas := coa.FindCoasChoicesByTenant()

		org := rqAuth.NewOrgs(d.AuthOltp)
		org.FindCompanyByTenantCode(user.TenantCode)

		return views.RenderTenantAdminManualJournalEdit(ctx, M.SX{
			`title`:              `Tenant Admin Manual Journal Edit`,
			`user`:               user,
			`segments`:           segments,
			`org`:                org,
			`coas`:               coas,
			`transactionJournal`: out.TransactionJournal,
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

		tenant := rqAuth.NewTenants(d.AuthOltp)
		tenant.TenantCode = user.TenantCode
		tenant.FindByTenantCode()

		coaChoices := rqFinance.NewCoa(d.AuthOltp)
		coaChoices.TenantCode = tenant.TenantCode
		cChoices := coaChoices.FindCoasChoicesByTenant()
		return views.RenderTenantAdminCoa(ctx, M.SX{
			`title`:      `Tenant Admin Coa`,
			`user`:       user,
			`segments`:   segments,
			`coas`:       out.Coas,
			`tenant`:     tenant,
			`coaChoices`: cChoices,
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

		tenant := rqAuth.NewTenants(d.AuthOltp)
		tenant.TenantCode = user.TenantCode
		tenant.FindByTenantCode()

		coa := rqFinance.NewCoa(d.AuthOltp)
		coa.TenantCode = user.TenantCode
		coa.Id = tenant.BanksCoaId
		coas := coa.FindCoasChoicesChildByParentByTenant()

		L.Print(`coas:`, coas)

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
			`coas`:     coas,
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
			`products`:         products,
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

		r := rqFinance.NewCoa(d.AuthOltp)
		r.TenantCode = user.TenantCode
		coas := r.FindCoasChoicesByTenant()
		return views.RenderTenantAdminTransactionTemplate(ctx, M.SX{
			`title`:                `Tenant Admin Transaction Template`,
			`user`:                 user,
			`segments`:             segments,
			`transactionTemplates`: out.TransactionTemplates,
			`coas`:                 coas,
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
			`logs`:     out.Logs,
			`fields`:   out.Meta.Fields,
			`pager`:    out.Pager,
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

	fw.Get(`/`+domain.UserPaymentResultAction, func(ctx *fiber.Ctx) error {
		var in domain.UserPaymentResultIn
		err := webApiParseInput(ctx, &in.RequestCommon, &in, domain.UserPaymentResultAction)
		if err != nil {
			return err
		}
		if notLogin(d, in.RequestCommon, false) {
			return ctx.Redirect(`/`, 302)
		}
		user, segments := userInfoFromRequest(in.RequestCommon, d)

		if in.InvoiceNumber != `` {
			in.InvoiceNumber = `"` + in.InvoiceNumber + `"`
		}

		d.UserPaymentResult(&in)
		return views.RenderUserPaymentResult(ctx, M.SX{
			`title`:    `Payment Result ` + in.InvoiceNumber,
			`segments`: segments,
			`user`:     user,
		})
	})

	fw.Get(`/`+domain.UserPaymentCancelAction, func(ctx *fiber.Ctx) error {
		return views.RenderUserPaymentCancel(ctx, M.SX{
			`title`: `Payment Cancelled`,
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

func notReportViewer(d *domain.Domain, in domain.RequestCommon) bool {
	var check domain.ResponseCommon

	if sess := d.MustReportViewer(in, &check); sess == nil {
		return true
	}

	return false
}

func notFieldSupervisorLogin(d *domain.Domain, in domain.RequestCommon) bool {
	var check domain.ResponseCommon

	if sess := d.MustFieldSupervisor(in, &check); sess == nil {
		return true
	}

	return false
}
