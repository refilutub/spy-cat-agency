package main

import (
	"log"
	"os"
	"spy-cat-agency/internal/shared/db"
	"spy-cat-agency/internal/shared/router"

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

	r := router.SetUpRouter()

	apiPort := os.Getenv("API_PORT")
	if err := r.Run(":" + apiPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
