package domain

import (
	"benakun/model/mAuth/rqAuth"

	"github.com/kokizzu/gotro/X"
)

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
	orgs := rqAuth.NewOrgs(d.AuthOltp)
	tenantsHost := orgs.FindTenantsHost()

	if len(tenantsHost) > 0 {
		for _, v := range tenantsHost {
			subdomain := `https://`+ X.ToS(v[0]) +`.benakun.com`
			hostmap[subdomain] = TenantCodeId{
				TenantCode: X.ToS(v[0]),
				OrgId: X.ToU(v[1]),
			}
		}
	}
}
