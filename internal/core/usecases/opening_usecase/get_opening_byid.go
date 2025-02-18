package opening_usecase

import (
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
)

func (uc *OpeningUseCase) GetByID(id uint) (*schemas.Opening, *internal_error.InternalError) {
	opening, err := uc.repo.FindByID(id)
	if err != nil {
		return nil, internal_error.NewNotFoundError("opening not found")
	}

	return opening, nil
}
