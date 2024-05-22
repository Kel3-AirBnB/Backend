package handler

import "airbnb/features/review"

type ReviewRequest struct {
	PenginapanID uint   `json:"penginapan_id" form:"penginapan_id"`
	UserID       uint   `json:"user_id" form:"user_id"`
	Rating       uint   `json:"rating" form:"rating"`
	Komentar     string `json:"komentar" form:"komentar"`
	Foto         string `json:"review_picture" form:"review_picture"`
}

func RequestToCore(input ReviewRequest) review.Core {
	inputCore := review.Core{
		PenginapanID: input.PenginapanID,
		UserID:       input.UserID,
		Komentar:     input.Komentar,
		Rating:       input.Rating,
		Foto:         input.Foto,
	}
	return inputCore
}
