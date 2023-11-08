package router

import (
	"github.com/gin-gonic/gin"
)

func Inicializacao() {
	//Initialize Router
	router := gin.Default()

	//Initialize Routes/ inicilaização das rotas
	initializeRoutes(router)
	//Run the server
	router.Run(":9090") //O padrão da rota é 8080 mas já está em uso por isso mudei
}
