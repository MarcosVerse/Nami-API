package dto

type CreateUserDTO struct {
	Nome  string `json:"nome"  binding:"required,min=3"`
	Email string `json:"email" binding:"required,email"`
	Senha string `json:"senha" binding:"required,min=6"`
}
