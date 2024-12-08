package email

import (
	"github.com/LucienLSA/go-gin-mall/conf"
	"gopkg.in/mail.v2"
)

type EmailSender struct {
	SmtpHost      string `json:"smtp_host"`
	SmtpEmailFrom string `json:"smtp_email_from"`
	SmtpPass      string `json:"smtp_pass"`
}

func NewEmailBinder() *EmailSender {
	eConfig := conf.Config.Email
	return &EmailSender{
		SmtpHost:      eConfig.SmtpHost,
		SmtpEmailFrom: eConfig.SmtpEmail,
		SmtpPass:      eConfig.SmtpPass,
	}
}

// Bind 绑定发送邮件
func (e *EmailSender) Bind(data, emailTo, subject string) error {
	m := mail.NewMessage()
	m.SetHeader("From", e.SmtpEmailFrom)
	m.SetHeader("To", emailTo)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", data)
	d := mail.NewDialer(e.SmtpHost, 465, e.SmtpEmailFrom, e.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
