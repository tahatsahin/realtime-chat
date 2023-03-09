package redisrepo

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
	"os"
)

var redisClient *redis.Client

func InitialiseRedis() *redis.Client {
	conn := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_CONNECTION_STRING"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       0,
	})

	// checking if redis is connected
	pong, err := conn.Ping(context.Background()).Result()
	if err != nil {
		log.Fatalf("redis connection failed... %v", err)
	}
	log.Printf("connected to redis, ping %v", pong)

	redisClient = conn
	return redisClient
}
