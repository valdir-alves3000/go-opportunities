package opening_usecase_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"github.com/valdir-alves3000/go-opportunities/test/mocks"
	"gorm.io/gorm"
)

func TestListOpeningUsecase(t *testing.T) {
	openingUsecase, openingRepo := setupUsecaseTest()

	t.Run("ShouldReturnErrorWhenNoOpeningsAreFound", func(t *testing.T) {
		mockErr := internal_error.NewNotFoundError("opening record not found")
		openingRepo.On("FindAll", 10, 0).Return([]schemas.Opening{}, nil).Once()

		result, err := openingUsecase.ListOpenings(0)

		assert.Error(t, err, "expected an error")
		assert.Equal(t, mockErr, err, "The error message must be specific")

		openingRepo.AssertCalled(t, "FindAll", 10, 0)
		assert.Nil(t, result)
	})

	t.Run("ShouldReturnErrorWhenThereIsAnErrorInTheDB", func(t *testing.T) {
		mockErr := internal_error.NewNotFoundError("opening record not found")
		openingRepo.On("FindAll", 10, 0).Return([]schemas.Opening{}, gorm.ErrRecordNotFound).Once()

		result, err := openingUsecase.ListOpenings(0)

		assert.Error(t, err, "expected an error")
		assert.EqualError(t, mockErr, err.Message, "The error message must be specific")
		assert.Equal(t, mockErr, err, "The error message must be specific")

		openingRepo.AssertCalled(t, "FindAll", 10, 0)
		assert.Nil(t, result)
	})

	t.Run("ShouldReturnTheFirst10OpeningsIfPage0IsPassed", func(t *testing.T) {
		mockListOpenings := mocks.GenerateListOpenings(10)
		mockOpenings := make([]schemas.Opening, 10)
		for i := 0; i < 10; i++ {
			mockOpenings[i] = mockListOpenings[i]
		}

		openingRepo.On("FindAll", 10, 0).Return(mockOpenings, nil).Once()

		result, err := openingUsecase.ListOpenings(0)

		assert.Nil(t, err)
		assert.Len(t, result, 10, "Should return exactly 10 results")
		openingRepo.AssertCalled(t, "FindAll", 10, 0)
		assert.Equal(t, result[0].ID, uint(1))
		assert.Equal(t, result[9].ID, uint(10))
		assert.Equal(t, result[0], mockOpenings[0])
	})

	t.Run("ShouldReturnTheFirst10OpeningsIfPage1IsPassed", func(t *testing.T) {
		mockListOpenings := mocks.GenerateListOpenings(10)
		mockOpenings := make([]schemas.Opening, 10)
		for i := 0; i < 10; i++ {
			mockOpenings[i] = mockListOpenings[i]
		}

		openingRepo.On("FindAll", 10, 0).Return(mockOpenings, nil).Once()

		result, err := openingUsecase.ListOpenings(1)

		assert.Nil(t, err)
		assert.Len(t, result, 10, "Should return exactly 10 results")
		openingRepo.AssertCalled(t, "FindAll", 10, 0)
		assert.Equal(t, result[0].ID, uint(1))
		assert.Equal(t, result[9].ID, uint(10))
		assert.Equal(t, result[0], mockOpenings[0])
	})

	t.Run("ShouldReturnOpeningsFrom21To30", func(t *testing.T) {
		mockListOpenings := mocks.GenerateListOpenings(30)
		mockOpenings := make([]schemas.Opening, 10)
		for i := 20; i < 30; i++ {
			mockOpenings[i-20] = mockListOpenings[i]
		}

		openingRepo.On("FindAll", 10, 20).Return(mockOpenings, nil).Once()

		result, err := openingUsecase.ListOpenings(3)

		assert.Nil(t, err)
		assert.Len(t, result, 10, "Should return exactly 10 results")
		openingRepo.AssertCalled(t, "FindAll", 10, 20)
		assert.Equal(t, result[0].ID, uint(21))
		assert.Equal(t, result[9].ID, uint(30))
		assert.Equal(t, result[0], mockOpenings[0])
	})

	t.Run("ShouldReturnOpeningsFrom41To50", func(t *testing.T) {
		mockListOpenings := mocks.GenerateListOpenings(50)
		mockOpenings := make([]schemas.Opening, 10)
		for i := 40; i < 50; i++ {
			mockOpenings[i-40] = mockListOpenings[i]
		}

		openingRepo.On("FindAll", 10, 40).Return(mockOpenings, nil).Once()

		result, err := openingUsecase.ListOpenings(5)

		assert.Nil(t, err)
		assert.Len(t, result, 10, "Should return exactly 10 results")
		openingRepo.AssertCalled(t, "FindAll", 10, 40)
		assert.Equal(t, result[0].ID, uint(41))
		assert.Equal(t, result[9].ID, uint(50))
		assert.Equal(t, result[0], mockOpenings[0])
	})
}
