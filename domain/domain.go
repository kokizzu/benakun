package domain

import (
	"errors"
	"fmt"
	"net"
	"time"

	chBuffer "github.com/kokizzu/ch-timed-buffer"
	"github.com/kokizzu/gotro/D/Ch"
	"github.com/kokizzu/gotro/D/Tt"
	"github.com/kokizzu/gotro/M"
	"github.com/kokizzu/gotro/S"
	"github.com/kokizzu/gotro/T"
	"github.com/rs/zerolog"

	"benakun/conf"
	"benakun/model/mAuth"
	"benakun/model/mAuth/saAuth"
	"benakun/model/xMailer"
)

type Domain struct {
	AuthOltp *Tt.Adapter
	AuthOlap *Ch.Adapter

	StorOltp *Tt.Adapter

	Mailer xMailer.Mailer

	IsBgSvc bool // long-running program

	// 3rd party
	Oauth conf.OauthConf

	// oauth related cache
	googleUserInfoEndpointCache string

	// timed buffer
	authLogs *chBuffer.TimedBuffer

	// logger
	Log *zerolog.Logger

	// list of superadmin emails
	Superadmins M.SB
	UploadDir   string
}

// will run in background if background service
func (d *Domain) runSubtask(subTask func()) {
	if d.IsBgSvc {
		go subTask()
	} else {
		subTask()
	}
}

func (d *Domain) InitTimedBuffer() {
	d.authLogs = chBuffer.NewTimedBuffer(d.AuthOlap.DB, 100_000, 1*time.Second, saAuth.Preparators[mAuth.TableActionLogs])
}

func (d *Domain) WaitTimedBufferFinalFlush() {
	<-d.authLogs.WaitFinalFlush
	d.Log.Debug().Msg(`timed buffer flushed`)
}

var defaultIP4 = net.ParseIP(`0.0.0.0`).To4()
var defaultIP6 = net.ParseIP(`0:0:0:0:0:0:0:0`).To16()

func (d *Domain) InsertActionLog(in *RequestCommon, out *ResponseCommon) bool {
	ip := net.ParseIP(in.IpAddress)
	ip4 := ip.To4()
	if ip4 == nil {
		ip4 = defaultIP4
	}
	ip6 := ip.To16()
	if ip6 == nil {
		ip6 = defaultIP6
	}
	row := saAuth.ActionLogs{
		CreatedAt:  in.TimeNow(),
		RequestId:  in.RequestId,
		ActorId:    in.SessionUser.UserId,
		Action:     in.Action,
		StatusCode: int16(out.StatusCode),
		Traces:     out.Traces(),
		Error:      out.Error,
		IpAddr4:    ip4,
		IpAddr6:    ip6,
		UserAgent:  in.UserAgent,
		TenantCode: in.SessionUser.TenantCode,
		Latency:    in.Latency(),
	}
	return d.authLogs.Insert([]any{
		row.CreatedAt,
		row.RequestId,
		row.ActorId,
		row.Action,
		row.StatusCode,
		row.Traces,
		row.Error,
		row.IpAddr4,
		row.IpAddr6,
		row.UserAgent,
		row.Latency,
		row.TenantCode,
		row.RefId,
	})
}

func (d *Domain) CloseTimedBuffer() {
	go d.authLogs.Close()
	d.WaitTimedBufferFinalFlush()
}

const (
	InvitationStateInvited    = `invited`
	InvitationStateRevoked    = `revoked`
	InvitationStateAccepted   = `accepted`
	InvitationStateRejected   = `rejected`
	InvitationStateTerminated = `terminated`
	InvitationStateLeft       = `left`

	InvitationStateRespAccept = `accept`
	InvitationStateRespReject = `reject`
)

type (
	InviteState struct {
		TenantCode string
		State      string
		Date       string
	}
	StateMap map[string]InviteState
)

func (is InviteState) ToStateString() (str string) {
	return fmt.Sprintf("tenant:%s:%s:%v", is.TenantCode, is.State, is.Date)
}

func (s StateMap) ModifyState(tenantCode, newState string) {
	if sn, ok := s[tenantCode]; ok {
		if sn.State != newState {
			sn.TenantCode = tenantCode
			sn.State = newState
			sn.Date = T.DateStr()
			s[tenantCode] = sn
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
}

func (s StateMap) ToStateString() (str string) {
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

func ToStateMap(states string) (out StateMap, err error) {
	out = StateMap{}
	statesArray := S.Split(states, ` `)
	if len(statesArray) == 0 {
		return StateMap{}, errors.New(`States empty`)
	}
	for _, state := range statesArray {
		parts := S.Split(state, `:`)
		if len(parts) > 0 {
			out[parts[1]] = InviteState{
				TenantCode: parts[1],
				State:      parts[2],
				Date:       parts[3],
			}
		}
	}
	return out, nil
}
