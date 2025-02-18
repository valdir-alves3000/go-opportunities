package handler_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdir-alves3000/go-opportunities/internal/handler"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"github.com/valdir-alves3000/go-opportunities/test/mocks"
)

func TestDeleteOpeningHandler(t *testing.T) {
	t.Run("ShouldDeleteOpeningSuccessfully", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.DELETE("/openings/:id", handler.Delete)

		ID := 3000
		mockUseCase.On("DeleteByID", uint(ID)).Return((*internal_error.InternalError)(nil)).Once()

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/openings/%d", ID), nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var resp struct {
			Message string `json:"message"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Nil(t, err)
		assert.Equal(t, fmt.Sprintf("opening with id: %d deleted successfully", ID), resp.Message)
	})

	t.Run("ShouldReturnErrorWhenTryingToDeleteNotFoundOpening", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.DELETE("/openings/:id", handler.Delete)

		ID := 999
		mockUseCase.On("DeleteByID", uint(ID)).Return(internal_error.NewNotFoundError("opening not found")).Once()

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/openings/%d", ID), nil)
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
		assert.Equal(t, "opening not found", response.Message)
	})

	t.Run("ShouldReturnErrorWhenTryingToDeleteOpeningWithInvalidID", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.DELETE("/openings/:id", handler.Delete)

		req, _ := http.NewRequest("DELETE", "/openings/invalid", nil)
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

	t.Run("ShouldReturnAnErrorIfTheDBFails", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.DELETE("/openings/:id", handler.Delete)

		ID := 999
		errDB := internal_error.NewInternalServerError("error deleting opening")
		mockUseCase.On("DeleteByID", uint(ID)).Return(errDB).Once()

		req, _ := http.NewRequest("DELETE", fmt.Sprintf("/openings/%d", ID), nil)
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var response struct {
			Message   string `json:"message"`
			ErrorCode string `json:"errorCode"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &response)
		assert.Error(t, err)
		assert.Equal(t, "error deleting opening", response.Message)

		mockUseCase.AssertCalled(t, "DeleteByID", uint(ID))
		mockUseCase.AssertExpectations(t)
	})
}
