package data

import (
	"airbnb/features/review"

	"time"

	"gorm.io/gorm"
)

type reviewQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) review.DataInterface {
	return &reviewQuery{
		db: db,
	}
}

// SelectByUserID implements review.DataInterface.
func (r *reviewQuery) SelectByUserID(userid uint) ([]review.Core, error) {
	var reviewData []Review

	if tx := r.db.Where("user_id = ?", userid).Find(&reviewData); tx.Error != nil {
		return nil, tx.Error
	}

	var allProjectCore []review.Core
	for _, v := range reviewData {
		allProjectCore = append(allProjectCore, review.Core{
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
	return allProjectCore, nil
}

// SelectByPenginapanID implements review.DataInterface.
func (r *reviewQuery) SelectByPenginapanID(penginapanID uint) ([]review.Core, error) {
	var reviews []Review
	if err := r.db.Where("penginapan_id = ?", penginapanID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	var reviewCores []review.Core
	for _, reviewData := range reviews {
		reviewCores = append(reviewCores, ReviewGormToReviewCore(reviewData))
	}
	return reviewCores, nil
}

// EditById implements review.DataInterface.
func (r *reviewQuery) EditById(id uint, input review.Core) error {

	var existingReview Review
	if err := r.db.First(&existingReview, id).Error; err != nil {
		return err
	}

	// Mengonversi review.Core menjadi ReviewGorm
	inputGorm := ReviewCoreToReviewGorm(input)

	// Memperbarui nilai-nilai yang diperlukan dengan menggunakan updates
	updateValues := map[string]interface{}{
		"penginapan_id": inputGorm.PenginapanID,
		"user_id":       inputGorm.UserID,
		"pesanan_id":    inputGorm.PesananID,
		"komentar":      inputGorm.Komentar,
		"rating":        inputGorm.Rating,
		"foto":          inputGorm.Foto,
		"updated_at":    time.Now(),
	}

	// Melakukan pembaruan data ulasan menggunakan updates
	if err := r.db.Model(&existingReview).Updates(updateValues).Error; err != nil {
		return err
	}

	return nil

}

// Delete implements review.DataInterface.
func (r *reviewQuery) Delete(id uint) error {
	tx := r.db.Delete(&Review{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil

}

// SelectById implements review.DataInterface.
func (r *reviewQuery) SelectById(id uint) (*review.Core, error) {
	var reviewData Review
	tx := r.db.First(&reviewData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var reviewcore = ReviewGormToReviewCore(reviewData)

	return &reviewcore, nil
}

// Insert implements review.DataInterface.
func (r *reviewQuery) Insert(input review.Core) error {
	userGorm := ReviewCoreToReviewGorm(input)
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
