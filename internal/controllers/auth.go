package controllers

import (
	"net/http"
	"time"

	"github.com/MarcosVerse/nami/internal/config"
	"github.com/MarcosVerse/nami/internal/dto"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/MarcosVerse/nami/internal/utils"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var input dto.LoginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, dto.LoginResponse{Message: "Dados inválidos"})
		return
	}

	var usuario models.User
	if err := repository.DB.Where("email = ?", input.Email).First(&usuario).Error; err != nil {
		c.JSON(http.StatusUnauthorized, dto.LoginResponse{Message: "Email ou senha incorretos"})
		return
	}

	if err := utils.CheckPassword(usuario.Password, input.Senha); err != nil {
		c.JSON(http.StatusUnauthorized, dto.LoginResponse{Message: "Email ou senha incorretos"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  usuario.ID,
		"exp": time.Now().Add(72 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(config.JWTSecret)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.LoginResponse{Message: "Erro ao gerar token"})
		return
	}

	c.JSON(http.StatusOK, dto.LoginResponse{
		Message: "Login realizado com sucesso",
		Token:   tokenString,
	})
}
