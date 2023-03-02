package email

import (
	"crypto/tls"

	"gopkg.in/gomail.v2"
)

type Email struct {
	*SMTPInfo
}

// 发送邮箱必须参数
type SMTPInfo struct {
	Host     string
	Port     int
	IsSSL    bool
	UserName string
	Password string
	From     string
}

func NewEmail(info *SMTPInfo) *Email {
	return &Email{SMTPInfo: info}
}

// SendMail
/**
 * @description: 发送邮件
 * @param {[]string} to 收件人
 * @param {*} subject 主题
 * @param {string} body 正文
 * @return {*}
 */
func (e *Email) SendMail(to []string, subject, body string) error {
	// 消息实例
	m := gomail.NewMessage()
	m.SetHeader("From", e.From)     // 发件人
	m.SetHeader("To", to...)        // 收件人
	m.SetHeader("Subject", subject) // 主题
	m.SetBody("text/html", body)    // 正文

	// 拨号实例
	dialer := gomail.NewDialer(e.Host, e.Port, e.UserName, e.Password)
	// 发送邮件
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: e.IsSSL}
	return dialer.DialAndSend(m)
}
