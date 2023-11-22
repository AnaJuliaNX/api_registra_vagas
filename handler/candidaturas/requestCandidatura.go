package handler

import (
	"fmt"
)

func emBrancoNao(name, tipo string) error {
	return fmt.Errorf("campo: %s do tipo %s é obrigatório", name, tipo)
}

type CreateCandidaturaRequest struct {
	Nome              string `json:"nome"`
	Idade             int64  `json:"idade"`
	Nivel             string `json:"nivel"`
	Linkedin          string `json:"linkedin"`
	LinguaEstrangeira *bool  `json:"lingua_estrangeira"`
}

func (r *CreateCandidaturaRequest) Validade() error {
	if r.Nome == "" {
		return emBrancoNao("nome", "string")
	}
	if r.Idade < 18 {
		return fmt.Errorf("Idade minima não atingida")
	}
	if r.Nivel == "" {
		return emBrancoNao("nivel", "string")
	}
	if r.Linkedin == "" {
		return emBrancoNao("linkedin", "string")
	}
	if r.LinguaEstrangeira == nil {
		return emBrancoNao("lingua_estrangeira", "bool")
	}
	return nil
}

type UpdateCurriculoRequest struct {
	Nome              string `json:"nome"`
	Idade             int64  `json:"idade"`
	Nivel             string `json:"nivel"`
	Linkedin          string `json:"linkedin"`
	LinguaEstrangeira *bool  `json:"lingua_estrangeira"`
}

func (r *UpdateCurriculoRequest) Validade() error {
	if r.Nome != "" || r.Idade > 18 || r.Nivel != "" || r.Linkedin != "" || r.LinguaEstrangeira != nil {
		return nil
	}
	return fmt.Errorf("Pelo menos um campo válido deve ser fornecido")
}
