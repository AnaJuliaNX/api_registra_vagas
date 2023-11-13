package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cria uma nova vaga
func CreateOpening(ctx *gin.Context) {
	request := CreateOpeningRequest{}

	ctx.BindJSON(&request)

	//Faz as validações que escrevi
	err1 := request.Validate()
	if err1 != nil {
		logger.Errorf("erro de validação: %v", err1.Error())
		enviarMensagem(ctx, http.StatusBadRequest, err1.Error())
		return
	}

	err := db.Create(&request).Error
	if err != nil {
		logger.Errorf("erro ao criar opening: %v", err.Error())
		return
	}
}
