package application

import (
	"github.com/google/uuid"
	"spy-cat-agency/internal/cats/dtos"
)

type SpyCatService interface {
	CreateSpyCat(spyCatReq dtos.SpyCatRequest) (dtos.SpyCatSingleResponseDTO, error)

	DeleteSpyCat(spyCatId uuid.UUID) error
	DeleteBatchSpyCats(spyCatIds dtos.DeleteIds) error

	UpdateSpyCatSalary(spyCatId uuid.UUID, salaryReq dtos.SalaryUpdateRequest) (dtos.SpyCatSingleResponseDTO, error)
}
