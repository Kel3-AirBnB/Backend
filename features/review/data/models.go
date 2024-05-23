package data

import (
	"airbnb/features/review"

	"gorm.io/gorm"
)

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
