package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"

	"github.com/gofiber/fiber/v2"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminOrganization.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminOrganization.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminOrganization.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminOrganization.go
//go:generate farify doublequote --file TenantAdminOrganization.go

type (
	TenantAdminOrganizationIn struct {
		RequestCommon
		Cmd        string      `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		Org        rqAuth.Orgs `json:"org" form:"org" query:"org" long:"org" msg:"org"`
		MoveToIdx  int         `json:"moveToIdx" form:"moveToIdx" query:"moveToIdx" long:"moveToIdx" msg:"moveToIdx"`
		ToParentId uint64      `json:"toParentId" form:"toParentId" query:"toParentId" long:"toParentId" msg:"toParentId"`
	}

	TenantAdminOrganizationOut struct {
		ResponseCommon
		Org  *rqAuth.Orgs   `json:"org" form:"org" query:"org" long:"org" msg:"org"`
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
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

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(fiber.StatusBadRequest, ``)
		return
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore, zCrud.CmdMove:
		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.TenantCode = user.TenantCode
		if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
			out.SetError(400, ``)
			return
		}

		var parent *wcAuth.OrgsMutator
		if in.Org.ParentId > 0 {
			parent = wcAuth.NewOrgsMutator(d.AuthOltp)
			parent.Id = in.Org.ParentId
			if !parent.FindById() {
				out.SetError(400, ``)
				return
			}
		}

		org := wcAuth.NewOrgsMutator(d.AuthOltp)
		org.Id = in.Org.Id
		if org.Id > 0 {
			if !org.FindById() {
				out.SetError(400, ``)
				return
			}

			switch in.Cmd {
			case zCrud.CmdUpsert:
				if in.Org.HeadTitle != org.HeadTitle {
					org.SetHeadTitle(in.Org.HeadTitle)
				}
				if in.Org.Name != org.Name {
					org.SetName(in.Org.Name)
				}
			case zCrud.CmdDelete:
				if org.DeletedAt == 0 {
					org.SetDeletedAt(in.UnixNow())
					org.SetDeletedBy(sess.UserId)
				}
			case zCrud.CmdRestore:
				if org.DeletedAt > 0 {
					org.SetDeletedAt(0)
					org.SetRestoredBy(sess.UserId)
				}
			case zCrud.CmdMove:
				if in.ToParentId == parent.Id {
					if len(parent.Children) >= 2 {
						children, err := moveChildToIndex(parent.Children, org.Id, in.MoveToIdx)
						if err != nil {
							out.SetError(400, err.Error())
							return
						}
						parent.SetChildren(children)
						if !parent.DoUpsertById() {
							out.SetError(400, ``)
							return
						}
					}
				} else {
					toParent := wcAuth.NewOrgsMutator(d.AuthOltp)
					toParent.Id = in.ToParentId
					if !toParent.FindById() {
						out.SetError(400, ``)
						return
					}

					children := insertChildToIndex(toParent.Children, org.Id, in.MoveToIdx)

					org.SetParentId(toParent.Id)
					if !org.DoUpdateById() {
						out.SetError(400, ``)
						return
					}

					children, err := removeChild(parent.Children, org.Id)
					if err != nil {
						out.SetError(400, err.Error())
						return
					}
					parent.SetChildren(children)
					if !parent.DoUpdateById() {
						out.SetError(400, ``)
						return
					}
				}
			}
		} else {
			if in.Org.HeadTitle != `` {
				org.SetHeadTitle(in.Org.HeadTitle)
			}
			if in.Org.Name != `` {
				org.SetName(in.Org.Name)
			}

			org.SetTenantCode(tenant.TenantCode)
		}

		if org.HaveMutation() {
			org.SetUpdatedAt(in.UnixNow())
			org.SetUpdatedBy(sess.UserId)
			if org.Id == 0 {
				org.SetCreatedAt(in.UnixNow())
				org.SetCreatedBy(sess.UserId)
			}
		}

		if !org.DoUpsertById() {
			out.SetError(400, ``)
			return
		}

		out.Org = &org.Orgs

		fallthrough
	case zCrud.CmdList:
		org := wcAuth.NewOrgsMutator(d.AuthOltp)
		org.TenantCode = user.TenantCode
		orgs := org.FindOrgsByTenant()
		out.Orgs = &orgs
	}

	return
}
