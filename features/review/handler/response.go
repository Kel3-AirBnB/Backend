package handler

type ReviewResponse struct {
	ID           uint   `json:"id"`
	PenginapanID uint   `json:"penginapan_id" form:"penginapan_id"`
	UserID       uint   `json:"user_id" form:"user_id"`
	Rating       uint   `json:"rating" form:"rating"`
	Komentar     string `json:"komentar" form:"komentar"`
	Foto         string `json:"foto" form:"foto"`
}
