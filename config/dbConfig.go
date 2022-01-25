package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	dbURL := "postgres://ssutonib:24TvZRAbayu6QSSlkAx7ctVjww83Ff7q@arjuna.db.elephantsql.com/ssutonib"

	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(db)
	// db.AutoMigrate(&models.Todo{})
	// db.AutoMigrate(&models.User{})
}
