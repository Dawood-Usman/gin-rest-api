package config

import (
	"log"

	"github.com/dawood-usman/gorest/models"
)

func MigrateDB() {
	if err := DB.AutoMigrate(&models.Book{}); err != nil {
		log.Println("Database Migration Failed: ", err)
	}
	log.Println("Database Migrated Successfully!")
}
