package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	vagas "api_registraVagas/schemas/vagas"
)

// @BasePath /vagas/v1

// @Summary Update vaga
// @Description código para deletar uma vaga
// @Tags vaga
// @Accept json
// @Produce json
// @Param ID path string true "label:ID|type:int64|required"
// @Success 200 {object} CreateVagaResponse
// @Failure 400 {object} ErrorResponse
// @Failure 404 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /vagas/update [put]
// Atualiza/edita uma vaga
func VagaUpdate(ctx *gin.Context) {
	request := UpdateVagaRequest{}

	//Populo o meu request
	ctx.BindJSON(&request)

	//Faço as validações necessárias (estão no arquivo request)
	err := request.Validate()
	if err != nil {
		handler.Logger.Errorf("Erro de validação: %v", err.Error())
		handler.ErrorMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	//Confere se o id não veio vazio
	id := ctx.Query("id")
	if id == "" {
		handler.ErrorMessage(ctx, http.StatusBadRequest, naoPodesernulo("id", "queryPArameter").Error())
		return
	}

	//populo vaga com as vagas do banco de dados
	vaga := vagas.Vagas{}
	//Procura o primeiro registro e ordena por chave primária
	err = handler.DB.First(&vaga, id).Error
	if err != nil {
		handler.ErrorMessage(ctx, http.StatusNotFound, "vaga não encontrada")
		return
	}

	//Faz a atualização da vaga conferindo o preenchimento dos campos
	if request.Nivel != "" {
		vaga.Nivel = request.Nivel
	}
	if request.Empresa != "" {
		vaga.Empresa = request.Empresa
	}
	if request.Localizacao != "" {
		vaga.Localizacao = request.Localizacao
	}
	if request.Presencial != nil {
		vaga.Presencial = *request.Presencial
	}
	if request.Link != "" {
		vaga.Link = request.Link
	}
	if request.Salario > 0 {
		vaga.Salario = request.Salario
	}

	//Salvar a vaga atualizada
	err = handler.DB.Save(&vaga).Error
	if err != nil {
		handler.Logger.Errorf("erro ao atualizar os dados da vaga %v", err.Error())
		handler.ErrorMessage(ctx, http.StatusInternalServerError, "erro ao atualizar os dados da vaga")
	}
	handler.SuccessMessage(ctx, "update-vaga", vaga)
}
