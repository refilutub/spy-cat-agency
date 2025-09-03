package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

// DeleteSpyCatHandler godoc
// @Summary Delete a spy cat
// @Description Remove a spy cat from the system by ID
// @Tags Cats
// @Accept json
// @Produce json
// @Param spycatId path string true "Spy Cat ID" format(uuid)
// @Success 200 {object} map[string]interface{} "Spy cat deleted successfully"
// @Failure 400 {object} map[string]interface{} "Bad request - invalid ID format or service error"
// @Failure 404 {object} map[string]interface{} "Spy cat not found"
// @Failure 500 {object} map[string]interface{} "Internal server error"
// @Router /spycats/{spycatId} [delete]
// @Security BearerAuth
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
