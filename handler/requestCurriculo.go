package handler

import (
	"fmt"
)

func emBrancoNao(name, tipo string) error {
	return fmt.Errorf("campo: %s do tipo %s é obrigatório", name, tipo)
}

type CreateCurriculoRequest struct {
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Level    string `json:"level"`
	Linkedin string `json:"linkedin"`
	English  *bool  `json:"english"`
}

func (r *CreateCurriculoRequest) Validade() error {
	if r.Name == "" {
		return emBrancoNao("name", "string")
	}
	if r.Age <= 18 {
		return fmt.Errorf("Idade minima não atingida")
	}
	if r.Level == "" {
		return emBrancoNao("level", "string")
	}
	if r.Linkedin == "" {
		return emBrancoNao("linkedin", "string")
	}
	if r.English == nil {
		return emBrancoNao("english", "bool")
	}
	return nil
}

type UpdateCurriculoRequest struct {
	Name     string `json:"name"`
	Age      int64  `json:"age"`
	Level    string `json:"level"`
	Linkedin string `json:"linkedin"`
	English  *bool  `json:"english"`
}

func (r *UpdateCurriculoRequest) Validade() error {
	if r.Name != "" || r.Age > 15 || r.Level != "" || r.Linkedin != "" || r.English != nil {
		return nil
	}
	return fmt.Errorf("Pelo menos um campo válido deve ser fornecido")
}
