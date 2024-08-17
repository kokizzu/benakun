package rqFinance

import (
	"benakun/model/zCrud"
	"errors"
	"fmt"

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

type (
	coaWithNum struct {
		id 				uint64
		name			string
		parentId	uint64
		children	[]any
	}
	coaWithNumF struct {
		id 				uint64
		name			string
		parentId	uint64
		children	[]coaWithNumF
	}
)

func removeUnParent(cs []coaWithNumF, cond func(coaWithNumF) bool) []coaWithNumF {
	var filtered []coaWithNumF
	for _, c := range cs {
		if !cond(c) {
			filtered = append(filtered, c)
		}
	}

	return filtered
}

func (c *Coa) FindCoasChoicesByTenant() map[string]string {
	const comment = `-- Coa) FindCoasChoicesByTenant`

	whereAndSql := ` WHERE ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode)

	queryRows := comment + `
SELECT ` + c.SqlId() + `, ` + c.SqlName() + `, ` + c.SqlParentId() + `, ` + c.SqlChildren() + `
FROM SEQSCAN ` + c.SqlTableName() + whereAndSql

	coasWithNums := []coaWithNum{}

	c.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 4 {
			coasWithNums = append(coasWithNums, coaWithNum{
				id: X.ToU(row[0]),
				name: X.ToS(row[1]),
				parentId: X.ToU(row[2]),
				children: X.ToArr(row[3]),
			})
		}
	})

	coasWithNumsF := []coaWithNumF{}
	coaVisited := map[int]bool{0: true}
	var coaMaker func(id uint64) (coaWithNumF, error)
	coaMaker = func(id uint64) (cwnF coaWithNumF, err error) {
		coaError := errors.New(`coa is already visited`)

		if _, exist := coaVisited[int(id)]; exist {
			err = coaError
			return
		}

		coaVisited = map[int]bool{int(id): true}
		if len(coasWithNums) > 0 {
			for _, v := range coasWithNums {
				if v.id == id {
					cld := v.children
					if len(cld) > 0 {
						for _, cid := range cld {
							child, err := coaMaker(X.ToU(cid))
							if err == nil {
								cwnF.children = append(cwnF.children, child)
							} else {
								continue
							}
						}
					}

					cwnF.id = v.id
					cwnF.name = v.name
					cwnF.parentId = v.parentId

					return
				}
			}
		}

		return
	}

	for _, v := range coasWithNums {
		coa, err := coaMaker(v.id)
		if err == nil {
			coasWithNumsF = append(coasWithNumsF, coa)
		}
	}

	coasWithNumsF = removeUnParent(coasWithNumsF, func(cwnf coaWithNumF) bool {
		return cwnf.parentId != 0
	})

	coaChoices := make(map[string]string)

	var coaTree func(c coaWithNumF, num string, parentNum int64, idx int)
	coaTree = func(c coaWithNumF, num string, parentNum int64, idx int) {
		numStr := I.ToS(parentNum)
		if c.parentId != 0 {
			numStr = num
		}

		name := numStr + ` ` + c.name
		coaChoices[I.UToS(c.id)] = name
		for ix, v := range c.children {
			snum := fmt.Sprintf("%v.%v", num, ix+1)
			coaTree(v, snum, parentNum, ix)
		}

		_ = idx
	}
	
	for i, v := range coasWithNumsF {
		coaTree(v, I.ToS(int64(i)+1), int64(i)+1, i)
	}

	return coaChoices
}

func (c *Coa) FindCoasChoicesChildByParentByTenant() map[string]string {
	const comment = `-- Coa) FindCoasChoicesChildByParentByTenant`

	whereAndSql := ` WHERE ` + c.SqlParentId() + ` = ` + I.UToS(c.Id) + `
		AND `  + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode)

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

func (ttpl *TransactionTemplate) FindTransactionTamplatesChoicesByTenant() map[string]string {
	const comment = `-- TransactionTemplate) FindTransactionTamplatesChoicesByTenant`

	whereAndSql := ` WHERE ` + ttpl.SqlTenantCode() + ` = ` + S.Z(ttpl.TenantCode)

	queryRows := comment + `
SELECT ` + ttpl.SqlId() + `, ` + ttpl.SqlName() + `
FROM SEQSCAN ` + ttpl.SqlTableName() + whereAndSql

	ttplsChoices := make(map[string]string)
	ttpl.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			ttplsChoices[X.ToS(row[0])] = X.ToS(row[1])
		}
	})

	return ttplsChoices
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