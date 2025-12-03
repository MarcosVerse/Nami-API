package transaction

import (
	"net/http"

	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// DeleteTransaction godoc
// @Summary Deleta uma transação
// @Description Remove uma transação pelo ID
// @Tags Transactions
// @Produce json
// @Param id path int true "ID da Transação"
// @Success 200 {object} map[string]interface{} "Transação deletada com sucesso"
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{} "Transação não encontrada"
// @Failure 500 {object} map[string]interface{} "Erro interno"
// @Router /transactions/{id} [delete]
func DeleteTransaction(c *gin.Context) {
	id := c.Param("id")

	var tx models.Transaction
	if err := repository.DB.First(&tx, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transação não encontrada"})
		return
	}

	if err := repository.DB.Delete(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao deletar transação"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Transação deletada com sucesso",
	})
}
