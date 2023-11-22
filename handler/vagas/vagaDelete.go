package handler

import (
	"github.com/gin-gonic/gin"

	"fmt"
	"net/http"

	handler "api_registraVagas/handler"
	"api_registraVagas/schemas"
)

// Deleta uma vaga
func VagaDelete(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		handler.ErrorMessage(ctx, http.StatusBadRequest, naoPodesernulo("id", "queryParameter").Error())
		return
	}
	vaga := schemas.Vagas{}
	//busca os dados
	err := handler.DB.First(&vaga, id).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusNotFound, fmt.Sprintf("Vaga com o id %s não encontrado", id))
		return
	}

	err = handler.DB.Delete(&vaga).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusInternalServerError, fmt.Sprintf("Erro ao tentar excluir a vaga do id: %s", id))
		return
	}

	handler.SuccessMessage(ctx, "delete-vaga", vaga)
}
