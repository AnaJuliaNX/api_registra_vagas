package router

import (
	"api_registraVagas/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize handler
	handler.InitializeHandler()

	//Para agrupar as rotas
	v1 := router.Group("/api/v1")
	{
		v1.POST("/opening", handler.CreateOpening)
		v1.GET("/opening", handler.ShowOpening)
		v1.PUT("/opening", handler.UpdateOpening)
		v1.DELETE("/opening", handler.DeleteOpening)
		v1.GET("/openings", handler.ListOpenings)
	}
}
