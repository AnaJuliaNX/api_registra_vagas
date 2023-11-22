// +run !staging
//go:build !staging
// +build !staging

package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

const (
	API_PORT          int    = 8080
	DEBUG             bool   = true
	PER_PAGE          int    = 15
	MAIL_FROM_ADDRESS string = "noreply@treeunfe.com.br"
	AWS_S3_BUCKET     string = "fn-api-sandbox"
	JWT_SECRET        string = "nfhvlpc3m4Y3ozje"
	HOST_NFEAPI       string = "https://sandbox-api.freenfeweb.com.br"
	// HOST_NFEAPI       string = "https://api.freenfeweb.com.br"
	// HOST_NFEAPI string = "https://api-br.freenfeweb.com.br"
)

var (
	REDIS_HOST     string = "127.0.0.1"
	REDIS_PORT     string = "6379"
	REDIS_PASSWORD string = ""
	REDIS_DB       int    = 0
	//
	HOST string = "http://localhost:8080"
)

func init() {
	err := godotenv.Load(".env-general")
	if err != nil {
		log.Fatal(err)
	}

	if os.Getenv("REDIS_HOST") != "" {
		REDIS_HOST = os.Getenv("REDIS_HOST")
	}

	if os.Getenv("REDIS_PORT") != "" {
		REDIS_PORT = os.Getenv("REDIS_PORT")
	}

	if os.Getenv("REDIS_PASSWORD") != "" {
		REDIS_PASSWORD = os.Getenv("REDIS_PASSWORD")
	}

	if os.Getenv("REDIS_DB") != "" {
		db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
		if err == nil {
			REDIS_DB = db
		}
	}

	if os.Getenv("HOST") != "" {
		HOST = os.Getenv("HOST")
	}
}
