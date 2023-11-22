package handler

import (
	"api_registraVagas/config"

	"gorm.io/gorm"
)

var (
	Logger *config.Logger
	DB     *gorm.DB
)

func InitializeHandler() {
	Logger = config.GetLogger("handler")
	DB = config.GetSQLite()
}
