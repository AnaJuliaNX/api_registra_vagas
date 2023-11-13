package handler

import (
	"api_registraVagas/config"

	"gorm.io/gorm"
)

// variavel aplicada somente aos arquivos presentes no pkg handler
var (
	logger *config.Logger
	db     *gorm.DB
)

// Quando usar essa função vai automaticamente criar a database dele
func InitializeHandler() {
	//"handler" é o identificador
	logger = config.GetLogger("handler")
	db = config.GetSQLite()
}
