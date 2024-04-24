package rqBudget

import (
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/X"
)

func (p *Plans) FindPlansByOrg(orgId uint64) (plans []Plans) {
	const comment = "-- plans) FindPlansByOrg"

	whereAndSql := ` WHERE ` + p.SqlOrgId() + ` = '` + I.UToS(orgId) + `'`
	queryRows := comment +
		`
SELECT ` + p.SqlSelectAllFields() + `
FROM ` + p.SqlTableName() +
		whereAndSql

	p.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0])
		plans = append(plans, *p.FromArray(row))
	})

	return
}