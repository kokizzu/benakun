package rqAuth

import (
	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
	"github.com/tarantool/go-tarantool"

	"benakun/model/zCrud"
)

type StaffWithInvitation struct {
	Id              uint64 `json:"id" form:"id" query:"id" long:"id" msg:"id"`
	Email           string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	FullName        string `json:"fullName" form:"fullName" query:"fullName" long:"fullName" msg:"fullName"`
	InvitationState	string `json:"invitationState" form:"invitationState" query:"invitationState" long:"invitationState" msg:"invitationState"`
	Role            string `json:"role" form:"role" query:"role" long:"role" msg:"role"`
}

func (u *Users) CheckPassword(pass string) error {
	return S.CheckPassword(u.Password, pass)
}

func (s *Sessions) AllActiveSession(userId uint64, now int64) (res []*Sessions) {
	const comment = `-- Sessions) AllActiveSession`

	query := comment + `
SELECT ` + s.SqlSelectAllFields() + `
FROM ` + s.SqlTableName() + `
WHERE ` + s.SqlUserId() + ` = ` + I.UToS(userId) + `
	AND ` + s.SqlExpiredAt() + ` > ` + I.ToS(now)
	s.Adapter.QuerySql(query, func(row []any) {
		rec := &Sessions{}
		res = append(res, rec.FromArray(row))
	})
	return
}

func (u *Users) FindByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Users) FindByPagination`

	validFields := UsersFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)

	L.Print(`whereAndSQL:`, whereAndSql)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + u.SqlTableName() + whereAndSql + `
LIMIT 1`
	u.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + u.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	u.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}

func (t *Tenants) FindByPagination(z *zCrud.Meta, z2 *zCrud.PagerIn, z3 *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Tenants) FindByPagination`

	validFields := TenantsFieldTypeMap
	whereAndSql := z3.WhereAndSqlTt(z2.Filters, validFields)

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + t.SqlTableName() + whereAndSql + `
LIMIT 1`
	t.Adapter.QuerySql(queryCount, func(row []any) {
		z3.CalculatePages(z2.Page, z2.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := z3.OrderBySqlTt(z2.Order, validFields)
	limitOffsetSql := z3.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + z.ToSelect() + `
FROM ` + t.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	t.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}

func (o *Orgs) FindByCompanyName() bool {
	res, err := o.Adapter.Select(o.SpaceName(), o.IdxName(), 0, 1, tarantool.IterEq, A.X{o.Name})
	if L.IsError(err, `Orgs.FindByCompanyName failed:`+o.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		o.FromArray(rows[0])
		return true
	}
	return false
}

func (u *Users) FindByTenantCode() bool {
	res, err := u.Adapter.Select(u.SpaceName(), u.IdxTenantCode(), 0, 1, tarantool.IterEq, A.X{u.TenantCode})
	if L.IsError(err, `Users.FindByTenantCode failed:`+u.SpaceName()) {
		return false
	}
	rows := res.Tuples()
	if len(rows) == 1 {
		u.FromArray(rows[0])
		return true
	}
	return false
}

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

func (o *Orgs) FindOrgsByTenant(tenantCode string) (orgs []Orgs) {
	const comment = "-- orgs) FindOrgsByTenant"

	whereAndSql := ` WHERE ` + o.SqlTenantCode() + ` = '` + tenantCode + `'`
	queryRows := comment +
		`
SELECT ` + o.SqlSelectAllFields() + `
FROM ` + o.SqlTableName() +
		whereAndSql

	o.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0])
		orgs = append(orgs, *o.FromArray(row))
	})

	return
}

// TODO: find staff
func (u *Users) FindStaffByPagination(meta *zCrud.Meta, in *zCrud.PagerIn, out *zCrud.PagerOut) (res [][]any) {
	const comment = `-- Users) FindStaffByPagination`

	validFields := UsersFieldTypeMap
	whereAndSql := out.WhereAndSqlTt(in.Filters, validFields)
	whereAndSql2 := `AND (` + u.SqlInvitationState() + ` LIKE ` + S.Z(`%tenant:`+u.TenantCode+`:%`)
	if whereAndSql == `` {
		whereAndSql2 = ` WHERE ` + u.SqlInvitationState() + ` LIKE ` + S.Z(`%tenant:`+u.TenantCode+`:%`)
	}

	queryCount := comment + `
SELECT COUNT(1)
FROM ` + u.SqlTableName() + whereAndSql + whereAndSql2 + `
LIMIT 1`
	u.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM ` + u.SqlTableName() + whereAndSql  + whereAndSql2 + orderBySql + limitOffsetSql
	u.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		invState := staffState(X.ToS(row[4]), u.TenantCode)
		row[4] = invState 
		res = append(res, row)
	})

	return
}

func staffState(states, tenantCode string) (invState string) {
	sliceStates := S.Split(S.Trim(states), ` `)
	if len(sliceStates) == 0 || states == `` {
		return ``
	}
	for _, state := range sliceStates {
		if state == `` {
			continue
		}
		parts := S.Split(state, `:`)
		if len(parts) == 4 {
			if parts[1] == tenantCode {
				return parts[2]
			}
		} 
	}
	return ``
}
