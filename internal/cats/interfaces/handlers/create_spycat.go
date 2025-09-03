package handlers

import (
	"net/http"
	"spy-cat-agency/internal/cats/dtos"

	"github.com/gin-gonic/gin"
)

// CreateSpyCatsHandler godoc
// @Summary Create a new spy cat
// @Description Create a new spy cat in the system with the specified attributes
// @Tags Cats
// @Accept json
// @Produce json
// @Param cat body dtos.SpyCatRequest true "Spy Cat information"
// @Success 201 {object} dtos.SpyCatSingleResponseDTO "Spy cat created successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - validation error or service error"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /spycats [post]
// @Security BearerAuth
func (h *SpyCatHandler) CreateSpyCatsHandler(c *gin.Context) {
	var req dtos.SpyCatRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	resp, err := h.Service.CreateSpyCat(req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, resp)
}
