package mAuth

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	TableUsers Tt.TableName = `users`

	Id                 = `id`
	Email              = `email`
	Password           = `password`
	CreatedAt          = `createdAt`
	CreatedBy          = `createdBy`
	UpdatedAt          = `updatedAt`
	UpdatedBy          = `updatedBy`
	DeletedAt          = `deletedAt`
	PasswordSetAt      = `passwordSetAt`
	SecretCode         = `secretCode`
	SecretCodeAt       = `secretCodeAt`
	VerificationSentAt = `verificationSentAt`
	VerifiedAt         = `verifiedAt`
	LastLoginAt        = `lastLoginAt`
	FullName           = `fullName`
	TenantCode         = `tenantCode`
	Role               = `role`
	InvitedAt          = `invitedAt`
	InvitationState    = `invitationState`
)

const (
	TableSessions Tt.TableName = `sessions`

	SessionToken = `sessionToken`
	UserId       = `userId`
	ExpiredAt    = `expiredAt`
	Device       = `device`

	LoginAt  = `loginAt`
	LoginIPs = `loginIPs`
)

const (
	TableOrgs Tt.TableName = `orgs` // organization

	// OrgType is the type of organization:
	// 1: company (sub of a tenant)
	// 2: department
	// 3: division, have head title
	// 4: job
	OrgType   = `orgType`
	Name      = `name`
	HeadTitle = `headTitle`
	ParentId  = `parentId`
	Children  = `children`

	OrgTypeCompany  = 1
	OrgTypeDept     = 2
	OrgTypeDivision = 3
	OrgTypeJob      = 4
)

const (
	TableTenants Tt.TableName = `tenants`
)

const (
	TableCoa Tt.TableName = `coa`
	Level                 = `level`
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
	CompletedAt = `completedAt`
	CoaId = `coaId`
	Price = `price`
	Description = `descriptions`
	Qty = `qty`
)

var TarantoolTables = map[Tt.TableName]*Tt.TableProp{
	TableUsers: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{Email, Tt.String},
			{Password, Tt.String},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
			{PasswordSetAt, Tt.Integer},
			{SecretCode, Tt.String},
			{SecretCodeAt, Tt.Integer},
			{VerificationSentAt, Tt.Integer},
			{VerifiedAt, Tt.Integer},
			{LastLoginAt, Tt.Integer},
			{FullName, Tt.String},
			{TenantCode, Tt.String},
			{Role, Tt.String},
			{InvitationState, Tt.String},
		},
		AutoIncrementId:  true,
		Unique1:          Email,
		AutoCensorFields: []string{Password, SecretCode, SecretCodeAt},
		Engine:           Tt.Memtx,
	},
	TableSessions: {
		Fields: []Tt.Field{
			{SessionToken, Tt.String},
			{UserId, Tt.Unsigned},
			{ExpiredAt, Tt.Integer},
			{Device, Tt.String},
			{LoginAt, Tt.Integer},
			{LoginIPs, Tt.String},
			{TenantCode, Tt.String},
		},
		Unique1: SessionToken,
		Engine:  Tt.Memtx,
	},
	TableTenants: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			// TODO: add mapping to specific database
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
		},
		AutoIncrementId: true,
		Unique1:         TenantCode,
		Engine:          Tt.Memtx,
	},
	TableOrgs: {
		Fields: []Tt.Field{
			{Id, Tt.Unsigned},
			{TenantCode, Tt.String},
			{Name, Tt.String},
			{HeadTitle, Tt.String},
			{ParentId, Tt.Unsigned},
			{Children, Tt.Array},
			{OrgType, Tt.Unsigned},
			{CreatedAt, Tt.Integer},
			{CreatedBy, Tt.Unsigned},
			{UpdatedAt, Tt.Integer},
			{UpdatedBy, Tt.Unsigned},
			{DeletedAt, Tt.Integer},
		},
		AutoIncrementId: true,
	},
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
			{Price, Tt.Integer},
			{Description, Tt.String},
			{Qty, Tt.Integer},
		},
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
			{TenantCode, Ch.String},
			{RefId, Ch.UInt64},
		},
		Orders: []string{CreatedAt, RequestId, ActorId, Action},
	},
}
