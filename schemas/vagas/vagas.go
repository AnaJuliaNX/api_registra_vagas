package schemas

import (
	"time"

	"gorm.io/gorm"
)

// Estrutura do banco de dados
type Vagas struct {
	gorm.Model
	Nivel       string
	Empresa     string
	Localizacao string
	Presencial  bool
	Link        string
	Salario     int64
}

// Estrutura que retono do banco e transformo em json
type VagasResponse struct {
	ID          int64     `json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	DeletedAt   time.Time `json:"deletedAt,omitempty"` //se for false ou não tiver o campo ele é omitido
	Nivel       string    `json:"nivel"`
	Empresa     bool      `json:"empresa"`
	Localizacao string    `json:"localizacao"`
	Presencial  string    `json:"presencial"`
	Link        string    `json:"link"`
	Salario     int64     `json:"salario"`
}
