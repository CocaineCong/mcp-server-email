package main

import (
	"gopkg.in/mail.v2"
)

const (
	SMTPHost  = "smtp.qq.com"
	SMTPEmail = "xxx@qq.com" // 发送人的邮箱
	SMTPPass  = "xxx"        // 通行证密码
)

type EmailSender struct {
	SmtpHost      string `json:"smtp_host"`
	SmtpEmailFrom string `json:"smtp_email_from"`
	SmtpPass      string `json:"smtp_pass"`
}

func NewEmailSender() *EmailSender {
	return &EmailSender{
		SmtpHost:      SMTPHost,
		SmtpEmailFrom: SMTPEmail,
		SmtpPass:      SMTPPass,
	}
}

// Send 发送邮件
func (s *EmailSender) Send(email, subject, content string) error {
	m := mail.NewMessage()
	m.SetHeader("From", s.SmtpEmailFrom)
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", content)
	d := mail.NewDialer(s.SmtpHost, 465, s.SmtpEmailFrom, s.SmtpPass)
	d.StartTLSPolicy = mail.MandatoryStartTLS
	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return nil
}
