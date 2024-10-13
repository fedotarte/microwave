package interfaces

import (
	"microwave/constants"
	"microwave/models"
)

type MicrowaveInterface interface {
	Start() (error, *models.Microwave)
	Stop() (error, *models.Microwave)
	TurnOn() (error, *models.Microwave)
	TurnOff() (error, *models.Microwave)
	OpenDoor() (error, *models.Microwave)
	CloseDoor() (error, *models.Microwave)
	InsertFood() (error, *models.Microwave)
	GetFood() (error, *models.Microwave)
	SetPowerLevel(level constants.PowerLevel) (error, *models.Microwave)
	SetTimer(seconds int) (error, *models.Microwave)
	GetCurrentTime() (error, string)
}
