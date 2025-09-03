package main

import (
	"log"
	"os"
	"spy-cat-agency/internal/cats/application/services"
	"spy-cat-agency/internal/cats/interfaces/handlers"
	"spy-cat-agency/internal/cats/repository"
	services2 "spy-cat-agency/internal/missions/application/services"
	handlers2 "spy-cat-agency/internal/missions/interfaces/handlers"
	repository2 "spy-cat-agency/internal/missions/repository"
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

	db, err := db.InitDB()
	if err != nil {
		log.Fatalf("Error initializing database: %v", err)
	}

	spyCatsRepo := repository.NewSpyCatRepository(db)
	spyCatsService := services.NewSpyCatService(spyCatsRepo)
	spyCatsHandler := &handlers.SpyCatHandler{
		Service: spyCatsService,
	}

	missionsRepo := repository2.NewMissionsRepository(db)
	missionsService := services2.NewMissionsService(missionsRepo)
	missionsHandler := handlers2.NewMissionsHandler(missionsService)

	r := router.SetUpRouter(spyCatsHandler, missionsHandler)

	apiPort := os.Getenv("API_PORT")
	if err := r.Run(":" + apiPort); err != nil {
		log.Fatalf("Error starting server: %v", err)
	}

}
