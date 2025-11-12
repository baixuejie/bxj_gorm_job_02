package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	DbName   string `mapstructure:"db-name" json:"dbName" yaml:"db-name"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
}

func InitMysql() {
	DB, _ = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		CONFIG.Mysql.Username,
		CONFIG.Mysql.Password,
		CONFIG.Mysql.Host,
		CONFIG.Mysql.Port,
		CONFIG.Mysql.DbName)))
}
