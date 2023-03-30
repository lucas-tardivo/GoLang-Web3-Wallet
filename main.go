package main

import (
	"log"
	"net/http"
	"web3-wallet/router"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/wpcodevo/golang-gorm-postgres/initializers"
)

var (
	server              *gin.Engine
)

func init() {
	server = gin.Default()
}

func main() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("? Could not load environment variables", err)
	}

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOrigins = []string{"http://localhost:8000", config.ClientOrigin}

	server.Use(cors.New(corsConfig))

	serverRouter := server.Group("/api")
	serverRouter.GET("/healthchecker", func(ctx *gin.Context) {
		message := "Welcome to Golang with Gorm and Postgres"
		ctx.JSON(http.StatusOK, gin.H{"status": "success", "message": message})
	})

	router.AuthRoute(serverRouter)
	router.ERC20Route(serverRouter)
	router.ERC721Route(serverRouter)
	log.Fatal(server.Run(":" + config.ServerPort))
}