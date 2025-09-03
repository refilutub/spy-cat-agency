package domain

import "spy-cat-agency/internal/cats/domain/models"

type SpyCatRepository interface {
	CreateSpyCat(newSpyCat *models.Cat) (*models.Cat, error)
}
