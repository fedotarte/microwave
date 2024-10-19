package services

import (
	"errors"
	"microwave-service/internal/domain"
	"microwave-service/internal/repository"
	"time"
)

type MicrowaveService struct {
	repo repository.MicrowaveRepository
}

func NewMicrowaveService(repo repository.MicrowaveRepository) *MicrowaveService {
	return &MicrowaveService{repo: repo}
}

func (s *MicrowaveService) InitializeMicrowave() (domain.Microwave, error) {
	microwave := domain.Microwave{
		Status:     "active",
		DoorState:  "closed",
		PowerLevel: 0,
		IsOn:       false,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}
	err := s.repo.CreateMicrowave(&microwave)
	return microwave, err
}

func (s *MicrowaveService) GetMicrowaveById(id uint) (*domain.Microwave, error) {
	microwave, err := s.repo.FindMicrowaveByID(id)
	if err != nil {
		return nil, err
	}
	return microwave, nil
}

func (s *MicrowaveService) TurnOnMicrowave(id uint) error {
	microwave, err := s.repo.FindMicrowaveByID(id)
	if err != nil {
		return err
	}

	if microwave.DoorState == "open" {
		return errors.New("can't turn on microwave with the door open")
	}

	microwave.IsOn = true
	microwave.UpdatedAt = time.Now()
	return s.repo.UpdateMicrowave(microwave)
}

func (s *MicrowaveService) TurnOffMicrowave(id uint) error {
	microwave, err := s.repo.FindMicrowaveByID(id)
	if err != nil {
		return err
	}

	microwave.IsOn = false
	microwave.UpdatedAt = time.Now()
	return s.repo.UpdateMicrowave(microwave)
}

func (s *MicrowaveService) OpenDoor(id uint) error {
	microwave, err := s.repo.FindMicrowaveByID(id)
	if err != nil {
		return err
	}

	if microwave.IsOn {
		return errors.New("can't open the door while microwave is on")
	}

	microwave.DoorState = "open"
	microwave.UpdatedAt = time.Now()
	return s.repo.UpdateMicrowave(microwave)
}

func (s *MicrowaveService) CloseDoor(id uint) error {
	microwave, err := s.repo.FindMicrowaveByID(id)
	if err != nil {
		return err
	}

	microwave.DoorState = "closed"
	microwave.UpdatedAt = time.Now()
	return s.repo.UpdateMicrowave(microwave)
}

func (s *MicrowaveService) SetPowerLevel(id uint, level int) error {
	if level < 0 || level > 10 {
		return errors.New("invalid power level, must be between 0 and 10")
	}

	microwave, err := s.repo.FindMicrowaveByID(id)
	if err != nil {
		return err
	}

	microwave.PowerLevel = level
	microwave.UpdatedAt = time.Now()
	return s.repo.UpdateMicrowave(microwave)
}
