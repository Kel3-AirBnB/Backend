package handler

import (
	"airbnb/features/user"

	"gorm.io/gorm"
)

type UserRequest struct {
	gorm.Model
	Nama               string `json:"nama" form:"nama"`
	Email              string `gorm:"unique" json:"email" form:"email"`
	Password           string `json:"password" form:"password"`
	KetikUlangPassword string `json:"repeat_password" form:"repeat_password"`
	TanggalLahir       string `json:"tanggal_lahir" form:"tanggal_lahir"`
	Foto               string `json:"foto" form:"foto"`
}

// type RegisterRequest struct {
// 	gorm.Model
// 	Nama               string `json:"nama" form:"nama"`
// 	Email              string `gorm:"unique" json:"email" form:"email"`
// 	Password           string `json:"password" form:"password"`
// 	KetikUlangPassword string `json:"repeat_password" form:"repeat_password"`
// 	TanggalLahir       string `json:"tanggal_lahir" form:"tanggal_lahir"`
// 	Foto               string `json:"foto" form:"foto"`
// }

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
