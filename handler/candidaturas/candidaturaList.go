package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	"api_registraVagas/schemas"
)

func CandidaturaList(ctx *gin.Context) {
	candidaturas := []schemas.Candidaturas{}
	err := handler.DB.Find(&candidaturas).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusInternalServerError, "erro ao listar as candidaturas")
		return
	}
	handler.SuccessMessage(ctx, "list-Candidatura", candidaturas)
}
