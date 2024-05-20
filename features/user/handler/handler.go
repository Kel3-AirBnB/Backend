package handler

import (
	"airbnb/features/user"
)

type UserHandler struct {
	userService user.ServiceInterface
	hashService encrypts.HashInterface
}
