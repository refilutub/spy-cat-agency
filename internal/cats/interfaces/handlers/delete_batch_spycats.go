package handlers

import (
	"net/http"
	"spy-cat-agency/internal/cats/dtos"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// DeleteBatchSpyCatsHandler godoc
// @Summary Delete multiple spy cats
// @Description Remove multiple spy cats from the system by their IDs
// @Tags Cats
// @Accept json
// @Produce json
// @Param ids body dtos.DeleteIds true "Array of spy cat IDs to delete"
// @Success 204 "Spy cats deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid request body or service error"
// @Failure 404 {object} map[string]interface{} "No spy cats found to delete"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /spycats [delete]
// @Security BearerAuth
func (h *SpyCatHandler) DeleteBatchSpyCatsHandler(c *gin.Context) {
	var ids dtos.DeleteIds
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid spycats ids"})
		return
	}

	err := h.Service.DeleteBatchSpyCats(ids)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "no records found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
