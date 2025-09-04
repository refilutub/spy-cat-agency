package handlers

import (
	"net/http"
	"spy-cat-agency/internal/missions/dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UpdateMission godoc
// @Summary Update a mission
// @Description Update mission information including marking it as completed
// @Tags Missions
// @Accept json
// @Produce json
// @Param id path string true "Mission ID" format(uuid)
// @Param mission body dtos.UpdateMissionRequest true "Mission update information"
// @Success 200 {object} dtos.MissionResponseDTO "Mission updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid mission ID or request body"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /missions/{id} [put]
// @Security BearerAuth
func (h *MissionsHandler) UpdateMission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission ID"})
		return
	}

	var request dtos.UpdateMissionRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	mission, err := h.Service.UpdateMission(id, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update mission"})
		return
	}

	c.JSON(http.StatusOK, mission)
}
