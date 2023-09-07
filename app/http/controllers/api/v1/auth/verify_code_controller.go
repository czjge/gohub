package auth

import (
	"net/http"

	v1 "github.com/czjge/gohub/app/http/controllers/api/v1"
	"github.com/czjge/gohub/pkg/captcha"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
)

type VerifyCodeController struct {
	v1.BaseAPIControler
}

func (vc *VerifyCodeController) ShowCaptcha(c *gin.Context) {
	id, b64s, err := captcha.NewCaptcha().GenerateCaptcha()
	logger.LogIf(err)
	c.JSON(http.StatusOK, gin.H{
		"captcha_id":    id,
		"captcha_image": b64s,
	})
}
