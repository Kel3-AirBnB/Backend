package handler

import (
	"airbnb/features/user"

	"gorm.io/gorm"
)

type UserRequest struct {
	gorm.Model
	Nama               string `json:"name" form:"name"`
	Email              string `gorm:"unique" json:"email" form:"email"`
	Password           string `json:"password" form:"password"`
	KetikUlangPassword string `json:"retypepassword" form:"retypepassword"`
	TanggalLahir       string `json:"birth" form:"birth"`
	Foto               string `json:"profilepicture" form:"profilepicture"`
}

type LoginRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func RequestToCore(input UserRequest) user.Core {
	inputCore := user.Core{
		Nama:               input.Nama,
		Email:              input.Email,
		Password:           input.Password,
		KetikUlangPassword: input.KetikUlangPassword,
		TanggalLahir:       input.TanggalLahir,
		Foto:               input.Foto,
	}
	return inputCore
}
