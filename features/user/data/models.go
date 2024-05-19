package data

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID             uint   `json:"id" form:"id"`
	Name           string `json:"name" form:"name"`
	Email          string `gorm:"unique" json:"email" form:"email"`
	Password       string `json:"password" form:"password"`
	PhoneNumber    string `gorm:"unique" json:"phonenumber" form:"phonenumber"`
	ProfilePicture string `json:"profilepicture" form:"profilepicture"`
}
