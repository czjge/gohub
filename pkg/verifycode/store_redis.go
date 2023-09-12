package verifycode

import (
	"time"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/app"
	"github.com/czjge/gohub/pkg/redis"
)

// 需实现 verifycode.Store interface
type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func (s *RedisStore) Set(key string, value string) bool {

	config := config.GetConfig().Verifycode
	ExpireTime := time.Minute * time.Duration(config.ExpireTime)
	if app.IsLocal() {
		ExpireTime = time.Minute * time.Duration(config.DebugExpireTime)
	}

	return s.RedisClient.Set(s.KeyPrefix+key, value, ExpireTime)
}

func (s *RedisStore) Get(key string, clear bool) (value string) {
	key = s.KeyPrefix + key
	val := s.RedisClient.Get(key)
	if clear {
		s.RedisClient.Del(key)
	}
	return val
}

func (s *RedisStore) Verify(key, answer string, clear bool) bool {
	v := s.Get(key, clear)
	return v == answer
}
