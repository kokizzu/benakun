package domain

import (
	"benakun/model/mAuth"
	"benakun/model/mAuth/wcAuth"
	"benakun/model/zCrud"
	"testing"

	"github.com/kokizzu/gotro/T"
	"github.com/stretchr/testify/assert"
)

func TestDataEntryTransactionEntry(t *testing.T) {
	d, closer := testDomain()
	defer closer()

	user := wcAuth.NewUsersMutator(d.AuthOltp)
	user.Email = testSuperAdminEmail
	isUserExist := user.FindByEmail()
	assert.True(t, isUserExist, `super admin user does not exist`)

	invState := mAuth.InviteState{
		TenantCode: testSuperAdminTenantCode,
		Role:       mAuth.RoleDataEntry,
		State:      mAuth.InvitationStateInvited,
		Date:       T.DateStr(),
	}
	user.SetInvitationState(invState.ToStateString())
	isUserUpdated := user.DoUpdateByEmail()
	assert.True(t, isUserUpdated, `failed to update super admin user`)

	t.Run(`insertMustSucceed`, func(t *testing.T) {
		in := DataEntryTransactionEntryIn{
			RequestCommon: testAdminRequestCommon(DataEntryTransactionEntryAction),
			Cmd:           zCrud.CmdUpsert,
		}

		out := d.DataEntryTransactionEntry(&in)
		assert.Empty(t, out.Error, `failed to insert transaction entry`)
	})
}
