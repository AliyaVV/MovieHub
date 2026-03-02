package redis

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type SearchLogger interface {
	Log(ctx context.Context, query string, resultCount int) error
}

type RedisSearchLogger struct {
	client *redis.Client
}

func NewRedisSearchLogger(client *redis.Client) SearchLogger {
	return &RedisSearchLogger{client: client}
}

func (redislog *RedisSearchLogger) Log(ctx context.Context, query string, count int) error {
	key := fmt.Sprintf("search time:%d", time.Now().UnixNano())
	value := map[string]interface{}{
		"query":     query,
		"count":     count,
		"timestamp": time.Now(),
	}
	data, _ := json.Marshal(value)
	return redislog.client.Set(ctx, key, data, 1*time.Hour).Err()
}

func InitRedis() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
}
