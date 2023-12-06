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

func (u *Users) FindUsersByTenant(tenantCode string) (res [][]any) {
	const comment = `-- Users) FindByTenant`

	whereAndSql := ` WHERE ` + u.SqlInvitationState() + `LIKE '%tenant:` + tenantCode + `:%'`

	queryRows := comment + `
SELECT *
FROM ` + u.SqlTableName() + whereAndSql

	u.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}
