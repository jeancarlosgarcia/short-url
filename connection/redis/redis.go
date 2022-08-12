package redis

import (
	"context"
	"fmt"
	"time"

	"short-url/config"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/gommon/log"
)

func NewRedisConnection() *redis.Client {
	redisConfig := config.Environments().Redis

	address := fmt.Sprintf("%s:%s", redisConfig.Host, redisConfig.Port)

	client := redis.NewClient(&redis.Options{
		Addr:         address,
		MinIdleConns: 3,
		IdleTimeout:  3 * time.Minute,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		log.Errorf("Error trying to connect to redis %v: %#v", err.Error())
	} else {
		log.Infof("Connected to redis %v successfully: %#v", address)
	}

	return client
}