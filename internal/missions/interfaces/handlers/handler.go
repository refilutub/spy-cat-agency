package handlers

import (
	"spy-cat-agency/internal/missions/application"
)

type MissionsHandler struct {
	Service application.MissionsService
}

func NewMissionsHandler(service application.MissionsService) *MissionsHandler {
	return &MissionsHandler{Service: service}
}
