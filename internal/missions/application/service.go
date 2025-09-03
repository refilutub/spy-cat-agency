package application

import "spy-cat-agency/internal/missions/dtos"

type MissionsService interface {
	CreateMission(request dtos.CreateMissionRequest) (dtos.MissionCreateResponseDTO, error)
}
