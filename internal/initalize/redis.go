package initalize

import (
	"context"
	"fmt"

	"github.com/huuloc2026/go-backend/global"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
)

var ctx = context.Background()

func InitRedis() {
	r := global.Config.Redis
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%v", r.Host, r.Port),
		Username: r.Username,
		Password: r.Password, // no password set
		DB:       r.Db,       // use default DB
		PoolSize: 10,
	})
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		global.Logger.Error("Failed to connect to Redis", zap.Error(err))
		panic(fmt.Sprintf("Failed to connect to Redis: %v", err))
	}
	global.Logger.Info("Redis connected")
	global.Rdb = rdb

}
func RedisExample() {
	key := "test_key"
	value := "Hello, Redis!"

	// Set a value
	err := global.Rdb.Set(ctx, key, value, 0).Err()
	if err != nil {
		global.Logger.Error("Failed to set value in Redis", zap.Error(err))
		return
	}

	// Get the value
	val, err := global.Rdb.Get(ctx, key).Result()
	if err != nil {
		global.Logger.Error("Failed to get value from Redis", zap.Error(err))
		return
	}

	global.Logger.Info("Successfully retrieved value from Redis",
		zap.String("key", key),
		zap.String("value", val),
	)
}
