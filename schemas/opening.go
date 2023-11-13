package schemas

import (
	"time"

	"gorm.io/gorm"
)

// Estrutura do banco de dados
type Openings struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Remote   bool
	Link     string
	Salary   int64
}

// Estrutura que retono do banco e transformo em json
type OpeningREsponse struct {
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
