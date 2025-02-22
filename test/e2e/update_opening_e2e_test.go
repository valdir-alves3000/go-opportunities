package e2e

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
)

func updateOpening(id uint, opening schemas.UpdateOpeningRequest) *httptest.ResponseRecorder {
	reqJsonBody, _ := json.Marshal(opening)
	req, _ := http.NewRequest("PUT", fmt.Sprintf(basePath+"/openings/%d", id), bytes.NewBuffer(reqJsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

func TestUodateOpeningE2E(t *testing.T) {
	clearDatabase := func() {
		db.Exec("DELETE FROM openings")
	}

	t.Run("ShouldReturnAnErrorIfTheOpeningIsNotFound", func(t *testing.T) {
		clearDatabase()
		openingReq := schemas.UpdateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "Silicon Valley",
			Link:     "http://example.com",
			Remote:   new(bool),
			Salary:   50000,
		}

		w := updateOpening(uint(0), openingReq)
		var resp struct {
			Message   string `json:"message"`
			ErrorCode int    `json:"errorCode"`
		}
		err := json.Unmarshal(w.Body.Bytes(), &resp)

		assert.Equal(t, http.StatusNotFound, w.Code)
		assert.Nil(t, err)
		assert.Equal(t, "opening not found", resp.Message)
		assert.Equal(t, http.StatusNotFound, resp.ErrorCode)
	})

	t.Run("ShouldReturnAnErrorIfAnEmptyBodyIsRequiredForTheUpdate", func(t *testing.T) {
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
			Message   string `json:"message"`
			ErrorCode int    `json:"errorCode"`
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

		upOpening := schemas.UpdateOpeningRequest{}
		w = updateOpening(opening.ID, upOpening)

		err = json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NoError(t, err)
		assert.Equal(t, "at least one valid field must be provided", resp.Message)
		assert.Equal(t, http.StatusBadRequest, resp.ErrorCode)
	})

	t.Run("ShouldReturnAnErrorWhenTryingToUpdateTheSalaryToAValueLessThan3k",func(t *testing.T) {
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
			Message   string `json:"message"`
			ErrorCode int    `json:"errorCode"`
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

		upOpening := schemas.UpdateOpeningRequest{
			Salary: 2999,
		}
		w = updateOpening(opening.ID, upOpening)

		err = json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.NoError(t, err)
		assert.Equal(t, "salary must be at least 3k", resp.Message)
		assert.Equal(t, http.StatusBadRequest, resp.ErrorCode)
	})

	t.Run("ShouldUpdateAllJobOpeningParameters", func(t *testing.T) {
		clearDatabase()
		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "Silicon Valley",
			Link:     "http://example.com",
			Remote:   &remote,
			Salary:   50000,
		}

		w := createOpening(openingReq)
		var resp struct {
			Message   string `json:"message"`
			ErrorCode int    `json:"errorCode"`
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

		remote = false
		upOpening := schemas.UpdateOpeningRequest{
			Role:     "Javascript Developer",
			Company:  "+3000 DEV",
			Location: "New York",
			Link:     "http://+3000dev.com/job",
			Remote:   &remote,
			Salary:   30000,
		}
		w = updateOpening(opening.ID, upOpening)

		err = json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusOK, w.Code)
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("opening with id: %d updated successfully", opening.ID), resp.Message)

		result = db.First(&opening, opening.ID)
		assert.NoError(t, result.Error)
		assert.Equal(t, "Javascript Developer", opening.Role)
		assert.Equal(t, "+3000 DEV", opening.Company)
		assert.Equal(t, upOpening.Location, opening.Location)
		assert.Equal(t, "http://+3000dev.com/job", opening.Link)
		assert.Equal(t, remote, opening.Remote)
		assert.Equal(t, int64(30000), opening.Salary)
	})

}
