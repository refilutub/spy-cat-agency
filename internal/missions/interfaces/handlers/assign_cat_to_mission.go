package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
