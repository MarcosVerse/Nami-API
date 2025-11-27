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
// @Param user body dto.UpdateDTO true "Dados atualizados"
// @Success 200 {object} dto.ResponseDTO
// @Failure 400 {object} dto.ResponseDTO
// @Failure 404 {object} dto.ResponseDTO "Usuário não encontrado"
// @Failure 500 {object} dto.ResponseDTO
// @Router /usuarios/{id} [put]
func UpdateUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseDTO{Message: config.InvalidUserID})
		return
	}

	var usuario models.User
	if err := repository.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ResponseDTO{Message: config.UserNotFound})
		return
	}

	var input dto.UpdateDTO
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseDTO{Message: config.BadRequest})
		return
	}

	usuario.Name = input.Nome
	usuario.Email = input.Email

	if input.Senha != "" {
		hashed, err := utils.HashPassword(input.Senha)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ResponseDTO{Message: config.InternalServerError})
			return
		}
		usuario.Password = hashed
	}

	if err := repository.DB.Save(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseDTO{Message: config.InternalServerError})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseDTO{Message: config.UserUpdated})
}
