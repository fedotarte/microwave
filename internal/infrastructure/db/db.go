package db

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"microwave-service/internal/domain"
	"os"
	"time"
)

func InitDB() *gorm.DB {
	dsn := os.Getenv("DB_DSN")
	var db *gorm.DB
	var err error

	for i := 0; i < 10; i++ {
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			break
		}
		log.Printf("failed to connect to database, retrying... (%d/10)\n", i+1)
		time.Sleep(2 * time.Second)
	}

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	db.AutoMigrate(&domain.Microwave{}, &domain.CommandHistory{})

	return db
}
