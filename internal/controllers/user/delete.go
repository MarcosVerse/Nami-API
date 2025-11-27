package user

import (
	"net/http"
	"strconv"

	"github.com/MarcosVerse/nami/internal/config"
	dto "github.com/MarcosVerse/nami/internal/dto/user"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// DeletarUsuario godoc
// @Summary Deleta um usuário
// @Description Remove um usuário pelo ID
// @Tags Users
// @Produce json
// @Param id path int true "ID do Usuário"
// @Success 200 {object} dto.ResponseDTO "Usuário deletado"
// @Failure 400 {object} dto.ResponseDTO
// @Failure 404 {object} dto.ResponseDTO "Usuário não encontrado"
// @Failure 500 {object} dto.ResponseDTO
// @Router /usuarios/{id} [delete]
func DeleteUser(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseDTO{Message: config.InvalidUserID})
		return
	}

	result := repository.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseDTO{Message: config.InternalServerError})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, dto.ResponseDTO{Message: config.UserNotFound})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseDTO{Message: config.UserDeleted})
}
