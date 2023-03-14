package main

import (
	"crypto/tls"
	"github.com/myoperator/messageoperator/pkg/sysconfig"
	"gopkg.in/mail.v2"
	"k8s.io/klog/v2"
	"log"
	"os"
)

func main() {

	// 加载配置文件
	if err := sysconfig.InitConfig(); err != nil {
		klog.Error(err, "unable to load sysconfig")
		os.Exit(1)
	}

	// 发送邮件
	s := NewSender()
	s.Send("再一次测试", "再一次测试")

}


// Sender 邮件发送器
type Send struct {
	dialer *mail.Dialer
}

// Send 发送邮件
func (sender *Send) Send(title, content string) {
	m := mail.NewMessage()
	m.SetHeader("From", sysconfig.SysConfig1.Sender.Email)
	m.SetHeader("To", sysconfig.SysConfig1.Sender.Targets)
	m.SetHeader("Subject", title)
	m.SetBody("text/plain", content)
	if err := sender.dialer.DialAndSend(m); err != nil {
		log.Print(err)
	}
}

// NewSender 创建邮件发送器
func NewSender() *Send {
	d := mail.NewDialer(sysconfig.SysConfig1.Sender.Remote, sysconfig.SysConfig1.Sender.Port,
		sysconfig.SysConfig1.Sender.Email, sysconfig.SysConfig1.Sender.Password)
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	return &Send{
		dialer: d,
	}
}

