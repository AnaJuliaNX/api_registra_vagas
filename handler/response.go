package handler

import (
	"api_registraVagas/schemas"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Código para exibir a mensagem de erro
func ErrorMessage(ctx *gin.Context, code int, msg string) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(code, gin.H{
		"message":        msg,
		"codigo de erro": code, // o mesmo que o ErrorCode
	})
}

// Código para exibir a mensagem de sucesso no fim do código
func SuccessMessage(ctx *gin.Context, op string, data interface{}) {
	ctx.Header("Content-type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"message": fmt.Sprintf("Operação %s bem sucedida", op),
		"data":    data,
	})
}

// Apenas para usar com o swagger (error message)
type ErrorResponse struct {
	Message   string `json:"message"`
	ErrorCode string `json:"errorCode"`
}

// Apenas para usar com o swagger (success message)
type CreateVagaResponse struct {
	Message string                `json:"message"`
	Data    schemas.VagasResponse `json:"data"`
}
