package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// DeleteMission godoc
// @Summary Delete a mission
// @Description Delete a mission from the system. A mission cannot be deleted if it is already assigned to a cat.
// @Tags Missions
// @Accept json
// @Produce json
// @Param id path string true "Mission ID" format(uuid)
// @Success 200 {object} map[string]interface{} "Mission deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid mission ID, mission assigned to cat, or service error"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /missions/{id} [delete]
// @Security BearerAuth
func (h *MissionsHandler) DeleteMission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission ID"})
		return
	}

	if err := h.Service.DeleteMission(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "mission deleted successfully"})
}
