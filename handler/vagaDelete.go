package handler

import (
	"api_registraVagas/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Deleta uma vaga
func VagaDelete(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		errorMessage(ctx, http.StatusBadRequest, naoPodesernulo("id", "queryParameter").Error())
		return
	}
	opening := schemas.Vagas{}
	//busca os dados
	err := db.First(&opening, id).Error
	if err != nil {
		errorMessage(ctx, http.StatusNotFound, fmt.Sprintf("Vaga com o id %s não encontrado", id))
		return
	}

	err = db.Delete(&opening).Error
	if err != nil {
		errorMessage(ctx, http.StatusInternalServerError, fmt.Sprintf("Erro ao tentar excluir a vaga do id: %s", id))
		return
	}

	successMessage(ctx, "delete-opening", opening)
}
