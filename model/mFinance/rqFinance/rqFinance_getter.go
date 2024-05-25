package rqFinance

import (
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

func (c *Coa) FindCoasByTenant(tenantCode string) (coas []Coa) {
	var res [][]any
	const comment = `-- Coa) FindCoasByTenant`

	whereAndSql := ` WHERE ` + c.SqlTenantCode() + ` = ` + S.Z(tenantCode)

	queryRows := comment + `
SELECT ` + c.SqlSelectAllFields() + `
FROM ` + c.SqlTableName() + whereAndSql

	c.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	if len(res) > 0 {
		for _, oa := range res {
			if len(oa) >= 5 {
				coas = append(coas, *c.FromArray(oa))
			}
		}
	} else {
		return []Coa{}
	}

	return
}

func (c *Coa) FindCoaIdByTenantByLevel() uint64 {
	var res [][]any
	const comment = `-- Coa) FindCoaIdByTenantByLevel`

	whereAndSql := ` WHERE ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode) + ` AND ` + c.SqlLevel() + ` = ` + I.ToS(int64(c.Level))  + ` AND "deletedAt" = 0`

	queryRow := comment + `
SELECT ` + c.SqlId() + `
FROM ` + c.SqlTableName() + whereAndSql

	c.Adapter.QuerySql(queryRow, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	if len(res) > 0 {
		pId := res[0][0].(string)
		return S.ToU(pId)
	}

	return 0
}