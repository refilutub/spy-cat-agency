package mappers

import (
	"spy-cat-agency/internal/missions/domain/models"
	"spy-cat-agency/internal/missions/dtos"
	"time"

	"github.com/google/uuid"
)

func CreateDTOToMission() *models.Mission {
	return &models.Mission{
		ID:            uuid.New(),
		CompleteState: models.InProgress,
	}
}

func CreateTargetDTOToMission(missionReq dtos.CreateMissionRequest, missionId uuid.UUID) []models.Target {
	targets := make([]models.Target, 0)

	for _, target := range missionReq.Targets {
		targets = append(targets, models.Target{
			ID:            uuid.New(),
			MissionID:     missionId,
			Name:          target.Name,
			Country:       target.Country,
			Notes:         target.Notes,
			CompleteState: models.InProgress,
			UpdatedAt:     time.Now(),
			CreatedAt:     time.Now(),
		})
	}
	return targets
}

func MissionCreateToDTO(mission *models.Mission, newMissionTargets []models.Target) dtos.MissionCreateResponseDTO {
	return dtos.MissionCreateResponseDTO{
		Id:            mission.ID,
		CompleteState: mission.CompleteState,
		Targets:       MissionTargetsToDTO(newMissionTargets),
		UpdatedAt:     mission.UpdatedAt,
	}
}

func MissionTargetsToDTO(targets []models.Target) []dtos.TargetDTO {
	targetDtos := make([]dtos.TargetDTO, 0, len(targets))

	for _, target := range targets {
		targetDtos = append(targetDtos, dtos.TargetDTO{
			Id:            target.ID,
			Name:          target.Name,
			Country:       target.Country,
			Notes:         target.Notes,
			CompleteState: target.CompleteState,
			UpdatedAt:     target.UpdatedAt,
		})
	}

	return targetDtos
}
