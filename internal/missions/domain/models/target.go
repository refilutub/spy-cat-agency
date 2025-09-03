package models

import "github.com/google/uuid"

type Target struct {
	ID            uuid.UUID     `gorm:"type:uuid;primaryKey" json:"id"`
	MissionID     uuid.UUID     `gorm:"type:uuid" json:"mission_id"`
	Name          string        `json:"name"`
	Country       string        `json:"country"`
	Notes         string        `json:"notes"`
	CompleteState CompleteState `json:"complete_state"`
	UpdatedAt     string        `json:"updated_at"`
	CreatedAt     string        `json:"created_at"`
}
