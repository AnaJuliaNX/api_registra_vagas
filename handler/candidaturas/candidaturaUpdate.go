package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	"api_registraVagas/schemas"
)

func CandidaturaUpdate(ctx *gin.Context) {
	request := UpdateCurriculoRequest{}

	ctx.BindJSON(&request)

	err := request.Validade()
	if err != nil {
		handler.Logger.Errorf("Erro de validação: %v", err.Error())
		handler.ErrorMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	id := ctx.Query("id")
	if id == "" {
		handler.ErrorMessage(ctx, http.StatusBadRequest, emBrancoNao("id", "queryParameter").Error())
		return
	}

	candidatura := schemas.Candidaturas{}
	err = handler.DB.First(&candidatura, id).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusNotFound, "candidatura não encontrada")
		return
	}

	if request.Nome != "" {
		candidatura.Nome = request.Nome
	}
	if request.Idade > 18 {
		candidatura.Idade = request.Idade
	}
	if request.Nivel != "" {
		candidatura.Nivel = request.Nivel
	}
	if request.Linkedin != "" {
		candidatura.Linkedin = request.Linkedin
	}
	if request.LinguaEstrangeira != nil {
		candidatura.LinguaEstrangeira = *request.LinguaEstrangeira
	}

	err = handler.DB.Save(&candidatura).Error
	if err != nil {
		handler.Logger.Errorf("Erro ao atualizar os dados da candidatura: %v", err.Error())
		handler.ErrorMessage(ctx, http.StatusInternalServerError, "Erro ao atulizar os dados da candidatura")
		return
	}
	handler.SuccessMessage(ctx, "update-candidatura", candidatura)
}
