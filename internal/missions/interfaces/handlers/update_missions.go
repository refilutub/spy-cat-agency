package handlers

import (
	"net/http"
	"spy-cat-agency/internal/missions/dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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
