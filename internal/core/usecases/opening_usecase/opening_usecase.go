package opening_usecase

import (
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"github.com/valdir-alves3000/go-opportunities/internal/repositories"
)

type OpeningUsecase interface {
	Create(co schemas.CreateOpeningRequest) *internal_error.InternalError
	GetByID(id uint) (*schemas.Opening, *internal_error.InternalError)
	Update(id uint, upo schemas.UpdateOpeningRequest) *internal_error.InternalError
	DeleteByID(id uint) *internal_error.InternalError
	ListOpenings(page int) ([]schemas.Opening, *internal_error.InternalError)
}

type OpeningUseCase struct {
	repo repositories.OpeningRepository
}

func NewOpeningUseCase(repo repositories.OpeningRepository) *OpeningUseCase {
	return &OpeningUseCase{repo: repo}
}
