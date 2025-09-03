package interfaces

import (
	"spy-cat-agency/internal/cats/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpCatRouter(r *gin.RouterGroup, h *handlers.SpyCatHandler) {
	spycats := r.Group("/spycats")
	{
		spycats.POST("/", h.CreateSpyCatsHandler)

		spycats.DELETE("/:spycatId", h.DeleteSpyCatHandler)
	}
}
