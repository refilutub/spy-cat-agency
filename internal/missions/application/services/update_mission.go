package services

import (
	"spy-cat-agency/internal/missions/dtos"
	"spy-cat-agency/internal/missions/mappers"

	"github.com/google/uuid"
)

// UpdateMission updates a mission
func (s *missionsService) UpdateMission(id uuid.UUID, request dtos.UpdateMissionRequest) (dtos.MissionResponseDTO, error) {
	mission, err := s.repo.GetMission(id)
	if err != nil {
		return dtos.MissionResponseDTO{}, err
	}

	if request.CompleteState != nil {
		mission.CompleteState = *request.CompleteState
	}

	if err := s.repo.UpdateMission(mission); err != nil {
		return dtos.MissionResponseDTO{}, err
	}

	return mappers.MissionToResponseDTO(mission), nil
}
