package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"spy-cat-agency/internal/cats/dtos"
)

func (h *SpyCatHandler) DeleteBatchSpyCatsHandler(c *gin.Context) {
	var ids dtos.DeleteIds
	if err := c.ShouldBindJSON(&ids); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid spycats ids"})
		return
	}

	err := h.Service.DeleteBatchSpyCats(ids)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusNoContent)
}
