package controllers

import (
	"github.com/gin-gonic/gin"
	"microwave-service/internal/app/services"
	"net/http"
	"strconv"
)

type MicrowaveController struct {
	microwaveService *services.MicrowaveService
}

func NewMicrowaveController(service *services.MicrowaveService) *MicrowaveController {
	return &MicrowaveController{microwaveService: service}
}

// InitializeMicrowave godoc
// @Summary Initialize a new microwave
// @Description Create a new microwave with active status
// @Tags microwave
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Router /microwave/init [post]
func (ctrl *MicrowaveController) InitializeMicrowave(c *gin.Context) {
	microwave, err := ctrl.microwaveService.InitializeMicrowave()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to initialize microwave"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"microwave_id": microwave.ID})
}

func (ctrl *MicrowaveController) TurnOnMicrowave(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.microwaveService.TurnOnMicrowave(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "microwave turned on"})
}

func (ctrl *MicrowaveController) TurnOffMicrowave(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.microwaveService.TurnOffMicrowave(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "microwave turned off"})
}

func (ctrl *MicrowaveController) OpenDoor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.microwaveService.OpenDoor(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "door opened"})
}

func (ctrl *MicrowaveController) CloseDoor(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	err := ctrl.microwaveService.CloseDoor(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "door closed"})
}

func (ctrl *MicrowaveController) SetPowerLevel(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	level, _ := strconv.Atoi(c.Param("level"))
	err := ctrl.microwaveService.SetPowerLevel(uint(id), level)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "power level set"})
}

func (ctrl *MicrowaveController) GetCurrentMicrowave(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	microwave, err := ctrl.microwaveService.GetMicrowaveById(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"microwave": microwave})

}
