package v1

import (
	"github.com/czjge/gohub/app/models/user"
	"github.com/czjge/gohub/pkg/auth"
	"github.com/czjge/gohub/pkg/response"
	"github.com/gin-gonic/gin"
)

type UsersController struct {
	BaseAPIControler
}

func (ctrl *UsersController) CurrentUser(c *gin.Context) {
	userModel := auth.CurrentUser(c)
	response.Data(c, userModel)
}

func (ctrl *UsersController) Index(c *gin.Context) {
	data, pager := user.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
