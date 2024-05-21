package data

import "gorm.io/gorm"

type homeStay struct {
	gorm.Model
	ID        uint
	Name      string
	Rating    string
	Foto      string
	Deskripsi string
	Harga     string
	Alamat    string
}
