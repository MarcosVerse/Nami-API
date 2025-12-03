package transaction

import (
	"net/http"

	"github.com/MarcosVerse/nami/internal/dto/transaction"
	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// CreateTransaction godoc
// @Summary Cria uma nova transação
// @Description Adiciona uma transação ao usuário
// @Tags Transactions
// @Accept json
// @Produce json
// @Param transaction body transaction.CreateTransactionDTO true "Dados de criação de transação"
// @Success 201 {object} models.Transaction
// @Failure 400 {object} map[string]interface{}
// @Failure 500 {object} map[string]interface{}
// @Router /transactions/ [post]
func CreateTransaction(c *gin.Context) {
	var body transaction.CreateTransactionDTO

	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	tran := models.Transaction{
		UserID:   body.UserID,
		Type:     body.Type,
		Category: body.Category,
		Value:    body.Value,
	}

	if err := repository.DB.Create(&tran).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao criar transação"})
		return
	}

	c.JSON(http.StatusCreated, tran)
}
