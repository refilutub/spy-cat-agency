package repository

import (
	"errors"
	"spy-cat-agency/internal/missions/domain/models"

	"github.com/google/uuid"
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

func (m *MissionsRepository) GetMission(id uuid.UUID) (*models.Mission, error) {
	var mission models.Mission
	err := m.db.Preload("Targets").Preload("Cat").First(&mission, id).Error
	if err != nil {
		return nil, err
	}
	return &mission, nil
}

func (m *MissionsRepository) ListMissions() ([]models.Mission, error) {
	var missions []models.Mission
	err := m.db.Preload("Targets").Preload("Cat").Find(&missions).Error
	if err != nil {
		return nil, err
	}
	return missions, nil
}

func (m *MissionsRepository) UpdateMission(mission *models.Mission) error {
	return m.db.Save(mission).Error
}

func (m *MissionsRepository) DeleteMission(id uuid.UUID) error {
	var mission models.Mission
	if err := m.db.First(&mission, id).Error; err != nil {
		return err
	}

	if err := mission.CanDelete(); err != nil {
		return err
	}

	if err := m.db.Where("mission_id = ?", id).Delete(&models.Target{}).Error; err != nil {
		return err
	}

	return m.db.Delete(&mission).Error
}

func (m *MissionsRepository) AssignCatToMission(missionID, catID uuid.UUID) error {
	var mission models.Mission
	if err := m.db.First(&mission, missionID).Error; err != nil {
		return err
	}

	if err := mission.CanAssignCat(); err != nil {
		return err
	}

	mission.CatID = &catID
	return m.db.Save(&mission).Error
}

func (m *MissionsRepository) AddTargetToMission(missionID uuid.UUID, target *models.Target) error {
	var mission models.Mission
	if err := m.db.First(&mission, missionID).Error; err != nil {
		return err
	}

	if err := mission.CanAddTarget(); err != nil {
		return err
	}

	target.MissionID = missionID
	return m.db.Create(target).Error
}

func (m *MissionsRepository) UpdateTarget(target *models.Target) error {
	var existingTarget models.Target
	if err := m.db.First(&existingTarget, target.ID).Error; err != nil {
		return err
	}

	if target.Notes != existingTarget.Notes {
		if err := existingTarget.CanUpdateNotes(); err != nil {
			return err
		}

		var mission models.Mission
		if err := m.db.First(&mission, existingTarget.MissionID).Error; err != nil {
			return err
		}
		if mission.CompleteState == models.Completed {
			return errors.New("cannot update notes: mission is completed")
		}
	}

	return m.db.Save(target).Error
}

func (m *MissionsRepository) DeleteTarget(id uuid.UUID) error {
	var target models.Target
	if err := m.db.First(&target, id).Error; err != nil {
		return err
	}

	if err := target.CanDelete(); err != nil {
		return err
	}

	return m.db.Delete(&target).Error
}

func (m *MissionsRepository) GetTarget(id uuid.UUID) (*models.Target, error) {
	var target models.Target
	err := m.db.First(&target, id).Error
	if err != nil {
		return nil, err
	}
	return &target, nil
}
