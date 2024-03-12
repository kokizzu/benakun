package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"

	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminOrganization.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminOrganization.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminOrganization.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminOrganization.go
//go:generate farify doublequote --file TenantAdminOrganization.go

type (
	TenantAdminOrganizationIn struct {
		RequestCommon
		Cmd         string       `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		User        rqAuth.Users `json:"user" form:"user" query:"user" long:"user" msg:"user"`
		Id          uint64       `json:"id" form:"id" query:"id" long:"id" msg:"id"`
		NewParentId uint64       `json:"newParentId" form:"newParentId" query:"newParentId" long:"newParentId" msg:"newParentId"`
		OldParentId uint64       `json:"oldParentId" form:"oldParentId" query:"oldParentId" long:"oldParentId" msg:"oldParentId"`
	}

	TenantAdminOrganizationOut struct {
		ResponseCommon
		Orgs []rqAuth.Orgs
	}
)

const (
	TenantAdminOrganizationAction = `tenantAdmin/organization`

	ErrTenantAdminOrganizationUnauthorized                 = `unauthorized user`
	ErrTenantAdminOrganizationTenantNotFound               = `tenant admin not found`
	ErrTenantAdminOrganizationOldParentNotFound            = `old parent of the organization cannot be found`
	ErrTenantAdminOrganizationNewParentNotFound            = `new parent of the organization cannot be found`
	ErrTenantAdminOrganizationUpdatedOldParentChildren     = `cannot update childs of old parent`
	ErrTenantAdminOrganizationUpdatedNewParentChildren     = `cannot update childs of new parent`
	ErrTenantAdminOrganizationUpdatedParentForOrganization = `cannot update parent for organization`
	ErrTenantAdminOrganizationNotFound                     = `cannot find the organization`
	ErrTenantAdminOrganizationChildsNotFound               = `cannot find childs for update`
	ErrTenantAdminOrganizationUpdatedChilds                = `cannot update childs`
)

func (d *Domain) TenantAdminOrganization(in *TenantAdminOrganizationIn) (out TenantAdminOrganizationOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := rqAuth.NewUsers(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminOrganizationUnauthorized)
		return
	}

	tenant := rqAuth.NewTenants(d.AuthOltp)
	tenant.TenantCode = user.TenantCode
	if !tenant.FindByTenantCode() {
		out.SetError(400, ErrTenantAdminDashboardTenantNotFound)
		return
	}

	org := rqAuth.NewOrgs(d.AuthOltp)
	out.Orgs = org.FindOrgsByTenantCode(tenant.TenantCode)

	switch in.Cmd {
	case zCrud.CmdUpsert:
		orgW := wcAuth.NewOrgsMutator(d.AuthOltp)
		isFamilyWithNewParent := isFamily(in.Id, in.NewParentId, out.Orgs)
		currentOrg := findLocalOrgById(out.Orgs, in.Id)

		if in.NewParentId == in.OldParentId {
			return
		}

		// add current org as child of new parent
		if in.NewParentId > 0 {
			org.Id = in.NewParentId
			if !org.FindById() {
				out.SetError(400, ErrTenantAdminOrganizationNewParentNotFound)
				return
			}

			if !checkChildExistById(org.Children, in.Id) {
				orgW.SetAll(*org, M.SB{}, M.SB{})
				orgW.SetChildren(append(org.Children, in.Id))
				if orgW.HaveMutation() {
					orgW.SetUpdatedAt(in.UnixNow())
					orgW.SetUpdatedBy(sess.UserId)
					if orgW.Id == 0 {
						orgW.SetCreatedAt(in.UnixNow())
					}
				}
				if !orgW.DoUpsert() {
					out.SetError(500, ErrTenantAdminOrganizationUpdatedNewParentChildren)
					return
				}

			}
		}

		// remove current org from child of old parent
		if in.OldParentId > 0 {
			org.Id = in.OldParentId
			if !org.FindById() {
				out.SetError(400, ErrTenantAdminOrganizationOldParentNotFound)
				return
			}

			// if new parent is family of current org
			if isFamilyWithNewParent {
				org.Children = append(org.Children, currentOrg.Children...)
			}

			childs := removeAChildById(org.Children, in.Id)

			orgW.SetAll(*org, M.SB{}, M.SB{})
			orgW.SetParentId(org.ParentId)
			orgW.SetChildren(childs)
			if orgW.HaveMutation() {
				orgW.SetUpdatedAt(in.UnixNow())
				orgW.SetUpdatedBy(sess.UserId)
				if orgW.Id == 0 {
					orgW.SetCreatedAt(in.UnixNow())
				}
			}
			if !orgW.DoUpsert() {
				out.SetError(500, ErrTenantAdminOrganizationUpdatedOldParentChildren)
				return
			}

		} else {
			// if new parent is family of current org
			if isFamilyWithNewParent {
				for _, childId := range currentOrg.Children {
					org.Id = childId.(uint64)
					if !org.FindById() {
						out.SetError(400, ErrTenantAdminOrganizationChildsNotFound)
						return
					}

					orgW.SetAll(*org, M.SB{}, M.SB{})
					orgW.SetParentId(0)
					if orgW.HaveMutation() {
						orgW.SetUpdatedAt(in.UnixNow())
						orgW.SetUpdatedBy(sess.UserId)
						if orgW.Id == 0 {
							orgW.SetCreatedAt(in.UnixNow())
						}
					}
					if !orgW.DoUpsert() {
						out.SetError(500, ErrTenantAdminOrganizationUpdatedChilds)
						return
					}
				}
			}
		}

		// update parent id of current org
		org.Id = in.Id
		if !org.FindById() {
			out.SetError(400, ErrTenantAdminOrganizationNotFound)
			return
		}

		orgW.SetAll(*org, M.SB{}, M.SB{})
		orgW.SetParentId(in.NewParentId)

		if isFamilyWithNewParent {
			orgW.SetChildren([]any{})
		}

		if orgW.HaveMutation() {
			orgW.SetUpdatedAt(in.UnixNow())
			orgW.SetUpdatedBy(sess.UserId)
			if orgW.Id == 0 {
				orgW.SetCreatedAt(in.UnixNow())
			}
		}

		if !orgW.DoUpsert() {
			out.SetError(500, ErrTenantAdminOrganizationUpdatedParentForOrganization)
			return
		}

		out.Orgs = org.FindOrgsByTenantCode(tenant.TenantCode)

	}
	return
}

// remove a child id
func removeAChildById(childs []any, removedChildId uint64) (res []any) {
	for _, child := range childs {
		if child.(uint64) != removedChildId {
			res = append(res, child)
		}
	}

	if len(res) < 1 {
		return []any{}
	}

	return
}

// check child exist
func checkChildExistById(childs []any, childId uint64) bool {
	for _, child := range childs {
		if child.(uint64) == childId {
			return true
		}
	}
	return false
}

// find local org by id
func findLocalOrgById(orgs []rqAuth.Orgs, orgId uint64) (res rqAuth.Orgs) {
	for _, org := range orgs {
		if org.Id == orgId {
			res = org
			return
		}
	}
	return
}

// check family
func isFamily(orgId uint64, newParentId uint64, orgs []rqAuth.Orgs) bool {
	orgEntry := findLocalOrgById(orgs, orgId)
	for _, childId := range orgEntry.Children {
		if childId.(uint64) == newParentId || isFamily(childId.(uint64), newParentId, orgs) {
			return true
		}
	}
	return false
}
