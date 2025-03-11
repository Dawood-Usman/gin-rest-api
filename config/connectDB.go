package config

import (
	"fmt"
	"os"

	"github.com/dawood-usman/gorest/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	postgres_dsn := os.Getenv("POSTGRES_DSN")
	DB, err = gorm.Open(postgres.Open(postgres_dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Database Connection Failed: ", err)
	}
	fmt.Println("Database Connected Successfully!")
	DB.AutoMigrate(&models.Book{})

}
