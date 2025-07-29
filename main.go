package main

import (
	conf "auth_service_medods/config"
	"auth_service_medods/internal/app"
	"auth_service_medods/transport/routes"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

// @title						auth service for medods
// @version					1.0
// @BasePath					/
// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func main() {
	app.InitUsers()
	server := gin.Default()
	routes.SetupRoutes(server)
	app.PrintUsers()
	server.Run(conf.Config.Server.Host + ":" + conf.Config.Server.Port)

}
