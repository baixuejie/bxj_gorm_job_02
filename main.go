package main

import (
	"bxj_gorm_job_02/config"
	"bxj_gorm_job_02/router"
)

func main() {
	initRun()
}

func initRun() {
	// 初始化配置
	err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	//初始化数据库
	config.InitMysql()
	//初始化redis
	config.InitRedis()
	//初始化路由
	router.InitRouter()
}
