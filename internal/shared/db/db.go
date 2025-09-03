package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
	"spy-cat-agency/internal/shared/models"
)

func InitDB() (*gorm.DB, error) {

	dbHost := os.Getenv("POSTGRES_HOST")
	dbPort := os.Getenv("POSTGRES_PORT")
	dbUser := os.Getenv("POSTGRES_USER")
	dbPassword := os.Getenv("POSTGRES_PASSWORD")
	dbName := os.Getenv("POSTGRES_DB")

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		dbHost, dbUser, dbPassword, dbName, dbPort,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
		return nil, err
	}

	for _, m := range models.AllModels {
		if err := db.AutoMigrate(m); err != nil {
			log.Fatal(err)
		}
	}
	if err != nil {
		log.Fatalf("Error migrating db: %v", err)
		return nil, err
	}

	return db, nil
}
