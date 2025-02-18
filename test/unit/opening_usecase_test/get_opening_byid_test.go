package opening_usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"gorm.io/gorm"
)

func TestGetOpeningByIDUsecase(t *testing.T) {
	openingUsecase, openingRepo := setupUsecaseTest()

	t.Run("ShouldReturnOpeningWhenFoundById", func(t *testing.T) {
		var ID uint = 3000
		opening := schemas.Opening{
			Model:    gorm.Model{ID: ID},
			Role:     "Junior Developer",
			Company:  "Tech Corp",
			Location: "Spain",
			Link:     "https://spain.com/job",
			Remote:   *new(bool),
			Salary:   25000,
		}
		openingRepo.On("FindByID", ID).Return(&opening, nil).Once()

		result, err := openingUsecase.GetByID(ID)

		assert.Nil(t, err)
		openingRepo.AssertCalled(t, "FindByID", ID)
		assert.Equal(t, &opening, result)
	})

	t.Run("ShouldReturnErrorWhenOpeningNotFound", func(t *testing.T) {
		var ID uint = 9999
		mockErr := internal_error.NewNotFoundError("opening not found")

		openingRepo.On("FindByID", ID).Return(&schemas.Opening{}, gorm.ErrRecordNotFound).Once()

		result, err := openingUsecase.GetByID(ID)

		assert.Error(t, err)
		assert.Equal(t, mockErr, err)
		assert.Nil(t, result)
	})
}
