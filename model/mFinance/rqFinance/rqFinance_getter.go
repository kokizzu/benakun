package rqFinance

import (
	"benakun/model/mFinance"
	"benakun/model/zCrud"
	"fmt"
	"sort"
	"strings"

	"github.com/kokizzu/gotro/I"
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

type ParsedAttribute struct {
	IsAutoSum   bool
	IsSales     bool
	IsChildOnly bool
	IsSelfSum   bool
}

func (ttpld *TransactionTplDetail) ParseAttributes() (pa ParsedAttribute) {
	for _, v := range ttpld.Attributes {
		if v == mFinance.AttributesAutoSum {
			pa.IsAutoSum = true
		}
		if v == mFinance.AttributesSales {
			pa.IsSales = true
		}
		if v == mFinance.AttributesChildOnly {
			pa.IsChildOnly = true
		}
		if v == mFinance.AttributeSelfSum {
			pa.IsSelfSum = true
		}
	}

	return
}

type (
	// obtain from db
	coaFromDB struct {
		id       uint64
		name     string
		parentId uint64
		children []any
	}
	// Convert coawithnum to Tree
	coaTreeNode struct {
		id       uint64
		name     string
		parentId uint64
		children []coaTreeNode
	}
)

func removeUnParent(cs []coaTreeNode, callback func(coaTreeNode) bool) (filtered []coaTreeNode) {
	for _, c := range cs {
		if callback(c) {
			filtered = append(filtered, c)
		}
	}

	return
}

func generateCoaChoicesMaps(coasFromDBs []coaFromDB) map[string]string {
	coaChoices := make(map[string]string, len(coasFromDBs))

	if len(coasFromDBs) <= 0 {
		return coaChoices
	}

	coasTreeNodes := []coaTreeNode{}
	coaVisited := map[int]bool{0: true}

	var coaMaker func(id uint64) (coaTreeNode, bool)
	coaMaker = func(id uint64) (ctn coaTreeNode, isVisited bool) {
		if coaVisited[int(id)] {
			isVisited = true
			return
		}

		coaVisited[int(id)] = true

		for _, v := range coasFromDBs {
			if v.id == id {
				cld := v.children
				if len(cld) > 0 {
					for _, cid := range cld {
						child, visited := coaMaker(X.ToU(cid))
						if visited {
							continue
						} else {
							ctn.children = append(ctn.children, child)
						}
					}
				}
				ctn.id = v.id
				ctn.name = v.name
				ctn.parentId = v.parentId
				return
			}
		}

		return
	}

	for _, v := range coasFromDBs {
		coa, visited := coaMaker(v.id)
		if !visited {
			coasTreeNodes = append(coasTreeNodes, coa)
		}
	}

	coasTreeNodes = removeUnParent(coasTreeNodes, func(cwnf coaTreeNode) bool {
		return cwnf.parentId == 0
	})

	var coaTree func(c coaTreeNode, num string, parentNum int64, idx int)
	coaTree = func(c coaTreeNode, num string, parentNum int64, idx int) {
		if c.id != 0 {
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
	}

	for i, v := range coasTreeNodes {
		coaTree(v, I.ToS(int64(i)+1), int64(i)+1, i)
	}

	return coaChoices
}

func extractNumParts(s string) []int {
	parts := strings.Fields(s)
	numberStr := parts[0]
	numberParts := S.Split(numberStr, `.`)

	var result []int
	for _, part := range numberParts {
		num := S.ToInt(part)
		result = append(result, num)
	}

	return result
}

func compareCoaNums(va, vb []int) bool {
	for i := 0; i < len(va) && i < len(vb); i++ {
		if va[i] != vb[i] {
			return va[i] < vb[i]
		}
	}

	return len(va) < len(vb)
}

func (c *Coa) FindCoasChoicesByTenant() map[string]string {
	const comment = `-- Coa) FindCoasChoicesByTenant`

	whereAndSql := ` WHERE ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode)

	queryRows := comment + `
SELECT ` + c.SqlId() + `, ` + c.SqlName() + `, ` + c.SqlParentId() + `, ` + c.SqlChildren() + `
FROM SEQSCAN ` + c.SqlTableName() + whereAndSql

	coasFromDBs := []coaFromDB{}

	c.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 4 {
			coasFromDBs = append(coasFromDBs, coaFromDB{
				id:       X.ToU(row[0]),
				name:     X.ToS(row[1]),
				parentId: X.ToU(row[2]),
				children: X.ToArr(row[3]),
			})
		}
	})

	coaChoices := generateCoaChoicesMaps(coasFromDBs)
	var keys []string = make([]string, 0, len(coaChoices))
	for key := range coaChoices {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		numsI := extractNumParts(coaChoices[keys[i]])
		numsJ := extractNumParts(coaChoices[keys[j]])

		return compareCoaNums(numsI, numsJ)
	})

	newCoaChoices := make(map[string]string)
	for _, k := range keys {
		newCoaChoices[k] = coaChoices[k]
	}

	return newCoaChoices
}

func (c *Coa) FindCoasChoicesChildByParentByTenant() map[string]string {
	const comment = `-- Coa) FindCoasChoicesChildByParentByTenant`

	whereAndSql := ` WHERE ` + c.SqlParentId() + ` = ` + I.UToS(c.ParentId) + `
		AND ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode)

	queryRows := comment + `
	SELECT ` + c.SqlId() + `, ` + c.SqlName() + `
	FROM SEQSCAN ` + c.SqlTableName() + whereAndSql

	coaChoices := make(map[string]string)
	var idx int64 = 1
	c.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			coaChoices[X.ToS(row[0])] = I.ToS(idx) + `. ` + X.ToS(row[1])
			idx++
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

func (ttm *TransactionTemplate) FindByNameByTenantCode() bool {
	const comment = `-- TransactionTemplate) FindByTenantCode`

	whereAndSql := ` WHERE ` + ttm.SqlTenantCode() + ` = ` + S.Z(ttm.TenantCode) + `
		AND ` + ttm.SqlName() + ` = ` + S.Z(ttm.Name)

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

func (tj *TransactionJournal) FindByIdByTenantCode() bool {
	const comment = `-- TransactionJournal) FindByTenantCode`

	whereAndSql := ` WHERE ` + tj.SqlTenantCode() + ` = ` + S.Z(tj.TenantCode) + `
		AND ` + tj.SqlId() + ` = ` + I.UToS(tj.Id)

	queryRows := comment + `
SELECT ` + tj.SqlSelectAllFields() + `
FROM SEQSCAN ` + tj.SqlTableName() + whereAndSql

	var res []any
	tj.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToU(row[0])
		res = row
	})

	if len(res) > 0 {
		tj.FromArray(res)
		return true
	}

	return false
}

func (tj *TransactionJournal) FindTrxJournalsByTenant() (trxJournals []TransactionJournal) {
	var res [][]any
	const comment = `-- TransactionJournal) FindTrxJournalsByTenant`

	whereAndSql := ` WHERE ` + tj.SqlTenantCode() + ` = ` + S.Z(tj.TenantCode)

	queryRows := comment + `
SELECT ` + tj.SqlSelectAllFields() + `
FROM SEQSCAN ` + tj.SqlTableName() + whereAndSql

	tj.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	if len(res) > 0 {
		for _, oa := range res {
			if len(oa) >= 5 {
				trxJournals = append(trxJournals, *tj.FromArray(oa))
			}
		}
	}

	return
}

func (tj *TransactionJournal) FindTrxJournalsByDateByTenant(startDate, endDate string) (trxJournals []TransactionJournal) {
	var res [][]any
	const comment = `-- TransactionJournal) FindTrxJournalsByDateByTenant`

	whereAndSql := ` WHERE ` + tj.SqlTenantCode() + ` = ` + S.Z(tj.TenantCode) + `
		AND ` + tj.SqlDate() + ` >= ` + S.Z(startDate) + `
		AND ` + tj.SqlDate() + ` <= ` + S.Z(endDate)

	queryRows := comment + `
SELECT ` + tj.SqlSelectAllFields() + `
FROM SEQSCAN ` + tj.SqlTableName() + whereAndSql

	L.Print(`Query:`, queryRows)

	tj.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	if len(res) > 0 {
		for _, oa := range res {
			if len(oa) >= 5 {
				trxJournals = append(trxJournals, *tj.FromArray(oa))
			}
		}
	}

	return
}

func (tj *TransactionJournal) FindTrxJournalsByCoaByTenant() (trxJournals []TransactionJournal) {
	var res [][]any
	const comment = `-- TransactionJournal) FindTrxJournalsByCoaByTenant`

	whereAndSql := ` WHERE ` + tj.SqlTenantCode() + ` = ` + S.Z(tj.TenantCode) + `
		AND ` + tj.SqlCoaId() + ` = ` + I.UToS(tj.CoaId)

	queryRows := comment + `
SELECT ` + tj.SqlSelectAllFields() + `
FROM SEQSCAN ` + tj.SqlTableName() + whereAndSql

	tj.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	if len(res) > 0 {
		for _, oa := range res {
			if len(oa) >= 5 {
				trxJournals = append(trxJournals, *tj.FromArray(oa))
			}
		}
	}

	return
}

func (bt *BusinessTransaction) FindByPagination(z *zCrud.Meta, z2 *zCrud.PagerIn, z3 *zCrud.PagerOut) (res [][]any) {
	const comment = `-- BusinessTransaction) FindByPagination`

	validFields := BusinessTransactionFieldTypeMap
	whereAndSql := z3.WhereAndSqlTt(z2.Filters, validFields)
	whereAndSql2 := `AND ` + bt.SqlTenantCode() + ` = ` + S.Z(bt.TenantCode)
	if whereAndSql == `` {
		whereAndSql2 = ` WHERE ` + bt.SqlTenantCode() + ` = ` + S.Z(bt.TenantCode)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM SEQSCAN ` + bt.SqlTableName() + whereAndSql + whereAndSql2 + `
LIMIT 1`
	bt.Adapter.QuerySql(queryCount, func(row []any) {
		z3.CalculatePages(z2.Page, z2.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := z3.OrderBySqlTt(z2.Order, validFields)
	limitOffsetSql := z3.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + z.ToSelect() + `
FROM SEQSCAN ` + bt.SqlTableName() + whereAndSql + whereAndSql2 + orderBySql + limitOffsetSql
	bt.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}

func (bt *BusinessTransaction) FindByIdByTenantCode() bool {
	const comment = `-- BusinessTransaction) FindByTenantCode`

	whereAndSql := ` WHERE ` + bt.SqlTenantCode() + ` = ` + S.Z(bt.TenantCode) + `
		AND ` + bt.SqlId() + ` = ` + I.UToS(bt.Id)

	queryRows := comment + `
SELECT ` + bt.SqlSelectAllFields() + `
FROM SEQSCAN ` + bt.SqlTableName() + whereAndSql

	var res []any
	bt.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToU(row[0])
		res = row
	})

	if len(res) > 0 {
		bt.FromArray(res)
		return true
	}

	return false
}

func (tj *TransactionJournal) FindTrxJournalsByTrxTemplateByTenant() (trxJournals []TransactionJournal) {
	var res [][]any
	const comment = `-- TransactionJournal) FindTrxJournalsByTrxTemplateByTenant`

	whereAndSql := ` WHERE ` + tj.SqlTenantCode() + ` = ` + S.Z(tj.TenantCode) + `
		AND ` + tj.SqlTransactionTemplateId() + ` = ` + I.UToS(tj.TransactionTemplateId)

	queryRows := comment + `
SELECT ` + tj.SqlSelectAllFields() + `
FROM SEQSCAN ` + tj.SqlTableName() + whereAndSql

	tj.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	if len(res) > 0 {
		for _, oa := range res {
			if len(oa) >= 5 {
				trxJournals = append(trxJournals, *tj.FromArray(oa))
			}
		}
	}

	return
}

func (tj *TransactionJournal) FindTrxJournalsLossIncomeByTenant() (trxJournals []TransactionJournal) {
	const comment = `-- TransactionJournal) FindTrxJournalsLossIncomeByTenant`

	coa := NewCoa(tj.Adapter)
	coa.TenantCode = tj.TenantCode
	coaIds := coa.FindCoaIDsIncomeExpensesByTenant()

	if len(coaIds) > 0 {
		var res [][]any
		for _, id := range coaIds {
			whereAndSql := ` WHERE ` + tj.SqlTenantCode() + ` = ` + S.Z(tj.TenantCode) + `
				AND ` + tj.SqlCoaId() + ` = ` + I.ToS(id)

			queryRows := comment + `
		SELECT ` + tj.SqlSelectAllFields() + `
		FROM SEQSCAN ` + tj.SqlTableName() + whereAndSql

			tj.Adapter.QuerySql(queryRows, func(row []any) {
				row[0] = X.ToS(row[0]) // ensure id is string
				res = append(res, row)
			})

			if len(res) > 0 {
				for _, oa := range res {
					if len(oa) >= 5 {
						trxJournals = append(trxJournals, *tj.FromArray(oa))
					}
				}
			}
		}
	}

	return
}

func (c *Coa) FindCoaIDsIncomeExpensesByTenant() (coaIds []int64) {
	const comment = `-- Coa) FindCoaIDsIncomeExpensesByTenant`

	coaParent1 := NewCoa(c.Adapter)
	coaParent1.TenantCode = c.TenantCode
	coaParent1.Name = mFinance.CoaExpense
	if coaParent1.FindByNameByTenant() {
		whereAndSql1 := ` WHERE ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode) + `
			AND ` + c.SqlParentId() + ` = ` + I.UToS(coaParent1.Id)

		queryRows1 := comment + `
	SELECT ` + c.SqlId() + `
	FROM SEQSCAN ` + c.SqlTableName() + whereAndSql1

		c.Adapter.QuerySql(queryRows1, func(row []any) {
			coaIds = append(coaIds, X.ToI(row[0]))
		})
	}

	coaParent2 := NewCoa(c.Adapter)
	coaParent2.TenantCode = c.TenantCode
	coaParent2.Name = mFinance.CoaIncome
	if coaParent2.FindByNameByTenant() {
		whereAndSql2 := ` WHERE ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode) + `
			AND ` + c.SqlParentId() + ` = ` + I.UToS(coaParent2.Id)

		queryRows2 := comment + `
	SELECT ` + c.SqlId() + `
	FROM SEQSCAN ` + c.SqlTableName() + whereAndSql2

		c.Adapter.QuerySql(queryRows2, func(row []any) {
			coaIds = append(coaIds, X.ToI(row[0]))
		})
	}

	return
}

func (c *Coa) FindByNameByTenant() bool {
	const comment = `-- Coa) FindByNameByTenant`
	whereAndSql := ` WHERE ` + c.SqlName() + ` = ` + S.Z(c.Name) + `
		AND ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode) + `
		LIMIT 1`

	queryRows := comment + `
	SELECT ` + c.SqlSelectAllFields() + `
	FROM SEQSCAN ` + c.SqlTableName() + whereAndSql

	var res [][]any
	c.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	if len(res) == 1 {
		c.FromArray(res[0])
		return true
	}
	return false
}

type (
	financialPosition struct {
		CoaId     uint64 `json:"coaId"`
		CoaName   string `json:"coaName"`
		AmountIDR int64  `json:"amountIDR"`
	}
	parentName string
)

const (
	rfpAsset     parentName = `asset`
	rfpEquity    parentName = `equity`
	rfpLiability parentName = `liability`
)

func (tj *TransactionJournal) FindReportOfFinancialPositionByTenant() map[parentName][]*financialPosition {
	var rfps = map[parentName][]*financialPosition{}

	const comment = `-- TransactionJournal) FindReportOfFinancialPositionByTenant`
	var findByCoas func(coas []Coa) (res []*financialPosition)
	findByCoas = func(coas []Coa) (res []*financialPosition) {
		cidsMap := map[uint64]bool{}
		cids := func() (res []string) {
			for _, c := range coas {
				cidsMap[c.Id] = false
				res = append(res, I.UToS(c.Id))
			}
			return
		}()

		whereAndSql := ` WHERE ` + tj.SqlTenantCode() + ` = ` + S.Z(tj.TenantCode) + `
			AND ` + tj.SqlCoaId() + ` IN (` + strings.Join(cids, `, `) + `)`

		queryRows := comment + `
			SELECT ` + tj.SqlSelectAllFields() + `
			FROM SEQSCAN ` + tj.SqlTableName() + whereAndSql

		tj.Adapter.QuerySql(queryRows, func(row []any) {
			coaId := X.ToU(row[2])
			debit := X.ToI(row[3])
			credit := X.ToI(row[4])
			amount := debit - credit

			if cidsMap[coaId] {
				for _, rc := range res {
					if rc.CoaId == coaId {
						rc.AmountIDR = rc.AmountIDR + amount
					}
				}

			} else {
				cName := func() (n string) {
					for _, c := range coas {
						if c.Id == coaId {
							n = c.Name
							return
						}
					}
					return
				}
				res = append(res, &financialPosition{
					CoaId:     coaId,
					CoaName:   cName(),
					AmountIDR: amount,
				})

				cidsMap[coaId] = true
			}
		})

		return
	}

	// Find asset
	coaAsset := NewCoa(tj.Adapter)
	coaAsset.TenantCode = tj.TenantCode
	coaAsset.Name = mFinance.CoaAsset
	coaAssetChildren := coaAsset.FindCoaChildrenByParentNameByTenant()
	rfps[rfpAsset] = findByCoas(coaAssetChildren)

	// Find equity
	coaEquity := NewCoa(tj.Adapter)
	coaEquity.TenantCode = tj.TenantCode
	coaEquity.Name = mFinance.CoaEquity
	coaEquityChildren := coaEquity.FindCoaChildrenByParentNameByTenant()
	rfps[rfpEquity] = findByCoas(coaEquityChildren)

	// Find equity
	coaLiability := NewCoa(tj.Adapter)
	coaLiability.TenantCode = tj.TenantCode
	coaLiability.Name = mFinance.CoaLiability
	coaLiabilityChildren := coaLiability.FindCoaChildrenByParentNameByTenant()
	rfps[rfpLiability] = findByCoas(coaLiabilityChildren)

	return rfps
}

func (c *Coa) FindCoaChildrenByParentNameByTenant() (coas []Coa) {
	const comment = `-- Coa) FindCoaChildrenByParentNameByTenant`

	parent := NewCoa(c.Adapter)
	parent.TenantCode = c.TenantCode
	parent.Name = c.Name
	if parent.FindByNameByTenant() {
		var res [][]any

		whereAndSql := ` WHERE ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode) + `
			AND ` + c.SqlParentId() + ` = ` + I.UToS(parent.Id)

		queryRows1 := comment + `
	SELECT ` + c.SqlSelectAllFields() + `
	FROM SEQSCAN ` + c.SqlTableName() + whereAndSql

		c.Adapter.QuerySql(queryRows1, func(row []any) {
			row[0] = X.ToS(row[0]) // ensure id is string
			res = append(res, row)
		})

		if len(res) > 0 {
			for _, v := range res {
				coas = append(coas, *c.FromArray(v))
			}
		}
	}

	return
}
