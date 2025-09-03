package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// GetMission godoc
// @Summary Get a mission by ID
// @Description Retrieve detailed information about a specific mission including its targets and assigned cat
// @Tags Missions
// @Accept json
// @Produce json
// @Param id path string true "Mission ID" format(uuid)
// @Success 200 {object} dtos.MissionResponseDTO "Mission retrieved successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid mission ID"
// @Failure 404 {object} map[string]interface{} "Mission not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /missions/{id} [get]
// @Security BearerAuth
func (h *MissionsHandler) GetMission(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission ID"})
		return
	}

	mission, err := h.Service.GetMission(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "mission not found"})
		return
	}

	c.JSON(http.StatusOK, mission)
}

// ListMissions godoc
// @Summary List all missions
// @Description Retrieve a list of all missions in the system
// @Tags Missions
// @Accept json
// @Produce json
// @Success 200 {array} dtos.MissionResponseDTO "List of missions retrieved successfully"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /missions [get]
// @Security BearerAuth
func (h *MissionsHandler) ListMissions(c *gin.Context) {
	missions, err := h.Service.ListMissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve missions"})
		return
	}

	c.JSON(http.StatusOK, missions)
}
