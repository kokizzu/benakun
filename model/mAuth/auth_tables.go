package mAuth

import (
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
)

const (
	RoleUser            = `user`
	RoleTenantAdmin     = `tenantAdmin`
	RoleDataEntry       = `dataEntry`
	RoleReportViewer    = `reportViewer`
	RoleFieldSupervisor = `fieldSupervisor`
	RoleSuperAdmin      = `superAdmin`
)

var RoleMap = map[string]string{
	RoleUser:            `User`,
	RoleTenantAdmin:     `Tenant Admin`,
	RoleDataEntry:       `Data Entry`,
	RoleReportViewer:    `Report Viewer`,
	RoleFieldSupervisor: `Field Supervisor`,
	RoleSuperAdmin:      `Super Admin`,
}

func IsValidRole(role string) bool {
	switch role {
	case RoleTenantAdmin:
		return true
	case RoleDataEntry:
		return true
	case RoleReportViewer:
		return true
	case RoleFieldSupervisor:
		return true
	default:
		return false
	}
}

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
	DeletedBy          = `deletedBy`
	RestoredBy         = `restoredBy`
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
	SupportExpiredAt   = `supportExpiredAt`
)

const DefaultPassword = `user12345678`

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

	ProductsCoaId            = `productsCoaId`
	SuppliersCoaId           = `suppliersCoaId`
	CustomersCoaId           = `customersCoaId`
	CustomerReceivablesCoaId = `customerReceivablesCoaId`
	StaffsCoaId              = `staffsCoaId`
	BanksCoaId               = `banksCoaId`
	FundersCoaId             = `fundersCoaId`
)

func IsCoaDifferent(values ...uint64) bool {
	valueMap := make(map[uint64]bool)
	for _, v := range values {
		if v == 0 {
			continue
		}
		if _, exist := valueMap[v]; exist {
			return false
		}
		valueMap[v] = true
	}
	return true
}

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
			{SupportExpiredAt, Tt.Integer},
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
			{ProductsCoaId, Tt.Unsigned},
			{SuppliersCoaId, Tt.Unsigned},
			{CustomersCoaId, Tt.Unsigned},
			{StaffsCoaId, Tt.Unsigned},
			{BanksCoaId, Tt.Unsigned},
			{CustomerReceivablesCoaId, Tt.Unsigned},
			{FundersCoaId, Tt.Unsigned},
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
			{DeletedBy, Tt.Unsigned},
			{RestoredBy, Tt.Unsigned},
		},
		AutoIncrementId: true,
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
