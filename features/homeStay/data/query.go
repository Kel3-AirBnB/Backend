package data

import (
	"airbnb/features/homestay"

	"gorm.io/gorm"
)

type homeStayQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) homestay.DataInterface {
	return &homeStayQuery{
		db: db,
	}
}

// Insert implements review.DataInterface.
func (h *homeStayQuery) Insert(input homestay.Core) error {
	userGorm := HomestayCoreToHomestayGorm(input)
	tx := h.db.Create(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

// SelectAll implements review.DataInterface.
func (h *homeStayQuery) SelectAll() ([]homestay.Core, error) {
	var allhomeStay []Homestay
	tx := h.db.Find(&allhomeStay)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var allhomestayCore []homestay.Core
	for _, v := range allhomeStay {
		allhomestayCore = append(allhomestayCore, homestay.Core{
			ID:        v.ID,
			Name:      v.Name,
			Rating:    v.Rating,
			Foto:      v.Foto,
			Deskripsi: v.Deskripsi,
			Harga:     v.Harga,
			Alamat:    v.Alamat,
			CreatedAt: v.CreatedAt,
			UpdatedAt: v.UpdatedAt,
		})
	}
	return allhomestayCore, nil
}

// SelectById implements review.DataInterface.
func (h *homeStayQuery) SelectById(id uint) (*homestay.Core, error) {
	var homeStayData Homestay
	tx := h.db.First(&homeStayData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}
	var reviewcore = HomestayGormToHomestayCore(homeStayData)

	return &reviewcore, nil
}

// Delete implements homestay.DataInterface.
func (h *homeStayQuery) Delete(id uint) error {
	tx := h.db.Delete(&Homestay{}, id)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
