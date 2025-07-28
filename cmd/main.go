package main

import (
	"docryze-backend/app/config"
	"docryze-backend/app/repositories"
	"docryze-backend/app/routes"
)

func main() {
	// 加载配置
	config, err := config.LoadConfig()
	if err != nil {
		panic(err)
	}

	println(config.Database.DSN)
	println(config.Server.Port)

	// 初始化数据库
	repositories.InitRepository(config.Database.DSN)

	// 配置路由
	router := routes.SetupRouter()

	// 启动服务器
	router.Run(":" + config.Server.Port)
}
