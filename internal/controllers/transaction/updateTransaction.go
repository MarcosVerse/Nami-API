package transaction

import (
	"net/http"

	dto "github.com/MarcosVerse/nami/internal/dto/transaction"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// UpdateTransaction godoc
// @Summary Atualiza uma transação
// @Description Atualiza valores da transação pelo ID
// @Tags Transactions
// @Accept json
// @Produce json
// @Param id path int true "ID da transação"
// @Param body body transaction.UpdateTransactionDTO true "Dados atualizados"
// @Success 200 {object} models.Transaction
// @Failure 400 {object} map[string]interface{}
// @Failure 404 {object} map[string]interface{} "Transação não encontrada"
// @Failure 500 {object} map[string]interface{}
// @Router /transactions/{id} [put]
func UpdateTransaction(c *gin.Context) {
	id := c.Param("id")

	var tx models.Transaction
	if err := repository.DB.First(&tx, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Transação não encontrada"})
		return
	}

	var body dto.UpdateTransactionDTO
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	// Atualiza somente campos enviados
	if body.Category != nil {
		tx.Category = *body.Category
	}

	if body.Value != nil {
		tx.Value = *body.Value
	}

	if body.Type != nil {
		tx.Type = *body.Type
	}

	if err := repository.DB.Save(&tx).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao atualizar transação"})
		return
	}

	c.JSON(http.StatusOK, tx)
}
