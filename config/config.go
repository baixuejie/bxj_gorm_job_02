package config

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

type Server struct {
	Jwt   Jwt
	Mysql Mysql
	Redis Redis
}

var DB *gorm.DB
var REDIS *redis.Client
