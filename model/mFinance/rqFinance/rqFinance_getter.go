package rqFinance

import (
	"benakun/model/mFinance"
	"benakun/model/zCrud"
	"fmt"
	"sort"
	"strings"

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

type ParsedAttribute struct {
	IsAutoSum bool
	IsSales bool
	IsChildOnly bool
	IsSelfSum bool
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
		id 				uint64
		name			string
		parentId	uint64
		children	[]any
	}
	// Convert coawithnum to Tree
	coaTreeNode struct {
		id 				uint64
		name			string
		parentId	uint64
		children	[]coaTreeNode
	}
)

func removeUnParent(cs []coaTreeNode, cond func(coaTreeNode) bool) (filtered []coaTreeNode) {
	for _, c := range cs {
		if cond(c) {
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
	coaMaker = func(id uint64) (cwnF coaTreeNode, isVisited bool) {
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
							cwnF.children = append(cwnF.children, child)
						}
					}
				}
				cwnF.id = v.id
				cwnF.name = v.name
				cwnF.parentId = v.parentId
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

func compareCoaNums(a, b []int) bool {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}

	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			return a[i] < b[i]
		}
	}

	return len(a) < len(b)
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
				id: X.ToU(row[0]),
				name: X.ToS(row[1]),
				parentId: X.ToU(row[2]),
				children: X.ToArr(row[3]),
			})
		}
	})

	coaChoices := generateCoaChoicesMaps(coasFromDBs)

	type kv struct {
		key string
		value string
	}
	
	kvsCoaChoices := []kv{}

	for k, v := range coaChoices {
		kvsCoaChoices = append(kvsCoaChoices, kv{k, v})
	}

	sort.Slice(kvsCoaChoices, func(i, j int) bool {
		numsI := extractNumParts(kvsCoaChoices[i].value)
		numsJ := extractNumParts(kvsCoaChoices[j].value)

		return compareCoaNums(numsI, numsJ)
	})

	if len(kvsCoaChoices) > 0 {
		coaChoices = map[string]string{}
		for _, v := range kvsCoaChoices {
			coaChoices[v.key] = v.value
		}
	}
	
	return coaChoices
}

func (c *Coa) FindCoasChoicesChildByParentByTenant() map[string]string {
	const comment = `-- Coa) FindCoasChoicesChildByParentByTenant`

	whereAndSql := ` WHERE ` + c.SqlParentId() + ` = ` + I.UToS(c.ParentId) + `
		AND `  + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode)

	queryRows := comment + `
	SELECT ` + c.SqlId() + `, ` + c.SqlName() + `
	FROM SEQSCAN ` + c.SqlTableName() + whereAndSql

	coaChoices := make(map[string]string)
	var idx int64 = 1
	c.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			coaChoices[X.ToS(row[0])] = I.ToS(idx) + `. `+ X.ToS(row[1])
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

func (tj *TransactionJournal) FindTrxJournalsByDateByTenant() (trxJournals []TransactionJournal) {
	var res [][]any
	const comment = `-- TransactionJournal) FindTrxJournalsByDateByTenant`

	whereAndSql := ` WHERE ` + tj.SqlTenantCode() + ` = ` + S.Z(tj.TenantCode) + `
		AND ` + tj.SqlDate() + ` = ` + S.Z(tj.Date)

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