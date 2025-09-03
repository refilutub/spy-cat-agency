package application

import (
	"spy-cat-agency/internal/missions/dtos"

	"github.com/google/uuid"
)

type MissionsService interface {
	CreateMission(request dtos.CreateMissionRequest) (dtos.MissionCreateResponseDTO, error)
	GetMission(id uuid.UUID) (dtos.MissionResponseDTO, error)
	ListMissions() ([]dtos.MissionResponseDTO, error)
	UpdateMission(id uuid.UUID, request dtos.UpdateMissionRequest) (dtos.MissionResponseDTO, error)
	DeleteMission(id uuid.UUID) error
	AssignCatToMission(missionID, catID uuid.UUID) (dtos.MissionResponseDTO, error)

	AddTargetToMission(missionID uuid.UUID, request dtos.AddTargetRequest) (dtos.TargetResponseDTO, error)
	UpdateTarget(id uuid.UUID, request dtos.UpdateTargetRequest) (dtos.TargetResponseDTO, error)
	DeleteTarget(id uuid.UUID) error
}
