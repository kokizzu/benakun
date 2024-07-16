package rqFinance

import (
	"benakun/model/zCrud"

	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

func (c *Coa) FindCoasByTenant() (coas []Coa) {
	var res [][]any
	const comment = `-- Coa) FindCoasByTenant`

	whereAndSql := ` WHERE ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode)

	queryRows := comment + `
SELECT ` + c.SqlSelectAllFields() + `
FROM SEQSCAN ` + c.SqlTableName() + whereAndSql

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
	}

	return
}

func (ttm *TransactionTemplate) FindByPagination(z *zCrud.Meta, z2 *zCrud.PagerIn, z3 *zCrud.PagerOut) (res [][]any) {
	const comment = `-- TransactionTemplate) FindByPagination`

	validFields := TransactionTemplateFieldTypeMap
	whereAndSql := z3.WhereAndSqlTt(z2.Filters, validFields)
	whereAndSql2 := `AND ` + ttm.SqlTenantCode() + ` = ` + S.Z(ttm.TenantCode)
	if whereAndSql == `` {
		whereAndSql2 = ` WHERE ` + ttm.SqlTenantCode() + ` = ` + S.Z(ttm.TenantCode)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM SEQSCAN ` + ttm.SqlTableName() + whereAndSql + whereAndSql2 + `
LIMIT 1`
	ttm.Adapter.QuerySql(queryCount, func(row []any) {
		z3.CalculatePages(z2.Page, z2.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := z3.OrderBySqlTt(z2.Order, validFields)
	limitOffsetSql := z3.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + z.ToSelect() + `
FROM SEQSCAN ` + ttm.SqlTableName() + whereAndSql + whereAndSql2 + orderBySql + limitOffsetSql
	ttm.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}

func (ttm *TransactionTemplate) FindByTenantCode() (ttms *[]TransactionTemplate) {
	const comment = `-- TransactionTemplate) FindByTenantCode`

	whereAndSql := ` WHERE ` + ttm.SqlTenantCode() + ` = ` + S.Z(ttm.TenantCode)

	queryRows := comment + `
SELECT ` + ttm.SqlSelectAllFields() + `
FROM SEQSCAN ` + ttm.SqlTableName() + whereAndSql

	var rows []TransactionTemplate

	ttm.Adapter.QuerySql(queryRows, func(row []any) {
		ttm.FromArray(row)
		rows = append(rows, *ttm)
	})

	ttms = &rows
	return
}

func (tt *TransactionTemplate) FindTransactionTemplatesByTenant() (ttpls []TransactionTemplate) {
	var res [][]any
	const comment = `-- TransactionTemplate) FindTransactionTemplatesByTenant`

	whereAndSql := ` WHERE ` + tt.SqlTenantCode() + ` = ` + S.Z(tt.TenantCode)

	queryRows := comment + `
SELECT ` + tt.SqlSelectAllFields() + `
FROM SEQSCAN ` + tt.SqlTableName() + whereAndSql

	tt.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	L.Print(`query:`, queryRows)

	if len(res) > 0 {
		for _, v := range res {
			if len(v) >= 5 {
				ttpls = append(ttpls, *tt.FromArray(v))
			}
		}
	}

	return
}
