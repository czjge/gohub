package auth

import (
	v1 "github.com/czjge/gohub/app/http/controllers/api/v1"
	"github.com/czjge/gohub/app/requests"
	"github.com/czjge/gohub/pkg/captcha"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/czjge/gohub/pkg/response"
	"github.com/czjge/gohub/pkg/verifycode"
	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseAPIControler
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	response.JSON(c, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}

func (vc *VerifyCodeController) SendUsingPhone(c *gin.Context) {

	request := requests.VerifyCodePhoneRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodePhone); !ok {
		return
	}

	if ok := verifycode.NewVerifyCode().SendSMS(request.Phone); !ok {
		response.Abort500(c, "发送短信验证码失败")
	} else {
		response.Success(c)
	}
}

func (vc *VerifyCodeController) SendUsingEmail(c *gin.Context) {

	request := requests.VerifyCodeEmailRequest{}
	if ok := requests.Validate(c, &request, requests.VerifyCodeEmail); !ok {
		return
	}

	err := verifycode.NewVerifyCode().SendEmail(request.Email)
	if err != nil {
		response.Abort500(c, "发送 Email 验证码失败")
	} else {
		response.Success(c)
	}
}
