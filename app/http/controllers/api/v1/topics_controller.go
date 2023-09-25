package v1

import (
	"github.com/czjge/gohub/app/models/topic"
	"github.com/czjge/gohub/app/requests"
	"github.com/czjge/gohub/pkg/auth"
	"github.com/czjge/gohub/pkg/response"

	"github.com/gin-gonic/gin"
)

type TopicsController struct {
	BaseAPIControler
}

func (ctrl *TopicsController) Store(c *gin.Context) {

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel := topic.Topic{
		Title:      request.Title,
		Body:       request.Body,
		CategoryID: request.CategoryID,
		UserID:     auth.CurrentUID(c),
	}
	topicModel.Create()
	if topicModel.ID > 0 {
		response.Created(c, topicModel)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}
