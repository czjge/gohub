package bootstrap

import (
	"github.com/czjge/gohub/pkg/cache"
	"github.com/czjge/gohub/pkg/redis"
)

func SetupCache() {
	redis := redis.Redis("cache")
	driver := cache.NewRedisStore2(redis)

	cache.InitWithCacheStore(driver)
}
