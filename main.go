package main

import (
	"NTT_DATA/config"
	db "NTT_DATA/database"
	"NTT_DATA/routes"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	// "github.com/joho/godotenv"
)


func main() {
	// database connection and migration
	// run migrate when modification are done in database table
	db.Init()
	db.Migrate()
	router := initRouter()
	port := config.GetEnv("SERVER_PORT","8000")
	fmt.Println("service started on port : " + port)
	router.Run(":" + port)
}

func initRouter() *gin.Engine {
	router := gin.Default()
	corsConfig := corsHandler()
	router.Use(cors.New(corsConfig))
	router.Use(gin.Logger())
	api := router.Group("/api")

	// initialize those routers first, which does not need authentication

	routes.PlantRoutes(api)
	return router
}

func corsHandler()cors.Config{
		// Configure CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000","http://localhost:3001"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"}
	config.AllowCredentials = true // Allow credentials

	return config
}