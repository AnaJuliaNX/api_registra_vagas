package handler

import (
	"api_registraVagas/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cria uma nova vaga
func CreateOpening(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	//Popula o meu request
	ctx.BindJSON(&request)

	//Faz as validações que escrevi
	err1 := request.Validate()
	if err1 != nil {
		logger.Errorf("erro de validação: %v", err1.Error())
		errorMessage(ctx, http.StatusBadRequest, err1.Error())
		return
	}

	opening := schemas.Openings{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   int64(request.Salary),
	}

	//Faço a conexão com o banco
	err := db.Create(&opening).Error
	if err != nil {
		logger.Errorf("erro ao criar a vaga: %v", err.Error())
		errorMessage(ctx, http.StatusInternalServerError, "erro ao criar a vaga no banco de dados")
		return
	}
	successMessage(ctx, "create-Opening", opening)
}
