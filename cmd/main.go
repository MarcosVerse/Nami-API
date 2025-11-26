// @title Nami API
// @version 1.0
// @description API para controlar despesas
// @host localhost:8081
// @BasePath /
package main

import (
	"github.com/MarcosVerse/nami/internal/config"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/MarcosVerse/nami/internal/routes"
	"github.com/gin-gonic/gin"

	_ "github.com/MarcosVerse/nami/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	config.LoadConfig()
	repository.Connect()

	// Servidor principal da API

	r := gin.Default()
	routes.RegisterRoutes(r)

	// Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run(":8081") // API + Swagger
}
