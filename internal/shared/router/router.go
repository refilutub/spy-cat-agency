package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"spy-cat-agency/internal/cats/interfaces"
	"spy-cat-agency/internal/cats/interfaces/handlers"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "spy-cat-agency/docs"
)

func SetUpRouter(spyCatshandler *handlers.SpyCatHandler) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(cors.New(cors.Config{
		AllowAllOrigins: true,
		AllowMethods:    []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:   []string{"Content-Length"},
	}))

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := r.Group("/api")

	interfaces.SetUpCatRouter(api, spyCatshandler)

	return r
}
