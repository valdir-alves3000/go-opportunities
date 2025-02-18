package handler_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/handler"
	"github.com/valdir-alves3000/go-opportunities/internal/internal_error"
	"github.com/valdir-alves3000/go-opportunities/test/mocks"
)

func TestCreateOpeningHandler(t *testing.T) {
	t.Run("ShouldSuccessfullyCreateAJobOpening", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.POST("/openings", handler.Create)

		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "Remote",
			Link:     "http://example.com",
			Remote:   new(bool),
			Salary:   50000,
		}

		mockUseCase.On("Create", openingReq).Return((*internal_error.InternalError)(nil)).Once()

		reqJsonBody, _ := json.Marshal(openingReq)
		req, _ := http.NewRequest("POST", "/openings", bytes.NewBuffer(reqJsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusCreated, w.Code)
		mockUseCase.AssertCalled(t, "Create", openingReq)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfTheRoleIsEmpty", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.POST("/openings", handler.Create)

		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "",
			Company:  "Tech Corp",
			Location: "Remote",
			Link:     "http://example.com",
			Remote:   &remote,
			Salary:   50000,
		}
		expectedErr := "param: role (type: string) is required"
		mockUseCase.On("Create", openingReq).Return(internal_error.NewBadRequestError("param: role (type: string) is required")).Once()

		reqJsonBody, _ := json.Marshal(openingReq)
		req, _ := http.NewRequest("POST", "/openings", bytes.NewBuffer(reqJsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), expectedErr)
		mockUseCase.AssertCalled(t, "Create", openingReq)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfTheCompanyIsEmpty", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.POST("/openings", handler.Create)

		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "",
			Location: "Remote",
			Link:     "http://example.com",
			Remote:   &remote,
			Salary:   50000,
		}

		expectedErr := "param: company (type: string) is required"
		mockUseCase.On("Create", openingReq).Return(internal_error.NewBadRequestError("param: company (type: string) is required")).Once()

		reqJsonBody, _ := json.Marshal(openingReq)
		req, _ := http.NewRequest("POST", "/openings", bytes.NewBuffer(reqJsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), expectedErr)
		mockUseCase.AssertCalled(t, "Create", openingReq)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfTheLocationIsEmpty", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.POST("/openings", handler.Create)

		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "",
			Link:     "http://example.com",
			Remote:   &remote,
			Salary:   50000,
		}

		expectedErr := "param: location (type: string) is required"
		mockUseCase.On("Create", openingReq).Return(internal_error.NewBadRequestError("param: location (type: string) is required")).Once()

		reqJsonBody, _ := json.Marshal(openingReq)
		req, _ := http.NewRequest("POST", "/openings", bytes.NewBuffer(reqJsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), expectedErr)
		mockUseCase.AssertCalled(t, "Create", openingReq)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfTheLinkIsEmpty", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.POST("/openings", handler.Create)

		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "São Paulo",
			Link:     "",
			Remote:   &remote,
			Salary:   50000,
		}

		expectedErr := "param: link (type: string) is required"
		mockUseCase.On("Create", openingReq).Return(internal_error.NewBadRequestError("param: link (type: string) is required")).Once()

		reqJsonBody, _ := json.Marshal(openingReq)
		req, _ := http.NewRequest("POST", "/openings", bytes.NewBuffer(reqJsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), expectedErr)
		mockUseCase.AssertCalled(t, "Create", openingReq)
		mockUseCase.AssertExpectations(t)
	})

	t.Run("ShouldReturnAnErrorIfTheRemoteFieldIsMissing", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.POST("/openings", handler.Create)

		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "São Paulo",
			Link:     "http://example.com",
			Salary:   50000,
		}

		expectedErr := "param: remote (type: bool) is required"
		mockUseCase.On("Create", openingReq).Return(internal_error.NewBadRequestError("param: remote (type: bool) is required")).Once()

		reqJsonBody, _ := json.Marshal(openingReq)
		req, _ := http.NewRequest("POST", "/openings", bytes.NewBuffer(reqJsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), expectedErr)
	})

	t.Run("ShouldReturnAnErrorIfTheSalaryFieldIsLessThan3k", func(t *testing.T) {
		router := setupRouter()
		mockUseCase := new(mocks.OpeningUseCaseMock)
		handler := handler.NewOpeningHandler(mockUseCase)
		router.POST("/openings", handler.Create)

		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "São Paulo",
			Remote:   &remote,
			Link:     "http://example.com",
			Salary:   2999,
		}

		expectedErr := "salary must be at least 3k"
		mockUseCase.On("Create", openingReq).Return(internal_error.NewBadRequestError("salary must be at least 3k")).Once()

		reqJsonBody, _ := json.Marshal(openingReq)
		req, _ := http.NewRequest("POST", "/openings", bytes.NewBuffer(reqJsonBody))
		req.Header.Set("Content-Type", "application/json")

		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Contains(t, w.Body.String(), expectedErr)
	})
}
