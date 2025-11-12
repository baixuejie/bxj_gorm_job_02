package config

import (
	"fmt"

	"github.com/redis/go-redis/v9"
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

}
