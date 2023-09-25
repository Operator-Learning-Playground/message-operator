package send

import (
	"crypto/tls"
	. "github.com/myoperator/messageoperator/pkg/sysconfig"
	"gopkg.in/mail.v2"
	"k8s.io/klog/v2"
)

// EmailSender 邮件发送器
type EmailSender struct {
	dialer *mail.Dialer
	mess   *mail.Message
}

// Send 发送邮件
func (sender *EmailSender) Send(title, content string) error {
	if !SysConfig1.Sender.Open {
		klog.Info("config no support email sending")
		return nil
	}

	// TODO: 需要抛出错误
	sender.mess.SetHeader("From", SysConfig1.Sender.Email)
	sender.mess.SetHeader("To", SysConfig1.Sender.Targets)
	sender.mess.SetHeader("Subject", title)
	sender.mess.SetBody("text/plain", content)
	if err := sender.dialer.DialAndSend(sender.mess); err != nil {
		klog.Error("send err: ", err)
		return err
	}
	klog.Info("send email.....")
	sender.mess.Reset()
	return nil
}

// NewEmailSender 创建邮件发送器
func NewEmailSender() *EmailSender {
	d := mail.NewDialer(SysConfig1.Sender.Remote, SysConfig1.Sender.Port,
		SysConfig1.Sender.Email, SysConfig1.Sender.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &EmailSender{
		dialer: d,
		mess:   mail.NewMessage(),
	}
}
