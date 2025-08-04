package main

import (
	"github.com/LareinaCornia/api-gateway/handlers"
	"github.com/LareinaCornia/api-gateway/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// 注册中间件
	router.Use(middleware.CORSMiddleware())
	router.Use(middleware.RateLimitMiddleware())
	router.Use(middleware.JWTMiddleware())

	// 路由分组
	api := router.Group("/api")
	{
		api.POST("/login", handlers.LoginHandler)
		api.GET("/user/:id", handlers.UserHandler)
		api.POST("/tasks", handlers.CreateTaskHandler)
		api.GET("/notifications", handlers.NotificationsHandler)
	}

	router.Run(":8080")
}
