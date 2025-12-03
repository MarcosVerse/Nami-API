package transaction

import (
	"net/http"
	"time"

	"github.com/MarcosVerse/nami/internal/models"
	"github.com/MarcosVerse/nami/internal/repository"
	"github.com/gin-gonic/gin"
)

// GetTransactionsByMonth godoc
// @Summary Lista transações por mês
// @Description Lista todas as transações de um usuário em um mês específico
// @Tags Transactions
// @Produce json
// @Param user_id query int true "ID do usuário"
// @Param month query int true "Mês (1-12)"
// @Param year query int true "Ano (ex: 2025)"
// @Success 200 {array} models.Transaction
// @Failure 400 {object} map[string]interface{} "Parâmetros inválidos"
// @Failure 500 {object} map[string]interface{}
// @Router /transactions [get]
func GetTransactionsByMonth(c *gin.Context) {
	userID := c.Query("user_id")
	month := c.Query("month")
	year := c.Query("year")

	if userID == "" || month == "" || year == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Missing required parameters"})
		return
	}

	start, _ := time.Parse("2006-01-02", year+"-"+month+"-01")
	end := start.AddDate(0, 1, 0)

	var transactions []models.Transaction

	err := repository.DB.
		Where("user_id = ?", userID).
		Where("created_at >= ? AND created_at < ?", start, end).
		Find(&transactions).Error

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch transactions"})
		return
	}

	c.JSON(http.StatusOK, transactions)
}
