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

func TestListOpeningE2E(t *testing.T) {
	clearDatabase := func() {
		db.Exec("DELETE FROM openings")
	}

	t.Run("ShouldReturnAnErrorIfThePageIsInvalid", func(t *testing.T) {
		clearDatabase()
		req, _ := http.NewRequest("GET", basePath+"/openings?page=invalid", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var resp struct {
			Message   string `json:"message"`
			ErrorCode int    `json:"errorCode"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Nil(t, err)
		assert.Equal(t, "invalid page number", resp.Message)
		assert.Equal(t, http.StatusBadRequest, resp.ErrorCode)
	})

	t.Run("ShouldReturnTheFirst10OpeningsIfPage0IsPassed", func(t *testing.T) {
		clearDatabase()

		for i := 0; i < 15; i++ {
			openingReq := schemas.CreateOpeningRequest{
				Role:     fmt.Sprintf("Go Developer %d", i),
				Company:  "Tech Corp",
				Location: "Silicon Valley",
				Link:     "http://example.com",
				Remote:   new(bool),
				Salary:   50000,
			}

			w := createOpening(openingReq)
			assert.Equal(t, http.StatusCreated, w.Code)
		}

		req, _ := http.NewRequest("GET", basePath+"/openings?page=0", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp struct {
			Message string                    `json:"message"`
			Data    []schemas.OpeningResponse `json:"data"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.NotEmpty(t, resp.Message)
		assert.Len(t, resp.Data, 10)
		assert.Equal(t, "list-openings successfully", resp.Message)

		assert.Equal(t, "Go Developer 0", resp.Data[0].Role)
		assert.Equal(t, "Go Developer 9", resp.Data[9].Role)
	})

	t.Run("ShouldReturnTheFirst10OpeningsIfPage1IsPassed", func(t *testing.T) {
		clearDatabase()

		for i := 0; i < 15; i++ {
			openingReq := schemas.CreateOpeningRequest{
				Role:     fmt.Sprintf("Go Developer %d", i),
				Company:  "Tech Corp",
				Location: "Silicon Valley",
				Link:     "http://example.com",
				Remote:   new(bool),
				Salary:   50000,
			}

			w := createOpening(openingReq)
			assert.Equal(t, http.StatusCreated, w.Code)
		}

		req, _ := http.NewRequest("GET", basePath+"/openings?page=1", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp struct {
			Message string                    `json:"message"`
			Data    []schemas.OpeningResponse `json:"data"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.NotEmpty(t, resp.Message)
		assert.Len(t, resp.Data, 10)
		assert.Equal(t, "list-openings successfully", resp.Message)

		assert.Equal(t, "Go Developer 0", resp.Data[0].Role)
		assert.Equal(t, "Go Developer 9", resp.Data[9].Role)
	})

	t.Run("ShouldReturnTheFirst7OpeningsIfPage1IsPassed", func(t *testing.T) {
		clearDatabase()

		for i := 0; i < 7; i++ {
			openingReq := schemas.CreateOpeningRequest{
				Role:     fmt.Sprintf("Go Developer %d", i),
				Company:  "Tech Corp",
				Location: "Silicon Valley",
				Link:     "http://example.com",
				Remote:   new(bool),
				Salary:   50000,
			}

			w := createOpening(openingReq)
			assert.Equal(t, http.StatusCreated, w.Code)
		}

		req, _ := http.NewRequest("GET", basePath+"/openings?page=1", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp struct {
			Message string                    `json:"message"`
			Data    []schemas.OpeningResponse `json:"data"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.NotEmpty(t, resp.Message)
		assert.Len(t, resp.Data, 7)
		assert.Equal(t, "list-openings successfully", resp.Message)

		assert.Equal(t, "Go Developer 0", resp.Data[0].Role)
		assert.Equal(t, "Go Developer 6", resp.Data[6].Role)
	})

	t.Run("ShouldReturnFrom11To15OpeningsIfPage2IsPassed", func(t *testing.T) {
		clearDatabase()

		for i := 1; i <= 15; i++ {
			openingReq := schemas.CreateOpeningRequest{
				Role:     fmt.Sprintf("Go Developer %d", i),
				Company:  "Tech Corp",
				Location: "Silicon Valley",
				Link:     "http://example.com",
				Remote:   new(bool),
				Salary:   50000,
			}

			w := createOpening(openingReq)
			assert.Equal(t, http.StatusCreated, w.Code)
		}

		req, _ := http.NewRequest("GET", basePath+"/openings?page=2", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp struct {
			Message string                    `json:"message"`
			Data    []schemas.OpeningResponse `json:"data"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.NotEmpty(t, resp.Message)
		assert.Len(t, resp.Data, 5)
		assert.Equal(t, "list-openings successfully", resp.Message)

		assert.Equal(t, "Go Developer 11", resp.Data[0].Role)
		assert.Equal(t, "Go Developer 15", resp.Data[4].Role)
	})

	t.Run("ShouldReturnFrom41To48OpeningsIfPage5IsPassed", func(t *testing.T) {
		clearDatabase()

		for i := 1; i <= 48; i++ {
			openingReq := schemas.CreateOpeningRequest{
				Role:     fmt.Sprintf("Go Developer %d", i),
				Company:  "Tech Corp",
				Location: "Silicon Valley",
				Link:     "http://example.com",
				Remote:   new(bool),
				Salary:   50000,
			}

			w := createOpening(openingReq)
			assert.Equal(t, http.StatusCreated, w.Code)
		}

		req, _ := http.NewRequest("GET", basePath+"/openings?page=5", nil)
		req.Header.Set("Content-Type", "application/json")
		w := httptest.NewRecorder()
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)

		var resp struct {
			Message string                    `json:"message"`
			Data    []schemas.OpeningResponse `json:"data"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.NoError(t, err)
		assert.NotEmpty(t, resp.Message)
		assert.Len(t, resp.Data, 8)
		assert.Equal(t, "list-openings successfully", resp.Message)

		assert.Equal(t, "Go Developer 41", resp.Data[0].Role)
		assert.Equal(t, "Go Developer 48", resp.Data[7].Role)
	})

}
