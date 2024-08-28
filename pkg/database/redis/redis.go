package redis

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

type Interface interface {
	SetData(key, value string, ttl time.Duration) error
	GetData(key string) (any,error)
}

type redisClient struct {
	Client *redis.Client
}

func Init() Interface {
	addr := os.Getenv("REDIS_HOST")
	password := os.Getenv("REDIS_PASSWORD")

	ctx := context.Background()
	rdb := redis.NewClient(&redis.Options{
		Addr:    addr,
		Password: password,
	})

	err := rdb.Ping(ctx).Err()

	if err != nil {
		log.Fatalf("Error connecting to redis: %v", err)
	}

	return &redisClient{Client: rdb}
}

func (r *redisClient) SetData(key, value string, ttl time.Duration) error {
	err := r.Client.Set(context.Background(), key, value, ttl).Err()

	if err != nil {
		return err
	}
	return nil
}

func (r *redisClient) GetData(key string) (any,error) {
	data, err := r.Client.Get(context.Background(), key).Result()

	if err != nil {
		return nil ,err
	}
	return data, nil
}