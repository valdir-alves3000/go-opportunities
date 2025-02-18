package opening_usecase_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"gorm.io/gorm"
)

func TestUpdateOpeningUsecase(t *testing.T) {
	openingUsecase, openingRepo := setupUsecaseTest()
	remote := true
	var ID uint = 3000

	upOpeningMock := schemas.UpdateOpeningRequest{
		Role:     "JavaScript Developer",
		Company:  "Tech Corp",
		Location: "Spain",
		Link:     "https://spain.com/job",
		Remote:   &remote,
		Salary:   50000,
	}

	t.Run("ShouldReturnAnErrorIfTheDBFails", func(t *testing.T) {
		openingExist := schemas.Opening{
			Model:    gorm.Model{ID: ID},
			Role:     "Java Developer",
			Company:  "Tech Corp New York",
			Location: "USA",
			Link:     "https://global.com/job/usa",
			Remote:   false,
			Salary:   80000,
		}
		expectedOpening := schemas.Opening{
			Model:    gorm.Model{ID: ID},
			Role:     upOpeningMock.Role,
			Company:  upOpeningMock.Company,
			Location: upOpeningMock.Location,
			Remote:   *upOpeningMock.Remote,
			Link:     upOpeningMock.Link,
			Salary:   upOpeningMock.Salary,
		}

		openingRepo.On("Update", expectedOpening).Return(gorm.ErrRegistered).Once()
		openingRepo.On("FindByID", ID).Return(&openingExist, nil).Once()
		mockErr := internal_error.NewInternalServerError("error updating opening")

		err := openingUsecase.Update(ID, upOpeningMock)

		assert.Error(t, err)
		assert.EqualError(t, mockErr, err.Error())
		openingRepo.AssertCalled(t, "Update", expectedOpening)
		openingRepo.AssertExpectations(t)
	})

	t.Run("ShouldUpdateAllJobOpeningParameters", func(t *testing.T) {
		openingExist := schemas.Opening{
			Model:    gorm.Model{ID: ID},
			Role:     "Java Developer",
			Company:  "Tech Corp New York",
			Location: "USA",
			Link:     "https://global.com/job/usa",
			Remote:   false,
			Salary:   80000,
		}
		expectedOpening := schemas.Opening{
			Model:    gorm.Model{ID: ID},
			Role:     upOpeningMock.Role,
			Company:  upOpeningMock.Company,
			Location: upOpeningMock.Location,
			Remote:   *upOpeningMock.Remote,
			Link:     upOpeningMock.Link,
			Salary:   upOpeningMock.Salary,
		}

		openingRepo.On("FindByID", ID).Return(&openingExist, nil).Once()
		openingRepo.On("Update", expectedOpening).Return(nil).Once()

		err := openingUsecase.Update(ID, upOpeningMock)

		assert.Nil(t, err)
		openingRepo.AssertCalled(t, "Update", expectedOpening)
		openingRepo.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfTheOpeningIsNotFound", func(t *testing.T) {
		mockErr := errors.New("opening not found")
		openingRepo.On("FindByID", ID).Return(&schemas.Opening{}, gorm.ErrRecordNotFound).Once()

		err := openingUsecase.Update(ID, upOpeningMock)

		openingRepo.AssertCalled(t, "FindByID", ID)
		openingRepo.AssertNotCalled(t, "Update")
		assert.Error(t, err, "I expected error when updating if opening not found")
		assert.EqualError(t, err, mockErr.Error(), "The error message must be specific")
	})

	t.Run("ShouldReturnAnErrorIfAnEmptyBodyIsRequiredForTheUpdate", func(t *testing.T) {
		upOpening := schemas.UpdateOpeningRequest{}

		err := openingUsecase.Update(ID, upOpening)

		assert.Error(t, err, "I expected error when updating if body is empty")
		assert.EqualError(t, err, "at least one valid field must be provided", "The error message must be specific")
	})

	t.Run("ShouldReturnAnErrorWhenTryingToUpdateTheSalaryToAValueLessThan3k", func(t *testing.T) {
		upOpeningMock := schemas.UpdateOpeningRequest{
			Salary: 2999,
		}

		err := openingUsecase.Update(ID, upOpeningMock)

		assert.Error(t, err, "I expected an error when updating a salary less than 3k")
		assert.EqualError(t, err, "salary must be at least 3k", "The error message must be specific")
	})
}
