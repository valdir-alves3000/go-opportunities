package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/handler"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"github.com/valdir-alves3000/go-opportunities/test/mocks"
	"gorm.io/gorm"
)

func TestShowOpeningHandler(t *testing.T) {
	t.Run("ShouldSuccessfullyGetAJobOpeningByID", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.GET("/openings/:id", handler.ShowOpening)

		ID := 1
		opening := schemas.Opening{
			Model:    gorm.Model{ID: uint(ID)},
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "Remote",
			Link:     "http://example.com",
			Remote:   true,
			Salary:   50000,
		}

		mockUseCase.On("GetByID", uint(ID)).Return(&opening, (*internal_error.InternalError)(nil)).Once()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/openings/%d", ID), nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var resp struct {
			Message string          `json:"message"`
			Data    schemas.Opening `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NoError(t, err)
		assert.Equal(t, "show-opening successfully", resp.Message)
		assert.Equal(t, opening, resp.Data)

		mockUseCase.AssertCalled(t, "GetByID", uint(ID))
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnBadRequestWhenIDIsInvalid", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.GET("/openings/:id", handler.ShowOpening)
		req, _ := http.NewRequest("GET", "/openings/invalid", nil)
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
		assert.Equal(t, "invalid ID", resp.Message)
		assert.Equal(t, resp.ErrorCode, w.Code)
	})

	t.Run("ShouldReturnNotFoundWhenOpeningDoesNotExist", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.GET("/openings/:id", handler.ShowOpening)
		ID := 999
		mockUseCase.On("GetByID", uint(ID)).Return((*schemas.Opening)(nil), internal_error.NewNotFoundError("opening not found")).Once()

		req, _ := http.NewRequest("GET", fmt.Sprintf("/openings/%d", ID), nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code)

		var response struct {
			Message   string `json:"message"`
			ErrorCode string `json:"errorCode"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Error(t, err)
		assert.Contains(t, "opening not found", response.Message)
	})
}
