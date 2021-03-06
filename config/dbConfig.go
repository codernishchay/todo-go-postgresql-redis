package config

import (
	"log"
	"os"
	"todo/models"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

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
	DB.AutoMigrate(&models.Todo{})

}

// https://talks.golang.org/2013/bestpractices.slide#9
