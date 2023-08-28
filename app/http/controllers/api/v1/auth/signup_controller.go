package auth

import (
	"fmt"
	"net/http"

	"github.com/czjge/gohub/app/models/user"
	"github.com/gin-gonic/gin"
)

type SignupController struct {
}

func (sc *SignupController) IsPhoneExist(c *gin.Context) {

	type PhoneExistRequest struct {
		Phone string `json:"phone"`
	}
	request := PhoneExistRequest{}

	if err := c.ShouldBindJSON(&request); err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
