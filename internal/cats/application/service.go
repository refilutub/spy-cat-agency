package application

import (
	"spy-cat-agency/internal/cats/dtos"
)

type SpyCatService interface {
	CreateSpyCat(spyCatReq dtos.SpyCatRequest) (dtos.SpyCatSingleResponseDTO, error)
}
