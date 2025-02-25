package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var (
	RedisClient *redis.Client
	Ctx         = context.Background()
)

func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"), // "localhost:6379"
		Password: "",                      // No password by default
		DB:       0,                        // Default DB
	})

	_, err := RedisClient.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("❌ Failed to connect to Redis:", err)
	}

	log.Println("🔥 Connected to Redis!")
}