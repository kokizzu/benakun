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
	Id              string `json:"id" form:"id" query:"id" long:"id" msg:"id"`
	Email           string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	FullName        string `json:"fullName" form:"fullName" query:"fullName" long:"fullName" msg:"fullName"`
	InvitationState string `json:"invitationState" form:"invitationState" query:"invitationState" long:"invitationState" msg:"invitationState"`
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

func (u *Users) FindUsersByTenant(tenantCode string) (staffs []StaffWithInvitation) {
	var res [][]any
	const comment = `-- Users) FindByTenant`

	whereAndSql := ` WHERE ` + u.SqlInvitationState() + `LIKE '%tenant:` + tenantCode + `:%'`

	queryRows := comment + `
SELECT ` + u.SqlId() + `, ` + u.SqlEmail() + `, ` + u.SqlFullName() + `, ` + u.SqlInvitationState() + `, ` + u.SqlRole() + `
FROM ` + u.SqlTableName() + whereAndSql

	u.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	if len(res) > 0 {
		for _, stf := range res {
			if len(stf) >= 5 {
				st := StaffWithInvitation{
					Id:              stf[0].(string),
					Email:           stf[1].(string),
					FullName:        stf[2].(string),
					InvitationState: stf[3].(string),
					Role:            stf[4].(string),
				}
				staffs = append(staffs, st)
			}
		}
	} else {
		return []StaffWithInvitation{}
	}

	return
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

	whereAndSql := ` WHERE ` + c.SqlTenantCode() + ` = ` + S.Z(c.TenantCode) + ` AND ` + c.SqlLevel() + ` = ` + I.ToS(int64(c.Level))

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

func (c *Coa) FindCoaChildIdByParentIdByName() uint64 {
	var res [][]any
	const comment = `-- Coa) FindCoaChildIdByParentIdByName`

	whereAndSql := ` WHERE ` + c.SqlParentId() + ` = ` + I.ToS(int64(c.ParentId)) + ` AND ` + c.SqlName() + ` = ` + S.Z(c.Name)

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

func (o *Orgs) FindOrgsByTenantCode(tenantCode string) (orgs []Orgs) {
	const comment = "-- orgs) FindByTenant"

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