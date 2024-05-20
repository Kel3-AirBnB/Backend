package user

import (
	"time"
)

type Core struct {
	ID           uint
	Name         string
	Email        string
	Password     string
	NomorTelepon string
	Foto         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type DataInterface interface {
}

type ServiceInterface interface {
}
