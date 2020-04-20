package database

import (
	"time"

	"github.com/fuadajip/stripe-react-go/server/shared/config"
	"github.com/go-redis/redis"
)

type (

	// RedisInterface returns implementations of redis database methods
	RedisInterface interface {
		OpenRedisConn() (*redis.Client, error)
	}

	// redisConf is a struct that map given config
	redisConf struct {
		SharedConfig config.ImmutableConfigInterface
	}
)

func (d *redisConf) OpenRedisConn() (*redis.Client, error) {
	logger.Info("Start open redis connection...")

	client := redis.NewClient(&redis.Options{
		Addr:        d.SharedConfig.GetRedisHost(),
		Password:    d.SharedConfig.GetRedisPassword(),
		DB:          0,
		PoolSize:    64,
		ReadTimeout: 10 * time.Second,
	})

	_, err := client.Ping().Result()
	if err != nil {
		return client, err
	}

	return client, nil
}

// NewRedis is a factory that return interface of its implementation
func NewRedis(config config.ImmutableConfigInterface) RedisInterface {
	if config == nil {
		panic("[CONFIG] immutable config is required")
	}

	return &redisConf{config}
}
