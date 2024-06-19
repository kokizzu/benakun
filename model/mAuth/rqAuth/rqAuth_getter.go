package rqAuth

import (
	"benakun/model/mAuth"

	"benakun/model/zCrud"

	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

type Staff struct {
	Id       string `json:"id" form:"id" query:"id" long:"id" msg:"id"`
	Email    string `json:"email" form:"email" query:"email" long:"email" msg:"email"`
	FullName string `json:"fullName" form:"fullName" query:"fullName" long:"fullName" msg:"fullName"`
	Role     string `json:"role" form:"role" query:"role" long:"role" msg:"role"`
}

func (u *Users) FindStaffsByTenantCode(tenantCode string) (staffs []Staff) {
	var res [][]any
	const comment = `-- Users) FindStaffByTenantCode`

	whereAndSql := ` WHERE ` + u.SqlInvitationState() + ` LIKE ` + S.Z(`%tenant:`+tenantCode+`:accepted%`)

	queryRows := comment + `
SELECT ` + u.SqlId() + `, ` + u.SqlEmail() + `, ` + u.SqlFullName() + `, ` + u.SqlRole() + `
FROM SEQSCAN ` + u.SqlTableName() + whereAndSql

	u.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	if len(res) > 0 {
		for _, stf := range res {
			if len(stf) == 4 {
				st := Staff{
					Id:       X.ToS(stf[0]),
					Email:    X.ToS(stf[1]),
					FullName: X.ToS(stf[2]),
					Role:     X.ToS(stf[3]),
				}

				staffs = append(staffs, st)
			}
		}
	}

	L.Print(`Staffs:`, staffs)

	return
}

func (u *Users) FindStaffsChoicesByTenantCode(tenantCode string) map[string]string {
	const comment = `-- Users) FindStaffByTenantCode`

	whereAndSql := ` WHERE ` + u.SqlInvitationState() + ` LIKE ` + S.Z(`%tenant:`+tenantCode+`:accepted%`)

	queryRows := comment + `
SELECT ` + u.SqlId() + `, ` + u.SqlFullName() + `, ` + u.SqlEmail() + `, ` + u.SqlRole() + `
FROM SEQSCAN ` + u.SqlTableName() + whereAndSql

	staffChoices := make(map[string]string)
	u.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 4 {
			fullName := X.ToS(row[1])
			if fullName == `` {
				fullName = `(--)`
			} else {
				fullName = `(` + fullName + `)`
			}
			email := X.ToS(row[2])
			role := X.ToS(row[3])

			staffChoices[X.ToS(row[0])] = fullName + ` ` + email + ` [` + role + `]`
		}
	})

	return staffChoices
}

func (u *Users) FindStaffByIdByTenantCode(id uint64, tenantCode string) bool {
	var res [][]any
	const comment = `-- Users) FindStaffByIdByTenantCode`

	whereAndSql := ` WHERE ` + u.SqlId() + ` = ` + I.UToS(id) + ` AND ` + u.SqlInvitationState() + ` LIKE ` + S.Z(`%tenant:`+tenantCode+`:accepted%`)

	queryRows := comment + `
SELECT ` + u.SqlSelectAllFields() + `
FROM SEQSCAN ` + u.SqlTableName() + whereAndSql + ` LIMIT 1`


	u.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0])
		res = append(res, row)
	})

	if len(res) == 1 {
		u.FromArray(res[0])
		return true
	}

	return false
}

func (u *Users) CheckPassword(pass string) error {
	return S.CheckPassword(u.Password, pass)
}

func (s *Sessions) AllActiveSession(userId uint64, now int64) (res []*Sessions) {
	const comment = `-- Sessions) AllActiveSession`

	query := comment + `
SELECT ` + s.SqlSelectAllFields() + `
FROM SEQSCAN ` + s.SqlTableName() + `
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
FROM SEQSCAN ` + u.SqlTableName() + whereAndSql + `
LIMIT 1`
	u.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM SEQSCAN ` + u.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
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
FROM SEQSCAN ` + t.SqlTableName() + whereAndSql + `
LIMIT 1`
	t.Adapter.QuerySql(queryCount, func(row []any) {
		z3.CalculatePages(z2.Page, z2.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := z3.OrderBySqlTt(z2.Order, validFields)
	limitOffsetSql := z3.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + z.ToSelect() + `
FROM SEQSCAN ` + t.SqlTableName() + whereAndSql + orderBySql + limitOffsetSql
	t.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0]) // ensure id is string
		res = append(res, row)
	})

	return
}

func (o *Orgs) FindByCompanyName() bool {
	selectSql := `-- Orgs) FindByCompanyName
SELECT ` + o.SqlSelectAllFields() + `
FROM SEQSCAN ` + o.SqlTableName() + `
WHERE ` + o.SqlName() + ` = ` + S.Z(o.Name) + `
LIMIT 1`
	ok := false
	o.Adapter.QuerySql(selectSql, func(row []any) {
		o.FromArray(row)
		ok = true
	})
	return ok
}

func (u *Users) FindByTenantCode() bool {
	selectSql := `-- Users) FindByTenantCode
SELECT ` + u.SqlSelectAllFields() + `
FROM SEQSCAN ` + u.SqlTableName() + `
WHERE ` + u.SqlTenantCode() + ` = ` + S.Z(u.TenantCode) + `
LIMIT 1`
	ok := false
	u.Adapter.QuerySql(selectSql, func(row []any) {
		u.FromArray(row)
		ok = true
	})
	return ok
}

func (o *Orgs) FindOrgsByTenant(tenantCode string) (orgs []Orgs) {
	const comment = "-- orgs) FindOrgsByTenant"

	whereAndSql := ` 
WHERE ` + o.SqlTenantCode() + ` = ` + S.Z(tenantCode)
	queryRows := comment + `
SELECT ` + o.SqlSelectAllFields() + `
FROM SEQSCAN ` + o.SqlTableName() +
		whereAndSql

	o.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0])
		orgs = append(orgs, *o.FromArray(row))
	})

	return
}

func (o *Orgs) FindCompanyByTenantCode(tenantCode string) bool {
	const comment = "-- orgs) FindCompanyByTenantCode"

	whereAndSql := ` WHERE ` + o.SqlTenantCode() + ` = ` + S.Z(tenantCode) + `
AND ` + o.SqlOrgType() + ` = ` + I.ToS(mAuth.OrgTypeCompany) + ` LIMIT 1`

	queryRows := comment +
`
SELECT ` + o.SqlSelectAllFields() + `
FROM SEQSCAN ` + o.SqlTableName() + whereAndSql

	o.Adapter.QuerySql(queryRows, func(row []any) {
		o.FromArray(row)
	})

	return o.Id > 0
}

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
FROM SEQSCAN ` + u.SqlTableName() + whereAndSql + whereAndSql2 + `
LIMIT 1`
	u.Adapter.QuerySql(queryCount, func(row []any) {
		out.CalculatePages(in.Page, in.PerPage, int(X.ToI(row[0])))
	})

	orderBySql := out.OrderBySqlTt(in.Order, validFields)
	limitOffsetSql := out.LimitOffsetSql()

	queryRows := comment + `
SELECT ` + meta.ToSelect() + `
FROM SEQSCAN ` + u.SqlTableName() + whereAndSql + whereAndSql2 + orderBySql + limitOffsetSql
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
		if len(parts) == 5 {
			if parts[1] == tenantCode {
				return parts[3]
			}
		}
	}
	return ``
}

func (o *Orgs) FindTenantsHost() (tenants [][]any) {
	const comment = "-- orgs) FindTenantsHost"

	queryRows := comment + `
SELECT orgs.tenantCode, orgs.id
	FROM SEQSCAN orgs
	JOIN SEQSCAN tenants
		ON (orgs.orgType = 1 AND orgs.tenantCode = tenants.tenantCode)`

	o.Adapter.QuerySql(queryRows, func(row []any) {
		if len(row) == 2 {
			row[1] = X.ToS(row[1])

			tenants = append(tenants, row)
		}
	})

	return
}