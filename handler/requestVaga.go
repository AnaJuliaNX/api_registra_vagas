package handler

import "fmt"

func naoPodesernulo(name, tipo string) error {
	return fmt.Errorf("campo: %s do tipo %s é obrigatório", name, tipo)
}

type CreateVagaRequest struct {
	Role     string  `json:"role"`
	Company  string  `json:"company"`
	Location string  `json:"location"`
	Remote   *bool   `json:"remote"`
	Link     string  `json:"link"`
	Salary   float64 `json:"salary"`
}

func (r *CreateVagaRequest) Validate() error {

	if r.Role == "" {
		return naoPodesernulo("role", "string")
	}
	if r.Company == "" {
		return naoPodesernulo("company", "string")
	}
	if r.Location == "" {
		return naoPodesernulo("location", "string")
	}
	if r.Link == "" {
		return naoPodesernulo("Link", "string")
	}
	if r.Remote == nil {
		return naoPodesernulo("remote", "bool")
	}
	if r.Salary <= 0 {
		return naoPodesernulo("salary", "int64")
	}
	return nil
}

// Validação para o update
type UpdateVagaRequest struct {
	Role     string `json:"role"`
	Company  string `json:"company"`
	Location string `json:"location"`
	Remote   *bool  `json:"remote"`
	Link     string `json:"link"`
	Salary   int64  `json:"salary"`
}

func (r *UpdateVagaRequest) Validate() error {
	if r.Role != "" || r.Company != "" || r.Location != "" || r.Remote != nil || r.Link != "" || r.Salary > 0 {
		return nil
	}
	return fmt.Errorf("Pelo menos um campo válido deve ser fornecido")
}