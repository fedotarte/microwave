package interfaces

import "microwave-service/constants"

type MicrowaveInterface interface {
	Start() error
	Stop() error
	TurnOn() error
	TurnOff() error
	OpenDoor() error
	CloseDoor() error
	InsertFood() error
	GetFood() error
	SetPowerLevel(level constants.PowerLevel) error
	SetTimer(seconds int) error
	GetCurrentTime() (error, string)
}
