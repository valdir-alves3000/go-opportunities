package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valdir-alves3000/go-opportunities/config/rest_err"
)

// @BasePath /api/v1

// @Summary Delete opening
// @Description Delete a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path string true "Opening ID"
// @Success 200 {object} DeleteOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /openings/{id} [delete]
func (h *OpeningHandler) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendError(c, http.StatusBadRequest, "invalid ID")
		return
	}

	errCase := h.useCase.DeleteByID(uint(id))
	if errCase != nil {
		rest_err := rest_err.ConvertError(errCase)
		sendError(c, rest_err.Code, rest_err.Message)
		return
	}

	sendSuccess(c, fmt.Sprintf("opening with id: %d deleted", id), nil)
}
