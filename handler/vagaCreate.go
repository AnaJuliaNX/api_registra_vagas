package handler

import (
	"api_registraVagas/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cria uma nova vaga
func VagaCreate(ctx *gin.Context) {
	request := CreateVagaRequest{}

	//Popula o meu request
	ctx.BindJSON(&request)

	//Faz as validações que escrevi
	err := request.Validate()
	if err != nil {
		logger.Errorf("erro de validação: %v", err.Error())
		errorMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	vaga := schemas.Vagas{
		Role:     request.Role,
		Company:  request.Company,
		Location: request.Location,
		Remote:   *request.Remote,
		Link:     request.Link,
		Salary:   int64(request.Salary),
	}

	//Faço a conexão com o banco
	err = db.Create(&vaga).Error
	if err != nil {
		logger.Errorf("erro ao criar a vaga: %v", err.Error())
		errorMessage(ctx, http.StatusInternalServerError, "erro ao criar a vaga no banco de dados")
		return
	}
	successMessage(ctx, "create-Vaga", vaga)
}
