package link

import (
	"time"

	"github.com/czjge/gohub/pkg/app"
	"github.com/czjge/gohub/pkg/cache"
	"github.com/czjge/gohub/pkg/database"
	"github.com/czjge/gohub/pkg/helpers"
	"github.com/czjge/gohub/pkg/paginator"

	"github.com/gin-gonic/gin"
)

func Get(idstr string) (link Link) {
	database.DB().Where("id", idstr).First(&link)
	return
}

func GetBy(field, value string) (link Link) {
	// database.DB().Where("? = ?", field, value).First(&link)
	database.DB().Where(field+" = ?", value).First(&link)
	return
}

func All() (links []Link) {
	database.DB().Find(&links)
	return
}

func IsExist(field, value string) bool {
	var count int64
	// database.DB().Model(Link{}).Where("? = ?", field, value).Count(&count)
	database.DB().Model(Link{}).Where(field+" = ?", value).Count(&count)
	return count > 0
}

func Paginate(c *gin.Context, perPage int) (links []Link, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB().Model(Link{}),
		&links,
		app.V1URL(database.TableName(&Link{})),
		perPage,
	)
	return
}

func AllCached() (links []Link) {

	cacheKey := "links:all"
	expireTime := 120 * time.Minute
	cache.GetObject(cacheKey, &links)

	if helpers.Empty(links) {
		links = All()
		if helpers.Empty(links) {
			return links
		}
		cache.Set(cacheKey, links, expireTime)
	}
	return
}
