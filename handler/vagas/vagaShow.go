package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	vagas "api_registraVagas/schemas/vagas"
)

// @BasePath /vagas/v1

// @Summary Mostra uma vaga pelo ID
// @Description código para buscar uma vaga pelo ID especificado na rota
// @Tags vaga
// @Accept json
// @Produce json
// @Param ID path string true "label:ID|type:int64|required"
// @Success 200 {object} CreateVagaResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Router /vagas/show [get]
// Mostra uma vaga
func VagaShow(ctx *gin.Context) {
	//busca o ID na query
	id := ctx.Query("id")

	if id == "" {
		handler.ErrorMessage(ctx, http.StatusBadRequest, naoPodesernulo("id", "queryParameter").Error())
		return
	}

	vaga := vagas.Vagas{}
	err := handler.DB.First(&vaga, id).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusNotFound, "vaga não encontrada")
		return
	}

	handler.SuccessMessage(ctx, "show-vaga", vaga)
}
