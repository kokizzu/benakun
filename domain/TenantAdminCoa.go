package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
	"benakun/model/zCrud"

	"github.com/kokizzu/gotro/L"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminCoa.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminCoa.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminCoa.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminCoa.go
//go:generate farify doublequote --file TenantAdminCoa.go

type (
	TenantAdminCoaIn struct {
		RequestCommon
		Cmd        string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		Coa        rqFinance.Coa `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
		MoveToIdx  int           `json:"moveToIdx" form:"moveToIdx" query:"moveToIdx" long:"moveToIdx" msg:"moveToIdx"`
		ToParentId uint64        `json:"toParentId" form:"toParentId" query:"toParentId" long:"toParentId" msg:"toParentId"`
	}
	TenantAdminCoaOut struct {
		ResponseCommon
		Coa  *rqFinance.Coa   `json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
		Coas *[]rqFinance.Coa `json:"coas" form:"coas" query:"coas" long:"coas" msg:"coas"`
	}
)

const (
	TenantAdminCoaAction = `tenantAdmin/coa`

	ErrTenantAdminCoaUnauthorized   = `unauthorized user for coa`
	ErrTenantAdminCoaTenantNotFound = `tenant admin not found for coa`
)

func (d *Domain) TenantAdminCoa(in *TenantAdminCoaIn) (out TenantAdminCoaOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustLogin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Id = sess.UserId
	if !user.FindById() {
		out.SetError(400, ErrTenantAdminCoaUnauthorized)
		return
	}

	switch in.Cmd {
	case zCrud.CmdForm:
	case zCrud.CmdUpsert, zCrud.CmdDelete, zCrud.CmdRestore, zCrud.CmdMove:
		tenant := wcAuth.NewTenantsMutator(d.AuthOltp)
		tenant.TenantCode = user.TenantCode
		if !tenant.FindByTenantCode() && !sess.IsSuperAdmin {
			out.SetError(400, ErrTenantAdminCoaTenantNotFound)
			return
		}

		coa := wcFinance.NewCoaMutator(d.AuthOltp)
		coa.Id = in.Coa.Id

		var parent *wcFinance.CoaMutator

		if in.Coa.ParentId > 0 {
			parent = wcFinance.NewCoaMutator(d.AuthOltp)
			parent.Id = in.Coa.ParentId
			if !parent.FindById() {
				out.SetError(400, `coa parent not found`)
				return
			}
		}

		// If coa id > 0
		// update coa
		if coa.Id > 0 {
			if !coa.FindById() {
				out.SetError(400, `coa not found`)
				return
			}

			switch in.Cmd {
			case zCrud.CmdUpsert:
				if in.Coa.Name != coa.Name {
					coa.SetName(in.Coa.Name)
				}

				if in.Coa.Label != coa.Label {
					coa.SetLabel(in.Coa.Label)
				}
			case zCrud.CmdDelete:
				if coa.DeletedAt == 0 {
					coa.SetDeletedAt(in.UnixNow())
					coa.SetDeletedBy(sess.UserId)
				}
			case zCrud.CmdRestore:
				if coa.DeletedAt > 0 {
					coa.SetDeletedAt(0)
					coa.SetRestoredBy(sess.UserId)
				}
			case zCrud.CmdMove:
				if in.Coa.ParentId == 0 || parent == nil {
					out.SetError(400, `must include parent id`)
					return
				}
				if parent.Id == in.ToParentId {
					if len(parent.Children) >= 2 {
						children, err := moveChildToIndex(parent.Children, coa.Id, in.MoveToIdx)
						if err != nil {
							L.Print(err)
							out.SetError(400, `child not found in parent`)
							return
						}

						parent.SetChildren(children)
						if !parent.DoUpdateById() {
							out.SetError(400, `failed reorder children`)
							return
						}
					}
				} else {
					toParent := wcFinance.NewCoaMutator(d.AuthOltp)
					toParent.Id = in.ToParentId
					if !toParent.FindById() {
						out.SetError(400, `cannot find parent to move coa`)
						return
					}

					children := insertChildToIndex(toParent.Children, coa.Id, in.MoveToIdx)

					coa.SetParentId(toParent.Id)
					if !coa.DoUpdateById() {
						out.SetError(400, `child not found in target parent`)
						return
					}

					toParent.SetChildren(children)
					if !toParent.DoUpdateById() {
						out.SetError(400, `failed to update destination parent`)
						return
					}

					children, err := removeChild(parent.Children, in.Coa.Id)
					if err != nil {
						out.SetError(400, `failed to remove child from source parent`)
						return
					}

					parent.SetChildren(children)
					if !parent.DoUpdateById() {
						out.SetError(400, `failed to update source parent`)
						return
					}
				}
			}
		} else {
			// If coa id == 0
			// create a new coa

			if in.Coa.Name != `` {
				coa.SetName(in.Coa.Name)
			}

			if in.Coa.Label != `` {
				coa.SetLabel(in.Coa.Label)
			}

			coa.SetTenantCode(tenant.TenantCode)
		}

		if coa.HaveMutation() {
			coa.SetUpdatedAt(in.UnixNow())
			coa.SetUpdatedBy(sess.UserId)
			if coa.Id == 0 {
				coa.SetCreatedAt(in.UnixNow())
				coa.SetCreatedBy(sess.UserId)
			}
		}

		if !coa.DoUpsertById() {
			out.SetError(400, `failed to update coa`)
			return
		}

		if in.Coa.Id == 0 {
			if in.Coa.ParentId > 0 {
				children := parent.Children
				children = append(children, coa.Id)
				parent.SetChildren(children)
				parent.SetUpdatedAt(in.UnixNow())
				parent.SetUpdatedBy(sess.UserId)
				if !parent.DoUpdateById() {
					out.SetError(400, `failed to update parent of new coa`)
					return
				}

				coa.SetParentId(parent.Id)
				if !coa.DoUpsertById() {
					out.SetError(400, `failed to create a new coa`)
					return
				}
			}
		}

		out.Coa = &coa.Coa

		fallthrough
	case zCrud.CmdList:
		coa := wcFinance.NewCoaMutator(d.AuthOltp)
		coa.TenantCode = user.TenantCode
		coas := coa.FindCoasByTenant()
		out.Coas = &coas
	}

	return
}
