package router

import (
	"api_registraVagas/handler"

	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize handler
	handler.InitializeHandler()

	//Para agrupar as rotas
	v1 := router.Group("/v1")
	{
		v1.POST("/vagas/create", handler.VagaCreate)
		v1.GET("/vagas/show", handler.VagaShow)
		v1.GET("/vagas/list", handler.VagaList)
		v1.PUT("/vagas/update", handler.VagaUpdate)
		v1.DELETE("/vagas/delete", handler.VagaDelete)

		v1.POST("/curriculos/create", handler.CurriculoCreate)
	}
}
