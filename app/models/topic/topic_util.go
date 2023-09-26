package topic

import (
	"github.com/czjge/gohub/pkg/app"
	"github.com/czjge/gohub/pkg/database"
	"github.com/czjge/gohub/pkg/paginator"
	"gorm.io/gorm/clause"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (topic Topic) {
	database.DB().Preload(clause.Associations).Where("id", idstr).First(&topic)
	return
}

func GetBy(field, value string) (topic Topic) {
	// database.DB().Where("? = ?", field, value).First(&topic)
	database.DB().Where(field+" = ?", value).First(&topic)
	return
}

func GetByIDWithAssociation(id uint64) (topic Topic) {
	database.DB().Preload("User").Preload("Category").First(&topic, id)
	return
}

func All() (topics []Topic) {
	database.DB().Find(&topics)
	return
}

func IsExist(field, value string) bool {
	var count int64
	// database.DB().Model(Topic{}).Where("? = ?", field, value).Count(&count)
	database.DB().Model(Topic{}).Where(field+" = ?", value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (topics []Topic, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB().Model(Topic{}),
		&topics,
		app.V1URL(database.TableName(&Topic{})),
		perPage,
	)
	return
}
