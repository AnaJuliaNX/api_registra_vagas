package main

import (
	"api_registraVagas/config"
	"api_registraVagas/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("erro na inicialçização da config: %v", err)
		return
	}

	//Initialize Router
	router.Inicializacao()
}
