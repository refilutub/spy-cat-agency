package dtos

import (
	"time"

	"github.com/google/uuid"
)

type SpyCatRequest struct {
	Name            string  `json:"name"`
	ExperienceYears int     `json:"experience_years"`
	Breed           string  `json:"breed"`
	Salary          float64 `json:"salary"`
}

type SpyCatSingleResponseDTO struct {
	Id              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	ExperienceYears int       `json:"experience_years"`
	Breed           string    `json:"breed"`
	Salary          float64   `json:"salary"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type BreedName struct {
	Name string `json:"name"`
}
