package router

import (
	"api_registraVagas/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize handler
	handler.InitializeHandler()

	//Para agrupar as rotas
	v1 := router.Group("/api")
	{
		v1.POST("/opening/create", handler.CreateOpening)
		v1.GET("/opening/show", handler.ShowOpening)
		v1.GET("/openings/list", handler.ListOpenings)
		v1.PUT("/opening/update", handler.UpdateOpening)
		v1.DELETE("/opening/delete", handler.DeleteOpening)
	}
}
