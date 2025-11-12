package config

import (
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Server struct {
	Jwt   Jwt
	Mysql Mysql
	Redis Redis
}

var DB *gorm.DB
var REDIS redis.UniversalClient
