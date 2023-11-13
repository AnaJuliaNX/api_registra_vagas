package config

import (
	"fmt"

	"gorm.io/gorm"
)

// Local
var (
	db     *gorm.DB
	logger *Logger
)

func Init() error {
	var err error

	//Inicializando o SQLite
	db, err = DatabaseInitialize()
	if err != nil {
		return fmt.Errorf("Erro ao inicializar a database: %v", err)
	}
	return nil
}

// Retorna a database
func GetSQLite() *gorm.DB {
	return db
}

func GetLogger(prefix string) *Logger {
	logger = NewLogger(prefix)
	return logger
}
