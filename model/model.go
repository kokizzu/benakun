package model

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/L"
	"github.com/kokizzu/gotro/W2/example/model/mAuth"
)

type Migrator struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter
}

func RunMigration(authOltp *Tt.Adapter, authOlap *Ch.Adapter, propOltp *Tt.Adapter, propOlap *Ch.Adapter, storOltp *Tt.Adapter) {
	Tt.DEBUG = true
	Ch.DEBUG = true
	L.Print(`run migration..`)
	m := Migrator{
		AuthOltp: authOltp,
		AuthOlap: authOlap,
	}
	m.AuthOltp.MigrateTables(mAuth.TarantoolTables)
	m.AuthOlap.MigrateTables(mAuth.ClickhouseTables)
}

// VerifyTables function to check whether tables are there or not
// go run main.go migrate
func VerifyTables(
	authOltp *Tt.Adapter,
	authOlap *Ch.Adapter,
) {
	Ch.CheckClickhouseTables(authOlap, mAuth.ClickhouseTables)
	Tt.CheckTarantoolTables(authOltp, mAuth.TarantoolTables)
}
