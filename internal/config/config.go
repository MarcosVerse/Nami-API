package config

import (
	"log"

	"github.com/spf13/viper"
)

var JWTSecret []byte

func LoadConfig() {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		log.Println("Arquivo .env não encontrado, usando variáveis de ambiente")
	}

	secret := viper.GetString("JWT_SECRET")
	if secret == "" {
		log.Fatal("JWT_SECRET não definido no ambiente")
	}

	JWTSecret = []byte(secret)
}
