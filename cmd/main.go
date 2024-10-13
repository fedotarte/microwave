package main

import (
	"microwave/controllers"
	"microwave/models"
	"microwave/services"
)

func main() {
	microwave := models.NewMicrowave()

	microwaveService := services.NewMicrowaveService(microwave)

	commandController := controllers.NewCommandController(microwaveService)

	commandController.Run()
}
