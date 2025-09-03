package services

import (
	"spy-cat-agency/internal/missions/dtos"
	"spy-cat-agency/internal/missions/mappers"

	"github.com/google/uuid"
)

// UpdateTarget updates a target
func (s *missionsService) UpdateTarget(id uuid.UUID, request dtos.UpdateTargetRequest) (dtos.TargetResponseDTO, error) {
	target, err := s.repo.GetTarget(id)
	if err != nil {
		return dtos.TargetResponseDTO{}, err
	}

	if request.Name != nil {
		target.Name = *request.Name
	}
	if request.Country != nil {
		target.Country = *request.Country
	}
	if request.Notes != nil {
		target.Notes = *request.Notes
	}
	if request.CompleteState != nil {
		target.CompleteState = *request.CompleteState
	}

	if err := s.repo.UpdateTarget(target); err != nil {
		return dtos.TargetResponseDTO{}, err
	}

	return mappers.TargetToResponseDTO(target), nil
}
