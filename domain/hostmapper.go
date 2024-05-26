package domain

// TODO: add mapping to database (translates domain to proper tenant)

type TenantCodeId struct {
	TenantCode string
	OrgId      uint64
}

var hostmap = map[string]TenantCodeId{
	`http://127.0.0.1:1235`: {
		TenantCode: `test_6727`,
		OrgId:      6727,
	},
}
