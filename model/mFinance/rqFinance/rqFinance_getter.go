package rqFinance

import (
	"benakun/model/zCrud"

	"github.com/kokizzu/gotro/I"
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

func (c *Coa) FindCoasChoicesByTenant() map[string]string {
	const comment = `-- Coa) FindCoasChoicesByTenant`

	whereAndSql := ` WHERE ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode)

	queryRows := comment + `
SELECT ` + c.SqlId() + `, ` + c.SqlName() + `
FROM SEQSCAN ` + c.SqlTableName() + whereAndSql

	coaChoices := make(map[string]string)
	c.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			coaChoices[X.ToS(row[0])] = X.ToS(row[1])
		}
	})

	return coaChoices
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

func (ttm *TransactionTemplate) FindByIdByTenantCode() bool {
	const comment = `-- TransactionTemplate) FindByTenantCode`

	whereAndSql := ` WHERE ` + ttm.SqlTenantCode() + ` = ` + S.Z(ttm.TenantCode) + `
		AND ` + ttm.SqlId() + ` = ` + I.UToS(ttm.Id)

	queryRows := comment + `
SELECT ` + ttm.SqlSelectAllFields() + `
FROM SEQSCAN ` + ttm.SqlTableName() + whereAndSql

	var res []any
	ttm.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToU(row[0])
		res = row
	})

	if len(res) > 0 {
		ttm.FromArray(res)
		return true
	}

	return false
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

	if len(res) > 0 {
		for _, v := range res {
			if len(v) >= 5 {
				ttpls = append(ttpls, *tt.FromArray(v))
			}
		}
	}

	return
}

func (ttpld *TransactionTplDetail) FindTrxTplDetailsByTenantByTrxTplId() (ttplds []TransactionTplDetail) {
	var res [][]any
	const comment = `-- TransactionTplDetail) FindTrxTplDetailsByTenantByTrxTplId`

	whereAndSql := ` WHERE ` + ttpld.SqlTenantCode() + ` = ` + S.Z(ttpld.TenantCode) + `
		AND ` + ttpld.SqlParentId() + ` = ` + I.UToS(ttpld.ParentId)

	queryRows := comment + `
SELECT ` + ttpld.SqlSelectAllFields() + `
FROM SEQSCAN ` + ttpld.SqlTableName() + whereAndSql

	ttpld.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	if len(res) > 0 {
		for _, v := range res {
			if len(v) >= 5 {
				ttplds = append(ttplds, *ttpld.FromArray(v))
			}
		}
	}

	return
}

func (tj *TransactionJournal) FindByPagination(z *zCrud.Meta, z2 *zCrud.PagerIn, z3 *zCrud.PagerOut) (res [][]any) {
	const comment = `-- TransactionJournal) FindByPagination`

	validFields := TransactionJournalFieldTypeMap
	whereAndSql := z3.WhereAndSqlTt(z2.Filters, validFields)
	whereAndSql2 := `AND ` + tj.SqlTenantCode() + ` = ` + S.Z(tj.TenantCode)
	if whereAndSql == `` {
		whereAndSql2 = ` WHERE ` + tj.SqlTenantCode() + ` = ` + S.Z(tj.TenantCode)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM SEQSCAN ` + tj.SqlTableName() + whereAndSql + whereAndSql2 + `
LIMIT 1`
	tj.Adapter.QuerySql(queryCount, func(row []any) {
		z3.CalculatePages(z2.Page, z2.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := z3.OrderBySqlTt(z2.Order, validFields)
	limitOffsetSql := z3.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + z.ToSelect() + `
FROM SEQSCAN ` + tj.SqlTableName() + whereAndSql + whereAndSql2 + orderBySql + limitOffsetSql
	tj.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}