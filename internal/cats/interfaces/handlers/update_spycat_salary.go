package handlers

import (
	"net/http"
	"spy-cat-agency/internal/cats/dtos"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UpdateSpyCatSalaryHandler godoc
// @Summary Update spy cat salary
// @Description Update the salary of an existing spy cat
// @Tags Cats
// @Accept json
// @Produce json
// @Param spycatId path string true "Spy Cat ID" format(uuid)
// @Param salary body dtos.SalaryUpdateRequest true "New salary information"
// @Success 200 {object} models.Cat "Salary updated successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid ID format, invalid salary, or service error"
// @Failure 404 {object} map[string]interface{} "Spy cat not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /spycats/{spycatId} [patch]
// @Security BearerAuth
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
