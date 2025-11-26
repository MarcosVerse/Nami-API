package dto

type UpdateUsuarioInput struct {
	Nome  string `json:"nome"  binding:"omitempty,min=3"`
	Email string `json:"email" binding:"omitempty,email"`
	Senha string `json:"senha" binding:"omitempty,min=6"`
}
