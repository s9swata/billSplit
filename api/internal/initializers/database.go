package initializers

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	databaseUrl := os.Getenv("POSTGRES_URL")
	if databaseUrl == "" {
		log.Fatal("POSTGRES_URL is not set in .env file")
	}
	var err error
	DB, err = gorm.Open(postgres.Open(databaseUrl), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Connected to the database successfully")
}
