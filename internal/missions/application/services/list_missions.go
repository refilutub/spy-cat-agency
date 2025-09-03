package services

import (
	"spy-cat-agency/internal/missions/dtos"
	"spy-cat-agency/internal/missions/mappers"
)

// ListMissions retrieves all missions
func (s *missionsService) ListMissions() ([]dtos.MissionResponseDTO, error) {
	missions, err := s.repo.ListMissions()
	if err != nil {
		return nil, err
	}

	var missionDTOs []dtos.MissionResponseDTO
	for _, mission := range missions {
		missionDTOs = append(missionDTOs, mappers.MissionToResponseDTO(&mission))
	}
	return missionDTOs, nil
}
