package send


import (
	"crypto/tls"
	. "github.com/myoperator/messageoperator/pkg/sysconfig"
	"gopkg.in/mail.v2"
	"log"
)

var (
	GlobalSend  *Send
)


func init() {
	GlobalSend = NewSender()
}

// Sender 邮件发送器
type Send struct {
	dialer *mail.Dialer
}

// Send 发送邮件
func (sender *Send) Send(title, content string) {
	m := mail.NewMessage()
	m.SetHeader("From", SysConfig1.Sender.Email)
	m.SetHeader("To", SysConfig1.Sender.Targets)
	m.SetHeader("Subject", title)
	m.SetBody("text/plain", content)
	if err := sender.dialer.DialAndSend(m); err != nil {
		log.Print(err)
	}
}

// NewSender 创建邮件发送器
func NewSender() *Send {
	d := mail.NewDialer(SysConfig1.Sender.Remote, SysConfig1.Sender.Port,
		SysConfig1.Sender.Email, SysConfig1.Sender.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Send{
		dialer: d,
	}
}
