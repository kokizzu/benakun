package xMailer

import (
	"github.com/kokizzu/gotro/L"

	"benakun/conf"
)

type SendMailFunc func(toEmailName map[string]string, subject, text, html string) error

type Mailer struct {
	SendMailFunc SendMailFunc
}

func (m *Mailer) SendRegistrationEmail(email string, verifyEmailUrl string) error {
	if conf.IsDebug() {
		L.Print(`SendRegistrationEmail`, email, verifyEmailUrl)
	}
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Verify Registration Link`,
		`Hi `+email+`, 

please click this link to verify your registration: 
  `+verifyEmailUrl+`

please ignore this email if you didn't register`,
		`Hi `+email+`, <br><br>
please click this link to verify your registration: <br/>
  <a href="`+verifyEmailUrl+`">`+verifyEmailUrl+`</a><br/><br/>
please ignore this email if you didn't register<br/>`,
	)
}

func (m *Mailer) SendResetPasswordEmail(email string, resetPassUrl string) error {
	if conf.IsDebug() {
		L.Print(`SendResetPasswordEmail`, email, resetPassUrl)
	}
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Reset Password Link`,
		`Hi `+email+`,

please click this link to reset your password: 
`+resetPassUrl+`

please ignore this email if you didn't request reset password link`,
		`Hi `+email+`, <br><br>
please click this link to reset your password: <br/>
<a href="`+resetPassUrl+`">`+resetPassUrl+`</a><br/><br/>
please ignore this email if you didn't request reset password link<br/>`,
	)
}

func (m *Mailer) SendInviteUserEmail(tenantCode, email, invitationUrl string) error {
	acceptUrl := invitationUrl + `accept`
	rejectUrl := invitationUrl + `reject`
	if conf.IsDebug() {
		L.Print(`SendInviteUserEmail`, email, invitationUrl)
	}
	return m.SendMailFunc(
		map[string]string{email: ``},
		`Invitation Company`,
		`Hi `+email+`,

please click one of these links to response our invitation to join our company:

Accept: `+acceptUrl+`
Reject: `+rejectUrl,
		`Hi `+email+`, <br><br>
please click one of these links to response our invitation to join our company: <br/>
<br/>
Accept: <a href="`+acceptUrl+`">`+acceptUrl+`</a><br/><br/>
Reject: <a href="`+rejectUrl+`">`+rejectUrl+`</a><br/><br/>`,
	)
}

func (m *Mailer) SendResponseStateTenantEmail(tenantEmail, userEmail, state string) error {
	if conf.IsDebug() {
		L.Print(`SendResponseStateTenantEmail`, tenantEmail, userEmail, state)
	}
	return m.SendMailFunc(
		map[string]string{tenantEmail: ``},
		`User status changed`,
		`Hi `+tenantEmail+`,

A user with email `+userEmail+` has `+state+` from your company:`,
		`Hi `+tenantEmail+`, <br><br>
A user with email <a href="mailto:`+userEmail+`">`+userEmail+`</a> has `+state+` from your company: <br/>`,
	)
}
