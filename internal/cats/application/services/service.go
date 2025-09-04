package services

import (
	"spy-cat-agency/internal/cats/application"
	"spy-cat-agency/internal/cats/domain"
)

type spyCatService struct {
	repo domain.SpyCatRepository
}

func NewSpyCatService(repo domain.SpyCatRepository) application.SpyCatService {
	return &spyCatService{repo: repo}
}
