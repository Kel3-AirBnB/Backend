package data

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	// ID           uint
	PenginapanID uint
	UserID       uint
	PesananID    uint
	Komentar     string
	Rating       uint
	Foto         string
}
