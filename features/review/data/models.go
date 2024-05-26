package data

import (
	"airbnb/features/review"

	"gorm.io/gorm"
)

type Review struct {
	gorm.Model
	PenginapanID uint   `json:"penginapan_id" form:"penginapan_id"`
	UserID       uint   `json:"user_id" form:"user_id"`
	PesananID    uint   `json:"pesanan_id" form:"pesanan_id"`
	Komentar     string `json:"komentar" form:"komentar"`
	Rating       uint   `json:"rating" form:"rating"`
	Foto         string `json:"foto" form:"foto"`
	// User         data.User `gorm:"foreignKey:PenginapanID"`
}

func ReviewCoreToReviewGorm(reviewCore review.Core) Review {
	reviewGorm := Review{

		PenginapanID: reviewCore.PenginapanID,
		UserID:       reviewCore.UserID,
		PesananID:    reviewCore.PesananID,
		Komentar:     reviewCore.Komentar,
		Rating:       reviewCore.Rating,
		Foto:         reviewCore.Foto,
	}
	return reviewGorm
}
func ReviewGormToReviewCore(reviewGorm Review) review.Core {
	reviewCore := review.Core{
		ID:           reviewGorm.ID,
		PenginapanID: reviewGorm.PenginapanID,
		UserID:       reviewGorm.UserID,
		PesananID:    reviewGorm.PesananID,
		Komentar:     reviewGorm.Komentar,
		Rating:       reviewGorm.Rating,
		Foto:         reviewGorm.Foto,
		CreatedAt:    reviewGorm.CreatedAt,
		UpdatedAt:    reviewGorm.UpdatedAt,
	}
	return reviewCore
}
