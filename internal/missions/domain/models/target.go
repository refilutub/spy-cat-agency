package models

import (
	"time"

	"github.com/google/uuid"
)

type Target struct {
	ID            uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	MissionID     uuid.UUID     `gorm:"type:uuid" json:"mission_id"`
	Name          string        `json:"name"`
	Country       string        `json:"country"`
	Notes         string        `json:"notes"`
	CompleteState CompleteState `gorm:"default:in_progress" json:"complete_state"`
	UpdatedAt     time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
	CreatedAt     time.Time     `gorm:"autoCreateTime" json:"created_at"`
}
