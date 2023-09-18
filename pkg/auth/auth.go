package auth

import (
	"errors"

	"github.com/czjge/gohub/app/models/user"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/gin-gonic/gin"
)

func Attempt(loginID string, password string) (user.User, error) {

	userModel := user.GetByMulti(loginID)
	if userModel.ID == 0 {
		return user.User{}, errors.New("账号不存在")
	}

	if !userModel.ComparePassword(password) {
		return user.User{}, errors.New("密码错误")
	}

	return userModel, nil
}

func LoginByPhone(phone string) (user.User, error) {

	userModel := user.GetByPhone(phone)
	if userModel.ID == 0 {
		return user.User{}, errors.New("手机号未注册")
	}

	return userModel, nil
}

func CurrentUser(c *gin.Context) user.User {
	userModel, ok := c.MustGet("current_user").(user.User)
	if !ok {
		logger.LogIf(errors.New("无法获取当前用户"))
		return user.User{}
	}

	return userModel
}

func CurrentUID(c *gin.Context) string {
	return c.GetString("current_user_id")
}
