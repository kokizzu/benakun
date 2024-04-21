package wcAuth

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

import (
	"benakun/model/mAuth/rqAuth"

	"github.com/kokizzu/gotro/A"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/X"
)

// CoaMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcAuth__ORM.GEN.go
type CoaMutator struct {
	rqAuth.Coa
	mutations []A.X
	logs      []A.X
}

// NewCoaMutator create new ORM writer/command object
func NewCoaMutator(adapter *Tt.Adapter) (res *CoaMutator) {
	res = &CoaMutator{Coa: rqAuth.Coa{Adapter: adapter}}
	res.Children = []any{}
	return
}

// Logs get array of logs [field, old, new]
func (c *CoaMutator) Logs() []A.X { //nolint:dupl false positive
	return c.logs
}

// HaveMutation check whether Set* methods ever called
func (c *CoaMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(c.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (c *CoaMutator) ClearMutations() { //nolint:dupl false positive
	c.mutations = []A.X{}
	c.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (c *CoaMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := c.Adapter.Update(c.SpaceName(), c.UniqueIndexId(), A.X{c.Id}, c.ToUpdateArray())
	return !L.IsError(err, `Coa.DoOverwriteById failed: `+c.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (c *CoaMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !c.HaveMutation() {
		return true
	}
	_, err := c.Adapter.Update(c.SpaceName(), c.UniqueIndexId(), A.X{c.Id}, c.mutations)
	return !L.IsError(err, `Coa.DoUpdateById failed: `+c.SpaceName())
}

// DoDeletePermanentById permanent delete
func (c *CoaMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := c.Adapter.Delete(c.SpaceName(), c.UniqueIndexId(), A.X{c.Id})
	return !L.IsError(err, `Coa.DoDeletePermanentById failed: `+c.SpaceName())
}

// func (c *CoaMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := c.ToArray()
//	_, err := c.Adapter.Upsert(c.SpaceName(), arr, A.X{
//		A.X{`=`, 0, c.Id},
//		A.X{`=`, 1, c.TenantCode},
//		A.X{`=`, 2, c.Name},
//		A.X{`=`, 3, c.Level},
//		A.X{`=`, 4, c.ParentId},
//		A.X{`=`, 5, c.Children},
//		A.X{`=`, 6, c.CreatedAt},
//		A.X{`=`, 7, c.CreatedBy},
//		A.X{`=`, 8, c.UpdatedAt},
//		A.X{`=`, 9, c.UpdatedBy},
//		A.X{`=`, 10, c.DeletedAt},
//	})
//	return !L.IsError(err, `Coa.DoUpsert failed: `+c.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (c *CoaMutator) DoInsert() bool { //nolint:dupl false positive
	arr := c.ToArray()
	row, err := c.Adapter.Insert(c.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			c.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Coa.DoInsert failed: `+c.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (c *CoaMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := c.ToArray()
	row, err := c.Adapter.Replace(c.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			c.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Coa.DoUpsert failed: `+c.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (c *CoaMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != c.Id {
		c.mutations = append(c.mutations, A.X{`=`, 0, val})
		c.logs = append(c.logs, A.X{`id`, c.Id, val})
		c.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (c *CoaMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != c.TenantCode {
		c.mutations = append(c.mutations, A.X{`=`, 1, val})
		c.logs = append(c.logs, A.X{`tenantCode`, c.TenantCode, val})
		c.TenantCode = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (c *CoaMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != c.Name {
		c.mutations = append(c.mutations, A.X{`=`, 2, val})
		c.logs = append(c.logs, A.X{`name`, c.Name, val})
		c.Name = val
		return true
	}
	return false
}

// SetLevel create mutations, should not duplicate
func (c *CoaMutator) SetLevel(val float64) bool { //nolint:dupl false positive
	if val != c.Level {
		c.mutations = append(c.mutations, A.X{`=`, 3, val})
		c.logs = append(c.logs, A.X{`level`, c.Level, val})
		c.Level = val
		return true
	}
	return false
}

// SetParentId create mutations, should not duplicate
func (c *CoaMutator) SetParentId(val uint64) bool { //nolint:dupl false positive
	if val != c.ParentId {
		c.mutations = append(c.mutations, A.X{`=`, 4, val})
		c.logs = append(c.logs, A.X{`parentId`, c.ParentId, val})
		c.ParentId = val
		return true
	}
	return false
}

// SetChildren create mutations, should not duplicate
func (c *CoaMutator) SetChildren(val []any) bool { //nolint:dupl false positive
	c.mutations = append(c.mutations, A.X{`=`, 5, val})
	c.logs = append(c.logs, A.X{`children`, c.Children, val})
	c.Children = val
	return true
}

// SetCreatedAt create mutations, should not duplicate
func (c *CoaMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != c.CreatedAt {
		c.mutations = append(c.mutations, A.X{`=`, 6, val})
		c.logs = append(c.logs, A.X{`createdAt`, c.CreatedAt, val})
		c.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (c *CoaMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != c.CreatedBy {
		c.mutations = append(c.mutations, A.X{`=`, 7, val})
		c.logs = append(c.logs, A.X{`createdBy`, c.CreatedBy, val})
		c.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (c *CoaMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != c.UpdatedAt {
		c.mutations = append(c.mutations, A.X{`=`, 8, val})
		c.logs = append(c.logs, A.X{`updatedAt`, c.UpdatedAt, val})
		c.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (c *CoaMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != c.UpdatedBy {
		c.mutations = append(c.mutations, A.X{`=`, 9, val})
		c.logs = append(c.logs, A.X{`updatedBy`, c.UpdatedBy, val})
		c.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (c *CoaMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != c.DeletedAt {
		c.mutations = append(c.mutations, A.X{`=`, 10, val})
		c.logs = append(c.logs, A.X{`deletedAt`, c.DeletedAt, val})
		c.DeletedAt = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (c *CoaMutator) SetAll(from rqAuth.Coa, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		c.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		c.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`name`] && (forceMap[`name`] || from.Name != ``) {
		c.Name = S.Trim(from.Name)
		changed = true
	}
	if !excludeMap[`level`] && (forceMap[`level`] || from.Level != 0) {
		c.Level = from.Level
		changed = true
	}
	if !excludeMap[`parentId`] && (forceMap[`parentId`] || from.ParentId != 0) {
		c.ParentId = from.ParentId
		changed = true
	}
	if !excludeMap[`children`] && (forceMap[`children`] || from.Children != nil) {
		c.Children = from.Children
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		c.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		c.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		c.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		c.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		c.DeletedAt = from.DeletedAt
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// OrgsMutator DAO writer/command struct
type OrgsMutator struct {
	rqAuth.Orgs
	mutations []A.X
	logs      []A.X
}

// NewOrgsMutator create new ORM writer/command object
func NewOrgsMutator(adapter *Tt.Adapter) (res *OrgsMutator) {
	res = &OrgsMutator{Orgs: rqAuth.Orgs{Adapter: adapter}}
	res.Children = []any{}
	return
}

// Logs get array of logs [field, old, new]
func (o *OrgsMutator) Logs() []A.X { //nolint:dupl false positive
	return o.logs
}

// HaveMutation check whether Set* methods ever called
func (o *OrgsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(o.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (o *OrgsMutator) ClearMutations() { //nolint:dupl false positive
	o.mutations = []A.X{}
	o.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (o *OrgsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := o.Adapter.Update(o.SpaceName(), o.UniqueIndexId(), A.X{o.Id}, o.ToUpdateArray())
	return !L.IsError(err, `Orgs.DoOverwriteById failed: `+o.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (o *OrgsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !o.HaveMutation() {
		return true
	}
	_, err := o.Adapter.Update(o.SpaceName(), o.UniqueIndexId(), A.X{o.Id}, o.mutations)
	return !L.IsError(err, `Orgs.DoUpdateById failed: `+o.SpaceName())
}

// DoDeletePermanentById permanent delete
func (o *OrgsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := o.Adapter.Delete(o.SpaceName(), o.UniqueIndexId(), A.X{o.Id})
	return !L.IsError(err, `Orgs.DoDeletePermanentById failed: `+o.SpaceName())
}

// func (o *OrgsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := o.ToArray()
//	_, err := o.Adapter.Upsert(o.SpaceName(), arr, A.X{
//		A.X{`=`, 0, o.Id},
//		A.X{`=`, 1, o.TenantCode},
//		A.X{`=`, 2, o.Name},
//		A.X{`=`, 3, o.HeadTitle},
//		A.X{`=`, 4, o.ParentId},
//		A.X{`=`, 5, o.Children},
//		A.X{`=`, 6, o.OrgType},
//		A.X{`=`, 7, o.CreatedAt},
//		A.X{`=`, 8, o.CreatedBy},
//		A.X{`=`, 9, o.UpdatedAt},
//		A.X{`=`, 10, o.UpdatedBy},
//		A.X{`=`, 11, o.DeletedAt},
//	})
//	return !L.IsError(err, `Orgs.DoUpsert failed: `+o.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (o *OrgsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := o.ToArray()
	row, err := o.Adapter.Insert(o.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			o.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Orgs.DoInsert failed: `+o.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (o *OrgsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := o.ToArray()
	row, err := o.Adapter.Replace(o.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			o.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Orgs.DoUpsert failed: `+o.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (o *OrgsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != o.Id {
		o.mutations = append(o.mutations, A.X{`=`, 0, val})
		o.logs = append(o.logs, A.X{`id`, o.Id, val})
		o.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (o *OrgsMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != o.TenantCode {
		o.mutations = append(o.mutations, A.X{`=`, 1, val})
		o.logs = append(o.logs, A.X{`tenantCode`, o.TenantCode, val})
		o.TenantCode = val
		return true
	}
	return false
}

// SetName create mutations, should not duplicate
func (o *OrgsMutator) SetName(val string) bool { //nolint:dupl false positive
	if val != o.Name {
		o.mutations = append(o.mutations, A.X{`=`, 2, val})
		o.logs = append(o.logs, A.X{`name`, o.Name, val})
		o.Name = val
		return true
	}
	return false
}

// SetHeadTitle create mutations, should not duplicate
func (o *OrgsMutator) SetHeadTitle(val string) bool { //nolint:dupl false positive
	if val != o.HeadTitle {
		o.mutations = append(o.mutations, A.X{`=`, 3, val})
		o.logs = append(o.logs, A.X{`headTitle`, o.HeadTitle, val})
		o.HeadTitle = val
		return true
	}
	return false
}

// SetParentId create mutations, should not duplicate
func (o *OrgsMutator) SetParentId(val uint64) bool { //nolint:dupl false positive
	if val != o.ParentId {
		o.mutations = append(o.mutations, A.X{`=`, 4, val})
		o.logs = append(o.logs, A.X{`parentId`, o.ParentId, val})
		o.ParentId = val
		return true
	}
	return false
}

// SetChildren create mutations, should not duplicate
func (o *OrgsMutator) SetChildren(val []any) bool { //nolint:dupl false positive
	o.mutations = append(o.mutations, A.X{`=`, 5, val})
	o.logs = append(o.logs, A.X{`children`, o.Children, val})
	o.Children = val
	return true
}

// SetOrgType create mutations, should not duplicate
func (o *OrgsMutator) SetOrgType(val uint64) bool { //nolint:dupl false positive
	if val != o.OrgType {
		o.mutations = append(o.mutations, A.X{`=`, 6, val})
		o.logs = append(o.logs, A.X{`orgType`, o.OrgType, val})
		o.OrgType = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (o *OrgsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != o.CreatedAt {
		o.mutations = append(o.mutations, A.X{`=`, 7, val})
		o.logs = append(o.logs, A.X{`createdAt`, o.CreatedAt, val})
		o.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (o *OrgsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != o.CreatedBy {
		o.mutations = append(o.mutations, A.X{`=`, 8, val})
		o.logs = append(o.logs, A.X{`createdBy`, o.CreatedBy, val})
		o.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (o *OrgsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != o.UpdatedAt {
		o.mutations = append(o.mutations, A.X{`=`, 9, val})
		o.logs = append(o.logs, A.X{`updatedAt`, o.UpdatedAt, val})
		o.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (o *OrgsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != o.UpdatedBy {
		o.mutations = append(o.mutations, A.X{`=`, 10, val})
		o.logs = append(o.logs, A.X{`updatedBy`, o.UpdatedBy, val})
		o.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (o *OrgsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != o.DeletedAt {
		o.mutations = append(o.mutations, A.X{`=`, 11, val})
		o.logs = append(o.logs, A.X{`deletedAt`, o.DeletedAt, val})
		o.DeletedAt = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (o *OrgsMutator) SetAll(from rqAuth.Orgs, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		o.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		o.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`name`] && (forceMap[`name`] || from.Name != ``) {
		o.Name = S.Trim(from.Name)
		changed = true
	}
	if !excludeMap[`headTitle`] && (forceMap[`headTitle`] || from.HeadTitle != ``) {
		o.HeadTitle = S.Trim(from.HeadTitle)
		changed = true
	}
	if !excludeMap[`parentId`] && (forceMap[`parentId`] || from.ParentId != 0) {
		o.ParentId = from.ParentId
		changed = true
	}
	if !excludeMap[`children`] && (forceMap[`children`] || from.Children != nil) {
		o.Children = from.Children
		changed = true
	}
	if !excludeMap[`orgType`] && (forceMap[`orgType`] || from.OrgType != 0) {
		o.OrgType = from.OrgType
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		o.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		o.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		o.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		o.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		o.DeletedAt = from.DeletedAt
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// SessionsMutator DAO writer/command struct
type SessionsMutator struct {
	rqAuth.Sessions
	mutations []A.X
	logs      []A.X
}

// NewSessionsMutator create new ORM writer/command object
func NewSessionsMutator(adapter *Tt.Adapter) (res *SessionsMutator) {
	res = &SessionsMutator{Sessions: rqAuth.Sessions{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (s *SessionsMutator) Logs() []A.X { //nolint:dupl false positive
	return s.logs
}

// HaveMutation check whether Set* methods ever called
func (s *SessionsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(s.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (s *SessionsMutator) ClearMutations() { //nolint:dupl false positive
	s.mutations = []A.X{}
	s.logs = []A.X{}
}

// func (s *SessionsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := s.ToArray()
//	_, err := s.Adapter.Upsert(s.SpaceName(), arr, A.X{
//		A.X{`=`, 0, s.SessionToken},
//		A.X{`=`, 1, s.UserId},
//		A.X{`=`, 2, s.ExpiredAt},
//		A.X{`=`, 3, s.Device},
//		A.X{`=`, 4, s.LoginAt},
//		A.X{`=`, 5, s.LoginIPs},
//		A.X{`=`, 6, s.TenantCode},
//	})
//	return !L.IsError(err, `Sessions.DoUpsert failed: `+s.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteBySessionToken update all columns, error if not exists, not using mutations/Set*
func (s *SessionsMutator) DoOverwriteBySessionToken() bool { //nolint:dupl false positive
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexSessionToken(), A.X{s.SessionToken}, s.ToUpdateArray())
	return !L.IsError(err, `Sessions.DoOverwriteBySessionToken failed: `+s.SpaceName())
}

// DoUpdateBySessionToken update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (s *SessionsMutator) DoUpdateBySessionToken() bool { //nolint:dupl false positive
	if !s.HaveMutation() {
		return true
	}
	_, err := s.Adapter.Update(s.SpaceName(), s.UniqueIndexSessionToken(), A.X{s.SessionToken}, s.mutations)
	return !L.IsError(err, `Sessions.DoUpdateBySessionToken failed: `+s.SpaceName())
}

// DoDeletePermanentBySessionToken permanent delete
func (s *SessionsMutator) DoDeletePermanentBySessionToken() bool { //nolint:dupl false positive
	_, err := s.Adapter.Delete(s.SpaceName(), s.UniqueIndexSessionToken(), A.X{s.SessionToken})
	return !L.IsError(err, `Sessions.DoDeletePermanentBySessionToken failed: `+s.SpaceName())
}

// DoInsert insert, error if already exists
func (s *SessionsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	_, err := s.Adapter.Insert(s.SpaceName(), arr)
	return !L.IsError(err, `Sessions.DoInsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (s *SessionsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := s.ToArray()
	_, err := s.Adapter.Replace(s.SpaceName(), arr)
	return !L.IsError(err, `Sessions.DoUpsert failed: `+s.SpaceName()+`\n%#v`, arr)
}

// SetSessionToken create mutations, should not duplicate
func (s *SessionsMutator) SetSessionToken(val string) bool { //nolint:dupl false positive
	if val != s.SessionToken {
		s.mutations = append(s.mutations, A.X{`=`, 0, val})
		s.logs = append(s.logs, A.X{`sessionToken`, s.SessionToken, val})
		s.SessionToken = val
		return true
	}
	return false
}

// SetUserId create mutations, should not duplicate
func (s *SessionsMutator) SetUserId(val uint64) bool { //nolint:dupl false positive
	if val != s.UserId {
		s.mutations = append(s.mutations, A.X{`=`, 1, val})
		s.logs = append(s.logs, A.X{`userId`, s.UserId, val})
		s.UserId = val
		return true
	}
	return false
}

// SetExpiredAt create mutations, should not duplicate
func (s *SessionsMutator) SetExpiredAt(val int64) bool { //nolint:dupl false positive
	if val != s.ExpiredAt {
		s.mutations = append(s.mutations, A.X{`=`, 2, val})
		s.logs = append(s.logs, A.X{`expiredAt`, s.ExpiredAt, val})
		s.ExpiredAt = val
		return true
	}
	return false
}

// SetDevice create mutations, should not duplicate
func (s *SessionsMutator) SetDevice(val string) bool { //nolint:dupl false positive
	if val != s.Device {
		s.mutations = append(s.mutations, A.X{`=`, 3, val})
		s.logs = append(s.logs, A.X{`device`, s.Device, val})
		s.Device = val
		return true
	}
	return false
}

// SetLoginAt create mutations, should not duplicate
func (s *SessionsMutator) SetLoginAt(val int64) bool { //nolint:dupl false positive
	if val != s.LoginAt {
		s.mutations = append(s.mutations, A.X{`=`, 4, val})
		s.logs = append(s.logs, A.X{`loginAt`, s.LoginAt, val})
		s.LoginAt = val
		return true
	}
	return false
}

// SetLoginIPs create mutations, should not duplicate
func (s *SessionsMutator) SetLoginIPs(val string) bool { //nolint:dupl false positive
	if val != s.LoginIPs {
		s.mutations = append(s.mutations, A.X{`=`, 5, val})
		s.logs = append(s.logs, A.X{`loginIPs`, s.LoginIPs, val})
		s.LoginIPs = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (s *SessionsMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != s.TenantCode {
		s.mutations = append(s.mutations, A.X{`=`, 6, val})
		s.logs = append(s.logs, A.X{`tenantCode`, s.TenantCode, val})
		s.TenantCode = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (s *SessionsMutator) SetAll(from rqAuth.Sessions, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`sessionToken`] && (forceMap[`sessionToken`] || from.SessionToken != ``) {
		s.SessionToken = S.Trim(from.SessionToken)
		changed = true
	}
	if !excludeMap[`userId`] && (forceMap[`userId`] || from.UserId != 0) {
		s.UserId = from.UserId
		changed = true
	}
	if !excludeMap[`expiredAt`] && (forceMap[`expiredAt`] || from.ExpiredAt != 0) {
		s.ExpiredAt = from.ExpiredAt
		changed = true
	}
	if !excludeMap[`device`] && (forceMap[`device`] || from.Device != ``) {
		s.Device = S.Trim(from.Device)
		changed = true
	}
	if !excludeMap[`loginAt`] && (forceMap[`loginAt`] || from.LoginAt != 0) {
		s.LoginAt = from.LoginAt
		changed = true
	}
	if !excludeMap[`loginIPs`] && (forceMap[`loginIPs`] || from.LoginIPs != ``) {
		s.LoginIPs = S.Trim(from.LoginIPs)
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		s.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// TenantsMutator DAO writer/command struct
type TenantsMutator struct {
	rqAuth.Tenants
	mutations []A.X
	logs      []A.X
}

// NewTenantsMutator create new ORM writer/command object
func NewTenantsMutator(adapter *Tt.Adapter) (res *TenantsMutator) {
	res = &TenantsMutator{Tenants: rqAuth.Tenants{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (t *TenantsMutator) Logs() []A.X { //nolint:dupl false positive
	return t.logs
}

// HaveMutation check whether Set* methods ever called
func (t *TenantsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(t.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (t *TenantsMutator) ClearMutations() { //nolint:dupl false positive
	t.mutations = []A.X{}
	t.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (t *TenantsMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.ToUpdateArray())
	return !L.IsError(err, `Tenants.DoOverwriteById failed: `+t.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TenantsMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id}, t.mutations)
	return !L.IsError(err, `Tenants.DoUpdateById failed: `+t.SpaceName())
}

// DoDeletePermanentById permanent delete
func (t *TenantsMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := t.Adapter.Delete(t.SpaceName(), t.UniqueIndexId(), A.X{t.Id})
	return !L.IsError(err, `Tenants.DoDeletePermanentById failed: `+t.SpaceName())
}

// func (t *TenantsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := t.ToArray()
//	_, err := t.Adapter.Upsert(t.SpaceName(), arr, A.X{
//		A.X{`=`, 0, t.Id},
//		A.X{`=`, 1, t.TenantCode},
//		A.X{`=`, 2, t.CreatedAt},
//		A.X{`=`, 3, t.CreatedBy},
//		A.X{`=`, 4, t.UpdatedAt},
//		A.X{`=`, 5, t.UpdatedBy},
//		A.X{`=`, 6, t.DeletedAt},
//	})
//	return !L.IsError(err, `Tenants.DoUpsert failed: `+t.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByTenantCode update all columns, error if not exists, not using mutations/Set*
func (t *TenantsMutator) DoOverwriteByTenantCode() bool { //nolint:dupl false positive
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexTenantCode(), A.X{t.TenantCode}, t.ToUpdateArray())
	return !L.IsError(err, `Tenants.DoOverwriteByTenantCode failed: `+t.SpaceName())
}

// DoUpdateByTenantCode update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (t *TenantsMutator) DoUpdateByTenantCode() bool { //nolint:dupl false positive
	if !t.HaveMutation() {
		return true
	}
	_, err := t.Adapter.Update(t.SpaceName(), t.UniqueIndexTenantCode(), A.X{t.TenantCode}, t.mutations)
	return !L.IsError(err, `Tenants.DoUpdateByTenantCode failed: `+t.SpaceName())
}

// DoDeletePermanentByTenantCode permanent delete
func (t *TenantsMutator) DoDeletePermanentByTenantCode() bool { //nolint:dupl false positive
	_, err := t.Adapter.Delete(t.SpaceName(), t.UniqueIndexTenantCode(), A.X{t.TenantCode})
	return !L.IsError(err, `Tenants.DoDeletePermanentByTenantCode failed: `+t.SpaceName())
}

// DoInsert insert, error if already exists
func (t *TenantsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.Insert(t.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Tenants.DoInsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (t *TenantsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	row, err := t.Adapter.Replace(t.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			t.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Tenants.DoUpsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (t *TenantsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != t.Id {
		t.mutations = append(t.mutations, A.X{`=`, 0, val})
		t.logs = append(t.logs, A.X{`id`, t.Id, val})
		t.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (t *TenantsMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != t.TenantCode {
		t.mutations = append(t.mutations, A.X{`=`, 1, val})
		t.logs = append(t.logs, A.X{`tenantCode`, t.TenantCode, val})
		t.TenantCode = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (t *TenantsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.CreatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 2, val})
		t.logs = append(t.logs, A.X{`createdAt`, t.CreatedAt, val})
		t.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (t *TenantsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.CreatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 3, val})
		t.logs = append(t.logs, A.X{`createdBy`, t.CreatedBy, val})
		t.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (t *TenantsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.UpdatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 4, val})
		t.logs = append(t.logs, A.X{`updatedAt`, t.UpdatedAt, val})
		t.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (t *TenantsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.UpdatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 5, val})
		t.logs = append(t.logs, A.X{`updatedBy`, t.UpdatedBy, val})
		t.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (t *TenantsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != t.DeletedAt {
		t.mutations = append(t.mutations, A.X{`=`, 6, val})
		t.logs = append(t.logs, A.X{`deletedAt`, t.DeletedAt, val})
		t.DeletedAt = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (t *TenantsMutator) SetAll(from rqAuth.Tenants, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		t.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		t.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		t.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		t.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		t.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		t.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		t.DeletedAt = from.DeletedAt
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// TransactionsMutator DAO writer/command struct
type TransactionsMutator struct {
	rqAuth.Transactions
	mutations []A.X
	logs      []A.X
}

// NewTransactionsMutator create new ORM writer/command object
func NewTransactionsMutator(adapter *Tt.Adapter) (res *TransactionsMutator) {
	res = &TransactionsMutator{Transactions: rqAuth.Transactions{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (t *TransactionsMutator) Logs() []A.X { //nolint:dupl false positive
	return t.logs
}

// HaveMutation check whether Set* methods ever called
func (t *TransactionsMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(t.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (t *TransactionsMutator) ClearMutations() { //nolint:dupl false positive
	t.mutations = []A.X{}
	t.logs = []A.X{}
}

// func (t *TransactionsMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := t.ToArray()
//	_, err := t.Adapter.Upsert(t.SpaceName(), arr, A.X{
//		A.X{`=`, 0, t.Id},
//		A.X{`=`, 1, t.TenantCode},
//		A.X{`=`, 2, t.CreatedAt},
//		A.X{`=`, 3, t.CreatedBy},
//		A.X{`=`, 4, t.UpdatedAt},
//		A.X{`=`, 5, t.UpdatedBy},
//		A.X{`=`, 6, t.DeletedAt},
//		A.X{`=`, 7, t.CompletedAt},
//		A.X{`=`, 8, t.CoaId},
//		A.X{`=`, 9, t.Price},
//		A.X{`=`, 10, t.Descriptions},
//		A.X{`=`, 11, t.Qty},
//	})
//	return !L.IsError(err, `Transactions.DoUpsert failed: `+t.SpaceName()+ `\n%#v`, arr)
// }

// DoInsert insert, error if already exists
func (t *TransactionsMutator) DoInsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	_, err := t.Adapter.Insert(t.SpaceName(), arr)
	return !L.IsError(err, `Transactions.DoInsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (t *TransactionsMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := t.ToArray()
	_, err := t.Adapter.Replace(t.SpaceName(), arr)
	return !L.IsError(err, `Transactions.DoUpsert failed: `+t.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (t *TransactionsMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != t.Id {
		t.mutations = append(t.mutations, A.X{`=`, 0, val})
		t.logs = append(t.logs, A.X{`id`, t.Id, val})
		t.Id = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (t *TransactionsMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != t.TenantCode {
		t.mutations = append(t.mutations, A.X{`=`, 1, val})
		t.logs = append(t.logs, A.X{`tenantCode`, t.TenantCode, val})
		t.TenantCode = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (t *TransactionsMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.CreatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 2, val})
		t.logs = append(t.logs, A.X{`createdAt`, t.CreatedAt, val})
		t.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (t *TransactionsMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.CreatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 3, val})
		t.logs = append(t.logs, A.X{`createdBy`, t.CreatedBy, val})
		t.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (t *TransactionsMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != t.UpdatedAt {
		t.mutations = append(t.mutations, A.X{`=`, 4, val})
		t.logs = append(t.logs, A.X{`updatedAt`, t.UpdatedAt, val})
		t.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (t *TransactionsMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != t.UpdatedBy {
		t.mutations = append(t.mutations, A.X{`=`, 5, val})
		t.logs = append(t.logs, A.X{`updatedBy`, t.UpdatedBy, val})
		t.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (t *TransactionsMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != t.DeletedAt {
		t.mutations = append(t.mutations, A.X{`=`, 6, val})
		t.logs = append(t.logs, A.X{`deletedAt`, t.DeletedAt, val})
		t.DeletedAt = val
		return true
	}
	return false
}

// SetCompletedAt create mutations, should not duplicate
func (t *TransactionsMutator) SetCompletedAt(val int64) bool { //nolint:dupl false positive
	if val != t.CompletedAt {
		t.mutations = append(t.mutations, A.X{`=`, 7, val})
		t.logs = append(t.logs, A.X{`completedAt`, t.CompletedAt, val})
		t.CompletedAt = val
		return true
	}
	return false
}

// SetCoaId create mutations, should not duplicate
func (t *TransactionsMutator) SetCoaId(val uint64) bool { //nolint:dupl false positive
	if val != t.CoaId {
		t.mutations = append(t.mutations, A.X{`=`, 8, val})
		t.logs = append(t.logs, A.X{`coaId`, t.CoaId, val})
		t.CoaId = val
		return true
	}
	return false
}

// SetPrice create mutations, should not duplicate
func (t *TransactionsMutator) SetPrice(val int64) bool { //nolint:dupl false positive
	if val != t.Price {
		t.mutations = append(t.mutations, A.X{`=`, 9, val})
		t.logs = append(t.logs, A.X{`price`, t.Price, val})
		t.Price = val
		return true
	}
	return false
}

// SetDescriptions create mutations, should not duplicate
func (t *TransactionsMutator) SetDescriptions(val string) bool { //nolint:dupl false positive
	if val != t.Descriptions {
		t.mutations = append(t.mutations, A.X{`=`, 10, val})
		t.logs = append(t.logs, A.X{`descriptions`, t.Descriptions, val})
		t.Descriptions = val
		return true
	}
	return false
}

// SetQty create mutations, should not duplicate
func (t *TransactionsMutator) SetQty(val int64) bool { //nolint:dupl false positive
	if val != t.Qty {
		t.mutations = append(t.mutations, A.X{`=`, 11, val})
		t.logs = append(t.logs, A.X{`qty`, t.Qty, val})
		t.Qty = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (t *TransactionsMutator) SetAll(from rqAuth.Transactions, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		t.Id = from.Id
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		t.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		t.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		t.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		t.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		t.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		t.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`completedAt`] && (forceMap[`completedAt`] || from.CompletedAt != 0) {
		t.CompletedAt = from.CompletedAt
		changed = true
	}
	if !excludeMap[`coaId`] && (forceMap[`coaId`] || from.CoaId != 0) {
		t.CoaId = from.CoaId
		changed = true
	}
	if !excludeMap[`price`] && (forceMap[`price`] || from.Price != 0) {
		t.Price = from.Price
		changed = true
	}
	if !excludeMap[`descriptions`] && (forceMap[`descriptions`] || from.Descriptions != ``) {
		t.Descriptions = S.Trim(from.Descriptions)
		changed = true
	}
	if !excludeMap[`qty`] && (forceMap[`qty`] || from.Qty != 0) {
		t.Qty = from.Qty
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go

// UsersMutator DAO writer/command struct
type UsersMutator struct {
	rqAuth.Users
	mutations []A.X
	logs      []A.X
}

// NewUsersMutator create new ORM writer/command object
func NewUsersMutator(adapter *Tt.Adapter) (res *UsersMutator) {
	res = &UsersMutator{Users: rqAuth.Users{Adapter: adapter}}
	return
}

// Logs get array of logs [field, old, new]
func (u *UsersMutator) Logs() []A.X { //nolint:dupl false positive
	return u.logs
}

// HaveMutation check whether Set* methods ever called
func (u *UsersMutator) HaveMutation() bool { //nolint:dupl false positive
	return len(u.mutations) > 0
}

// ClearMutations clear all previously called Set* methods
func (u *UsersMutator) ClearMutations() { //nolint:dupl false positive
	u.mutations = []A.X{}
	u.logs = []A.X{}
}

// DoOverwriteById update all columns, error if not exists, not using mutations/Set*
func (u *UsersMutator) DoOverwriteById() bool { //nolint:dupl false positive
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexId(), A.X{u.Id}, u.ToUpdateArray())
	return !L.IsError(err, `Users.DoOverwriteById failed: `+u.SpaceName())
}

// DoUpdateById update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (u *UsersMutator) DoUpdateById() bool { //nolint:dupl false positive
	if !u.HaveMutation() {
		return true
	}
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexId(), A.X{u.Id}, u.mutations)
	return !L.IsError(err, `Users.DoUpdateById failed: `+u.SpaceName())
}

// DoDeletePermanentById permanent delete
func (u *UsersMutator) DoDeletePermanentById() bool { //nolint:dupl false positive
	_, err := u.Adapter.Delete(u.SpaceName(), u.UniqueIndexId(), A.X{u.Id})
	return !L.IsError(err, `Users.DoDeletePermanentById failed: `+u.SpaceName())
}

// func (u *UsersMutator) DoUpsert() bool { //nolint:dupl false positive
//	arr := u.ToArray()
//	_, err := u.Adapter.Upsert(u.SpaceName(), arr, A.X{
//		A.X{`=`, 0, u.Id},
//		A.X{`=`, 1, u.Email},
//		A.X{`=`, 2, u.Password},
//		A.X{`=`, 3, u.CreatedAt},
//		A.X{`=`, 4, u.CreatedBy},
//		A.X{`=`, 5, u.UpdatedAt},
//		A.X{`=`, 6, u.UpdatedBy},
//		A.X{`=`, 7, u.DeletedAt},
//		A.X{`=`, 8, u.PasswordSetAt},
//		A.X{`=`, 9, u.SecretCode},
//		A.X{`=`, 10, u.SecretCodeAt},
//		A.X{`=`, 11, u.VerificationSentAt},
//		A.X{`=`, 12, u.VerifiedAt},
//		A.X{`=`, 13, u.LastLoginAt},
//		A.X{`=`, 14, u.FullName},
//		A.X{`=`, 15, u.TenantCode},
//		A.X{`=`, 16, u.Role},
//		A.X{`=`, 17, u.InvitationState},
//	})
//	return !L.IsError(err, `Users.DoUpsert failed: `+u.SpaceName()+ `\n%#v`, arr)
// }

// DoOverwriteByEmail update all columns, error if not exists, not using mutations/Set*
func (u *UsersMutator) DoOverwriteByEmail() bool { //nolint:dupl false positive
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexEmail(), A.X{u.Email}, u.ToUpdateArray())
	return !L.IsError(err, `Users.DoOverwriteByEmail failed: `+u.SpaceName())
}

// DoUpdateByEmail update only mutated fields, error if not exists, use Find* and Set* methods instead of direct assignment
func (u *UsersMutator) DoUpdateByEmail() bool { //nolint:dupl false positive
	if !u.HaveMutation() {
		return true
	}
	_, err := u.Adapter.Update(u.SpaceName(), u.UniqueIndexEmail(), A.X{u.Email}, u.mutations)
	return !L.IsError(err, `Users.DoUpdateByEmail failed: `+u.SpaceName())
}

// DoDeletePermanentByEmail permanent delete
func (u *UsersMutator) DoDeletePermanentByEmail() bool { //nolint:dupl false positive
	_, err := u.Adapter.Delete(u.SpaceName(), u.UniqueIndexEmail(), A.X{u.Email})
	return !L.IsError(err, `Users.DoDeletePermanentByEmail failed: `+u.SpaceName())
}

// DoInsert insert, error if already exists
func (u *UsersMutator) DoInsert() bool { //nolint:dupl false positive
	arr := u.ToArray()
	row, err := u.Adapter.Insert(u.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			u.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Users.DoInsert failed: `+u.SpaceName()+`\n%#v`, arr)
}

// DoUpsert upsert, insert or overwrite, will error only when there's unique secondary key being violated
// replace = upsert, only error when there's unique secondary key
// previous name: DoReplace
func (u *UsersMutator) DoUpsert() bool { //nolint:dupl false positive
	arr := u.ToArray()
	row, err := u.Adapter.Replace(u.SpaceName(), arr)
	if err == nil {
		tup := row.Tuples()
		if len(tup) > 0 && len(tup[0]) > 0 && tup[0][0] != nil {
			u.Id = X.ToU(tup[0][0])
		}
	}
	return !L.IsError(err, `Users.DoUpsert failed: `+u.SpaceName()+`\n%#v`, arr)
}

// SetId create mutations, should not duplicate
func (u *UsersMutator) SetId(val uint64) bool { //nolint:dupl false positive
	if val != u.Id {
		u.mutations = append(u.mutations, A.X{`=`, 0, val})
		u.logs = append(u.logs, A.X{`id`, u.Id, val})
		u.Id = val
		return true
	}
	return false
}

// SetEmail create mutations, should not duplicate
func (u *UsersMutator) SetEmail(val string) bool { //nolint:dupl false positive
	if val != u.Email {
		u.mutations = append(u.mutations, A.X{`=`, 1, val})
		u.logs = append(u.logs, A.X{`email`, u.Email, val})
		u.Email = val
		return true
	}
	return false
}

// SetPassword create mutations, should not duplicate
func (u *UsersMutator) SetPassword(val string) bool { //nolint:dupl false positive
	if val != u.Password {
		u.mutations = append(u.mutations, A.X{`=`, 2, val})
		u.Password = val
		return true
	}
	return false
}

// SetCreatedAt create mutations, should not duplicate
func (u *UsersMutator) SetCreatedAt(val int64) bool { //nolint:dupl false positive
	if val != u.CreatedAt {
		u.mutations = append(u.mutations, A.X{`=`, 3, val})
		u.logs = append(u.logs, A.X{`createdAt`, u.CreatedAt, val})
		u.CreatedAt = val
		return true
	}
	return false
}

// SetCreatedBy create mutations, should not duplicate
func (u *UsersMutator) SetCreatedBy(val uint64) bool { //nolint:dupl false positive
	if val != u.CreatedBy {
		u.mutations = append(u.mutations, A.X{`=`, 4, val})
		u.logs = append(u.logs, A.X{`createdBy`, u.CreatedBy, val})
		u.CreatedBy = val
		return true
	}
	return false
}

// SetUpdatedAt create mutations, should not duplicate
func (u *UsersMutator) SetUpdatedAt(val int64) bool { //nolint:dupl false positive
	if val != u.UpdatedAt {
		u.mutations = append(u.mutations, A.X{`=`, 5, val})
		u.logs = append(u.logs, A.X{`updatedAt`, u.UpdatedAt, val})
		u.UpdatedAt = val
		return true
	}
	return false
}

// SetUpdatedBy create mutations, should not duplicate
func (u *UsersMutator) SetUpdatedBy(val uint64) bool { //nolint:dupl false positive
	if val != u.UpdatedBy {
		u.mutations = append(u.mutations, A.X{`=`, 6, val})
		u.logs = append(u.logs, A.X{`updatedBy`, u.UpdatedBy, val})
		u.UpdatedBy = val
		return true
	}
	return false
}

// SetDeletedAt create mutations, should not duplicate
func (u *UsersMutator) SetDeletedAt(val int64) bool { //nolint:dupl false positive
	if val != u.DeletedAt {
		u.mutations = append(u.mutations, A.X{`=`, 7, val})
		u.logs = append(u.logs, A.X{`deletedAt`, u.DeletedAt, val})
		u.DeletedAt = val
		return true
	}
	return false
}

// SetPasswordSetAt create mutations, should not duplicate
func (u *UsersMutator) SetPasswordSetAt(val int64) bool { //nolint:dupl false positive
	if val != u.PasswordSetAt {
		u.mutations = append(u.mutations, A.X{`=`, 8, val})
		u.logs = append(u.logs, A.X{`passwordSetAt`, u.PasswordSetAt, val})
		u.PasswordSetAt = val
		return true
	}
	return false
}

// SetSecretCode create mutations, should not duplicate
func (u *UsersMutator) SetSecretCode(val string) bool { //nolint:dupl false positive
	if val != u.SecretCode {
		u.mutations = append(u.mutations, A.X{`=`, 9, val})
		u.SecretCode = val
		return true
	}
	return false
}

// SetSecretCodeAt create mutations, should not duplicate
func (u *UsersMutator) SetSecretCodeAt(val int64) bool { //nolint:dupl false positive
	if val != u.SecretCodeAt {
		u.mutations = append(u.mutations, A.X{`=`, 10, val})
		u.SecretCodeAt = val
		return true
	}
	return false
}

// SetVerificationSentAt create mutations, should not duplicate
func (u *UsersMutator) SetVerificationSentAt(val int64) bool { //nolint:dupl false positive
	if val != u.VerificationSentAt {
		u.mutations = append(u.mutations, A.X{`=`, 11, val})
		u.logs = append(u.logs, A.X{`verificationSentAt`, u.VerificationSentAt, val})
		u.VerificationSentAt = val
		return true
	}
	return false
}

// SetVerifiedAt create mutations, should not duplicate
func (u *UsersMutator) SetVerifiedAt(val int64) bool { //nolint:dupl false positive
	if val != u.VerifiedAt {
		u.mutations = append(u.mutations, A.X{`=`, 12, val})
		u.logs = append(u.logs, A.X{`verifiedAt`, u.VerifiedAt, val})
		u.VerifiedAt = val
		return true
	}
	return false
}

// SetLastLoginAt create mutations, should not duplicate
func (u *UsersMutator) SetLastLoginAt(val int64) bool { //nolint:dupl false positive
	if val != u.LastLoginAt {
		u.mutations = append(u.mutations, A.X{`=`, 13, val})
		u.logs = append(u.logs, A.X{`lastLoginAt`, u.LastLoginAt, val})
		u.LastLoginAt = val
		return true
	}
	return false
}

// SetFullName create mutations, should not duplicate
func (u *UsersMutator) SetFullName(val string) bool { //nolint:dupl false positive
	if val != u.FullName {
		u.mutations = append(u.mutations, A.X{`=`, 14, val})
		u.logs = append(u.logs, A.X{`fullName`, u.FullName, val})
		u.FullName = val
		return true
	}
	return false
}

// SetTenantCode create mutations, should not duplicate
func (u *UsersMutator) SetTenantCode(val string) bool { //nolint:dupl false positive
	if val != u.TenantCode {
		u.mutations = append(u.mutations, A.X{`=`, 15, val})
		u.logs = append(u.logs, A.X{`tenantCode`, u.TenantCode, val})
		u.TenantCode = val
		return true
	}
	return false
}

// SetRole create mutations, should not duplicate
func (u *UsersMutator) SetRole(val string) bool { //nolint:dupl false positive
	if val != u.Role {
		u.mutations = append(u.mutations, A.X{`=`, 16, val})
		u.logs = append(u.logs, A.X{`role`, u.Role, val})
		u.Role = val
		return true
	}
	return false
}

// SetInvitationState create mutations, should not duplicate
func (u *UsersMutator) SetInvitationState(val string) bool { //nolint:dupl false positive
	if val != u.InvitationState {
		u.mutations = append(u.mutations, A.X{`=`, 17, val})
		u.logs = append(u.logs, A.X{`invitationState`, u.InvitationState, val})
		u.InvitationState = val
		return true
	}
	return false
}

// SetAll set all from another source, only if another property is not empty/nil/zero or in forceMap
func (u *UsersMutator) SetAll(from rqAuth.Users, excludeMap, forceMap M.SB) (changed bool) { //nolint:dupl false positive
	if excludeMap == nil { // list of fields to exclude
		excludeMap = M.SB{}
	}
	if forceMap == nil { // list of fields to force overwrite
		forceMap = M.SB{}
	}
	if !excludeMap[`id`] && (forceMap[`id`] || from.Id != 0) {
		u.Id = from.Id
		changed = true
	}
	if !excludeMap[`email`] && (forceMap[`email`] || from.Email != ``) {
		u.Email = S.Trim(from.Email)
		changed = true
	}
	if !excludeMap[`password`] && (forceMap[`password`] || from.Password != ``) {
		u.Password = S.Trim(from.Password)
		changed = true
	}
	if !excludeMap[`createdAt`] && (forceMap[`createdAt`] || from.CreatedAt != 0) {
		u.CreatedAt = from.CreatedAt
		changed = true
	}
	if !excludeMap[`createdBy`] && (forceMap[`createdBy`] || from.CreatedBy != 0) {
		u.CreatedBy = from.CreatedBy
		changed = true
	}
	if !excludeMap[`updatedAt`] && (forceMap[`updatedAt`] || from.UpdatedAt != 0) {
		u.UpdatedAt = from.UpdatedAt
		changed = true
	}
	if !excludeMap[`updatedBy`] && (forceMap[`updatedBy`] || from.UpdatedBy != 0) {
		u.UpdatedBy = from.UpdatedBy
		changed = true
	}
	if !excludeMap[`deletedAt`] && (forceMap[`deletedAt`] || from.DeletedAt != 0) {
		u.DeletedAt = from.DeletedAt
		changed = true
	}
	if !excludeMap[`passwordSetAt`] && (forceMap[`passwordSetAt`] || from.PasswordSetAt != 0) {
		u.PasswordSetAt = from.PasswordSetAt
		changed = true
	}
	if !excludeMap[`secretCode`] && (forceMap[`secretCode`] || from.SecretCode != ``) {
		u.SecretCode = S.Trim(from.SecretCode)
		changed = true
	}
	if !excludeMap[`secretCodeAt`] && (forceMap[`secretCodeAt`] || from.SecretCodeAt != 0) {
		u.SecretCodeAt = from.SecretCodeAt
		changed = true
	}
	if !excludeMap[`verificationSentAt`] && (forceMap[`verificationSentAt`] || from.VerificationSentAt != 0) {
		u.VerificationSentAt = from.VerificationSentAt
		changed = true
	}
	if !excludeMap[`verifiedAt`] && (forceMap[`verifiedAt`] || from.VerifiedAt != 0) {
		u.VerifiedAt = from.VerifiedAt
		changed = true
	}
	if !excludeMap[`lastLoginAt`] && (forceMap[`lastLoginAt`] || from.LastLoginAt != 0) {
		u.LastLoginAt = from.LastLoginAt
		changed = true
	}
	if !excludeMap[`fullName`] && (forceMap[`fullName`] || from.FullName != ``) {
		u.FullName = S.Trim(from.FullName)
		changed = true
	}
	if !excludeMap[`tenantCode`] && (forceMap[`tenantCode`] || from.TenantCode != ``) {
		u.TenantCode = S.Trim(from.TenantCode)
		changed = true
	}
	if !excludeMap[`role`] && (forceMap[`role`] || from.Role != ``) {
		u.Role = S.Trim(from.Role)
		changed = true
	}
	if !excludeMap[`invitationState`] && (forceMap[`invitationState`] || from.InvitationState != ``) {
		u.InvitationState = S.Trim(from.InvitationState)
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
