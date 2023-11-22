package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	vagas "api_registraVagas/schemas/vagas"
)

// @Summary Listar vagas
// @Description código para listar as vagas
// @Tags vaga
// @Accept json
// @Produce json
// @Success 200 {object} CreateVagaResponse
// @Failure 500 {object} ErrorResponse
// @Router /vagas/list [get]
// Mostra uma vaga
func VagaList(ctx *gin.Context) {
	vagas := []vagas.Vagas{}
	err := handler.DB.Find(&vagas).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusInternalServerError, "erro ao listar as vagas")
		return
	}
	handler.SuccessMessage(ctx, "list-vagas", vagas)
}
