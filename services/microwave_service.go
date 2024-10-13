package services

import (
	"errors"
	"microwave/constants"
	"microwave/models"
)

type MicrowaveService struct {
	Microwave *models.Microwave
}

func NewMicrowaveService(m *models.Microwave) *MicrowaveService {
	return &MicrowaveService{
		Microwave: m,
	}
}

func (ms *MicrowaveService) TurnOn() (error, *models.Microwave) {
	if ms.Microwave.StandBy {
		return errors.New("микроволновка уже включена"), nil
	}
	ms.Microwave.StandBy = true
	return nil, ms.Microwave
}

func (ms *MicrowaveService) TurnOff() (error, *models.Microwave) {
	if !ms.Microwave.StandBy {
		return errors.New("микроволновка уже выклбчена"), nil
	}
	ms.Microwave.IsRunning = false
	ms.Microwave.IsLight = false
	ms.Microwave.StandBy = false
	return nil, ms.Microwave
}

func (ms *MicrowaveService) Start() (error, *models.Microwave) {
	if !ms.Microwave.StandBy {
		return errors.New("микроволновку невозмодно запустить, так как она не включена"), nil
	}
	ms.Microwave.IsRunning = true
	return nil, ms.Microwave
}

func (ms *MicrowaveService) Stop() (error, *models.Microwave) {
	if !ms.Microwave.StandBy {
		return errors.New("микроволновку невозмодно запустить, так как она не включена"), nil
	} else if !ms.Microwave.IsRunning {
		return errors.New("микроволновка не запущена"), nil
	}
	ms.Microwave.IsRunning = false
	return nil, ms.Microwave
}

func (ms *MicrowaveService) OpenDoor() (error, *models.Microwave) {

	if ms.Microwave.IsDoorOpen {
		return errors.New("Дверь уже открыта"), nil
	}

	if !ms.Microwave.StandBy {
		if ms.Microwave.IsRunning {
			ms.Microwave.IsRunning = false
		}
		ms.Microwave.IsLight = true
	}

	ms.Microwave.IsDoorOpen = true

	return nil, ms.Microwave
}

func (ms *MicrowaveService) CloseDoor() (error, *models.Microwave) {

	if !ms.Microwave.IsDoorOpen {
		return errors.New("Дверь уже закрыта"), nil
	}

	if ms.Microwave.StandBy {
		if ms.Microwave.Timer > 0 {
			ms.Microwave.IsRunning = true
		}
		ms.Microwave.IsLight = false
	}

	ms.Microwave.IsDoorOpen = false

	return nil, ms.Microwave
}

func (ms *MicrowaveService) InsertFood() (error, *models.Microwave) {
	if ms.Microwave.HasItem {
		return errors.New("Больше уже не вместится, надо сначала убрать"), nil
	}
	ms.Microwave.HasItem = true

	return nil, ms.Microwave
}

func (ms *MicrowaveService) GetFood() (error, *models.Microwave) {
	if !ms.Microwave.HasItem {
		return errors.New("Нечего вынимать, Внутри пусто"), nil
	}

	return nil, ms.Microwave
}

func (ms *MicrowaveService) SetPowerLevel(level constants.PowerLevel) (error, *models.Microwave) {
	ms.Microwave.PowerLevel = level

	return nil, ms.Microwave
}

func (ms *MicrowaveService) SetTimer(seconds int) (error, *models.Microwave) {
	if seconds < 0 {
		return errors.New("только полодительные числа"), nil
	}
	if seconds == 0 && ms.Microwave.IsRunning {
		ms.Microwave.IsRunning = false
	}

	return nil, ms.Microwave
}

func (ms *MicrowaveService) GetCurrentTime() (error, string) {
	if !ms.Microwave.StandBy {
		return errors.New("микроволновка выключена"), ""
	}
	return nil, ms.Microwave.GetCurrentTime()
}
