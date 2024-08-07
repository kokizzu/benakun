package mBusiness

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	// Rule type
	RuleTypeFIFO	 	= `fifo`		// First in First Out
	RuleTypeLIFO 		= `lifo`		// Life in First Out
	RuleTypeAVERAGE = `average` // Rata - rata dari struk kita

	// Kind type
	KindTypeGOODS 	= `goods`
	KindTypeService = `service`
)

func IsValidProductRule(rule string) bool {
	switch rule {
	case RuleTypeFIFO:
		return true
	case RuleTypeLIFO:
		return true
	case RuleTypeAVERAGE:
		return true
	default:
		return false
	}
}

func IsValidProductKind(kind string) bool {
	switch kind {
	case KindTypeGOODS:
		return true
	case KindTypeService:
		return true
	default:
		return false
	}
}

const (
	TableProducts Tt.TableName = `products`

	Id		      			= `id`
	TenantCode				= `tenantCode`
	CreatedAt   			= `createdAt`
	CreatedBy   			= `createdBy`
	UpdatedAt   			= `updatedAt`
	UpdatedBy   			= `updatedBy`
	DeletedAt   			= `deletedAt`
	DeletedBy					= `deletedBy`
	RestoredBy				= `restoredBy`
	Name							= `name`
	Detail						= `detail`
	Rule							= `rule`
	Kind							= `kind`
	CogsIDR						= `cogsIDR`
	ProfitPercentage	= `profitPercentage`
)

const (
	TableLocations Tt.TableName = `locations`

	Country				= `country`
	StateProvince = `stateProvice`
	CityRegency		= `cityRegency`
	Subdistrict 	= `subdistrict`
	Village 			= `village`
	RwBanjar 			= `rwBanjar`
	RtNeigb 			= `rtNeigb`
	Address 			= `address`
	Lat 					= `lat`
	Lng 					= `lng`
)

const (
	TableInventoryChanges Tt.TableName = `inventoryChanges`

	StockDelta		= `stockDelta`
	ProductId 		= `productId`
	LocationId		= `locationId`
	SpendingId		= `spendingId`
	ExpenseId			= `expenseId`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableProducts: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
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
			{CogsIDR, Tt.Unsigned},
			{ProfitPercentage, Tt.Double},
		},
		AutoIncrementId: true,
		Engine:          Tt.Vinyl,
	},
	TableLocations: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
			{Name, Tt.String},
			{Country, Tt.String},
			{StateProvince, Tt.String},
			{CityRegency, Tt.String},
			{Subdistrict, Tt.String},
			{Village, Tt.String},
			{RwBanjar, Tt.String},
			{RtNeigb, Tt.String},
			{Address, Tt.String},
			{Lat, Tt.Double},
			{Lng, Tt.Double},
		},
		AutoIncrementId: true,
		Engine: Tt.Vinyl,
	},
	TableInventoryChanges: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
			{StockDelta, Tt.Unsigned},
			{ProductId, Tt.Unsigned},
			{LocationId, Tt.Unsigned},
			{SpendingId, Tt.Unsigned},
			{ExpenseId, Tt.Unsigned},
		},
		AutoIncrementId: true,
		Engine: Tt.Vinyl,
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
