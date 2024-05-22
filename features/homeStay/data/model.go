package data

import "gorm.io/gorm"

type HomeStay struct {
	gorm.Model
	ID        uint
	UserID    uint
	Name      string
	Rating    string
	Foto      string
	Deskripsi string
	Harga     string
	Alamat    string
}
