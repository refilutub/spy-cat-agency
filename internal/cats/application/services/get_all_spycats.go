package services

import "spy-cat-agency/internal/cats/dtos"

func (s *spyCatService) GetAllSpyCats() ([]dtos.SpyCatListResponseDTO, error) {
	spyCats, err := s.repo.GetAllSpyCats()
	if err != nil {
		return nil, err
	}

	var spyCatDTOs []dtos.SpyCatListResponseDTO
	for _, cat := range spyCats {
		spyCatDTO := dtos.SpyCatListResponseDTO{
			Id:              cat.ID,
			Name:            cat.Name,
			ExperienceYears: cat.ExperienceYears,
			Breed:           cat.Breed,
			Salary:          cat.Salary,
		}
		spyCatDTOs = append(spyCatDTOs, spyCatDTO)
	}

	return spyCatDTOs, nil
}
