package repository

import (
	"spy-cat-agency/internal/missions/domain/models"

	"gorm.io/gorm"
)

type MissionsRepository struct {
	db *gorm.DB
}

func NewMissionsRepository(db *gorm.DB) *MissionsRepository {
	return &MissionsRepository{db: db}
}

func (m *MissionsRepository) CreateMission(newMission *models.Mission,
	newMissionTargets []models.Target) (*models.Mission, error) {
	tx := m.db.Begin()

	if err := tx.Create(newMission).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	for i := range newMissionTargets {
		if err := tx.Create(&newMissionTargets[i]).Error; err != nil {
			tx.Rollback()
			return nil, err
		}
	}

	if err := tx.Commit().Error; err != nil {
		return nil, err
	}

	if err := m.db.First(newMission, newMission.ID).Error; err != nil {
		return nil, err
	}

	return newMission, nil
}
