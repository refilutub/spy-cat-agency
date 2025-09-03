package services

import "github.com/google/uuid"

func (s *spyCatService) DeleteSpyCat(spycatId uuid.UUID) error {
	return s.repo.DeleteSpyCat(spycatId)
}
