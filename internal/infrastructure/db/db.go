package db

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"microwave-service/internal/domain"
	"os"
	"time"
)

func InitDB() *gorm.DB {
	//dsn := os.Getenv("DB_DSN"

	dotEnvErr := godotenv.Load()
	if dotEnvErr != nil {
		log.Println("No .env file found, continuing with system environment variables")
	}

	dsn := os.Getenv("DB_DSN")

	fmt.Printf("this is DSN: %s\n", dsn)

	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			log.Println("db connect error")
			break
		}
		log.Printf("failed to connect to database, retrying... (%d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	migrateErr := db.AutoMigrate(&domain.Microwave{}, &domain.CommandHistory{})
	if migrateErr != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
