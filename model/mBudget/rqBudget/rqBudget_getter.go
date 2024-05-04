package rqBudget

import (
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/X"
)

func (p *Plans) FindPlansByOrg(orgId uint64) (plans []Plans) {
	const comment = "-- plans) FindPlansByOrg"

	whereAndSql := ` WHERE ` + p.SqlOrgId() + ` = ` + I.UToS(orgId)
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

func (p *Plans) IsVisionExist() bool {
	var res []any
	const comment = `-- Plans) IsVisionExist`

	whereAndSql := ` WHERE ` + p.SqlPlanType() + ` = 'vision' AND ` + p.SqlOrgId() + ` = ` + I.UToS(p.OrgId)

	queryRows := comment + `
SELECT ` + p.SqlPlanType() + `
FROM ` + p.SqlTableName() + whereAndSql

	p.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return len(res) > 0
}

func (p *Plans) IsMissionExist() bool {
	var res []any
	const comment = `-- Plans) IsMissionExist`

	whereAndSql := ` WHERE ` + p.SqlPlanType() + ` = 'mission' AND ` + p.SqlOrgId() + ` = ` + I.UToS(p.OrgId)

	queryRows := comment + `
SELECT ` + p.SqlPlanType() + `
FROM ` + p.SqlTableName() + whereAndSql

	p.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return len(res) > 0
}