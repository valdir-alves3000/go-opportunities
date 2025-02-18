package repositories

import (
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
)

type OpeningRepository interface {
	Create(opening schemas.Opening) error
	FindByID(id uint) (*schemas.Opening, error)
	Update(opening schemas.Opening) error
	Delete(id uint) error
	FindAll(limit, offset int) ([]schemas.Opening, error)
}
