package handler

import "fmt"

func naoPodesernulo(name, tipo string) error {
	return fmt.Errorf("campo: %s do tipo %s é obrigatório", name, tipo)
}

type CreateVagaRequest struct {
	Nivel       string  `json:"nivel"`
	Empresa     string  `json:"empresa"`
	Localizacao string  `json:"localizacao"`
	Presencial  *bool   `json:"presencial"`
	Link        string  `json:"link"`
	Salario     float64 `json:"salario"`
}

func (r *CreateVagaRequest) Validate() error {

	if r.Nivel == "" {
		return naoPodesernulo("nivel", "string")
	}
	if r.Empresa == "" {
		return naoPodesernulo("empresa", "string")
	}
	if r.Localizacao == "" {
		return naoPodesernulo("localizacao", "string")
	}
	if r.Link == "" {
		return naoPodesernulo("link", "string")
	}
	if r.Presencial == nil {
		return naoPodesernulo("presencial", "bool")
	}
	if r.Salario <= 0 {
		return naoPodesernulo("salario", "int64")
	}
	return nil
}

// Validação para o update
type UpdateVagaRequest struct {
	Nivel       string `json:"nivel"`
	Empresa     string `json:"empresa"`
	Localizacao string `json:"localizacao"`
	Presencial  *bool  `json:"presencial"`
	Link        string `json:"link"`
	Salario     int64  `json:"salario"`
}

func (r *UpdateVagaRequest) Validate() error {
	if r.Nivel != "" || r.Empresa != "" || r.Localizacao != "" || r.Presencial != nil || r.Link != "" || r.Salario > 0 {
		return nil
	}
	return fmt.Errorf("Pelo menos um campo válido deve ser fornecido")
}
