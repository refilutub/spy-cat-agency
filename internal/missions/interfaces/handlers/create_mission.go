package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spy-cat-agency/internal/missions/dtos"
)

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
