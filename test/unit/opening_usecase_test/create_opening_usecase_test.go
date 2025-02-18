package opening_usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"gorm.io/gorm"
)

func TestCreateUsecase(t *testing.T) {
	openingUsecase, openingRepo := setupUsecaseTest()

	t.Run("ShouldSuccessfullyCreateAJobOpening", func(t *testing.T) {
		request := schemas.CreateOpeningRequest{
			Role:     "Software Engineer",
			Company:  "TechCorp",
			Location: "Remote",
			Remote:   boolPtr(true),
			Link:     "https://example.com/job",
			Salary:   5000,
		}

		opening := schemas.Opening{
			Role:     request.Role,
			Company:  request.Company,
			Location: request.Location,
			Remote:   *request.Remote,
			Link:     request.Link,
			Salary:   request.Salary,
		}

		openingRepo.On("Create", opening).Return(nil).Once()
		err := openingUsecase.Create(request)

		assert.Nil(t, err)

		openingRepo.AssertCalled(t, "Create", opening)
		openingRepo.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfTheDBFails", func(t *testing.T) {
		request := schemas.CreateOpeningRequest{
			Role:     "Software Engineer",
			Company:  "TechCorp",
			Location: "Remote",
			Remote:   boolPtr(true),
			Link:     "https://example.com/job",
			Salary:   5000,
		}

		opening := schemas.Opening{
			Role:     request.Role,
			Company:  request.Company,
			Location: request.Location,
			Remote:   *request.Remote,
			Link:     request.Link,
			Salary:   request.Salary,
		}

		mockErr := internal_error.NewInternalServerError("error creating opening")
		openingRepo.On("Create", opening).Return(gorm.ErrRegistered).Once()

		err := openingUsecase.Create(request)

		assert.Error(t, err)
		assert.EqualError(t, mockErr, err.Error())
		openingRepo.AssertCalled(t, "Create", opening)
		openingRepo.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfTheRoleIsEmpty", func(t *testing.T) {
		openingMockWithEmptyRole := schemas.CreateOpeningRequest{
			Role:     "",
			Company:  "Tech Corp",
			Location: "Spain",
			Link:     "https://spain.com/job",
			Remote:   new(bool),
			Salary:   50000,
		}

		mockErr := internal_error.NewInternalServerError("param: role (type: string) is required")

		err := openingUsecase.Create(openingMockWithEmptyRole)

		assert.Error(t, err, "expected an error when creating the opening without Role")
		assert.EqualError(t, mockErr, err.Error(), "The error message must be specific")

		openingRepo.AssertNotCalled(t, "Create")
	})

	t.Run("ShouldReturnAnErrorIfTheCompanyIsEmpty", func(t *testing.T) {
		openingMockWithEmptyCompany := schemas.CreateOpeningRequest{
			Role:     "Developer JavaScript",
			Company:  "",
			Location: "Spain",
			Link:     "https://spain.com/job",
			Remote:   new(bool),
			Salary:   50000,
		}

		mockErr := internal_error.NewInternalServerError("param: company (type: string) is required")

		err := openingUsecase.Create(openingMockWithEmptyCompany)

		assert.Error(t, err, "expected an error when creating the opening without company")
		assert.EqualError(t, mockErr, err.Error(), "The error message must be specific")

		openingRepo.AssertNotCalled(t, "Create")
	})

	t.Run("ShouldReturnAnErrorIfTheLocationIsEmpty", func(t *testing.T) {
		openingMockWithEmptyLocation := schemas.CreateOpeningRequest{
			Role:     "Developer JavaScript",
			Company:  "Tech Corp",
			Location: "",
			Link:     "https://spain.com/job",
			Remote:   new(bool),
			Salary:   50000,
		}
		mockErr := internal_error.NewInternalServerError("param: location (type: string) is required")

		err := openingUsecase.Create(openingMockWithEmptyLocation)

		assert.Error(t, err, "expected an error when creating the opening without company")
		assert.EqualError(t, mockErr, err.Error(), "The error message must be specific")

		openingRepo.AssertNotCalled(t, "Create")
	})

	t.Run("ShouldReturnAnErrorIfTheLinkIsEmpty", func(t *testing.T) {
		openingMockWithEmptyLink := schemas.CreateOpeningRequest{
			Role:     "Developer JavaScript",
			Company:  "Tech Corp",
			Location: "USA",
			Link:     "",
			Remote:   new(bool),
			Salary:   50000,
		}
		mockErr := internal_error.NewInternalServerError("param: link (type: string) is required")

		err := openingUsecase.Create(openingMockWithEmptyLink)

		assert.Error(t, err, "expected an error when creating the opening without company")
		assert.EqualError(t, mockErr, err.Error(), "The error message must be specific")

		openingRepo.AssertNotCalled(t, "Create")
	})

	t.Run("ShouldReturnAnErrorIfTheRemoteFieldIsMissing", func(t *testing.T) {
		openingMockWithoutRemote := schemas.CreateOpeningRequest{
			Role:     "Developer JavaScript",
			Company:  "Tech Corp",
			Location: "USA",
			Link:     "https://global.com/job",
			Salary:   50000,
		}
		mockErr := internal_error.NewInternalServerError("param: remote (type: bool) is required")

		err := openingUsecase.Create(openingMockWithoutRemote)

		assert.Error(t, err, "expected an error when creating the opening without company")
		assert.EqualError(t, mockErr, err.Error(), "The error message must be specific")

		openingRepo.AssertNotCalled(t, "Create")
	})

	t.Run("ShouldReturnAnErrorIfTheSalaryFieldIsLessThan3k", func(t *testing.T) {
		openingMockWithLowSalary := schemas.CreateOpeningRequest{
			Role:     "Junior Developer",
			Company:  "Tech Corp",
			Location: "Spain",
			Link:     "https://spain.com/job",
			Remote:   new(bool),
			Salary:   2500,
		}

		mockErr := internal_error.NewInternalServerError("salary must be at least 3k")

		err := openingUsecase.Create(openingMockWithLowSalary)

		assert.Error(t, err, "expected an error when the salary is less than 3K")
		assert.EqualError(t, mockErr, err.Error())
		openingRepo.AssertNotCalled(t, "Create")
	})
}
