package repository

import (
	"spy-cat-agency/internal/cats/domain/models"

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
