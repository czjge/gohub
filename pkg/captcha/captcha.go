package captcha

import (
	"sync"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/app"
	"github.com/czjge/gohub/pkg/redis"
	"github.com/mojocn/base64Captcha"
)

type Captcha struct {
	Base64Captcha *base64Captcha.Captcha
}

// 单例模式
var once sync.Once

var internalCaptcha *Captcha

func NewCaptcha() *Captcha {

	once.Do(func() {
		internalCaptcha = &Captcha{}

		store := RedisStore{
			RedisClient: redis.Redis(),
			KeyPrefix:   config.GetConfig().App.Name + ":captcha",
		}

		captchaConfig := config.GetConfig().Captcha
		driver := base64Captcha.NewDriverDigit(
			captchaConfig.Height,
			captchaConfig.Width,
			captchaConfig.Length,
			captchaConfig.Maxskew,
			captchaConfig.Dotcount,
		)

		internalCaptcha.Base64Captcha = base64Captcha.NewCaptcha(driver, &store)
	})

	return internalCaptcha
}

func (c *Captcha) GenerateCaptcha() (id string, b64s string, err error) {
	return c.Base64Captcha.Generate()
}

func (c *Captcha) VerifyCaptcha(id string, answer string) (match bool) {

	if !app.IsProduction() && id == config.GetConfig().Captcha.TestingKey {
		return true
	}

	return c.Base64Captcha.Verify(id, answer, false) // 多次提交
}
