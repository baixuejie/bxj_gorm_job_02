package config

import (
	"context"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	DB       int    `mapstructure:"db" json:"db" yaml:"db"`
}

func InitRedis() {
	REDIS = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", CONFIG.Redis.Host, CONFIG.Redis.Port),
		Password: CONFIG.Redis.Password,
		DB:       CONFIG.Redis.DB,
	})
	_, err := REDIS.Ping(context.Background()).Result()
	if err != nil {
		log.Println("redis connect ping failed, err:", err)
	}

}
