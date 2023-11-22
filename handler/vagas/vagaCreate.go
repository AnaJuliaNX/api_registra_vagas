package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	handler "api_registraVagas/handler"
	"api_registraVagas/schemas"
)

// @BasePath /vagas/v1

// @Summary Create vagas
// @Description código para a criação de novas vagas
// @Tags vaga
// @Accept json
// @Produce json
// @Parama request body VagaCreate true "Vaga Request Body"
// @Success 200 {object} CreateVagaResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /vagas/create [post]
func VagaCreate(ctx *gin.Context) {
	request := CreateVagaRequest{}

	//Popula o meu request
	ctx.BindJSON(&request)

	//Faz as validações que escrevi
	err := request.Validate()
	if err != nil {
		handler.Logger.Errorf("erro de validação: %v", err.Error())
		handler.ErrorMessage(ctx, http.StatusBadRequest, err.Error())
		return
	}

	vaga := schemas.Vagas{
		Nivel:       request.Nivel,
		Empresa:     request.Empresa,
		Localizacao: request.Localizacao,
		Presencial:  *request.Presencial,
		Link:        request.Link,
		Salario:     int64(request.Salario),
	}

	//Faço a conexão com o banco
	err = handler.DB.Create(&vaga).Error
	if err != nil {
		handler.Logger.Errorf("erro ao criar a vaga: %v", err.Error())
		handler.ErrorMessage(ctx, http.StatusInternalServerError, "erro ao criar a vaga no banco de dados")
		return
	}
	handler.SuccessMessage(ctx, "create-Vaga", vaga)
}
