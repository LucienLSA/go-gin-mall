package cache

import (
	"context"
	"fmt"
	"ginmall/conf"
	"ginmall/pkg/util/logging"

	"github.com/redis/go-redis/v9"
)

// RedisClient Redis缓存客户端单例
var RedisClient *redis.Client
var RedisContext = context.Background()

// Redis 在中间件中初始化redis链接
func InitRedis() {
	rConfig := conf.Config.Redis
	// db, _ := strconv.ParseUint(os.Getenv("REDIS_DB"), 10, 64)
	addr := rConfig.RedisHost + ":" + rConfig.RedisPort
	client := redis.NewClient(&redis.Options{
		Addr:            addr,
		Password:        rConfig.RedisPassword,
		DB:              rConfig.RedisDbName,
		MaxRetries:      rConfig.RedisMaxRetries,
		MaxIdleConns:    rConfig.RedisMaxIdle,
		MinIdleConns:    rConfig.RedisMinIdle,
		ConnMaxIdleTime: rConfig.RedisIdleTimeout,
	})

	_, err := client.Ping(RedisContext).Result()

	if err != nil {
		logging.LogrusObj.Panic("连接Redis失败", err)
	}
	RedisClient = client
	fmt.Println("连接Redis成功")
}
