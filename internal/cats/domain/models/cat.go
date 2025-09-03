package models

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type Cat struct {
	ID              uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	Name            string         `json:"name"`
	ExperienceYears int            `json:"experience_years"`
	Breed           string         `json:"breed"`
	Salary          float64        `gorm:"type:numeric(12,2)" json:"salary"`
	UpdatedAt       time.Time      `json:"updated_at"`
	CreatedAt       time.Time      `json:"-"`
	DeletedAt       gorm.DeletedAt `gorm:"index" json:"-"`
}
