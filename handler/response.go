package handler

import "github.com/gin-gonic/gin"

func enviarMensagem(ctx *gin.Context, code int, msg string) {
	ctx.JSON(code, gin.H{
		"message":        msg,
		"codigo de erro": code,
	})
}
