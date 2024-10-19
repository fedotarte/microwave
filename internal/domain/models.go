package domain

import "time"

type Command struct {
	ID        uint   `gorm:"primary_key"`
	Name      string `gorm:"not null"`
	Status    string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
	History   []CommandHistory
}

type CommandHistory struct {
	ID        uint      `gorm:"primary_key"`
	CommandID uint      `gorm:"not null"`
	Status    string    `gorm:"not null"`
	ChangeAt  time.Time `gorm:"autoCreateTime"`
}

type Microwave struct {
	ID         uint   `gorm:"primaryKey"`
	Status     string `gorm:"not null"`
	DoorState  string `gorm:"not null"`
	PowerLevel int    `gorm:"not null"`
	IsOn       bool   `gorm:"not null"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
