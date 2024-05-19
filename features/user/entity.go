package user

import (
	"time"
)

type Core struct {
	ID             uint
	Name           string
	Email          string
	Password       string
	PhoneNumber    string
	ProfilePicture string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type DataInterface interface {
}

type ServiceInterface interface {
}
