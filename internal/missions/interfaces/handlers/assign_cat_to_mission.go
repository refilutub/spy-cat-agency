package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// AssignCatToMission godoc
// @Summary Assign a cat to a mission
// @Description Assign a spy cat to a specific mission. The mission must not be completed.
// @Tags Missions
// @Accept json
// @Produce json
// @Param id path string true "Mission ID" format(uuid)
// @Param assignment body object true "Cat assignment information" schema=object{cat_id=string}
// @Success 200 {object} dtos.MissionResponseDTO "Cat assigned to mission successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid mission ID, invalid request body, or business rule violation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /missions/{id}/assign-cat [post]
// @Security BearerAuth
func (h *MissionsHandler) AssignCatToMission(c *gin.Context) {
	missionIDStr := c.Param("id")
	missionID, err := uuid.Parse(missionIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid mission ID"})
		return
	}

	var request struct {
		CatID uuid.UUID `json:"cat_id"`
	}
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	mission, err := h.Service.AssignCatToMission(missionID, request.CatID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mission)
}
