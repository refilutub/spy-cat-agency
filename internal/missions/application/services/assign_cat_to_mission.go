package services

import (
	"spy-cat-agency/internal/missions/dtos"
	"spy-cat-agency/internal/missions/mappers"

	"github.com/google/uuid"
)

// AssignCatToMission assigns a cat to a mission
func (s *missionsService) AssignCatToMission(missionID, catID uuid.UUID) (dtos.MissionResponseDTO, error) {
	if err := s.repo.AssignCatToMission(missionID, catID); err != nil {
		return dtos.MissionResponseDTO{}, err
	}

	mission, err := s.repo.GetMission(missionID)
	if err != nil {
		return dtos.MissionResponseDTO{}, err
	}

	return mappers.MissionToResponseDTO(mission), nil
}
