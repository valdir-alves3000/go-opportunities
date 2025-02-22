package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valdir-alves3000/go-opportunities/config/rest_err"
)

// @BasePath /api/v1

// @Summary List openings
// @Description Get a list of all openings with pagination
// @Tags Openings
// @Accept json
// @Produce json
// @Param page query int false "Page number"
// @Success 200 {object} ListOpeningsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [get]
func (h *OpeningHandler) List(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "0"))
	if err != nil {
		sendError(c, http.StatusBadRequest, "invalid page number")
		return
	}

	openings, errCase := h.useCase.ListOpenings(page)
	if errCase != nil {
		restErr := rest_err.ConvertError(errCase)
		sendError(c, restErr.Code, errCase.Message)
		return
	}

	sendSuccess(c, "list-openings", openings)
}
