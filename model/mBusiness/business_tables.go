package mBusiness

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	// Rule type
	RuleTypeFIFO	 	= `fifo`
	RuleTypeLIFO 		= `lifo`
	RuleTypeAVERAGE = `average`

	// Kind type
	KindTypeGOOD 		= `good`
	KindTypeService = `service`
)

const (
	TableProducts Tt.TableName = `products`

	Id		      = `id`
	CreatedAt   = `createdAt`
	CreatedBy   = `createdBy`
	UpdatedAt   = `updatedAt`
	UpdatedBy   = `updatedBy`
	DeletedAt   = `deletedAt`
	DeletedBy		= `deletedBy`
	RestoredBy	= `restoredBy`
	Name				= `name`
	Detail			= `detail`
	Rule				= `rule`
	Kind				= `kind`
	CogsIDR			= `cogsIDR`
)


var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableProducts: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
			{Name, Tt.String},
			{Detail, Tt.String},
			{Rule, Tt.String},
			{Kind, Tt.String},
			{CogsIDR, Tt.String},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
	},
}

const (
	TableActionLogs Ch.TableName = `actionLogs`

	RequestId  = `requestId`
	Error      = `error`
	ActorId    = `actorId`
	IpAddr4    = `ipAddr4`
	IpAddr6    = `ipAddr6`
	UserAgent  = `userAgent`
	Action     = `action`
	Traces     = `traces`
	StatusCode = `statusCode`

	Latency = `latency` // in seconds

	RefId = `refId`
)

var ClickhouseTables = map[Ch.TableName]*Ch.TableProp{
	TableActionLogs: {
		Fields: []Ch.Field{
			{CreatedAt, Ch.DateTime},
			{RequestId, Ch.String},
			{ActorId, Ch.UInt64},
			{Action, Ch.String},
			{StatusCode, Ch.Int16},
			{Traces, Ch.String},
			{Error, Ch.String},
			{IpAddr4, Ch.IPv4},
			{IpAddr6, Ch.IPv6},
			{UserAgent, Ch.String},
			{Latency, Ch.Float64},
			{RefId, Ch.UInt64},
		},
		Orders: []string{CreatedAt, RequestId, ActorId, Action},
	},
}
