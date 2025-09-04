package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// GetSpyCat godoc
// @Summary Get a spy cat by ID
// @Description Retrieve detailed information about a specific spy cat
// @Tags Cats
// @Accept json
// @Produce json
// @Param spycatId path string true "Spy Cat ID" format(uuid)
// @Success 200 {object} dtos.SpyCatSingleResponseDTO "Spy cat retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid ID format or service error"
// @Failure 404 {object} map[string]interface{} "Spy cat not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /spycats/{spycatId} [get]
// @Security BearerAuth
func (h *SpyCatHandler) GetSpyCat(c *gin.Context) {
	idStr := c.Param("spycatId")
	spycatId, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid spycat id"})
		return
	}

	spycat, err := h.Service.GetSpyCat(spycatId)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "spycat not found"})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, spycat)

}
