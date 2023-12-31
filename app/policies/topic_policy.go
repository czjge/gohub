package policies

import (
	"github.com/czjge/gohub/app/models/topic"
	"github.com/czjge/gohub/pkg/auth"
	"github.com/gin-gonic/gin"
)

func CanModifyTopic(c *gin.Context, _topic topic.Topic) bool {
	return auth.CurrentUID(c) == _topic.UserID
}
