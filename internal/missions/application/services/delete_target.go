package services

import (
	"github.com/google/uuid"
)

// DeleteTarget deletes a target
func (s *missionsService) DeleteTarget(id uuid.UUID) error {
	return s.repo.DeleteTarget(id)
}
