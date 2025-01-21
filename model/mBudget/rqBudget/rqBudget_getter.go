package rqBudget

import (
	"benakun/model/zCrud"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

func (p *Plans) FindPlansByOrg(orgId uint64) (plans []Plans) {
	const comment = "-- plans) FindPlansByOrg"

	whereAndSql := ` WHERE ` + p.SqlOrgId() + ` = ` + I.UToS(orgId)
	queryRows := comment +
		`
SELECT ` + p.SqlSelectAllFields() + `
FROM /* /* SEQSCAN */CAN */ */ ` + p.SqlTableName() +
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
FROM /* /* SEQSCAN */CAN */ */ ` + p.SqlTableName() + whereAndSql

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
FROM /* /* SEQSCAN */CAN */ */ ` + p.SqlTableName() + whereAndSql

	p.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return len(res) > 0
}

func (b *BankAccounts) FindByPagination(z *zCrud.Meta, z2 *zCrud.PagerIn, z3 *zCrud.PagerOut) (res [][]any) {
	const comment = `-- BankAccounts) FindByPagination`

	validFields := BankAccountsFieldTypeMap
	whereAndSql := z3.WhereAndSqlTt(z2.Filters, validFields)
	whereAndSql2 := `
AND ` + b.SqlTenantCode() + ` = ` + S.Z(b.TenantCode)
	if whereAndSql == `` {
		whereAndSql2 = ` 
WHERE ` + b.SqlTenantCode() + ` = ` + S.Z(b.TenantCode)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM /* /* SEQSCAN */CAN */ */ ` + b.SqlTableName() + whereAndSql + whereAndSql2 + ` 
LIMIT 1`
	b.Adapter.QuerySql(queryCount, func(row []any) {
		z3.CalculatePages(z2.Page, z2.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := z3.OrderBySqlTt(z2.Order, validFields)
	limitOffsetSql := z3.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + z.ToSelect() + `
FROM /* /* SEQSCAN */CAN */ */ ` + b.SqlTableName() + whereAndSql + whereAndSql2 + orderBySql + limitOffsetSql
	b.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}
