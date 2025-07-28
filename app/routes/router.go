package routes

import (
	"docryze-backend/app/middleware"

	"github.com/gin-gonic/gin"
)

// router.go
func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 全局中间件
	r.Use(middleware.Logger(), middleware.Recovery())

	// API 路由分组
	// api := r.Group("/api/v1")
	// {
	// userCtrl := controllers.NewUserController()
	// api.POST("/users", userCtrl.CreateUser)
	// api.GET("/users/:id", userCtrl.GetUser)
	// }

	return r
}
