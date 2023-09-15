package auth

import (
	"errors"

	"github.com/czjge/gohub/app/models/user"
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
