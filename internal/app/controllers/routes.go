package controllers

import (
	"github.com/gin-gonic/gin"
	"microwave-service/internal/app/services"
)

func RegisterRoutes(router *gin.Engine, microwaveService *services.MicrowaveService) {
	microwaveController := NewMicrowaveController(microwaveService)

	router.POST("/microwave/init", microwaveController.InitializeMicrowave)
	router.POST("/microwave/:id/on", microwaveController.TurnOnMicrowave)
	router.POST("/microwave/:id/off", microwaveController.TurnOffMicrowave)
	router.POST("/microwave/:id/open", microwaveController.OpenDoor)
	router.POST("/microwave/:id/close", microwaveController.CloseDoor)
	router.POST("/microwave/:id/power/:level", microwaveController.SetPowerLevel)
	router.GET("/microwave/:id", microwaveController.GetCurrentMicrowave)
}
