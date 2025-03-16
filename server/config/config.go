package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	WebAddress string
}

func LoadConfig() Config {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORLD"),
		DBName:     os.Getenv("DB_NAME"),
		WebAddress: os.Getenv("WEB_ADDRESS"),
	}

	fmt.Println("Database config loaded: ", cfg.DBHost+":"+cfg.DBPort)

	return cfg
}
