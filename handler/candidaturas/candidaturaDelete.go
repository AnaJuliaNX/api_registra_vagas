package candidaturas

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	candidatura "api_registraVagas/schemas/candidaturas"
)

func CandidaturaDelete(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		handler.ErrorMessage(ctx, http.StatusBadRequest, emBrancoNao("id", "queryParameter").Error())
		return
	}
	candidatura := candidatura.Candidaturas{}
	//Busca os dados
	err := handler.DB.First(&candidatura, id).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusNotFound, fmt.Sprintf("Candidatura com id %s não encontrada", id))
		return
	}

	err = handler.DB.Delete(&candidatura).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusInternalServerError,
			fmt.Sprintf("Erro ao tentar excluir a candidatura do id: %s", id))
		return
	}

	handler.SuccessMessage(ctx, "delete-candidatura", candidatura)
}
