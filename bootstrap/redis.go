package bootstrap

import (
	"fmt"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/redis"
)

func SetupRedis() {

	configs := config.GetConfig().Redis
	redis.Once.Do(func() {
		if redis.RedisConllections == nil {
			redis.RedisConllections = make(map[string]*redis.RedisClient, len(redis.RedisConllections))
		}
		for name, config := range configs {
			redis.RedisConllections[name] = redis.NewClient(
				fmt.Sprintf("%v:%v", config.Host, config.Port),
				config.Username,
				config.Password,
				config.DB,
			)
		}
	})
}
