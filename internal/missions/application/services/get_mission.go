package services

import (
	"spy-cat-agency/internal/missions/dtos"
	"spy-cat-agency/internal/missions/mappers"

	"github.com/google/uuid"
)

// GetMission retrieves a single mission by ID
func (s *missionsService) GetMission(id uuid.UUID) (dtos.MissionResponseDTO, error) {
	mission, err := s.repo.GetMission(id)
	if err != nil {
		return dtos.MissionResponseDTO{}, err
	}
	return mappers.MissionToResponseDTO(mission), nil
}
