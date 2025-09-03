package services

import (
	"github.com/google/uuid"
	"spy-cat-agency/internal/cats/dtos"
	"spy-cat-agency/internal/cats/mappers"
)

func (s *spyCatService) GetSpyCat(spyCatId uuid.UUID) (dtos.SpyCatSingleResponseDTO, error) {
	spycat, err := s.repo.GetSpyCat(spyCatId)
	if err != nil {
		return dtos.SpyCatSingleResponseDTO{}, err
	}
	return mappers.SpyCatSingleToDTO(spycat), nil
}
