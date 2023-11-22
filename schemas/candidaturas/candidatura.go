package candidaturas

import (
	"time"

	"gorm.io/gorm"
)

type Candidaturas struct {
	gorm.Model
	Nome              string
	Idade             int64
	Nivel             string
	Linkedin          string
	LinguaEstrangeira bool
}

type CandidaturaResponse struct {
	ID                int64     `json:"id"`
	CreatedAt         time.Time `json:"createdAt"`
	UpdatedAt         time.Time `json:"updatedAt"`
	DeletedAt         time.Time `json:"deletedAt,omitempty"`
	Nome              string    `json:"nome"`
	Idade             int64     `json:"idade"`
	Nivel             string    `json:"nivel"`
	Linkedin          string    `json:"linkedin"`
	LinguaEstrangeira string    `json:"lingua_estrangeira"`
}
