package handler

import (
	"api_registraVagas/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Atualiza/edita uma vaga
func UpdateOpening(ctx *gin.Context) {
	request := UpdateOpeningRequest{}

	//Populo o meu request
	ctx.BindJSON(&request)

	//Faço as validações necessárias (estão no arquivo request)
	err := request.Validate()
	if err != nil {
		logger.Errorf("Erro de validação: %v", err.Error())
		errorMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	//Confere se o id não veio vazio
	id := ctx.Query("id")
	if id == "" {
		errorMessage(ctx, http.StatusBadRequest, naoPodesernulo("id", "queryPArameter").Error())
		return
	}

	//populo o opening com as vagas
	opening := schemas.Openings{}
	//Procura o primeiro registro e ordena por chave primária
	err = db.First(&opening, id).Error
	if err != nil {
		errorMessage(ctx, http.StatusNotFound, "opening não encontrada")
		return
	}

	//Faz a atualização da vaga conferindo o preenchimento dos campos
	if request.Role != "" {
		opening.Role = request.Role
	}
	if request.Company != "" {
		opening.Company = request.Company
	}
	if request.Location != "" {
		opening.Location = request.Location
	}
	if request.Remote != nil {
		opening.Remote = *request.Remote
	}
	if request.Link != "" {
		opening.Link = request.Link
	}
	if request.Salary > 0 {
		opening.Salary = request.Salary
	}

	//Salvar a vaga atualizada
	err = db.Save(&opening).Error
	if err != nil {
		logger.Errorf("erro ao atualizar a vaga %v", err.Error())
		errorMessage(ctx, http.StatusInternalServerError, "erro ao atualizar vaga")
	}
	successMessage(ctx, "update-opening", opening)
}
