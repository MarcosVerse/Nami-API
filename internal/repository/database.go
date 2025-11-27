package repository

import (
	"fmt"
	"log"
	"os"

	"github.com/MarcosVerse/nami/internal/models"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Aviso: não foi possível carregar o .env, usando variáveis de ambiente do sistema")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao PostgreSQL")
	}

	DB = db
	fmt.Println("Conectado ao PostgreSQL com sucesso!")

	err = db.AutoMigrate(
		&models.User{},
		&models.Transaction{},
		&models.Goal{},
		&models.MonthlySummary{},
	)
	if err != nil {
		panic("Falha ao migrar tabelas")
	}

	fmt.Println("Tabelas migradas com sucesso!")
}
