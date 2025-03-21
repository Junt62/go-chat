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
	DBName     string
	DBUsername string
	DBPassword string
	WebAddress string
}

var defaultConfig = []struct {
	Key   string
	Value string
}{
	{"DB_HOST", "localhost"},
	{"DB_PORT", "3306"},
	{"DB_NAME", "go-chat"},
	{"DB_USERNAME", "root"},
	{"DB_PASSWORD", "123456"},
	{"WEB_ADDRESS", "http://localhost:5173,http://127.0.0.1:5173"},
}

func LoadConfig() Config {
	if _, err := os.Stat(".env"); os.IsNotExist(err) {
		CreateConfig()
		log.Println(".env file no exist, create now")
	} else if err != nil {
		log.Fatal("Error checking .env file:", err)
	}

	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg := Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBName:     os.Getenv("DB_NAME"),
		DBUsername: os.Getenv("DB_USERNAME"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		WebAddress: os.Getenv("WEB_ADDRESS"),
	}

	fmt.Println("Database config loaded:", cfg.DBHost+":"+cfg.DBPort)

	return cfg
}

func CreateConfig() {
	file, err := os.OpenFile(".env", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		log.Fatal("Error opening .env file:", err)
		return
	}
	defer file.Close()

	for _, field := range defaultConfig {
		if _, err := fmt.Fprintf(file, "%s=%s\n", field.Key, field.Value); err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
}
