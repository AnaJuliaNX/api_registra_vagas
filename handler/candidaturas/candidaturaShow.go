package candidaturas

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	candidatura "api_registraVagas/schemas/candidaturas"
)

func CandidaturaShow(ctx *gin.Context) {
	id := ctx.Query("id")
	if id == "" {
		handler.ErrorMessage(ctx, http.StatusBadRequest, emBrancoNao("id", "queryParameter").Error())
		return
	}

	candidatura := candidatura.Candidaturas{}
	err := handler.DB.First(&candidatura, id).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusNotFound, "candidatura não encontrada")
		return
	}

	handler.SuccessMessage(ctx, "show-candidatura", candidatura)
}
