package domain

import (
	"benakun/model/mFinance/rqFinance"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file FieldSupervisorBusinessTransactionEdit.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type FieldSupervisorBusinessTransactionEdit.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type FieldSupervisorBusinessTransactionEdit.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type FieldSupervisorBusinessTransactionEdit.go
//go:generate farify doublequote --file FieldSupervisorBusinessTransactionEdit.go

type (
	FieldSupervisorBusinessTransactionEditIn struct {
		RequestCommon
	}
	FieldSupervisorBusinessTransactionEditOut struct {
		ResponseCommon
		TransactionJournals []rqFinance.TransactionJournal `json:"transactionJournals" form:"transactionJournals" query:"transactionJournals" long:"transactionJournals" msg:"transactionJournals"`
	}
)

const (
	FieldSupervisorBusinessTransactionEditAction = `fieldSupervisor/businessTransaction/`

	ErrFieldSupervisorBusinessTransactionEditUnauthorized = `you are unauthorized as field supervisor`
	ErrFieldSupervisorBusinessTransactionEditNotFieldSupervisor = `must be field supervisor to do this operation`
)

func (d *Domain) FieldSupervisorBusinessTransactionEdit(in *FieldSupervisorBusinessTransactionEditIn) (out FieldSupervisorBusinessTransactionEditOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustFieldSupervisor(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		out.SetError(400, ErrFieldSupervisorBusinessTransactionEditUnauthorized)
		return
	}

	if !sess.IsFieldSupervisor {
		out.SetError(400, ErrFieldSupervisorBusinessTransactionEditNotFieldSupervisor)
		return
	}

	return
}
