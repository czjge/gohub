package limiter

import (
	"strings"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/logger"
	"github.com/czjge/gohub/pkg/redis"
	"github.com/gin-gonic/gin"
	limiterlib "github.com/ulule/limiter/v3"
	sredis "github.com/ulule/limiter/v3/drivers/store/redis"
)

func GetKeyIP(c *gin.Context) string {
	return c.ClientIP()
}

func GetKeyRouteWithIP(c *gin.Context) string {
	return routeToKeyString(c.FullPath()) + c.ClientIP()
}

// 检查请求是否超额
func CheckRate(c *gin.Context, key string, formatted string) (limiterlib.Context, error) {

	var context limiterlib.Context
	rate, error := limiterlib.NewRateFromFormatted(formatted)
	if error != nil {
		logger.LogIf(error)
		return context, error
	}

	store, err := sredis.NewStoreWithOptions(redis.Redis().Client, limiterlib.StoreOptions{
		Prefix: config.GetConfig().App.Name + ":limiter",
	})
	if err != nil {
		logger.LogIf(error)
		return context, error
	}

	limiterObj := limiterlib.New(store, rate)

	if c.GetBool("limiter-once") {
		// 不增加访问次数
		return limiterObj.Peek(c, key)
	} else {
		// 确保多个路由组里调用 LimitIP 进行限流时，只增加一次访问次数
		c.Set("limiter-once", true)
		// 取结果且增加访问次数
		return limiterObj.Get(c, key)
	}

}

func routeToKeyString(routeName string) string {
	routeName = strings.ReplaceAll(routeName, "/", "-")
	routeName = strings.ReplaceAll(routeName, ":", "_")
	return routeName
}
