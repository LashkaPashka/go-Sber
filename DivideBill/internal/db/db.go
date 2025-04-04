package db

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
)

type DbRedis struct {
	Pool *redis.Client
}

func New() *DbRedis {
	redis := DbRedis{
		Pool: redis.NewClient(&redis.Options{
			Addr: "localhost:6379",
			Password: "mypassword",
			DB: 0,
		}),
	}

	if err := redis.Pool.Ping(context.Background()).Err(); err != nil {
		log.Fatal("не удалось подключиться к БД redis")
		return nil
	}

	return &redis
}