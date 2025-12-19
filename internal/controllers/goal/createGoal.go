package goal

import (
	"net/http"

	"github.com/MarcosVerse/nami/internal/dto/goal"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// CreateGoal godoc
// @Summary Cria uma nova meta
// @Description Cria uma meta financeira para o usuário
// @Tags Goals
// @Accept json
// @Produce json
// @Param goal body goal.CreateGoalDTO true "Dados de criação da meta"
// @Success 201 {object} models.Goal
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /goals/ [post]
func CreateGoal(c *gin.Context) {
	var body goal.CreateGoalDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newGoal := models.Goal{
		UserID:       body.UserID,
		Title:        body.Title,
		TargetValue:  body.TargetValue,
		CurrentValue: 0,
	}

	if err := repository.DB.Create(&newGoal).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar meta"})
		return
	}

	c.JSON(http.StatusCreated, newGoal)
}
