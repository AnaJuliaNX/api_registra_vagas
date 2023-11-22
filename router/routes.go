package router

import (
	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	handlerCandidaturas "api_registraVagas/handler/candidaturas"
	handlerVagas "api_registraVagas/handler/vagas"

	docs "api_registraVagas/docs"

	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func initializeRoutes(router *gin.Engine) {
	//Initialize handler
	handler.InitializeHandler()

	basePath := "/v1"
	docs.SwaggerInfo.BasePath = basePath
	//Para agrupar as rotas
	v1 := router.Group(basePath)
	{
		v1.POST("/vagas/create", handlerVagas.VagaCreate)
		v1.GET("/vagas/show", handlerVagas.VagaShow)
		v1.GET("/vagas/list", handlerVagas.VagaList)
		v1.PUT("/vagas/update", handlerVagas.VagaUpdate)
		v1.DELETE("/vagas/delete", handlerVagas.VagaDelete)

		v1.POST("/candidaturas/create", handlerCandidaturas.CurriculoCreate)
		v1.GET("/candidaturas/show", handlerCandidaturas.CandidaturaShow)
		v1.GET("/candidaturas/list", handlerCandidaturas.CandidaturaList)
		v1.PUT("/candidaturas/update", handlerCandidaturas.CandidaturaUpdate)
		v1.DELETE("/candidaturas/delete", handlerCandidaturas.CandidaturaDelete)
	}

	//Rota para pode adicionar o Swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
}
