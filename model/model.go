package model

import (
	"benakun/model/mBudget"

	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"

	"benakun/model/mAuth"
)

type Migrator struct {
	AuthOltp   *Tt.Adapter
	AuthOlap   *Ch.Adapter
	BudgetOltp *Tt.Adapter
}

func RunMigration(authOltp *Tt.Adapter, authOlap *Ch.Adapter, budgetOltp *Tt.Adapter) {
	Tt.DEBUG = true
	Ch.DEBUG = true
	L.Print(`run migration..`)
	m := Migrator{
		AuthOltp:   authOltp,
		AuthOlap:   authOlap,
		BudgetOltp: budgetOltp,
	}
	m.AuthOltp.MigrateTables(mAuth.TarantoolTables)
	m.AuthOlap.MigrateTables(mAuth.ClickhouseTables)
	m.BudgetOltp.MigrateTables(mBudget.TarantoolTables)
}

// VerifyTables function to check whether tables are there or not
// go run main.go migrate
func VerifyTables(
	authOltp *Tt.Adapter,
	authOlap *Ch.Adapter,
	budgetOltp *Tt.Adapter,
) {
	Ch.CheckClickhouseTables(authOlap, mAuth.ClickhouseTables)
	Tt.CheckTarantoolTables(authOltp, mAuth.TarantoolTables)
	Tt.CheckTarantoolTables(budgetOltp, mBudget.TarantoolTables)
}
