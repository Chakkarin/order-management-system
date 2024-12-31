package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

func ConnectRedis() *redis.Client {
	dsn := os.Getenv("REDIS_URL")
	if dsn == "" {
		panic("❌ REDIS_URL is not set")
	}

	opt, err := redis.ParseURL(dsn)
	if err != nil {
		panic(fmt.Sprintf("❌ Invalid Redis DSN: %v", err))
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
