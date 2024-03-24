package domain

import (
	"benakun/model/mAuth/rqAuth"
)

//go:generate gomodifytags -all -add-tags json,form,query,long,msg -transform camelcase --skip-unexported -w -file TenantAdminDeleteOrganizationChild.go
//go:generate replacer -afterprefix "Id\" form" "Id,string\" form" type TenantAdminDeleteOrganizationChild.go
//go:generate replacer -afterprefix "json:\"id\"" "json:\"id,string\"" type TenantAdminDeleteOrganizationChild.go
//go:generate replacer -afterprefix "By\" form" "By,string\" form" type TenantAdminDeleteOrganizationChild.go
//go:generate farify doublequote --file TenantAdminDeleteOrganizationChild.go

type (
	TenantAdminDeleteOrganizationChildIn struct {
		RequestCommon
		Id         uint64      `json:"id" form:"id" query:"id" long:"id" msg:"id"`
	}
	TenantAdminDeleteOrganizationChildOut struct {
		ResponseCommon
		Org *rqAuth.Orgs `json:"org" form:"org" query:"org" long:"org" msg:"org"`
		Orgs *[]rqAuth.Orgs `json:"orgs" form:"orgs" query:"orgs" long:"orgs" msg:"orgs"`
	}
)

const (
	TenantAdminDeleteOrganizationChildAction = `tenantAdmin/deleteOrganizationChild`
)