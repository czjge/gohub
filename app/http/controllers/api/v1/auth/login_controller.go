package auth

import (
	v1 "github.com/czjge/gohub/app/http/controllers/api/v1"
	"github.com/czjge/gohub/app/requests"
	"github.com/czjge/gohub/pkg/auth"
	"github.com/czjge/gohub/pkg/jwt"
	"github.com/czjge/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type LoginController struct {
	v1.BaseAPIControler
}

func (lc *LoginController) LoginByPhone(c *gin.Context) {

	request := requests.LoginByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPhone); !ok {
		return
	}

	user, err := auth.LoginByPhone(request.Phone)
	if err != nil {
		response.Error(c, err, "账号不存在")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)

		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

func (lc *LoginController) LoginByPassword(c *gin.Context) {

	request := requests.LoginByPaswordRequest{}
	if ok := requests.Validate(c, &request, requests.LoginByPassword); !ok {
		return
	}

	user, err := auth.Attempt(request.LoginID, request.Password)
	if err != nil {
		response.Unauthorized(c, "账号不存在或者密码错误")
	} else {
		token := jwt.NewJWT().IssueToken(user.GetStringID(), user.Name)
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}

func (lc *LoginController) RefreshToken(c *gin.Context) {

	token, err := jwt.NewJWT().RefreshToken(c)

	if err != nil {
		response.Error(c, err, "令牌刷新失败")
	} else {
		response.JSON(c, gin.H{
			"token": token,
		})
	}
}
