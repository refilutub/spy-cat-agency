package domain

import (
	"github.com/google/uuid"
	"spy-cat-agency/internal/cats/domain/models"
)

type SpyCatRepository interface {
	CreateSpyCat(newSpyCat *models.Cat) (*models.Cat, error)

	DeleteSpyCat(spyCatId uuid.UUID) error
	DeleteBatchSpyCats(spyCatIds []uuid.UUID) error
}
