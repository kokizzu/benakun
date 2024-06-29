package domain

import (
	"benakun/model/mAuth/rqAuth"
	"benakun/model/mAuth/saAuth"
)

type (
	SuperAdminDashboardIn struct {
		RequestCommon
	}

	SuperAdminDashboardOut struct {
		ResponseCommon
		RegisteredUserTotal int64 `json:"registeredUserTotal" form:"registeredUserTotal" query:"registeredUserTotal" long:"registeredUserTotal" msg:"registeredUserTotal"`
		RegisteredUserToday int64 `json:"registeredUserToday" form:"registeredUserToday" query:"registeredUserToday" long:"registeredUserToday" msg:"registeredUserToday"`

		RequestsPerDate   map[string]int `json:"requestsPerDate" form:"requestsPerDate" query:"requestsPerDate" long:"requestsPerDate" msg:"requestsPerDate"`
		UniqueUserPerDate map[string]int `json:"uniqueUserPerDate" form:"uniqueUserPerDate" query:"uniqueUserPerDate" long:"uniqueUserPerDate" msg:"uniqueUserPerDate"`
		UniqueIpPerDate   map[string]int `json:"uniqueIpPerDate" form:"uniqueIpPerDate" query:"uniqueIpPerDate" long:"uniqueIpPerDate" msg:"uniqueIpPerDate"`

		CountPerActionsPerDate map[string]map[string]int `json:"countPerActionsPerDate" form:"countPerActionsPerDate" query:"countPerActionsPerDate" long:"countPerActionsPerDate" msg:"countPerActionsPerDate"`
	}
)

const (
	SuperAdminDashboardAction = `superAdmin/dashboard`
)

func (d *Domain) SuperAdminDashboard(in *SuperAdminDashboardIn) (out SuperAdminDashboardOut) {
	defer d.InsertActionLog(&in.RequestCommon, &out.ResponseCommon)

	sess := d.MustSuperAdmin(in.RequestCommon, &out.ResponseCommon)
	if sess == nil {
		return
	}

	rq := rqAuth.NewUsers(d.AuthOltp)
	out.RegisteredUserTotal = rq.Total()
	out.RegisteredUserToday = rq.CountUserRegisterToday()

	sa := saAuth.NewActionLogs(d.AuthOlap)
	out.RequestsPerDate = sa.StatRequestsPerDate()
	out.UniqueUserPerDate = sa.StatUniqueUserPerDate()
	out.UniqueIpPerDate = sa.StatUniqueIpPerDate()

	out.CountPerActionsPerDate = sa.StatPerActionsPerDate()
	return
}
