package dtos

import (
	"spy-cat-agency/internal/missions/domain/models"
	"time"

	"github.com/google/uuid"
)

type CreateMissionRequest struct {
	Targets []TargetDTO `json:"targets"`
}

type TargetDTO struct {
	Name          string               `json:"name"`
	Country       string               `json:"country"`
	Notes         string               `json:"notes"`
	CompleteState models.CompleteState `json:"complete_state"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

type MissionTargetRequestDTO struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Notes   string `json:"notes"`
}

type MissionCreateResponseDTO struct {
	Id            uuid.UUID            `json:"id"`
	CompleteState models.CompleteState `json:"complete_state"`
	Targets       []TargetDTO          `json:"targets"`
	UpdatedAt     time.Time            `json:"updated_at"`
}

type MissionResponseDTO struct {
	ID            uuid.UUID            `json:"id"`
	CatID         *uuid.UUID           `json:"cat_id"`
	Cat           *CatDTO              `json:"cat,omitempty"`
	Targets       []TargetResponseDTO  `json:"targets"`
	CompleteState models.CompleteState `json:"complete_state"`
	UpdatedAt     time.Time            `json:"updated_at"`
	CreatedAt     time.Time            `json:"created_at"`
}

type CatDTO struct {
	ID              uuid.UUID `json:"id"`
	Name            string    `json:"name"`
	ExperienceYears int       `json:"experience_years"`
	Breed           string    `json:"breed"`
	Salary          float64   `json:"salary"`
}

type TargetResponseDTO struct {
	ID            uuid.UUID            `json:"id"`
	MissionID     uuid.UUID            `json:"mission_id"`
	Name          string               `json:"name"`
	Country       string               `json:"country"`
	Notes         string               `json:"notes"`
	CompleteState models.CompleteState `json:"complete_state"`
	UpdatedAt     time.Time            `json:"updated_at"`
	CreatedAt     time.Time            `json:"created_at"`
}

type UpdateMissionRequest struct {
	CompleteState *models.CompleteState `json:"complete_state,omitempty"`
}

type AddTargetRequest struct {
	Name    string `json:"name"`
	Country string `json:"country"`
	Notes   string `json:"notes"`
}

type UpdateTargetRequest struct {
	Name          *string               `json:"name,omitempty"`
	Country       *string               `json:"country,omitempty"`
	Notes         *string               `json:"notes,omitempty"`
	CompleteState *models.CompleteState `json:"complete_state,omitempty"`
}
