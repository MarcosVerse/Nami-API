package user

import (
	"net/http"
	"strconv"

	"github.com/MarcosVerse/nami/internal/config"
	"github.com/MarcosVerse/nami/internal/dto/user"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/MarcosVerse/nami/internal/utils"
	"github.com/gin-gonic/gin"
)

// AtualizarUsuario godoc
// @Summary Atualiza um usuário existente
// @Description Atualiza os dados de um usuário pelo ID
// @Tags Users
// @Accept json
// @Produce json
// @Param id path int true "ID do Usuário"
// @Param user body dto.UpdateUserDTO true "Dados atualizados"
// @Success 200 {object} dto.ResponseUserDTO
// @Failure 400 {object} dto.ResponseUserDTO
// @Failure 404 {object} dto.ResponseUserDTO "Usuário não encontrado"
// @Failure 500 {object} dto.ResponseUserDTO
// @Router /usuarios/{id} [put]
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseUserDTO{Message: config.InvalidUserID})
		return
	}

	var usuario models.User
	if err := repository.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ResponseUserDTO{Message: config.UserNotFound})
		return
	}

	var input dto.UpdateUserDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseUserDTO{Message: config.BadRequest})
		return
	}

	usuario.Name = input.Nome
	usuario.Email = input.Email

	if input.Senha != "" {
		hashed, err := utils.HashPassword(input.Senha)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ResponseUserDTO{Message: config.InternalServerError})
			return
		}
		usuario.Password = hashed
	}

	if err := repository.DB.Save(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseUserDTO{Message: config.InternalServerError})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseUserDTO{Message: config.UserUpdated})
}
