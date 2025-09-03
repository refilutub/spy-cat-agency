package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
)

func (h *SpyCatHandler) DeleteSpyCatHandler(c *gin.Context) {
	spycatIdStr := c.Param("spycatId")
	spycatId, err := uuid.Parse(spycatIdStr)
	if err != nil {
		c.JSON(400, gin.H{"error": "invalid Id"})
		return
	}
	err = h.Service.DeleteSpyCat(spycatId)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "spycat not found"})
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}

	c.Status(http.StatusNoContent)
}
