package goal

import (
	"net/http"

	"github.com/MarcosVerse/nami/internal/dto/goal"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// UpdateGoal godoc
// @Summary Atualiza uma meta
// @Description Atualiza dados da meta ou adiciona progresso
// @Tags Goals
// @Accept json
// @Produce json
// @Param id path int true "ID da meta"
// @Param goal body goal.UpdateGoalDTO true "Dados de atualização da meta"
// @Success 200 {object} models.Goal
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /goals/{id} [put]
func UpdateGoal(c *gin.Context) {
	id := c.Param("id")

	var body goal.UpdateGoalDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.Goal
	if err := repository.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "meta não encontrada"})
		return
	}

	if body.Title != nil {
		existing.Title = *body.Title
	}

	if body.TargetValue != nil {
		if *body.TargetValue < existing.CurrentValue {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "target não pode ser menor que o progresso atual",
			})
			return
		}
		existing.TargetValue = *body.TargetValue
	}

	if err := repository.DB.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar meta"})
		return
	}

	c.JSON(http.StatusOK, existing)
}
