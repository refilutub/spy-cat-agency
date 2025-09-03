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
	Id            uuid.UUID            `json:"id"`
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
