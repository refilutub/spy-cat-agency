package services

import (
	"spy-cat-agency/internal/missions/application"
	"spy-cat-agency/internal/missions/domain"
)

type missionsService struct {
	repo domain.MissionsRepository
}

func NewMissionsService(repo domain.MissionsRepository) application.MissionsService {
	return &missionsService{repo: repo}
}
