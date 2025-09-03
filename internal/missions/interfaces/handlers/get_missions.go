package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

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

func (h *MissionsHandler) ListMissions(c *gin.Context) {
	missions, err := h.Service.ListMissions()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to retrieve missions"})
		return
	}

	c.JSON(http.StatusOK, missions)
}
