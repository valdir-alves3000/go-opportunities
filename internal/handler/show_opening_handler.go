package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valdir-alves3000/go-opportunities/config/rest_err"
)

// @BasePath /api/v1

// @Summary Show opening
// @Description Show a job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param id path int true "Opening Identification"
// @Success 200 {object} ShowOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /openings/{id} [get]
func (h *OpeningHandler) ShowOpening(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendError(c, http.StatusBadRequest, "invalid ID")
		return
	}

	op, errCase := h.useCase.GetByID(uint(id))
	if errCase != nil {
		restErr := rest_err.ConvertError(errCase)
		sendError(c, restErr.Code, restErr.Message)
		return
	}

	sendSuccess(c, "show-opening", op)
}
