package mAuth

import (
	"fmt"

	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/T"
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
	ErrInvitationStateEmpty           = `state is empty`
)

type (
	InviteState        struct{ TenantCode, State, Date string }
	InvitationStateMap map[string]InviteState
)

var wrapInvitationError = func(errMsg, oldState, newState string) error {
	return fmt.Errorf("%s: cannot change state from %s to %s", errMsg, oldState, newState)
}

func (is InviteState) ToStateString() (str string) {
	return fmt.Sprintf("tenant:%s:%s:%v", is.TenantCode, is.State, is.Date)
}

func (s InvitationStateMap) ModifyState(tenantCode, newState string) error {
	if sn, ok := s[tenantCode]; ok {
		if sn.State != newState {
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
					s[st.TenantCode] = st
				} else if st.State == InvitationStateInvited {
					st.State = InvitationStateRevoked
					s[st.TenantCode] = st
				}
			}
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

func StateField(tenantCode, state string) string {
	return fmt.Sprintf("tenant:%s:%s:%v", tenantCode, state, T.DateStr())
}

func ToInvitationStateMap(states string) (InvitationStateMap, error) {
	out := InvitationStateMap{}
	statesArray := S.Split(states, ` `)
	if len(statesArray) == 0 || states == `` {
		return InvitationStateMap{}, wrapInvitationError(ErrInvitationStateEmpty, states, states)
	}
	for _, state := range statesArray {
		parts := S.Split(state, `:`)
		if len(parts) > 0 {
			out[parts[1]] = InviteState{
				TenantCode: parts[1],
				State:      parts[2],
				Date:       parts[3],
			}
		} else {
			return out, wrapInvitationError(ErrInvitationStateEmpty, states, states)
		}
	}
	return out, nil
}
