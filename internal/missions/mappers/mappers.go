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
			Name:          target.Name,
			Country:       target.Country,
			Notes:         target.Notes,
			CompleteState: target.CompleteState,
			UpdatedAt:     target.UpdatedAt,
		})
	}

	return targetDtos
}

func MissionToResponseDTO(mission *models.Mission) dtos.MissionResponseDTO {
	var catDTO *dtos.CatDTO
	if mission.Cat != nil {
		catDTO = &dtos.CatDTO{
			ID:              mission.Cat.ID,
			Name:            mission.Cat.Name,
			ExperienceYears: mission.Cat.ExperienceYears,
			Breed:           mission.Cat.Breed,
			Salary:          mission.Cat.Salary,
		}
	}

	var targetDTOs []dtos.TargetResponseDTO
	for _, target := range mission.Targets {
		targetDTOs = append(targetDTOs, TargetToResponseDTO(&target))
	}

	return dtos.MissionResponseDTO{
		ID:            mission.ID,
		CatID:         mission.CatID,
		Cat:           catDTO,
		Targets:       targetDTOs,
		CompleteState: mission.CompleteState,
		UpdatedAt:     mission.UpdatedAt,
		CreatedAt:     mission.CreatedAt,
	}
}

func TargetToResponseDTO(target *models.Target) dtos.TargetResponseDTO {
	return dtos.TargetResponseDTO{
		ID:            target.ID,
		MissionID:     target.MissionID,
		Name:          target.Name,
		Country:       target.Country,
		Notes:         target.Notes,
		CompleteState: target.CompleteState,
		UpdatedAt:     target.UpdatedAt,
		CreatedAt:     target.CreatedAt,
	}
}
