package models

import (
	"spy-cat-agency/internal/cats/domain/models"
	"time"

	"github.com/google/uuid"
)

type CompleteState string

const (
	InProgress CompleteState = "in_progress"
	Completed  CompleteState = "completed"
)

type Mission struct {
	ID            uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	CatID         *uuid.UUID    `gorm:"type:uuid" json:"cat_id"`
	Cat           *models.Cat   `gorm:"foreignKey:CatID" json:"cat"`
	Targets       []Target      `gorm:"foreignKey:MissionID" json:"targets"`
	CompleteState CompleteState `json:"complete_state"`
	UpdatedAt     time.Time     `json:"updated_at"`
	CreatedAt     time.Time     `json:"created_at"`
}
