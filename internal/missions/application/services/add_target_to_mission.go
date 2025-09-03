package services

import (
	"spy-cat-agency/internal/missions/domain/models"
	"spy-cat-agency/internal/missions/dtos"
	"spy-cat-agency/internal/missions/mappers"

	"github.com/google/uuid"
)

func (s *missionsService) AddTargetToMission(missionID uuid.UUID, request dtos.AddTargetRequest) (dtos.TargetResponseDTO, error) {
	target := &models.Target{
		ID:            uuid.New(),
		Name:          request.Name,
		Country:       request.Country,
		Notes:         request.Notes,
		CompleteState: models.InProgress,
	}

	if err := s.repo.AddTargetToMission(missionID, target); err != nil {
		return dtos.TargetResponseDTO{}, err
	}

	return mappers.TargetToResponseDTO(target), nil
}
