package user

import (
	"time"
)

type Core struct {
	ID           uint
	Nama         string
	Email        string
	Password     string
	NomorTelepon string
	Foto         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

type DataInterface interface {
	Insert(input Core) error
}

type ServiceInterface interface {
	Create(input Core) error
}
