package goal

import (
	"net/http"

	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// GetGoals godoc
// @Summary Lista metas do usuário
// @Description Retorna todas as metas de um usuário
// @Tags Goals
// @Accept json
// @Produce json
// @Param user_id query int true "ID do usuário"
// @Success 200 {array} models.Goal
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /goals/ [get]
func GetGoals(c *gin.Context) {
	userID := c.Query("user_id")
	if userID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user_id é obrigatório"})
		return
	}

	var goals []models.Goal

	if err := repository.DB.
		Where("user_id = ?", userID).
		Find(&goals).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar metas"})
		return
	}

	c.JSON(http.StatusOK, goals)
}
