package domain

// TODO: add mapping to database (translates domain to proper tenant)

type TenantCodeId struct {
	TenantCode string
	OrgId      uint64
}

var hostmap = map[string]TenantCodeId{
	// Default hostmap for development
	// change it with the actual data from your database
	`http://admin-2642:1235`: {
		TenantCode: `admin-2642`,
		OrgId: 10,
	},
}

func (d *Domain) InitHostMapper() {
	// TODO query table tenants and orgs
	// left join, where tenants.tenantCode = orgs.tenantCode AND orgs.orgType = 1 (company)
	// then fill it into array of struct []TenantCodeId{}
}
