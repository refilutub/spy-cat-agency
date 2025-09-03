package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeleteTarget godoc
// @Summary Delete a target
// @Description Delete a target from a mission. A target cannot be deleted if it is already completed.
// @Tags Targets
// @Accept json
// @Produce json
// @Param id path string true "Target ID" format(uuid)
// @Success 200 {object} map[string]interface{} "Target deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid target ID or business rule violation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /targets/{id} [delete]
// @Security BearerAuth
func (h *MissionsHandler) DeleteTarget(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid target ID"})
		return
	}

	if err := h.Service.DeleteTarget(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "target deleted successfully"})
}
