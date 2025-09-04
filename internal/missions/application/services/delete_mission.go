package services

import (
	"github.com/google/uuid"
)

// DeleteMission deletes a mission
func (s *missionsService) DeleteMission(id uuid.UUID) error {
	return s.repo.DeleteMission(id)
}
