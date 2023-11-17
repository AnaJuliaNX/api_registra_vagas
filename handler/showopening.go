package handler

import (
	"api_registraVagas/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Mostra uma vaga
func ShowOpening(ctx *gin.Context) {
	//busca o ID na query (encontrado nos quer)
	id := ctx.Query("id")

	if id == "" {
		errorMessage(ctx, http.StatusBadRequest, naoPodesernulo("id", "queryParameter").Error())
		return
	}

	opening := schemas.Openings{}
	err := db.First(&opening, id).Error
	if err != nil {
		errorMessage(ctx, http.StatusNotFound, "vaga não encontrada")
		return
	}

	successMessage(ctx, "show-opening", opening)
}
