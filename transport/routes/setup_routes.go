package routes

import (
	"auth_service_medods/transport/middlewares"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func SetupRoutes(router *gin.Engine) {
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AddAllowHeaders("Authorization")
	router.Use(cors.New(config))
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	auth := router.Group("/auth")
	auth.POST("/login", Login)
	auth.POST("/refresh", Refresh)
	auth.POST("/logout", Logout)
	data := router.Group("/data")
	data.Use(middlewares.Jwt_middleware)
	data.GET("/me", GetUUID)
}
