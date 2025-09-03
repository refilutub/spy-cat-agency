package models

import (
	"errors"
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

func (t *Target) CanDelete() error {
	if t.CompleteState == Completed {
		return errors.New("cannot delete completed target")
	}
	return nil
}

func (t *Target) CanUpdateNotes() error {
	if t.CompleteState == Completed {
		return errors.New("cannot update notes of completed target")
	}
	return nil
}

func (t *Target) CanUpdate() error {
	return nil
}
