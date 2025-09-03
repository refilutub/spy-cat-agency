package interfaces

import (
	"github.com/gin-gonic/gin"
	"spy-cat-agency/internal/missions/interfaces/handlers"
)

func SetUpMissionsRouter(r *gin.RouterGroup, h *handlers.MissionsHandler) {
	missions := r.Group("/missions")
	{
		missions.POST("/", h.CreateMission)
	}
}
