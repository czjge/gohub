package v1

import (
	"github.com/czjge/gohub/app/models/topic"
	"github.com/czjge/gohub/app/policies"
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
		_topic := topic.GetByIDWithAssociation(topicModel.ID)
		response.Created(c, _topic)
	} else {
		response.Abort500(c, "创建失败，请稍后尝试~")
	}
}

func (ctrl *TopicsController) Update(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	request := requests.TopicRequest{}
	if ok := requests.Validate(c, &request, requests.TopicSave); !ok {
		return
	}

	topicModel.Title = request.Title
	topicModel.Body = request.Body
	topicModel.CategoryID = request.CategoryID
	rowsAffected := topicModel.Save()
	if rowsAffected > 0 {
		_topic := topic.GetByIDWithAssociation(topicModel.ID)
		response.Data(c, _topic)
	} else {
		response.Abort500(c, "更新失败")
	}
}

func (ctrl *TopicsController) Delete(c *gin.Context) {

	topicModel := topic.Get(c.Param("id"))
	if topicModel.ID == 0 {
		response.Abort404(c)
		return
	}

	if ok := policies.CanModifyTopic(c, topicModel); !ok {
		response.Abort403(c)
		return
	}

	rowsAffected := topicModel.Delete()
	if rowsAffected > 0 {
		response.Success(c)
		return
	}

	response.Abort500(c, "删除失败")
}

func (ctrl *TopicsController) Index(c *gin.Context) {

	request := requests.PaginationRequest{}
	if ok := requests.Validate(c, &request, requests.Pagination); !ok {
		return
	}

	data, pager := topic.Paginate(c, 10)
	response.JSON(c, gin.H{
		"data":  data,
		"pager": pager,
	})
}
