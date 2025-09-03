package repository

import (
	"github.com/google/uuid"
	"spy-cat-agency/internal/cats/domain/models"
	models2 "spy-cat-agency/internal/missions/domain/models"

	"gorm.io/gorm"
)

type SpyCatRepository struct {
	db *gorm.DB
}

func NewSpyCatRepository(db *gorm.DB) *SpyCatRepository {
	return &SpyCatRepository{db: db}
}

func (s *SpyCatRepository) CreateSpyCat(newSpyCat *models.Cat) (*models.Cat, error) {
	if err := s.db.Create(newSpyCat).Error; err != nil {
		return nil, err
	}
	return newSpyCat, nil
}

func (s *SpyCatRepository) DeleteSpyCat(spyCatId uuid.UUID) error {
	if err := s.db.Model(&models2.Mission{}).
		Where("id = ?", spyCatId).
		First(&models2.Mission{}).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return s.db.Unscoped().Delete(&models.Cat{}, "id = ?", spyCatId).Error
		}
		return err
	}
	return s.db.Delete(&models.Cat{}, "id = ?", spyCatId).Error
}
