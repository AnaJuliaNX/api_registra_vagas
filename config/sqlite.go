package config

import (
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	candidatura "api_registraVagas/schemas/candidaturas"
	vagas "api_registraVagas/schemas/vagas"
)

func DatabaseInitialize() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/main.db"

	//Criar e conectar na database
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Database inexistente, criando uma...")
	}
	err = os.MkdirAll("./db/", os.ModePerm)
	if err != nil {
		return nil, err
	}
	file, err := os.Create(dbPath)
	if err != nil {
		return nil, err
	}
	file.Close()

	// Create DB and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("sqlite opening error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&vagas.Vagas{}, &candidatura.Candidaturas{})
	if err != nil {
		logger.Errorf("Erro de automigração: %v", err)
		return nil, err
	}
	return db, nil
}
