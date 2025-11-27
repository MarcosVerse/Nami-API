package user

import (
	"net/http"

	"github.com/MarcosVerse/nami/internal/config"
	"github.com/MarcosVerse/nami/internal/dto/user"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/MarcosVerse/nami/internal/utils"
	"github.com/gin-gonic/gin"
)

// CriarUsuario godoc
// @Summary Cria um novo usuário
// @Description Adiciona um usuário ao sistema
// @Tags Users
// @Accept json
// @Produce json
// @Param user body dto.CreateDTO true "Dados de criação de usuário"
// @Success 201 {object} dto.ResponseDTO
// @Failure 400 {object} dto.ResponseDTO "Dados inválidos"
// @Failure 500 {object} dto.ResponseDTO "Erro ao criar usuário"
// @Router /usuarios/ [post]
func CreateUser(c *gin.Context) {
	var input dto.CreateDTO

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseDTO{Message: config.BadRequest})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Senha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseDTO{Message: config.InternalServerError})
		return
	}

	usuario := models.User{
		Name:     input.Nome,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := repository.DB.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseDTO{Message: config.InternalServerError})
		return
	}

	c.JSON(http.StatusCreated, usuario)
}
