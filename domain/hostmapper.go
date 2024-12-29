package domain

import (
	"errors"
	"regexp"

	"benakun/model/mAuth/rqAuth"

	"github.com/kokizzu/gotro/X"
)

type TenantHost struct {
	TenantCode string
	OrgId      uint64
}

var hostmap = map[string]TenantHost{
	// Default hostmap for development
	// change it with the actual data from your database
	`http://admin-2642:1235`: {
		TenantCode: `admin-2642`,
		OrgId:      10,
	},
}

var RgxFindTenantCode = regexp.MustCompile(`://([^:./]+)`)

func (d *Domain) InitHostMapper() {

	orgs := rqAuth.NewOrgs(d.AuthOltp)
	tenantsHost := orgs.FindTenantsHost()

	if len(tenantsHost) > 0 {
		for _, v := range tenantsHost {
			subdomain := generateTenantSubdomain(X.ToS(v[0]))
			hostmap[subdomain] = TenantHost{
				TenantCode: X.ToS(v[0]),
				OrgId:      X.ToU(v[1]),
			}
		}
	}
}

func GetTenantCodeByHost(hostname string) (string, error) {
	// Development	: http://admin-2642:1235
	// Production		: https://admin-2642.benakun.com
	// find tenantCode from url
	match := RgxFindTenantCode.FindStringSubmatch(hostname)
	if len(match) > 1 {
		return match[1], nil
	} else {
		return ``, errors.New(`invalid hostname`)
	}
}

func generateTenantSubdomain(tenantCode string) string {
	return `https://` + tenantCode + `.benakun.com`
}
