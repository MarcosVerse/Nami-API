package goal

import (
	"net/http"

	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// DeleteGoal godoc
// @Summary Remove uma meta
// @Description Deleta uma meta pelo ID
// @Tags Goals
// @Accept json
// @Produce json
// @Param id path int true "ID da meta"
// @Success 204
// @Failure 404 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /goals/{id} [delete]
func DeleteGoal(c *gin.Context) {
	id := c.Param("id")

	if err := repository.DB.Delete(&models.Goal{}, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao deletar meta"})
		return
	}

	c.Status(http.StatusNoContent)
}
