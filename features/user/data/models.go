package data

import (
	"airbnb/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama         string `json:"nama" form:"nama"`
	Email        string `gorm:"unique" json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	TanggalLahir string `json:"tanggal_lahir" form:"tanggal_lahir"`
	Foto         string `json:"foto" form:"foto"`
}

func UserCoreToUserGorm(userCore user.Core) User {
	userGorm := User{
		Nama:         userCore.Nama,
		Email:        userCore.Email,
		Password:     userCore.Password,
		TanggalLahir: userCore.TanggalLahir,
		Foto:         userCore.Foto,
	}
	return userGorm
}

func UserGormToUserCore(userGorm User) user.Core {
	userCore := user.Core{
		ID:           userGorm.ID,
		Nama:         userGorm.Nama,
		Email:        userGorm.Email,
		Password:     userGorm.Password,
		TanggalLahir: userGorm.TanggalLahir,
		Foto:         userGorm.Foto,
		CreatedAt:    userGorm.CreatedAt,
		UpdatedAt:    userGorm.UpdatedAt,
	}
	return userCore
}
