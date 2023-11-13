package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Mostra uma vaga
func ShowOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "GET opening",
	})
}
