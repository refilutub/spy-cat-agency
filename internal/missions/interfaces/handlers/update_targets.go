package handlers

import (
	"net/http"
	"spy-cat-agency/internal/missions/dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UpdateTarget godoc
// @Summary Update a target
// @Description Update target information including marking it as completed. Notes cannot be updated if either the target or the mission is completed.
// @Tags Targets
// @Accept json
// @Produce json
// @Param id path string true "Target ID" format(uuid)
// @Param target body dtos.UpdateTargetRequest true "Target update information"
// @Success 200 {object} dtos.TargetResponseDTO "Target updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid target ID, invalid request body, or business rule violation"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /targets/{id} [put]
// @Security BearerAuth
func (h *MissionsHandler) UpdateTarget(c *gin.Context) {
	idStr := c.Param("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid target ID"})
		return
	}

	var request dtos.UpdateTargetRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	target, err := h.Service.UpdateTarget(id, request)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, target)
}
