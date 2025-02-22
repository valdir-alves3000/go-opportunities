package e2e

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
)

func showOpening(id uint) *httptest.ResponseRecorder {
	req, _ := http.NewRequest("GET", fmt.Sprintf(basePath+"/openings/%d", id), nil)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

func TestShowOpeningE2E(t *testing.T) {
	clearDatabase := func() {
		db.Exec("DELETE FROM openings")
	}
	t.Run("ShouldReturnBadRequestWhenIDIsInvalid", func(t *testing.T) {
		clearDatabase()
		req, _ := http.NewRequest("GET", basePath+"/openings/invalid", nil)
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
	})

	t.Run("ShouldReturnNotFoundWhenOpeningDoesNotExist", func(t *testing.T) {
		clearDatabase()

		w := showOpening(1)
		var resp struct {
			Message   string `json:"message"`
			ErrorCode int    `json:"errorCode"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Equal(t, "opening not found", resp.Message)
		assert.Equal(t, http.StatusNotFound, resp.ErrorCode)
	})

	t.Run("ShouldSuccessfullyGetAJobOpeningByID", func(t *testing.T) {
		clearDatabase()
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "Silicon Valley",
			Link:     "http://example.com",
			Remote:   new(bool),
			Salary:   50000,
		}

		w := createOpening(openingReq)
		var resp struct {
			Message string          `json:"message"`
			Data    schemas.Opening `json:"data"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("opening %s created successfully", openingReq.Role), resp.Message)

		var opening schemas.Opening
		result := db.First(&opening)
		assert.NoError(t, result.Error)
		assert.NotZero(t, opening.ID)

		w = showOpening(opening.ID)

		err = json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NoError(t, err)
		assert.Equal(t, "show-opening successfully", resp.Message)
		assert.Equal(t, opening, resp.Data)
	})
}
