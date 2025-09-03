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

func (s *SpyCatRepository) DeleteBatchSpyCats(spyCatIds []uuid.UUID) error {
	var catsWithMissions []uuid.UUID

	if err := s.db.
		Model(&models2.Mission{}).
		Where("cat_id IN ?", spyCatIds).
		Pluck("cat_id", &catsWithMissions).Error; err != nil {
		return err
	}

	referenced := make(map[uuid.UUID]bool, len(catsWithMissions))
	for _, id := range catsWithMissions {
		referenced[id] = true
	}

	var softDeleteIds, hardDeleteIds []uuid.UUID
	for _, id := range spyCatIds {
		if referenced[id] {
			softDeleteIds = append(softDeleteIds, id)
		} else {
			hardDeleteIds = append(hardDeleteIds, id)
		}
	}

	if len(softDeleteIds) > 0 {
		if err := s.db.Delete(&models.Cat{}, "id IN ?", softDeleteIds).Error; err != nil {
			return err
		}
	}

	if len(hardDeleteIds) > 0 {
		if err := s.db.Unscoped().Delete(&models.Cat{}, "id IN ?", hardDeleteIds).Error; err != nil {
			return err
		}
	}

	return nil

}

func (s *SpyCatRepository) UpdateSpyCatSalary(spyCatId uuid.UUID, newSalary float64) (*models.Cat, error) {
	var cat models.Cat
	if err := s.db.First(&cat, "id = ?", spyCatId).Error; err != nil {
		return nil, err
	}

	cat.Salary = newSalary
	if err := s.db.Save(&cat).Error; err != nil {
		return nil, err
	}
	return &cat, nil
}

func (s *SpyCatRepository) GetAllSpyCats() ([]models.Cat, error) {
	var cats []models.Cat
	if err := s.db.Find(&cats).Error; err != nil {
		return nil, err
	}
	return cats, nil
}
