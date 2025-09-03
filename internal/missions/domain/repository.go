package domain

import (
	"spy-cat-agency/internal/missions/domain/models"

	"github.com/google/uuid"
)

type MissionsRepository interface {
	CreateMission(newMission *models.Mission, newMissionTargets []models.Target) (*models.Mission, error)

	GetMission(id uuid.UUID) (*models.Mission, error)
	ListMissions() ([]models.Mission, error)
	UpdateMission(mission *models.Mission) error
	DeleteMission(id uuid.UUID) error
	AssignCatToMission(missionID, catID uuid.UUID) error

	AddTargetToMission(missionID uuid.UUID, target *models.Target) error
	UpdateTarget(target *models.Target) error
	DeleteTarget(id uuid.UUID) error
	GetTarget(id uuid.UUID) (*models.Target, error)
}
