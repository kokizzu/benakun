package domain

// TODO: add mapping to database (translates domain to proper tenant)

type TenantCodeId struct {
	TenantCode string
	OrgId      uint64
}

var hostmap = map[string]TenantCodeId{
	`http://127.0.0.1:1235`: {
		TenantCode: `benalu-6405`,
		OrgId:      6727,
	},
	`https://local1:1235`: {
		TenantCode: `ertwywret-2504`,
		OrgId: 1,
	},
}
