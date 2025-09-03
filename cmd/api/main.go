package main

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
	"os"
	"spy-cat-agency/internal/shared/db"

	_ "spy-cat-agency/docs"
)

// @title Spy Cat Agency API
// @version 1.0
// @description API for managing spy cats and their missions.

// @host localhost:8080
// @BasePath /api
func main() {

	_, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	apiPort := os.Getenv("API_PORT")
	if err := r.Run(":" + apiPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
