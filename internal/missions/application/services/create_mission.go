package services

import (
	"fmt"
	"spy-cat-agency/internal/missions/dtos"
	"spy-cat-agency/internal/missions/mappers"
)

func (s *missionsService) CreateMission(request dtos.CreateMissionRequest) (dtos.MissionCreateResponseDTO, error) {
	if len(request.Targets) == 0 {
		return dtos.MissionCreateResponseDTO{}, fmt.Errorf("no targets provided")
	}
	newMissionModel := mappers.CreateDTOToMission()
	newMissionTargets := mappers.CreateTargetDTOToMission(request, newMissionModel.ID)
	newMission, err := s.repo.CreateMission(newMissionModel, newMissionTargets)
	if err != nil {
		return dtos.MissionCreateResponseDTO{}, err
	}
	return mappers.MissionCreateToDTO(newMission, newMissionTargets), nil
}
