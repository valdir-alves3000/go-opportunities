package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/handler"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"github.com/valdir-alves3000/go-opportunities/test/mocks"
)

func TestListOpeningUsecase(t *testing.T) {
	t.Run("ShouldReturnErrorWhenThereIsAnErrorInTheDB", func(t *testing.T) {})

	t.Run("ShouldReturnErrorWhenInvalidPageIsPassed", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.GET("/openings", handler.List)

		req, _ := http.NewRequest("GET", "/openings?page=invalid", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var resp struct {
			Message   string `json:"message"`
			ErrorCode int    `json:"errorCode"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Nil(t, err)
		assert.Equal(t, "invalid page number", resp.Message)
		assert.Equal(t, resp.ErrorCode, w.Code)

		mockUseCase.AssertNotCalled(t, "List", mock.Anything, mock.Anything)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnErrorWhenNoOpeningsAreFound", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.GET("/openings", handler.List)

		mockErr := internal_error.NewNotFoundError("opening record not found")
		mockUseCase.On("ListOpenings", 1).Return([]schemas.Opening{}, mockErr).Once()

		req, _ := http.NewRequest("GET", "/openings?page=1", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var resp struct {
			Message   string `json:"message"`
			ErrorCode int    `json:"errorCode"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Nil(t, err)
		assert.Equal(t, "opening record not found", resp.Message)
		assert.Equal(t, resp.ErrorCode, w.Code)

		mockUseCase.AssertCalled(t, "ListOpenings", 1)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnTheFirst10OpeningsIfPage0IsPassed", func(t *testing.T) {
		mockListOpenings := mocks.GenerateListOpenings(10)
		mockOpenings := make([]schemas.Opening, 10)
		for i := 0; i < 10; i++ {
			mockOpenings[i] = mockListOpenings[i]
		}
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.GET("/openings", handler.List)

		mockUseCase.On("ListOpenings", 0).Return(mockOpenings, (*internal_error.InternalError)(nil)).Once()

		req, _ := http.NewRequest("GET", "/openings?page=0", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var resp struct {
			Message string                    `json:"message"`
			Data    []schemas.OpeningResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Nil(t, err)
		assert.Equal(t, "list-openings successfully", resp.Message)
		assert.Equal(t, resp.Data[0].ID, uint(1))
		assert.Equal(t, resp.Data[9].ID, uint(10))
		assert.Len(t, resp.Data, 10)

		assert.Equal(t, resp.Data[0].ID, mockListOpenings[0].ID)
		assert.Equal(t, resp.Data[0].Role, mockListOpenings[0].Role)
		assert.Equal(t, resp.Data[0].Company, mockListOpenings[0].Company)
		assert.Equal(t, resp.Data[0].Location, mockListOpenings[0].Location)

		assert.Equal(t, resp.Data[9].Link, mockListOpenings[9].Link)
		assert.Equal(t, resp.Data[9].Remote, mockListOpenings[9].Remote)
		assert.Equal(t, resp.Data[9].Salary, mockListOpenings[9].Salary)

		mockUseCase.AssertCalled(t, "ListOpenings", 0)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnTheFirst10OpeningsIfPage1IsPassed", func(t *testing.T) {
		mockListOpenings := mocks.GenerateListOpenings(10)
		mockOpenings := make([]schemas.Opening, 10)
		for i := 0; i < 10; i++ {
			mockOpenings[i] = mockListOpenings[i]
		}
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.GET("/openings", handler.List)

		mockUseCase.On("ListOpenings", 1).Return(mockOpenings, (*internal_error.InternalError)(nil)).Once()

		req, _ := http.NewRequest("GET", "/openings?page=1", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var resp struct {
			Message string                    `json:"message"`
			Data    []schemas.OpeningResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Nil(t, err)
		assert.Equal(t, "list-openings successfully", resp.Message)
		assert.Equal(t, resp.Data[0].ID, uint(1))
		assert.Equal(t, resp.Data[9].ID, uint(10))
		assert.Len(t, resp.Data, 10)

		assert.Equal(t, resp.Data[0].ID, mockOpenings[0].ID)
		assert.Equal(t, resp.Data[0].Role, mockOpenings[0].Role)
		assert.Equal(t, resp.Data[0].Company, mockOpenings[0].Company)
		assert.Equal(t, resp.Data[0].Location, mockOpenings[0].Location)

		assert.Equal(t, resp.Data[9].Link, mockOpenings[9].Link)
		assert.Equal(t, resp.Data[9].Remote, mockOpenings[9].Remote)
		assert.Equal(t, resp.Data[9].Salary, mockOpenings[9].Salary)

		mockUseCase.AssertCalled(t, "ListOpenings", 1)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnOpeningsFrom21To30", func(t *testing.T) {
		mockListOpenings := mocks.GenerateListOpenings(30)
		mockOpenings := make([]schemas.Opening, 10)
		for i := 20; i < 30; i++ {
			mockOpenings[i-20] = mockListOpenings[i]
		}
		page := 3

		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.GET("/openings", handler.List)

		mockUseCase.On("ListOpenings", page).Return(mockOpenings, (*internal_error.InternalError)(nil)).Once()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/openings?page=%d", page), nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var resp struct {
			Message string                    `json:"message"`
			Data    []schemas.OpeningResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Nil(t, err)
		assert.Equal(t, "list-openings successfully", resp.Message)
		assert.Equal(t, resp.Data[0].ID, uint(21))
		assert.Equal(t, resp.Data[9].ID, uint(30))
		assert.Len(t, resp.Data, 10)

		assert.Equal(t, resp.Data[0].ID, mockListOpenings[20].ID)
		assert.Equal(t, resp.Data[0].Role, mockListOpenings[20].Role)
		assert.Equal(t, resp.Data[0].Company, mockListOpenings[20].Company)
		assert.Equal(t, resp.Data[0].Location, mockListOpenings[20].Location)

		assert.Equal(t, resp.Data[9].Link, mockListOpenings[29].Link)
		assert.Equal(t, resp.Data[9].Remote, mockListOpenings[29].Remote)
		assert.Equal(t, resp.Data[9].Salary, mockListOpenings[29].Salary)

		mockUseCase.AssertCalled(t, "ListOpenings", 3)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnOpeningsFrom41To50", func(t *testing.T) {
		mockListOpenings := mocks.GenerateListOpenings(50)
		mockOpenings := make([]schemas.Opening, 10)
		for i := 40; i < 50; i++ {
			mockOpenings[i-40] = mockListOpenings[i]
		}

		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.GET("/openings", handler.List)

		mockUseCase.On("ListOpenings", 4).Return(mockOpenings, (*internal_error.InternalError)(nil)).Once()

		req, _ := http.NewRequest("GET", "/openings?page=4", nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var resp struct {
			Message string                    `json:"message"`
			Data    []schemas.OpeningResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Nil(t, err)
		assert.Equal(t, "list-openings successfully", resp.Message)
		assert.Equal(t, "41", fmt.Sprintf("%d", resp.Data[0].ID))
		assert.Equal(t, "50", fmt.Sprintf("%d", resp.Data[9].ID))
		assert.Len(t, resp.Data, 10)

		assert.Equal(t, resp.Data[0].ID, mockListOpenings[40].ID)
		assert.Equal(t, resp.Data[0].Role, mockListOpenings[40].Role)
		assert.Equal(t, resp.Data[0].Company, mockListOpenings[40].Company)
		assert.Equal(t, resp.Data[0].Location, mockListOpenings[40].Location)

		assert.Equal(t, resp.Data[9].Link, mockListOpenings[49].Link)
		assert.Equal(t, resp.Data[9].Remote, mockListOpenings[49].Remote)
		assert.Equal(t, resp.Data[9].Salary, mockListOpenings[49].Salary)

		mockUseCase.AssertCalled(t, "ListOpenings", 4)
		mockUseCase.AssertExpectations(t)
	})
}
