package rqBusiness

import (
	"benakun/model/zCrud"

	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

func (pr *Products) FindByPagination(z *zCrud.Meta, z2 *zCrud.PagerIn, z3 *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Products) FindByPagination`

	validFields := ProductsFieldTypeMap
	whereAndSql := z3.WhereAndSqlTt(z2.Filters, validFields)
	whereAndSql2 := `AND ` + pr.SqlTenantCode() + ` = ` + S.Z(pr.TenantCode)
	if whereAndSql == `` {
		whereAndSql2 = ` WHERE ` + pr.SqlTenantCode() + ` = ` + S.Z(pr.TenantCode)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM SEQSCAN ` + pr.SqlTableName() + whereAndSql + whereAndSql2 + `
LIMIT 1`
	pr.Adapter.QuerySql(queryCount, func(row []any) {
		z3.CalculatePages(z2.Page, z2.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := z3.OrderBySqlTt(z2.Order, validFields)
	limitOffsetSql := z3.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + z.ToSelect() + `
FROM SEQSCAN ` + pr.SqlTableName() + whereAndSql + whereAndSql2 + orderBySql + limitOffsetSql
	pr.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}
