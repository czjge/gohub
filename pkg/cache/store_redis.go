package cache

import (
	"time"

	"github.com/czjge/gohub/config"
	"github.com/czjge/gohub/pkg/redis"
)

type RedisStore struct {
	RedisClient *redis.RedisClient
	KeyPrefix   string
}

func NewRedisStore(address string, username string, password string, db int) *RedisStore {
	rs := &RedisStore{}
	rs.RedisClient = redis.NewClient(address, username, password, db)
	rs.KeyPrefix = config.GetConfig().App.Name + ":cache:"
	return rs
}

func NewRedisStore2(redis *redis.RedisClient) *RedisStore {
	rs := &RedisStore{}
	rs.RedisClient = redis
	rs.KeyPrefix = config.GetConfig().App.Name + ":cache:"
	return rs
}

func (s *RedisStore) Set(key string, value string, expireTime time.Duration) {
	s.RedisClient.Set(s.KeyPrefix+key, value, expireTime)
}

func (s *RedisStore) Get(key string) string {
	return s.RedisClient.Get(s.KeyPrefix + key)
}

func (s *RedisStore) Has(key string) bool {
	return s.RedisClient.Has(s.KeyPrefix + key)
}

func (s *RedisStore) Forget(key string) {
	s.RedisClient.Del(s.KeyPrefix + key)
}

func (s *RedisStore) Forever(key string, value string) {
	s.RedisClient.Set(s.KeyPrefix+key, value, 0)
}

func (s *RedisStore) Flush() {
	s.RedisClient.FlushDB()
}

func (s *RedisStore) Increment(parameters ...any) {
	s.RedisClient.Increment(parameters...)
}

func (s *RedisStore) Decrement(parameters ...any) {
	s.RedisClient.Decrement(parameters...)
}

func (s *RedisStore) IsAlive() error {
	return s.RedisClient.Ping()
}
