package handlers

import (
	"net/http"
	"spy-cat-agency/internal/cats/dtos"

	"github.com/gin-gonic/gin"
)

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
