package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	"api_registraVagas/schemas"
)

// Mostra uma vaga
func VagaList(ctx *gin.Context) {
	vagas := []schemas.Vagas{}
	err := handler.DB.Find(&vagas).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusInternalServerError, "erro ao listar as vagas")
		return
	}
	handler.SuccessMessage(ctx, "list-vagas", vagas)
}
