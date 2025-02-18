package opening_usecase

import (
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
)

func (uc *OpeningUseCase) DeleteByID(id uint) *internal_error.InternalError {
	_, err := uc.repo.FindByID(id)
	if err != nil {
		return internal_error.NewNotFoundError("opening not found")
	}

	errRepo := uc.repo.Delete(id)
	if errRepo != nil {
		return internal_error.NewInternalServerError("error deleting opening")
	}

	return nil
}
