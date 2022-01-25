package config

import (
	"fmt"
	"log"
	"os"
	"todo/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// postgresql db connection

func Init() {
	er := godotenv.Load()
	if er != nil {
		log.Fatal("Error Loading .env file")
	}
	dbURL := os.Getenv("DB_URL")
	var err error
	DB, err = gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(DB)
	DB.AutoMigrate(&models.Todo{})
}
