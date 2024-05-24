package handler

import "airbnb/features/review"

type ReviewResponse struct {
	ID           uint   `json:"id"`
	PenginapanID uint   `json:"penginapan_id" form:"penginapan_id"`
	UserID       uint   `json:"userId" form:"userId"`
	Rating       uint   `json:"rating" form:"rating"`
	Komentar     string `json:"komentar" form:"komentar"`
	Foto         string `json:"foto" form:"foto"`
}

func CoreToGorm(reviewGorm review.Core) ReviewResponse {
	reviewCore := ReviewResponse{
		ID:           reviewGorm.ID,
		PenginapanID: reviewGorm.PenginapanID,
		UserID:       reviewGorm.UserID,
		Rating:       reviewGorm.Rating,
		Komentar:     reviewGorm.Komentar,
		Foto:         reviewGorm.Foto,
	}

	return reviewCore
}
