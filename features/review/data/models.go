package data

import "gorm.io/gorm"

type Review struct {
	gorm.Model
	ID           uint   `json:"id" form:"id"`
	PenginapanID uint   `json:"penginapan_id" form:"penginapan_id"`
	UserID       uint   `json:"user_id" form:"user_id"`
	PesananID    uint   `json:"pesanan_id" form:"pesanan_id"`
	Komentar     string `json:"komentar" form:"komentar"`
	Rating       uint   `json:"rating" form:"rating"`
	Foto         string `json:"foto" form:"foto"`
}
