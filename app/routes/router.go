package routes

import (
	"docryze-backend/app/controllers"
	"docryze-backend/app/middleware"

	"github.com/gin-gonic/gin"
)

// router.go
func SetupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 全局中间件
	r.Use(middleware.Logger(), middleware.Recovery())

	// API 路由分组
	api := r.Group("/api/v1")
	{
		api.POST("/createDocument", controllers.DocController.CreateDocument)
	}

	return r
}
