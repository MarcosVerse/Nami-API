package controllers

import (
	"net/http"
	"strconv"

	"github.com/MarcosVerse/nami/internal/config"
	"github.com/MarcosVerse/nami/internal/dto"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/MarcosVerse/nami/internal/utils"
	"github.com/gin-gonic/gin"
)

func CriarUsuario(c *gin.Context) {
	var input dto.CreateUsuarioInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseMessage{Message: config.BadRequest})
		return
	}

	hashedPassword, err := utils.HashPassword(input.Senha)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseMessage{Message: config.InternalServerError})
		return
	}

	usuario := models.User{
		Name:     input.Nome,
		Email:    input.Email,
		Password: hashedPassword,
	}

	if err := repository.DB.Create(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseMessage{Message: config.InternalServerError})
		return
	}

	c.JSON(http.StatusCreated, usuario)
}

func AtualizarUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseMessage{Message: config.InvalidUserID})
		return
	}

	var usuario models.User
	if err := repository.DB.First(&usuario, id).Error; err != nil {
		c.JSON(http.StatusNotFound, dto.ResponseMessage{Message: config.UserNotFound})
		return
	}

	var input dto.UpdateUsuarioInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseMessage{Message: config.BadRequest})
		return
	}

	usuario.Name = input.Nome
	usuario.Email = input.Email

	if input.Senha != "" {
		hashed, err := utils.HashPassword(input.Senha)
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.ResponseMessage{Message: config.InternalServerError})
			return
		}
		usuario.Password = hashed
	}

	if err := repository.DB.Save(&usuario).Error; err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseMessage{Message: config.InternalServerError})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseMessage{Message: config.UserUpdated})
}

func DeletarUsuario(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.ResponseMessage{Message: config.InvalidUserID})
		return
	}

	result := repository.DB.Delete(&models.User{}, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseMessage{Message: config.InternalServerError})
		return
	}

	if result.RowsAffected == 0 {
		c.JSON(http.StatusNotFound, dto.ResponseMessage{Message: config.UserNotFound})
		return
	}

	c.JSON(http.StatusOK, dto.ResponseMessage{Message: config.UserDeleted})
}
