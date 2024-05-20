package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint   `json:"id" form:"id"`
	Nama         string `json:"name" form:"name"`
	Email        string `gorm:"unique" json:"email" form:"email"`
	Password     string `json:"password" form:"password"`
	NomorTelepon string `gorm:"unique" json:"phonenumber" form:"phonenumber"`
	Foto         string `json:"profilepicture" form:"profilepicture"`
}
