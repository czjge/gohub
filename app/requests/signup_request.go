package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type SignupPhoneExistRequest struct {
	Phone string `json:"phone,omitempty" valid:"phone"`
}

type SignupEmailExistRequest struct {
	Email string `json:"email,omitempty" valid:"email"`
}

func SignupPhoneExist(data any, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"phone": []string{"required", "digits:11"},
	}

	messages := govalidator.MapData{
		"phone": []string{
			"required:手机号必填",
			"digits:手机号为11位数字",
		},
	}

	return validate(data, rules, messages)
}

func SignupEmailExist(data any, c *gin.Context) map[string][]string {

	rules := govalidator.MapData{
		"email": []string{"required", "min:4", "max:30", "email"},
	}

	messages := govalidator.MapData{
		"email": []string{
			"required:邮箱必填",
			"min:邮箱长度至少为4",
			"max:邮箱长度最多为30",
			"email:邮箱格式不正确",
		},
	}

	return validate(data, rules, messages)
}
