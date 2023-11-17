package schemas

import (
	"time"

	"gorm.io/gorm"
)

// Estrutura do banco de dados
type Vagas struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

// Estrutura que retono do banco e transformo em json
type VagasResponse struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt,omitempty"` //se for false ou não tiver o campo ele é omitido
	Role      string    `json:"role"`
	Company   bool      `json:"company"`
	Location  string    `json:"location"`
	Remote    string    `json:"remote"`
	Link      string    `json:"link"`
	Salary    int64     `json:"salary"`
}

type Curriculos struct {
	gorm.Model
	Name     string
	Age      int64
	Level    string
	Linkedin string
	English  bool
}
