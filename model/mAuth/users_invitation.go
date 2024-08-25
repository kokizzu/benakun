package mAuth

import (
	"fmt"

	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/T"
	"github.com/rs/zerolog/log"
)

const (
	InvitationStateInvited    = `invited`
	InvitationStateRevoked    = `revoked`
	InvitationStateAccepted   = `accepted`
	InvitationStateRejected   = `rejected`
	InvitationStateTerminated = `terminated`
	InvitationStateLeft       = `left`

	InvitationStateRespAccept = `accept`
	InvitationStateRespReject = `reject`

	ErrInvitationStateIsTheSame       = `state is same as previous state`
	ErrInvitationStateAlreadyAccepted = `state already accepted`
	ErrInvitationStateIsTheSameRole		= `role is same as previous role`
	ErrInvitationStateInvalidRole			= `invalid role (must be user/dataEntry/reportViewer)`
)

// Invitation State:
// tenant:xx-1234:accepted:2024-06-09				| before
// tenant:xx-1234:role:accepted:2024-06-09	| after

type (
	InviteState        struct{ TenantCode, Role, State, Date string }
	InvitationStateMap map[string]InviteState
)

var (
	ErrInvitationStateEmpty = fmt.Errorf(`invitation state is empty`)
)

func (is InviteState) ToStateString() (str string) {
	return fmt.Sprintf("tenant:%s:%s:%s:%v", is.TenantCode, is.Role, is.State, is.Date)
}

func (s InvitationStateMap) ModifyState(tenantCode, newState string) error {
	var wrapInvitationError = func(errMsg, oldState, newState string) error {
		return fmt.Errorf("%s: cannot change state from %s to %s", errMsg, oldState, newState)
	}

	if sn, ok := s[tenantCode]; ok {
		if sn.State != newState || (sn.State == InvitationStateInvited && newState == InvitationStateInvited){
			if sn.State == InvitationStateAccepted {
				if !(newState == InvitationStateTerminated || newState == InvitationStateLeft) {
					return wrapInvitationError(ErrInvitationStateAlreadyAccepted, sn.State, newState)
				}
			}
			sn.TenantCode = tenantCode
			sn.State = newState
			sn.Date = T.DateStr()
			s[tenantCode] = sn
		} else {
			return wrapInvitationError(ErrInvitationStateIsTheSame, sn.State, newState)
		}
	} else {
		s[tenantCode] = InviteState{
			TenantCode: tenantCode,
			State:      newState,
			Date:       T.DateStr(),
		}
	}

	for _, st := range s {
		if newState == InvitationStateAccepted { // Change Accepted to Left
			if st.TenantCode != tenantCode {
				if st.State == InvitationStateAccepted {
					st.State = InvitationStateLeft
					st.Date = T.DateStr()
					s[st.TenantCode] = st
				} else if st.State == InvitationStateInvited {
					st.State = InvitationStateRevoked
					st.Date = T.DateStr()
					s[st.TenantCode] = st
				}
			}
		}
	}

	return nil
}

func (s InvitationStateMap) ModifyRole(tenantCode, newRole string) error {
	var wrapInvitationRoleError = func(errMsg, newRole string) error {
		return fmt.Errorf("%s: cannot change role to %s", errMsg, newRole)
	}

	switch newRole {
	case RoleUser, RoleDataEntry, RoleReportViewer, RoleFieldSupervisor:
		break
	default:
		return wrapInvitationRoleError(ErrInvitationStateInvalidRole, newRole)
	}

	if sn, ok := s[tenantCode]; ok {
		if sn.Role != newRole {
			sn.Role = newRole
			sn.Date = T.DateStr()
			s[tenantCode] = sn
		} else {
			return wrapInvitationRoleError(ErrInvitationStateIsTheSameRole, newRole)
		}
	}

	return nil
}

func (s InvitationStateMap) ToStateString() (str string) {
	idx := 0
	str = ``
	for _, st := range s {
		if idx == 0 {
			str += st.ToStateString()
		} else {
			str += ` ` + st.ToStateString()
		}
		idx++
	}
	return str
}

func ToInvitationStateMap(states string) (InvitationStateMap, error) {
	out := InvitationStateMap{}
	states = S.Trim(states)
	statesArray := S.Split(states, ` `)
	if len(statesArray) == 0 || states == `` {
		return InvitationStateMap{}, ErrInvitationStateEmpty
	}
	for _, state := range statesArray {
		if state == `` {
			continue
		}
		parts := S.Split(state, `:`)
		if len(parts) == 5 {
			out[parts[1]] = InviteState{
				TenantCode: parts[1],
				Role: 			parts[2],
				State:      parts[3],
				Date:       parts[4],
			}
		} else {
			log.Printf("ToInvitationStateMap.WARN: %v, have invalid state: %s", states, state)
		}
	}
	if len(out) == 0 {
		return InvitationStateMap{}, ErrInvitationStateEmpty
	}
	return out, nil
}

func (s InvitationStateMap) GetStateByTenantCode(tenantCode string) string {
	return s[tenantCode].State
}

func (s InvitationStateMap) GetRoleByTenantCode(tenantCode string) string {
	return s[tenantCode].Role
}