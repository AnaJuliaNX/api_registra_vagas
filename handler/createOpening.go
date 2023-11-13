package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Cria uma nova vaga
func CreateOpening(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "POST opening",
	})
}
