package interfaces

import (
	"spy-cat-agency/internal/missions/interfaces/handlers"

	"github.com/gin-gonic/gin"
)

func SetUpMissionsRouter(r *gin.RouterGroup, h *handlers.MissionsHandler) {
	missions := r.Group("/missions")
	{
		missions.POST("/", h.CreateMission)
		missions.GET("/", h.ListMissions)
		missions.GET("/:id", h.GetMission)
		missions.PUT("/:id", h.UpdateMission)
		missions.DELETE("/:id", h.DeleteMission)
		missions.POST("/:id/assign-cat", h.AssignCatToMission)
		missions.POST("/:id/targets", h.AddTargetToMission) // in return does not return target ID
	}

	targets := r.Group("/targets")
	{
		targets.PUT("/:id", h.UpdateTarget)
		targets.DELETE("/:id", h.DeleteTarget)
	}
}
