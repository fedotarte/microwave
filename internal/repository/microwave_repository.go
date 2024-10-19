package repository

import (
	"gorm.io/gorm"
	"microwave-service/internal/domain"
)

type MicrowaveRepository interface {
	CreateMicrowave(microwave *domain.Microwave) error
	FindMicrowaveByID(id uint) (*domain.Microwave, error)
	UpdateMicrowave(microwave *domain.Microwave) error
}

type microwaveRepo struct {
	db *gorm.DB
}

func NewMicrowaveRepository(db *gorm.DB) MicrowaveRepository {
	return &microwaveRepo{db: db}
}

func (r *microwaveRepo) CreateMicrowave(microwave *domain.Microwave) error {
	return r.db.Create(microwave).Error
}

func (r *microwaveRepo) FindMicrowaveByID(id uint) (*domain.Microwave, error) {
	var microwave domain.Microwave
	err := r.db.First(&microwave, id).Error
	if err != nil {
		return nil, err
	}
	return &microwave, nil
}

func (r *microwaveRepo) UpdateMicrowave(microwave *domain.Microwave) error {
	return r.db.Save(microwave).Error
}
