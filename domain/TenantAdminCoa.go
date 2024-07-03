package domain

import (
	"benakun/model/mAuth/wcAuth"
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
	"benakun/model/zCrud"

	"github.com/gofiber/fiber/v2"
	"github.com/kokizzu/gotro/M"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminCoa.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminCoa.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminCoa.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminCoa.go
//go:generate farify doublequote --file TenantAdminCoa.go

type (
	TenantAdminCoaIn struct {
		RequestCommon
		Cmd 				string        `json:"cmd" form:"cmd" query:"cmd" long:"cmd" msg:"cmd"`
		Coa 				rqFinance.Coa	`json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
		MoveToIdx		int    				`json:"moveToIdx" form:"moveToIdx" query:"moveToIdx" long:"moveToIdx" msg:"moveToIdx"`
		ToParentId	uint64 				`json:"toParentId" form:"toParentId" query:"toParentId" long:"toParentId" msg:"toParentId"`
	}
	TenantAdminCoaOut struct {
		ResponseCommon
		Coa		*rqFinance.Coa			`json:"coa" form:"coa" query:"coa" long:"coa" msg:"coa"`
		Coas	*[]rqFinance.Coa	`json:"coas" form:"coas" query:"coas" long:"coas" msg:"coas"`
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
		out.SetError(fiber.StatusBadRequest, ErrTenantAdminCoaUnauthorized)
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

		var parent *wcFinance.CoaMutator
		if in.Coa.ParentId > 0 {
			parent = wcFinance.NewCoaMutator(d.AuthOltp)
			parent.Id = in.Coa.ParentId
			if !parent.FindById() {
				out.SetError(400, ``)
				return
			}
		}

		coa := wcFinance.NewCoaMutator(d.AuthOltp)
		coa.Id = in.Coa.Id
		if coa.Id > 0 {
			if !coa.FindById() {
				out.SetError(400, ``)
				return
			}

			switch in.Cmd {
				case zCrud.CmdUpsert:
					coa.SetAll(in.Coa, M.SB{}, M.SB{})
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
					if parent.Id == in.ToParentId {
						if len(parent.Children) >= 2 {
							children, err := moveChildToIndex(parent.Children, in.Coa.Id, in.MoveToIdx)
							if err != nil {
								out.SetError(400, `` )
								return
							}
				
							parent.SetChildren(children)
							if !parent.DoUpdateById() {
								out.SetError(400, ``)
								return
							}
						}
					} else {
						toParent := wcFinance.NewCoaMutator(d.AuthOltp)
						toParent.Id = in.ToParentId
						if !toParent.FindById() {
							out.SetError(400, ``)
							return
						}

						children := insertChildToIndex(toParent.Children, in.Coa.Id, in.MoveToIdx)

						coa.SetParentId(in.ToParentId)
						if !coa.DoUpdateById() {
							out.SetError(400, ``)
							return
						}

						toParent.SetChildren(children)
						if !toParent.DoUpdateById() {
							out.SetError(400, ``)
							return
						}

						children, err := removeChild(parent.Children, in.Coa.Id)
						if err != nil {
							out.SetError(400, ``)
							return
						}

						parent.SetChildren(children)
						if !parent.DoUpdateById() {
							out.SetError(400, ``)
							return
						}
					}
			}
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
			out.SetError(400, ``)
			return
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
