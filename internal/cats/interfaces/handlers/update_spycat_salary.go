package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"net/http"
	"spy-cat-agency/internal/cats/dtos"
)

func (h *SpyCatHandler) UpdateSpyCatSalaryHandler(c *gin.Context) {
	idStr := c.Param("spycatId")
	spycatIdParam, err := uuid.Parse(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid spycat id"})
		return
	}

	var salaryUpdate dtos.SalaryUpdateRequest
	if err := c.ShouldBindJSON(&salaryUpdate); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid salary update"})
		return
	}

	spycat, err := h.Service.UpdateSpyCatSalary(spycatIdParam, salaryUpdate)
	if err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.JSON(http.StatusNotFound, gin.H{"error": "spycat not found"})
			return
		default:
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, spycat)
}
