package candidaturas

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	candidatura "api_registraVagas/schemas/candidaturas"
)

func CandidaturaCreate(ctx *gin.Context) {
	request := CreateCandidaturaRequest{}

	ctx.BindJSON(&request)

	//Faço as minhas validações (arquivo curriculoREquest)
	err := request.Validade()
	if err != nil {
		handler.Logger.Errorf("erro de validação: %v", err.Error())
		handler.ErrorMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	candidatura := candidatura.Candidaturas{
		Nome:              request.Nome,
		Idade:             request.Idade,
		Nivel:             request.Nivel,
		Linkedin:          request.Linkedin,
		LinguaEstrangeira: *request.LinguaEstrangeira,
	}

	err = handler.DB.Create(&candidatura).Error
	if err != nil {
		handler.Logger.Errorf("erro ao criar a candidatura %v", err.Error())
		handler.ErrorMessage(ctx, http.StatusInternalServerError, "erro ao criar a candidatura no banco de dados")
		return
	}
	handler.SuccessMessage(ctx, "create-Candidatura", candidatura)
}
