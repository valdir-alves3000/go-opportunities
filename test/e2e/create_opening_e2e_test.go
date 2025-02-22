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

func createOpening(opening schemas.CreateOpeningRequest) *httptest.ResponseRecorder {
	reqJsonBody, _ := json.Marshal(opening)
	req, _ := http.NewRequest("POST", basePath+"/openings", bytes.NewBuffer(reqJsonBody))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	return w
}

func TestCreateOpeningE2E(t *testing.T) {
	clearDatabase := func() {
		db.Exec("DELETE FROM openings")
	}

	t.Run("ShouldSuccessfullyCreateAJobOpening", func(t *testing.T) {
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
			Message string `json:"message"`
		}

		err := json.Unmarshal(w.Body.Bytes(), &resp)
		if err != nil {
			t.Fatal(err)
		}

		assert.Equal(t, http.StatusCreated, w.Code)
		assert.NoError(t, err)
		assert.Equal(t, fmt.Sprintf("opening %s created successfully", openingReq.Role), resp.Message)
	})

	t.Run("ShouldReturnAnErrorIfTheRoleIsEmpty", func(t *testing.T) {
		clearDatabase()
		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "",
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

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "param: role (type: string) is required", resp.Message)
		assert.Equal(t, http.StatusBadRequest, resp.ErrorCode)
	})

	t.Run("ShouldReturnAnErrorIfTheCompanyIsEmpty", func(t *testing.T) {
		clearDatabase()
		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "",
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

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "param: company (type: string) is required", resp.Message)
		assert.Equal(t, http.StatusBadRequest, resp.ErrorCode)
	})

	t.Run("ShouldReturnAnErrorIfTheLocationIsEmpty", func(t *testing.T){
		clearDatabase()
		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "",
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

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "param: location (type: string) is required", resp.Message)
		assert.Equal(t, http.StatusBadRequest, resp.ErrorCode)
	})

	t.Run("ShouldReturnAnErrorIfTheLinkIsEmpty", func(t *testing.T){
		clearDatabase()
		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "Silicon Valley",
			Link:     "",
			Salary:   50000,
			Remote:   &remote,
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

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "param: link (type: string) is required", resp.Message)
		assert.Equal(t, http.StatusBadRequest, resp.ErrorCode)
	})

	t.Run("ShouldReturnAnErrorIfTheRemoteIsEmpty",func(t *testing.T) {
		clearDatabase()
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "Silicon Valley",
			Link:     "http://example.com",
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

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "param: remote (type: bool) is required", resp.Message)
		assert.Equal(t, http.StatusBadRequest, resp.ErrorCode)
	})

	t.Run("ShouldReturnAnErrorIfTheSalaryFieldIsLessThan3k", func(t *testing.T){
		clearDatabase()
		remote := true
		openingReq := schemas.CreateOpeningRequest{
			Role:     "Go Developer",
			Company:  "Tech Corp",
			Location: "Silicon Valley",
			Link:     "http://example.com",
			Remote:   &remote,
			Salary:   2999,
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

		assert.Equal(t, http.StatusBadRequest, w.Code)
		assert.Equal(t, "salary must be at least 3k", resp.Message)
		assert.Equal(t, http.StatusBadRequest, resp.ErrorCode)
	})
}
