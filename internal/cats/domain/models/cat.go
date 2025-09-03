package models

import (
	"github.com/google/uuid"
	"time"
)

type Cat struct {
	ID              uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`
	Name            string    `json:"name"`
	ExperienceYears int       `json:"experience_years"`
	Breed           string    `json:"breed"`
	Salary          float32   `json:"salary"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedAt       time.Time
}
