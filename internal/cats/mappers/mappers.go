package mappers

import (
	"github.com/google/uuid"
	"spy-cat-agency/internal/cats/domain/models"
	"spy-cat-agency/internal/cats/dtos"
)

func CreateDTOToSpyCat(spycatReq dtos.SpyCatRequest) *models.Cat {
	return &models.Cat{
		ID:              uuid.New(),
		Name:            spycatReq.Name,
		ExperienceYears: spycatReq.ExperienceYears,
		Breed:           spycatReq.Breed,
		Salary:          spycatReq.Salary,
	}
}

func SpyCatSingleToDTO(spycat *models.Cat) dtos.SpyCatSingleResponseDTO {
	return dtos.SpyCatSingleResponseDTO{
		Id:              spycat.ID,
		Name:            spycat.Name,
		ExperienceYears: spycat.ExperienceYears,
		Breed:           spycat.Breed,
		Salary:          spycat.Salary,
		UpdatedAt:       spycat.UpdatedAt,
	}
}
