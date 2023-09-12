package verifycode

import (
	"fmt"
	"strings"
	"sync"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/app"
	"github.com/czjge/gohub/pkg/helpers"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/czjge/gohub/pkg/mail"
	"github.com/czjge/gohub/pkg/redis"
	"github.com/czjge/gohub/pkg/sms"
)

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

// 单例模式
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			// 指针类型的接口变量 *T 可调用接收者是 *T 和 T 的方法
			// 值类型的接口变量 T 可调用接收者是 T 的方法
			// RedisStore 的方法接受者都是 *T
			Store: &RedisStore{
				RedisClient: redis.Redis(),
				KeyPrefix:   config.GetConfig().App.Name + ":verifycode:",
			},
		}
	})

	return internalVerifyCode
}

func (vc *VerifyCode) SendSMS(phone string) bool {

	code := vc.generateVerifyCode(phone)

	config := config.GetConfig()

	if !app.IsProduction() && strings.HasPrefix(phone, config.Verifycode.DebugPhonePrefix) {
		return true
	}

	return sms.NewSMS().Send(phone, sms.Mesage{
		Template: config.Sms.TemplateCode,
		Data:     map[string]string{"code": code},
	})
}

func (vc *VerifyCode) SendEmail(email string) error {

	code := vc.generateVerifyCode(email)

	config := config.GetConfig()

	if !app.IsProduction() && strings.HasSuffix(email, config.Verifycode.DebugEmailSuffix) {
		return nil
	}

	content := fmt.Sprintf("<h1>您的 Email 验证码是 %v</h1>", code)

	mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: config.Email.From.Address,
			Name:    config.Email.From.Name,
		},
		To:      []string{email},
		Subject: "Email 验证码",
		HTML:    []byte(content),
	})

	return nil
}

func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {

	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})

	config := config.GetConfig()

	if !app.IsProduction() && (strings.HasPrefix(key, config.Verifycode.DebugPhonePrefix) ||
		strings.HasSuffix(key, config.Verifycode.DebugEmailSuffix)) {
		return true
	}

	return vc.Store.Verify(key, answer, false)
}

func (vc *VerifyCode) generateVerifyCode(key string) string {

	config := config.GetConfig().Verifycode

	code := helpers.RandomNumber(config.CodeLength)
	if app.IsLocal() {
		code = config.DebugCode
	}

	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})

	vc.Store.Set(key, code)
	return code
}
