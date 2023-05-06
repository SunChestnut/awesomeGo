package redisdemo

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
)

const redisAddr = "127.0.0.1:6379"

func NewClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     redisAddr,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func SaveToRedis(client *redis.Client, key string, value string, ctx context.Context) {
	err := client.Set(ctx, key, value, 0).Err()
	if err != nil {
		log.Fatalf("[SaveToRedis] error save key to redis:%v", err)
	}
}

func GetFromRedis(client *redis.Client, key string, ctx context.Context) string {
	// 当待查询的 key 不存在时，redis 会返回 redis:nil 的错误信息
	result, err := client.Get(ctx, key).Result()
	if err != nil {
		return ""
	}
	return result
}
