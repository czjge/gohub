package bootstrap

import (
	"fmt"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/redis"
)

func SetupRedis() {

	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetConfig().Redis.Host, config.GetConfig().Redis.Port),
		config.GetConfig().Redis.Username,
		config.GetConfig().Redis.Password,
		config.GetConfig().Redis.DB,
	)
}
