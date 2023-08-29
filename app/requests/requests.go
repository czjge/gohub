package requests

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(any, *gin.Context) map[string][]string

func validate(data any, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {

	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		TagIdentifier: "valid",
		Messages:      messages,
	}

	return govalidator.New(opts).ValidateStruct()
}

func Validate(c *gin.Context, obj any, handler ValidatorFunc) bool {

	if err := c.ShouldBind(obj); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "请求解析错误",
			"error":   err.Error(),
		})
		fmt.Println(err.Error())
		return false
	}

	errs := handler(obj, c)

	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"message": "表单验证失败",
			"errors":  errs,
		})
		return false
	}

	return true
}
