package handlers

import (
	"net/http"
	"spy-cat-agency/internal/missions/dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AddTargetToMission godoc
// @Summary Add a target to a mission
// @Description Add a new target to an existing mission. A target cannot be added if the mission is already completed.
// @Tags Missions
// @Accept json
// @Produce json
// @Param id path string true "Mission ID" format(uuid)
// @Param target body dtos.AddTargetRequest true "Target information to add"
// @Success 201 {object} dtos.TargetResponseDTO "Target added to mission successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid mission ID, invalid request body, or business rule violation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /missions/{id}/targets [post]
// @Security BearerAuth
func (h *MissionsHandler) AddTargetToMission(c *gin.Context) {
	missionIDStr := c.Param("id")
	missionID, err := uuid.Parse(missionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission ID"})
		return
	}

	var request dtos.AddTargetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	target, err := h.Service.AddTargetToMission(missionID, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, target)
}
