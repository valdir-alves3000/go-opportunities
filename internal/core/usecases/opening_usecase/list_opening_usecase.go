package opening_usecase

import (
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
)

func (uc *OpeningUseCase) ListOpenings(page int) ([]schemas.Opening, *internal_error.InternalError) {
	if page <= 0 {
		page = 1
	}
	limit := 10
	offset := (page - 1) * limit

	openings, err := uc.repo.FindAll(limit, offset)
	if err != nil || len(openings) == 0 {
		return nil, internal_error.NewNotFoundError("opening record not found")
	}

	return openings, nil
}
