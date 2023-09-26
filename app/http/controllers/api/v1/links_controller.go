package v1

import (
	"github.com/czjge/gohub/app/models/link"
	"github.com/czjge/gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type LinksController struct {
	BaseAPIControler
}

func (ctrl *LinksController) Index(c *gin.Context) {
	links := link.All()
	response.Data(c, links)
}
