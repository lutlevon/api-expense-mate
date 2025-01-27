package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error al cargar .env: %v", err)
	}
}

func ConnectDB() {
	dns := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{})

	if err != nil {
		log.Fatal("Error trying to connect to Data Base")
	}

	DB = db
}
