package handlers

import (
	"net/http"
	"spy-cat-agency/internal/missions/dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
