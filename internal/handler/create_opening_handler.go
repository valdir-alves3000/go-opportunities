package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/valdir-alves3000/go-opportunities/config/rest_err"
	"github.com/valdir-alves3000/go-opportunities/internal/core/schemas"
	"github.com/valdir-alves3000/go-opportunities/internal/core/usecases/opening_usecase"
)

type OpeningHandler struct {
	useCase opening_usecase.OpeningUsecase
}

func NewOpeningHandler(useCase opening_usecase.OpeningUsecase) *OpeningHandler {
	return &OpeningHandler{useCase: useCase}
}

// @BasePath /api/v1

// @Summary Create opening
// @Description Create a new job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Param request body schemas.CreateOpeningRequest true "Request body"
// @Success 200 {object} CreateOpeningResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /openings [post]
func (h *OpeningHandler) Create(c *gin.Context) {
	var req schemas.CreateOpeningRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	errCase := h.useCase.Create(req)
	if errCase != nil {
		rest_err := rest_err.ConvertError(errCase)
		sendError(c, rest_err.Code, rest_err.Message)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"message": fmt.Sprintf("opening %s created successfully", req.Role),
	})
}
