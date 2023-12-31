package mail

import (
	"crypto/tls"
	"fmt"
	"net/smtp"

	"github.com/czjge/gohub/pkg/logger"
	emailPKG "github.com/jordan-wright/email"
)

// 实现 email.Driver interface
type SMTP struct{}

func (s *SMTP) Send(email Email, config map[string]string) bool {

	e := emailPKG.NewEmail()

	e.From = fmt.Sprintf("%v <%v>", email.From.Name, email.From.Address)
	e.To = email.To
	e.Bcc = email.Bcc
	e.Cc = email.Cc
	e.Subject = email.Subject
	e.Text = email.Text
	e.HTML = email.HTML

	logger.DebugJSON("发送邮件", "发送详情", e)

	logger.DebugJSON("发送邮件", "参数", config)

	// SendWithTLS
	// Send
	err := e.SendWithTLS(
		fmt.Sprintf("%v:%v", config["Host"], config["Port"]),
		smtp.PlainAuth(
			"",
			config["Username"],
			config["Password"],
			config["Host"],
		),
		&tls.Config{
			ServerName: config["Host"],
		},
	)
	if err != nil {
		logger.ErrorString("发送邮件", "发件出错", err.Error())
		return false
	}

	logger.DebugString("发送邮件", "发送成功", "")
	return true
}
