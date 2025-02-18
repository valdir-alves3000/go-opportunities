package opening_usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"gorm.io/gorm"
)

func TestDeleteOpeningByIDUsecase(t *testing.T) {
	openingUsecase, openingRepo := setupUsecaseTest()

	t.Run("ShouldReturnErrorWhenTryingToDeleteNotFoundOpening", func(t *testing.T) {
		var ID uint = 9999
		mockErr := internal_error.NewNotFoundError("opening not found")

		openingRepo.On("FindByID", ID).Return(&schemas.Opening{}, gorm.ErrRecordNotFound).Once()

		err := openingUsecase.DeleteByID(ID)

		assert.Error(t, err)
		assert.EqualError(t, mockErr, err.Error())
		openingRepo.AssertCalled(t, "FindByID", ID)
		openingRepo.AssertNotCalled(t, "Delete")

	})

	t.Run("ShouldDeleteanOpeningwithValidID", func(t *testing.T) {
		var ID uint = 3000
		openingMock := &schemas.Opening{
			Model:    gorm.Model{ID: ID},
			Role:     "GO developer",
			Company:  "+3000 DEV",
			Location: "Mauá - SP",
			Remote:   true,
			Link:     "https//+3000dev.opportunies.com",
			Salary:   30000,
		}
		openingRepo.On("FindByID", ID).Return(openingMock, nil).Once()
		openingRepo.On("Delete", ID).Return(nil).Once()

		err := openingUsecase.DeleteByID(ID)

		assert.Nil(t, err)
		openingRepo.AssertCalled(t, "FindByID", ID)
		openingRepo.AssertCalled(t, "Delete", ID)
		openingRepo.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfTheDBFails", func(t *testing.T) {
		var ID uint = 3000
		openingMock := &schemas.Opening{
			Model:    gorm.Model{ID: ID},
			Role:     "GO developer",
			Company:  "+3000 DEV",
			Location: "Mauá - SP",
			Remote:   true,
			Link:     "https//+3000dev.opportunies.com",
			Salary:   30000,
		}
		openingRepo.On("FindByID", ID).Return(openingMock, nil).Once()
		openingRepo.On("Delete", ID).Return(gorm.ErrMissingWhereClause).Once()

		expectedErr := internal_error.NewInternalServerError("error deleting opening")
		err := openingUsecase.DeleteByID(ID)

		assert.Error(t, err)
		assert.Equal(t, expectedErr, err)
		openingRepo.AssertCalled(t, "FindByID", ID)
		openingRepo.AssertCalled(t, "Delete", ID)
		openingRepo.AssertExpectations(t)
	})
}
