package data

import (
	"airbnb/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nama         string `json:"name" form:"name"`
	Email        string `gorm:"unique" json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	NomorTelepon string `gorm:"unique" json:"phonenumber" form:"phonenumber"`
	Foto         string `json:"profilepicture" form:"profilepicture"`
}

func UserCoreToUserGorm(userCore user.Core) User {
	userGorm := User{
		Nama:         userCore.Nama,
		Email:        userCore.Email,
		Password:     userCore.Password,
		NomorTelepon: userCore.NomorTelepon,
		Foto:         userCore.Foto,
	}
	return userGorm
}
