package handler

import (
	"api_registraVagas/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CurriculoCreate(ctx *gin.Context) {
	request := CreateCurriculoRequest{}

	ctx.BindJSON(&request)

	//Faço as minhas validações (arquivo curriculoREquest)
	err := request.Validade()
	if err != nil {
		logger.Errorf("erro de validação: %v", err.Error())
		errorMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	curriculo := schemas.Curriculos{
		Name:     request.Name,
		Age:      request.Age,
		Level:    request.Level,
		Linkedin: request.Linkedin,
		English:  *request.English,
	}

	err = db.Create(&curriculo).Error
	if err != nil {
		logger.Errorf("erro ao criar a candidatura %v", err.Error())
		errorMessage(ctx, http.StatusInternalServerError, "erro ao criar o curriculo no banco de dados")
		return
	}
	successMessage(ctx, "create-Curriculo", curriculo)
}
