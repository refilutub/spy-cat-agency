package services

import (
	"spy-cat-agency/internal/cats/dtos"
)

func (s *spyCatService) DeleteBatchSpyCats(spyCatIds dtos.DeleteIds) error {
	return s.repo.DeleteBatchSpyCats(spyCatIds.Ids)
}
