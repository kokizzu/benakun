package domain

import (
	"benakun/model/mFinance/rqFinance"
	"benakun/model/mFinance/wcFinance"
	"benakun/model/zCrud"
	"testing"
	"time"

	"github.com/kpango/fastime"
	"github.com/stretchr/testify/assert"
)

func TestReportViewer(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	t.Run(`insertCoaMustSucceed`, func(t *testing.T) {
		coa := wcFinance.NewCoaMutator(d.AuthOltp)
		coa.SetTenantCode(testSuperAdminTenantCode)
		coa.SetName(`For Report Viewer`)
		coa.SetCreatedAt(fastime.UnixNow())
		coa.SetUpdatedAt(fastime.UnixNow())
		assert.True(t, coa.DoInsert(), `failed to insert coa`)

		t.Run(`insertTransactionJournal`, func(t *testing.T) {
			in := TenantAdminManualJournalIn{
				RequestCommon: testAdminRequestCommon(TenantAdminManualJournalAction),
				Cmd:           zCrud.CmdUpsert,
				CoaId:         coa.Id,
				TransactionJournal: rqFinance.TransactionJournal{
					CreditIDR: 100_000,
					DebitIDR:  50_000,
					Date:      time.Now().Format(time.DateOnly),
				},
			}

			out := d.TenantAdminManualJournal(&in)
			assert.Empty(t, out.Error, `failed to insert transaction journal`)

			t.Run(`getGeneralLedger`, func(t *testing.T) {
				in := ReportViewerGeneralLedgerIn{
					RequestCommon: testAdminRequestCommon(ReportViewerGeneralLedgerAction),
					Coa: rqFinance.Coa{
						Id:         coa.Id,
						TenantCode: testSuperAdminTenantCode,
					},
				}

				out := d.ReportViewerGeneralLedger(&in)
				assert.Empty(t, out.Error, `failed to get general ledger`)
			})
		})
	})
}
