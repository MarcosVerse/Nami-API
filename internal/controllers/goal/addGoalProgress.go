package goal

import (
	"net/http"

	"github.com/MarcosVerse/nami/internal/dto/goal"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// AddGoalProgress godoc
// @Summary Adiciona progresso à meta
// @Description Soma um valor ao progresso atual da meta
// @Tags Goals
// @Accept json
// @Produce json
// @Param id path int true "ID da meta"
// @Param body body goal.AddProgressDTO true "Valor a adicionar"
// @Success 200 {object} models.Goal
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{}
// @Router /goals/{id}/progress [post]
func AddGoalProgress(c *gin.Context) {
	id := c.Param("id")

	var body goal.AddProgressDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var existing models.Goal
	if err := repository.DB.First(&existing, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "meta não encontrada"})
		return
	}

	newValue := existing.CurrentValue + body.Amount

	if newValue > existing.TargetValue {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "valor excede o objetivo da meta",
		})
		return
	}

	existing.CurrentValue = newValue

	if err := repository.DB.Save(&existing).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar progresso"})
		return
	}

	c.JSON(http.StatusOK, existing)
}
