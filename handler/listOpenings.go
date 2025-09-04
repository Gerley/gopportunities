package handler

import (
	"net/http"

	"github.com/Gerley/gopportunities/schemas"
	"github.com/gin-gonic/gin"
)

// @BasePah /api/v1

// @Summary List openings
// @Description List all job opening
// @Tags Openings
// @Accept json
// @Produce json
// @Success 200 {object} ListOpeningsResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /openings [get]
func ListOpeningsHandler(ctx *gin.Context) {
	openings := []schemas.Opening{}

	if err := db.Find(&openings).Error; err != nil {
		sendError(ctx, http.StatusInternalServerError, "error listing openings")
		return
	}
	sendSuccess(ctx, "list_openings", openings)
}
