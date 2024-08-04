package presentation

import (
	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"
)


// Code generated by 1_codegen_test.go DO NOT EDIT.


var viewList = map[string]string{
	`404`: `../svelte/404.html`, // ../svelte/404.svelte
	`ApidocsIndex`: `../svelte/apidocs/index.html`, // ../svelte/apidocs/index.svelte
	`DataEntryTemplatesTemplate`: `../svelte/dataEntry/templates/template.html`, // ../svelte/dataEntry/templates/template.svelte
	`DataEntryDashboard`: `../svelte/dataEntry_dashboard.html`, // ../svelte/dataEntry_dashboard.svelte
	`DataEntryTransactionEntry`: `../svelte/dataEntry_transactionEntry.html`, // ../svelte/dataEntry_transactionEntry.svelte
	`Debug`: `../svelte/debug.html`, // ../svelte/debug.svelte
	`GuestOauthCallback`: `../svelte/guest_oauthCallback.html`, // ../svelte/guest_oauthCallback.svelte
	`GuestResetPassword`: `../svelte/guest_resetPassword.html`, // ../svelte/guest_resetPassword.svelte
	`GuestVerifyEmail`: `../svelte/guest_verifyEmail.html`, // ../svelte/guest_verifyEmail.svelte
	`Index`: `../svelte/index.html`, // ../svelte/index.svelte
	`ReportViewerDashboard`: `../svelte/reportViewer_dashboard.html`, // ../svelte/reportViewer_dashboard.svelte
	`SuperAdminAccessLog`: `../svelte/superAdmin_accessLog.html`, // ../svelte/superAdmin_accessLog.svelte
	`SuperAdminDashboard`: `../svelte/superAdmin_dashboard.html`, // ../svelte/superAdmin_dashboard.svelte
	`SuperAdminTenantManagement`: `../svelte/superAdmin_tenantManagement.html`, // ../svelte/superAdmin_tenantManagement.svelte
	`SuperAdminUserManagement`: `../svelte/superAdmin_userManagement.html`, // ../svelte/superAdmin_userManagement.svelte
	`TenantAdminInventoryChangesProduct`: `../svelte/tenantAdmin/inventoryChanges/product.html`, // ../svelte/tenantAdmin/inventoryChanges/product.svelte
	`TenantAdminBankAccounts`: `../svelte/tenantAdmin_bankAccounts.html`, // ../svelte/tenantAdmin_bankAccounts.svelte
	`TenantAdminBudgeting`: `../svelte/tenantAdmin_budgeting.html`, // ../svelte/tenantAdmin_budgeting.svelte
	`TenantAdminCoa`: `../svelte/tenantAdmin_coa.html`, // ../svelte/tenantAdmin_coa.svelte
	`TenantAdminDashboard`: `../svelte/tenantAdmin_dashboard.html`, // ../svelte/tenantAdmin_dashboard.svelte
	`TenantAdminInventoryChanges`: `../svelte/tenantAdmin_inventoryChanges.html`, // ../svelte/tenantAdmin_inventoryChanges.svelte
	`TenantAdminLocations`: `../svelte/tenantAdmin_locations.html`, // ../svelte/tenantAdmin_locations.svelte
	`TenantAdminManualJournal`: `../svelte/tenantAdmin_manualJournal.html`, // ../svelte/tenantAdmin_manualJournal.svelte
	`TenantAdminOrganization`: `../svelte/tenantAdmin_organization.html`, // ../svelte/tenantAdmin_organization.svelte
	`TenantAdminProducts`: `../svelte/tenantAdmin_products.html`, // ../svelte/tenantAdmin_products.svelte
	`TenantAdminTransaction`: `../svelte/tenantAdmin_transaction.html`, // ../svelte/tenantAdmin_transaction.svelte
	`TenantAdminTransactionTemplate`: `../svelte/tenantAdmin_transactionTemplate.html`, // ../svelte/tenantAdmin_transactionTemplate.svelte
	`UserProfile`: `../svelte/user_profile.html`, // ../svelte/user_profile.svelte
	`UserResponsejoin`: `../svelte/user_responsejoin.html`, // ../svelte/user_responsejoin.svelte
}


func (v *Views) Render404(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`404`].Str(m))
}

func (v *Views) RenderApidocsIndex(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`ApidocsIndex`].Str(m))
}

func (v *Views) RenderDataEntryTemplatesTemplate(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`DataEntryTemplatesTemplate`].Str(m))
}

func (v *Views) RenderDataEntryDashboard(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`DataEntryDashboard`].Str(m))
}

func (v *Views) RenderDataEntryTransactionEntry(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`DataEntryTransactionEntry`].Str(m))
}

func (v *Views) RenderDebug(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Debug`].Str(m))
}

func (v *Views) RenderGuestOauthCallback(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`GuestOauthCallback`].Str(m))
}

func (v *Views) RenderGuestResetPassword(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`GuestResetPassword`].Str(m))
}

func (v *Views) RenderGuestVerifyEmail(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`GuestVerifyEmail`].Str(m))
}

func (v *Views) RenderIndex(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`Index`].Str(m))
}

func (v *Views) RenderReportViewerDashboard(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`ReportViewerDashboard`].Str(m))
}

func (v *Views) RenderSuperAdminAccessLog(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`SuperAdminAccessLog`].Str(m))
}

func (v *Views) RenderSuperAdminDashboard(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`SuperAdminDashboard`].Str(m))
}

func (v *Views) RenderSuperAdminTenantManagement(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`SuperAdminTenantManagement`].Str(m))
}

func (v *Views) RenderSuperAdminUserManagement(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`SuperAdminUserManagement`].Str(m))
}

func (v *Views) RenderTenantAdminInventoryChangesProduct(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminInventoryChangesProduct`].Str(m))
}

func (v *Views) RenderTenantAdminBankAccounts(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminBankAccounts`].Str(m))
}

func (v *Views) RenderTenantAdminBudgeting(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminBudgeting`].Str(m))
}

func (v *Views) RenderTenantAdminCoa(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminCoa`].Str(m))
}

func (v *Views) RenderTenantAdminDashboard(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminDashboard`].Str(m))
}

func (v *Views) RenderTenantAdminInventoryChanges(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminInventoryChanges`].Str(m))
}

func (v *Views) RenderTenantAdminLocations(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminLocations`].Str(m))
}

func (v *Views) RenderTenantAdminManualJournal(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminManualJournal`].Str(m))
}

func (v *Views) RenderTenantAdminOrganization(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminOrganization`].Str(m))
}

func (v *Views) RenderTenantAdminProducts(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminProducts`].Str(m))
}

func (v *Views) RenderTenantAdminTransaction(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminTransaction`].Str(m))
}

func (v *Views) RenderTenantAdminTransactionTemplate(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`TenantAdminTransactionTemplate`].Str(m))
}

func (v *Views) RenderUserProfile(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`UserProfile`].Str(m))
}

func (v *Views) RenderUserResponsejoin(c *fiber.Ctx, m M.SX) error {
	c.Set("Content-Type", "text/html; charset=utf-8")
	return c.SendString(v.cache[`UserResponsejoin`].Str(m))
}

// Code generated by 1_codegen_test.go DO NOT EDIT.
