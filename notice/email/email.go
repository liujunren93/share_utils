package email

import (
	"net/smtp"
	"strings"
)

type Config struct {
	Host     string
	UserName string
	Password string
}

type Email struct {
	Subject  string
	To       []string
	MailType string
	Body     string
}

func (c Config) SendMsg(e Email) error {
	contentType := "Content-Type: text/plain;charset=UTF-8"
	if e.MailType == "html" {
		contentType = "Content-Type: text/html;charset=UTF-8"
	}
	msg := []byte("To:aaa \r\nFrom: " + c.UserName + "<" + c.UserName + ">\r\nSubject: " + e.Subject + "\r\n" + contentType + "\r\n\r\n" + e.Body)

	hp := strings.Split(c.Host, ":")
	auth := smtp.PlainAuth("", c.UserName, c.Password, hp[0])
	return smtp.SendMail(c.Host, auth, c.UserName, e.To, msg)
}
