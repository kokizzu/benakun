package domain

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file DataEntryTransactionEntry.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type DataEntryTransactionEntry.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type DataEntryTransactionEntry.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type DataEntryTransactionEntry.go
//go:generate farify doublequote --file DataEntryTransactionEntry.go

type (
	DataEntryTransactionEntryIn struct {
		RequestCommon
	}
	DataEntryTransactionEntryOut struct {
		ResponseCommon
	}
)

const (
	DataEntryTransactionEntryAction = `dataEntry/transactionEntry`
)

func (d *Domain) DataEntryTransactionEntry(in *DataEntryTransactionEntryIn) (out DataEntryTransactionEntryOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)
	return
}
