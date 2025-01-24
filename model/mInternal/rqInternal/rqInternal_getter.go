package rqInternal

import (
	"github.com/kokizzu/gotro/I"
	"github.com/kokizzu/gotro/X"
)

func (ip *InvoicePayment) FindInvoicesByUserId(userId uint64) (res []InvoicePayment) {
	const comment = `-- InvoicePayment) FindInvoicesByUserId`

	whereAndSql := ` WHERE ` + ip.SqlUserId() + ` = ` + I.UToS(userId)

	queryRows := comment + `
SELECT ` + ip.SqlSelectAllFields() + `
FROM /* SEQSCAN */ ` + ip.SqlTableName() + whereAndSql

	ip.Adapter.QuerySql(queryRows, func(row []any) {
		row[0] = X.ToS(row[0])
		res = append(res, *ip.FromArray(row))
	})

	return
}
