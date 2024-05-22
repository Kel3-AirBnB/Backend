package data

import (
	"airbnb/features/review"

	"gorm.io/gorm"
)

type reviewQuery struct {
	db *gorm.DB
}

// Insert implements review.DataInterface.
func (r *reviewQuery) Insert(input review.Core) error {
	userGorm := ReviewCoreToUserGorm(input)
	tx := r.db.Create(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAll implements review.DataInterface.
func (r *reviewQuery) SelectAll() ([]review.Core, error) {
	var allReviews []Review
	tx := r.db.Find(&allReviews)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var allReviewCore []review.Core
	for _, v := range allReviews {
		allReviewCore = append(allReviewCore, review.Core{
			ID:           v.ID,
			PenginapanID: v.PenginapanID,
			UserID:       v.UserID,
			PesananID:    v.PesananID,
			Komentar:     v.Komentar,
			Rating:       v.Rating,
			Foto:         v.Foto,
			CreatedAt:    v.CreatedAt,
			UpdatedAt:    v.UpdatedAt,
		})
	}
	return allReviewCore, nil
}

func New(db *gorm.DB) review.DataInterface {
	return &reviewQuery{
		db: db,
	}
}
