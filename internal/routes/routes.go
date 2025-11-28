package routes

import (
	"github.com/MarcosVerse/nami/internal/controllers/auth"
	"github.com/MarcosVerse/nami/internal/controllers/transaction"
	"github.com/MarcosVerse/nami/internal/controllers/user"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine) {
	// rota de teste
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	// Rotas de usuários
	usuarios := r.Group("/usuarios")
	{
		usuarios.POST("/", user.CreateUser)
		usuarios.PUT("/:id", user.UpdateUser)
		usuarios.DELETE("/:id", user.DeleteUser)
	}

	// Login
	r.POST("/login", auth.Login)

	// Rotas de transações
	transactions := r.Group("/transactions")
	{
		transactions.POST("/", transaction.CreateTransaction)
		transactions.GET("/", transaction.GetTransactionsByMonth)
		transactions.PUT("/:id", transaction.UpdateTransaction)
		transactions.DELETE("/:id", transaction.DeleteTransaction)
	}

	// Rotas de categorias
	// categories := r.Group("/categories")
	// {
	// 	categories.POST("/", category.CreateCategory)
	// 	categories.GET("/", category.GetUserCategories)
	// 	categories.PUT("/:id", category.UpdateCategory)
	// 	categories.DELETE("/:id", category.DeleteCategory)
	// }

}
