package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/valdir-alves3000/go-opportunities/config/rest_err"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
)

func (h *OpeningHandler) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		sendError(c, http.StatusBadRequest, "invalid ID")
		return
	}

	var req schemas.UpdateOpeningRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	errCase := h.useCase.Update(uint(id), req)
	if errCase != nil {
		rest_err := rest_err.ConvertError(errCase)
		sendError(c, rest_err.Code, rest_err.Message)
		return
	}

	sendSuccess(c, fmt.Sprintf("opening with id: %d updated", id), nil)
}
