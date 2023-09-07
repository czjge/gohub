package routes

import (
	"github.com/czjge/gohub/app/http/controllers/api/v1/auth"
	"github.com/gin-gonic/gin"
)

// Register API routes.
func RegisterAPIRoutes(r *gin.Engine) {

	// v1 route group
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController) // 分配内存，返回指向该类型的零值的指针
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
			authGroup.POST("/signup/email/exist", suc.IsEmailExist)

			vcc := new(auth.VerifyCodeController)
			authGroup.POST("/verify-codes/captcha", vcc.ShowCaptcha)
		}
	}
}
