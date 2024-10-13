package models

import (
	"microwave/constants"
	"time"
)

type Microwave struct {
	StandBy    bool
	PowerLevel constants.PowerLevel
	Timer      int
	IsRunning  bool
	IsLight    bool
	IsDoorOpen bool
	HasItem    bool
	OnTime     time.Time
}

func NewMicrowave() *Microwave {
	return &Microwave{
		StandBy:    false,
		PowerLevel: constants.PowerLevel600,
		Timer:      0,
		IsRunning:  false,
		IsLight:    false,
		IsDoorOpen: false,
		OnTime:     time.Now(),
	}
}

func (m *Microwave) GetCurrentTime() string {
	currentTime := time.Now().Format("15:04")
	return currentTime
}
