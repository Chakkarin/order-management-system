package redisconn

import (
	"context"
	"fmt"
	"log"
	"services/auth-service/internal/config"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis(conf *config.Redis) *redis.Client {

	// Create Redis connection options
	opt := &redis.Options{
		Addr:     fmt.Sprintf("%s:%s", conf.Host, conf.Port),
		Username: conf.User,
		Password: conf.Password,
	}

	// สร้าง Redis Client
	rdb := redis.NewClient(opt)

	// ทดสอบการเชื่อมต่อ
	if err := rdb.Ping(context.Background()).Err(); err != nil {
		panic("❌ Failed to connect to Redis: " + err.Error())
	}

	log.Println("✅ connected to redis...")

	return rdb
}
