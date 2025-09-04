package models

import (
	"errors"
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

func (m *Mission) CanDelete() error {
	if m.CatID != nil {
		return errors.New("cannot delete mission: already assigned to a cat")
	}
	return nil
}

func (m *Mission) CanAssignCat() error {
	if m.CompleteState == Completed {
		return errors.New("cannot assign cat to completed mission")
	}
	if m.CatID != nil {
		return errors.New("mission is already assigned to a cat")
	}
	return nil
}

func (m *Mission) CanAddTarget() error {
	if m.CompleteState == Completed {
		return errors.New("cannot add target to completed mission")
	}
	return nil
}

func (m *Mission) CanUpdate() error {
	return nil
}
