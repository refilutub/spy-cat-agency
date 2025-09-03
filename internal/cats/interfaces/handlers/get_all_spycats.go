package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetAllSpyCats godoc
// @Summary Get all spy cats
// @Description Retrieve a list of all spy cats in the system
// @Tags Cats
// @Accept json
// @Produce json
// @Success 200 {array} dtos.SpyCatListResponseDTO "List of spy cats retrieved successfully"
// @Failure 404 {object} map[string]interface{} "No spy cats found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /spycats [get]
// @Security BearerAuth
func (h *SpyCatHandler) GetAllSpyCats(c *gin.Context) {
	spycats, err := h.Service.GetAllSpyCats()
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "no records found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, spycats)
}
