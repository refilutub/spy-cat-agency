package handlers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "no records found"})
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
