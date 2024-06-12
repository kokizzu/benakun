package mFinance

import "github.com/kokizzu/gotro/D/Tt"

const (
	Id                 = `id`
	CreatedAt          = `createdAt`
	CreatedBy          = `createdBy`
	UpdatedAt          = `updatedAt`
	UpdatedBy          = `updatedBy`
	DeletedAt          = `deletedAt`
	TenantCode         = `tenantCode`
)

const (
	TableCoa Tt.TableName = `coa`
	Name 									= `name`
	Level                 = `level`
	ParentId  						= `parentId`
	Children 	 						= `children`
)

const (
	CoaLevel1Name = `Aktiva`
	CoaLevel2Name = `Kewajiban`
	CoaLevel3Name = `Ekuitas`
	CoaLevel4Name = `Pendapatan`
	CoaLevel5Name = `Beban`
	CoaLevel6Name = `Pendapatan Lain-lain`
	CoaLevel7Name = `Beban Lain-lain`

	CoaLevel1ChildName1 = `Aktiva Lancar`
	CoaLevel1ChildName2 = `Aktiva Tetap`
	CoaLevel1ChildName3 = `Aktiva Tak Berwujud`
)

type CoaLevelDefault struct {
	Name          string
	ChildrenNames []string
}

var CoaLevelDefaultList = map[string]CoaLevelDefault{
	`1`: {
		Name: CoaLevel1Name,
		ChildrenNames: []string{
			CoaLevel1ChildName1,
			CoaLevel1ChildName2,
			CoaLevel1ChildName3,
		},
	},
	`2`: {
		Name:          CoaLevel2Name,
		ChildrenNames: []string{},
	},
	`3`: {
		Name:          CoaLevel3Name,
		ChildrenNames: []string{},
	},
	`4`: {
		Name:          CoaLevel4Name,
		ChildrenNames: []string{},
	},
	`5`: {
		Name:          CoaLevel5Name,
		ChildrenNames: []string{},
	},
	`6`: {
		Name:          CoaLevel6Name,
		ChildrenNames: []string{},
	},
	`7`: {
		Name:          CoaLevel7Name,
		ChildrenNames: []string{},
	},
}

const (
	TableTransactions Tt.TableName = `transactions`

	CompletedAt 	= `completedAt`
	CoaId 				= `coaId`
	Price 				= `price`
	Description 	= `descriptions`
	Qty 					= `qty`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableCoa: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			{Name, Tt.String},
			{Level, Tt.Double},
			{ParentId, Tt.Unsigned},
			{Children, Tt.Array},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
		},
		AutoIncrementId: true,
	},
	TableTransactions: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{CompletedAt, Tt.Integer},
			{Price, Tt.Unsigned},
			{Description, Tt.String},
			{Qty, Tt.Integer},
		},
	},
}