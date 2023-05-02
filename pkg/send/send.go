package send

import (
	"crypto/tls"
	. "github.com/myoperator/messageoperator/pkg/sysconfig"
	"gopkg.in/mail.v2"
	"k8s.io/klog/v2"
)

var (
	GlobalSend *Send
)

func init() {
	GlobalSend = NewSender() // FIXME: 目前有sender实例无法重复利用的bug
}

// Sender 邮件发送器
type Send struct {
	dialer *mail.Dialer
}

// Send 发送邮件
func (sender *Send) Send(title, content string) error {
	// TODO: 需要抛出错误
	m := mail.NewMessage()
	m.SetHeader("From", SysConfig1.Sender.Email)
	m.SetHeader("To", SysConfig1.Sender.Targets)
	m.SetHeader("Subject", title)
	m.SetBody("text/plain", content)
	if err := sender.dialer.DialAndSend(m); err != nil {
		klog.Error("send err: ", err)
		return err
	}
	klog.Info("send email.....")
	return nil

	//d, err := sender.dialer.Dial()
	//if err != nil {
	//	fmt.Println("这里报错")
	//	log.Print(err)
	//}
	//t := []string{SysConfig1.Sender.Targets}
	//d.Send(SysConfig1.Sender.Email, t, io.WriteString())

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
