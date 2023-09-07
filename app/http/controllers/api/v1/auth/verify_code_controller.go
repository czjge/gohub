package auth

import (
	v1 "github.com/czjge/gohub/app/http/controllers/api/v1"
	"github.com/czjge/gohub/pkg/captcha"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/czjge/gohub/pkg/response"
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
