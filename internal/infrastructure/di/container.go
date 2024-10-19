package di

import (
	"gorm.io/gorm"
	"microwave-service/internal/app/services"
	"microwave-service/internal/repository"
)

type Container struct {
	MicrowaveService *services.MicrowaveService
}

func NewContainer(db *gorm.DB) *Container {
	microwaveRepo := repository.NewMicrowaveRepository(db)
	microwaveService := services.NewMicrowaveService(microwaveRepo)

	return &Container{
		MicrowaveService: microwaveService,
	}
}
