package handlers

import (
	"net/http"
	"spy-cat-agency/internal/missions/dtos"

	"github.com/gin-gonic/gin"
)

// CreateMission godoc
// @Summary Create a new mission
// @Description Create a new mission in the system along with its targets
// @Tags Missions
// @Accept json
// @Produce json
// @Param mission body dtos.CreateMissionRequest true "Mission information with targets"
// @Success 201 {object} dtos.MissionCreateResponseDTO "Mission created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - validation error"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /missions [post]
// @Security BearerAuth
func (h *MissionsHandler) CreateMission(c *gin.Context) {
	var req dtos.CreateMissionRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newMission, err := h.Service.CreateMission(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, newMission)
}
