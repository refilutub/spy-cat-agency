package domain

import (
	"spy-cat-agency/internal/missions/domain/models"
)

type MissionsRepository interface {
	CreateMission(newMission *models.Mission, newMissionTargets []models.Target) (*models.Mission, error)
}
