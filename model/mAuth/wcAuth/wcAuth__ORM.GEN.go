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

// OrgsMutator DAO writer/command struct
//
//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type wcAuth__ORM.GEN.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type wcAuth__ORM.GEN.go
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
//		A.X{`=`, 12, o.DeletedBy},
//		A.X{`=`, 13, o.RestoredBy},
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

// SetDeletedBy create mutations, should not duplicate
func (o *OrgsMutator) SetDeletedBy(val uint64) bool { //nolint:dupl false positive
	if val != o.DeletedBy {
		o.mutations = append(o.mutations, A.X{`=`, 12, val})
		o.logs = append(o.logs, A.X{`deletedBy`, o.DeletedBy, val})
		o.DeletedBy = val
		return true
	}
	return false
}

// SetRestoredBy create mutations, should not duplicate
func (o *OrgsMutator) SetRestoredBy(val uint64) bool { //nolint:dupl false positive
	if val != o.RestoredBy {
		o.mutations = append(o.mutations, A.X{`=`, 13, val})
		o.logs = append(o.logs, A.X{`restoredBy`, o.RestoredBy, val})
		o.RestoredBy = val
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
	if !excludeMap[`deletedBy`] && (forceMap[`deletedBy`] || from.DeletedBy != 0) {
		o.DeletedBy = from.DeletedBy
		changed = true
	}
	if !excludeMap[`restoredBy`] && (forceMap[`restoredBy`] || from.RestoredBy != 0) {
		o.RestoredBy = from.RestoredBy
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
//		A.X{`=`, 7, t.ProductsCoaId},
//		A.X{`=`, 8, t.SuppliersCoaId},
//		A.X{`=`, 9, t.CustomersCoaId},
//		A.X{`=`, 10, t.StaffsCoaId},
//		A.X{`=`, 11, t.BanksCoaId},
//		A.X{`=`, 12, t.CustomerReceivablesCoaId},
//		A.X{`=`, 13, t.FundersCoaId},
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

// SetProductsCoaId create mutations, should not duplicate
func (t *TenantsMutator) SetProductsCoaId(val uint64) bool { //nolint:dupl false positive
	if val != t.ProductsCoaId {
		t.mutations = append(t.mutations, A.X{`=`, 7, val})
		t.logs = append(t.logs, A.X{`productsCoaId`, t.ProductsCoaId, val})
		t.ProductsCoaId = val
		return true
	}
	return false
}

// SetSuppliersCoaId create mutations, should not duplicate
func (t *TenantsMutator) SetSuppliersCoaId(val uint64) bool { //nolint:dupl false positive
	if val != t.SuppliersCoaId {
		t.mutations = append(t.mutations, A.X{`=`, 8, val})
		t.logs = append(t.logs, A.X{`suppliersCoaId`, t.SuppliersCoaId, val})
		t.SuppliersCoaId = val
		return true
	}
	return false
}

// SetCustomersCoaId create mutations, should not duplicate
func (t *TenantsMutator) SetCustomersCoaId(val uint64) bool { //nolint:dupl false positive
	if val != t.CustomersCoaId {
		t.mutations = append(t.mutations, A.X{`=`, 9, val})
		t.logs = append(t.logs, A.X{`customersCoaId`, t.CustomersCoaId, val})
		t.CustomersCoaId = val
		return true
	}
	return false
}

// SetStaffsCoaId create mutations, should not duplicate
func (t *TenantsMutator) SetStaffsCoaId(val uint64) bool { //nolint:dupl false positive
	if val != t.StaffsCoaId {
		t.mutations = append(t.mutations, A.X{`=`, 10, val})
		t.logs = append(t.logs, A.X{`staffsCoaId`, t.StaffsCoaId, val})
		t.StaffsCoaId = val
		return true
	}
	return false
}

// SetBanksCoaId create mutations, should not duplicate
func (t *TenantsMutator) SetBanksCoaId(val uint64) bool { //nolint:dupl false positive
	if val != t.BanksCoaId {
		t.mutations = append(t.mutations, A.X{`=`, 11, val})
		t.logs = append(t.logs, A.X{`banksCoaId`, t.BanksCoaId, val})
		t.BanksCoaId = val
		return true
	}
	return false
}

// SetCustomerReceivablesCoaId create mutations, should not duplicate
func (t *TenantsMutator) SetCustomerReceivablesCoaId(val uint64) bool { //nolint:dupl false positive
	if val != t.CustomerReceivablesCoaId {
		t.mutations = append(t.mutations, A.X{`=`, 12, val})
		t.logs = append(t.logs, A.X{`customerReceivablesCoaId`, t.CustomerReceivablesCoaId, val})
		t.CustomerReceivablesCoaId = val
		return true
	}
	return false
}

// SetFundersCoaId create mutations, should not duplicate
func (t *TenantsMutator) SetFundersCoaId(val uint64) bool { //nolint:dupl false positive
	if val != t.FundersCoaId {
		t.mutations = append(t.mutations, A.X{`=`, 13, val})
		t.logs = append(t.logs, A.X{`fundersCoaId`, t.FundersCoaId, val})
		t.FundersCoaId = val
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
	if !excludeMap[`productsCoaId`] && (forceMap[`productsCoaId`] || from.ProductsCoaId != 0) {
		t.ProductsCoaId = from.ProductsCoaId
		changed = true
	}
	if !excludeMap[`suppliersCoaId`] && (forceMap[`suppliersCoaId`] || from.SuppliersCoaId != 0) {
		t.SuppliersCoaId = from.SuppliersCoaId
		changed = true
	}
	if !excludeMap[`customersCoaId`] && (forceMap[`customersCoaId`] || from.CustomersCoaId != 0) {
		t.CustomersCoaId = from.CustomersCoaId
		changed = true
	}
	if !excludeMap[`staffsCoaId`] && (forceMap[`staffsCoaId`] || from.StaffsCoaId != 0) {
		t.StaffsCoaId = from.StaffsCoaId
		changed = true
	}
	if !excludeMap[`banksCoaId`] && (forceMap[`banksCoaId`] || from.BanksCoaId != 0) {
		t.BanksCoaId = from.BanksCoaId
		changed = true
	}
	if !excludeMap[`customerReceivablesCoaId`] && (forceMap[`customerReceivablesCoaId`] || from.CustomerReceivablesCoaId != 0) {
		t.CustomerReceivablesCoaId = from.CustomerReceivablesCoaId
		changed = true
	}
	if !excludeMap[`fundersCoaId`] && (forceMap[`fundersCoaId`] || from.FundersCoaId != 0) {
		t.FundersCoaId = from.FundersCoaId
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
//		A.X{`=`, 18, u.SupportExpiredAt},
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

// SetSupportExpiredAt create mutations, should not duplicate
func (u *UsersMutator) SetSupportExpiredAt(val int64) bool { //nolint:dupl false positive
	if val != u.SupportExpiredAt {
		u.mutations = append(u.mutations, A.X{`=`, 18, val})
		u.logs = append(u.logs, A.X{`supportExpiredAt`, u.SupportExpiredAt, val})
		u.SupportExpiredAt = val
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
	if !excludeMap[`supportExpiredAt`] && (forceMap[`supportExpiredAt`] || from.SupportExpiredAt != 0) {
		u.SupportExpiredAt = from.SupportExpiredAt
		changed = true
	}
	return
}

// DO NOT EDIT, will be overwritten by github.com/kokizzu/D/Tt/tarantool_orm_generator.go
