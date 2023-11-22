package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	"api_registraVagas/schemas"
)

// Mostra uma vaga
func VagaShow(ctx *gin.Context) {
	//busca o ID na query
	id := ctx.Query("id")

	if id == "" {
		handler.ErrorMessage(ctx, http.StatusBadRequest, naoPodesernulo("id", "queryParameter").Error())
		return
	}

	vaga := schemas.Vagas{}
	err := handler.DB.First(&vaga, id).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusNotFound, "vaga não encontrada")
		return
	}

	handler.SuccessMessage(ctx, "show-vaga", vaga)
}
