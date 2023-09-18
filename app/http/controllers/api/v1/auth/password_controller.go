package auth

import (
	"github.com/czjge/gohub/app/models/user"
	"github.com/czjge/gohub/app/requests"
	"github.com/czjge/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type PasswordController struct{}

func (pc *PasswordController) ResetByPhone(c *gin.Context) {

	request := requests.ResetByPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.ResetByPhone); !ok {
		return
	}

	userModel := user.GetByPhone(request.Phone)
	if userModel.ID == 0 {
		response.Abort404(c)
	} else {
		userModel.Password = request.Password
		userModel.Save()

		response.Success(c)
	}
}
