package services

import (
	"fmt"
	"github.com/google/uuid"
	"math"
	"spy-cat-agency/internal/cats/dtos"
	"spy-cat-agency/internal/cats/mappers"
)

func (s *spyCatService) UpdateSpyCatSalary(spyCatId uuid.UUID, salaryReq dtos.SalaryUpdateRequest) (dtos.SpyCatSingleResponseDTO, error) {
	switch {
	case salaryReq.Salary < 0:
		return dtos.SpyCatSingleResponseDTO{}, fmt.Errorf("salary must be positive")
	case salaryReq.Salary >= math.MaxFloat64:
		return dtos.SpyCatSingleResponseDTO{}, fmt.Errorf("salary is too large")
	}
	updatedSpyCat, err := s.repo.UpdateSpyCatSalary(spyCatId, salaryReq.Salary)
	if err != nil {
		return dtos.SpyCatSingleResponseDTO{}, err
	}
	return mappers.SpyCatSingleToDTO(updatedSpyCat), nil
}
