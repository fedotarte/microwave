package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"microwave-service/internal/app/controllers"
	"microwave-service/internal/infrastructure/db"
	"microwave-service/internal/infrastructure/di"
)

func main() {
	database := db.InitDB()

	container := di.NewContainer(database)

	router := gin.Default()

	controllers.RegisterRoutes(router, container.MicrowaveService)

	routerRunningErr := router.Run()

	if routerRunningErr != nil {
		log.Fatalf("failed to run handller: %v", routerRunningErr)
	}
}
