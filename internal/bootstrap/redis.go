package bootstrap

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/nurchulis/idempotency/config"
)

func RegistryRedis(conf config.RedisConfig) *redis.Client {
	rc := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", conf.Addr, conf.Port),
		Password: conf.Password,
		DB:       conf.DB,
	})

	return rc
}
