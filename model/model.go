package model

import (
	"benakun/model/mBudget"
	"benakun/model/mBusiness"
	"benakun/model/mFinance"

	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"

	"benakun/model/mAuth"
)

type Migrator struct {
	AuthOltp     *Tt.Adapter
	AuthOlap     *Ch.Adapter
	BudgetOltp   *Tt.Adapter
	BusinessOltp *Tt.Adapter
	FinanceOltp  *Tt.Adapter
	InternalOltp *Tt.Adapter
}

func RunMigration(
	authOltp *Tt.Adapter,
	authOlap *Ch.Adapter,
	budgetOltp *Tt.Adapter,
	businessOltp *Tt.Adapter,
	financeOltp *Tt.Adapter,
	internalOltp *Tt.Adapter,
) {
	Tt.DEBUG = true
	Ch.DEBUG = true
	L.Print(`run migration..`)
	m := Migrator{
		AuthOltp:     authOltp,
		AuthOlap:     authOlap,
		BudgetOltp:   budgetOltp,
		BusinessOltp: businessOltp,
		FinanceOltp:  financeOltp,
		InternalOltp: internalOltp,
	}
	m.AuthOltp.MigrateTables(mAuth.TarantoolTables)
	m.AuthOlap.MigrateTables(mAuth.ClickhouseTables)
	m.BudgetOltp.MigrateTables(mBudget.TarantoolTables)
	m.BusinessOltp.MigrateTables(mBusiness.TarantoolTables)
	m.FinanceOltp.MigrateTables(mFinance.TarantoolTables)
	m.InternalOltp.MigrateTables(mAuth.TarantoolTables)
}

// VerifyTables function to check whether tables are there or not
// go run main.go migrate
func VerifyTables(
	authOltp *Tt.Adapter,
	authOlap *Ch.Adapter,
	budgetOltp *Tt.Adapter,
	businessOltp *Tt.Adapter,
	financeOltp *Tt.Adapter,
	internalOltp *Tt.Adapter,
) {
	Ch.CheckClickhouseTables(authOlap, mAuth.ClickhouseTables)
	Tt.CheckTarantoolTables(authOltp, mAuth.TarantoolTables)
	Tt.CheckTarantoolTables(budgetOltp, mBudget.TarantoolTables)
	Tt.CheckTarantoolTables(businessOltp, mBusiness.TarantoolTables)
	Tt.CheckTarantoolTables(financeOltp, mBudget.TarantoolTables)
	Tt.CheckTarantoolTables(internalOltp, mAuth.TarantoolTables)
}
