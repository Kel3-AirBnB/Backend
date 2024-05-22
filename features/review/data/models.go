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

func ReviewCoreToUserGorm(reviewCore review.Core) Review {
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
