package handler_test

import (
	"bytes"
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

func TestUpdateOpeningHandler(t *testing.T) {
	t.Run("ShouldReturnBadRequestWhenIDIsInvalid", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.PUT("/openings/:id", handler.Update)

		req, _ := http.NewRequest("PUT", "/openings/invalid", bytes.NewBuffer([]byte("{}")))
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

		mockUseCase.AssertNotCalled(t, "Update", mock.Anything, mock.Anything)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfAnEmptyBodyIsRequiredForTheUpdate", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.PUT("/openings/:id", handler.Update)

		mockErr := internal_error.NewBadRequestError("at least one valid field must be provided")
		mockUseCase.On("Update", mock.Anything, mock.Anything).Return(mockErr)

		req, _ := http.NewRequest("PUT", "/openings/1", bytes.NewBuffer([]byte("{}")))
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
		assert.Equal(t, "at least one valid field must be provided", resp.Message)
		assert.Equal(t, resp.ErrorCode, w.Code)

		mockUseCase.AssertCalled(t, "Update", mock.Anything, mock.Anything)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfTheOpeningIsNotFound", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.PUT("/openings/:id", handler.Update)

		openingReq := schemas.UpdateOpeningRequest{
			Salary: 50000,
		}

		mockErr := internal_error.NewNotFoundError("opening not found")
		mockUseCase.On("Update", mock.Anything, mock.Anything).Return(mockErr)

		reqJsonBody, _ := json.Marshal(openingReq)
		req, _ := http.NewRequest("PUT", "/openings/3000", bytes.NewBuffer(reqJsonBody))
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
		assert.Equal(t, "opening not found", resp.Message)
		assert.Equal(t, resp.ErrorCode, w.Code)

		mockUseCase.AssertCalled(t, "Update", mock.Anything, mock.Anything)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorWhenTryingToUpdateTheSalaryToAValueLessThan3k", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.PUT("/openings/:id", handler.Update)

		openingReq := schemas.UpdateOpeningRequest{
			Salary: 2999,
		}
		mockErr := internal_error.NewBadRequestError("salary must be at least 3k")
		mockUseCase.On("Update", mock.Anything, mock.Anything).Return(mockErr)

		reqJsonBody, _ := json.Marshal(openingReq)
		req, _ := http.NewRequest("PUT", "/openings/1", bytes.NewBuffer(reqJsonBody))
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
		assert.Equal(t, "salary must be at least 3k", resp.Message)

		mockUseCase.AssertCalled(t, "Update", mock.Anything, mock.Anything)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldUpdateAllJobOpeningParameters", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.PUT("/openings/:id", handler.Update)

		ID := 3000
		remote := true
		openingReq := schemas.UpdateOpeningRequest{
			Role:     "JavaScript Developer",
			Company:  "Tech Corp",
			Location: "Remote",
			Link:     "http://example.com",
			Remote:   &remote,
			Salary:   50000,
		}
		mockUseCase.On("Update", mock.Anything, mock.Anything).Return((*internal_error.InternalError)(nil))

		reqJsonBody, _ := json.Marshal(openingReq)
		req, _ := http.NewRequest("PUT", fmt.Sprintf("/openings/%d", uint(ID)), bytes.NewBuffer(reqJsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		var resp struct {
			Message string `json:"message"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Nil(t, err)
		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, fmt.Sprintf("opening with id: %d updated successfully", ID), resp.Message)

		mockUseCase.AssertCalled(t, "Update", uint(ID), mock.Anything)
		mockUseCase.AssertExpectations(t)
	})
}
