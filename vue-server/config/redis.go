package initialize

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
)

func Redis(*RedisClient, error) {
	rdb := redis.NewClient(&redis.Options{
		Addr:     viper.GetString("redis.url"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	pong, err := client.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	} else {
		return &RedisClient{}, nil
	}
}
