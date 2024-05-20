package handler

import (
	"airbnb/features/user"

	"gorm.io/gorm"
)

type UserRequest struct {
	gorm.Model
	Nama         string `json:"name" form:"name"`
	Email        string `gorm:"unique" json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	NomorTelepon string `gorm:"unique" json:"phonenumber" form:"phonenumber"`
	Foto         string `json:"profilepicture" form:"profilepicture"`
}

func RequestToCore(input UserRequest) user.Core {
	inputCore := user.Core{
		ID:           0,
		Nama:         input.Nama,
		Email:        input.Email,
		Password:     input.Password,
		NomorTelepon: input.NomorTelepon,
		Foto:         input.Foto,
	}
	return inputCore
}
