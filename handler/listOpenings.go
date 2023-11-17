package handler

import (
	"api_registraVagas/schemas"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Mostra uma vaga
func ListOpenings(ctx *gin.Context) {
	openings := []schemas.Openings{}
	err := db.Find(&openings).Error
	if err != nil {
		errorMessage(ctx, http.StatusInternalServerError, "erro ao listar as vagas")
		return
	}
	successMessage(ctx, "list-openings", openings)
}
