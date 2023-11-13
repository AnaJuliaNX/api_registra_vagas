package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Deleta uma vaga
func DeleteOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "DELETE opening",
	})
}
