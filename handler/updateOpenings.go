package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Atualiza/edita uma vaga
func UpdateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "PUT opening",
	})
}
