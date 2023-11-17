package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Código para exibir a mensagem de erro
func errorMessage(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":        msg,
		"codigo de erro": code,
	})
}

// Código para exibir a mensagem de sucesso no fim do código
func successMessage(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operação %s bem sucedida", op),
		"data":    data,
	})
}
