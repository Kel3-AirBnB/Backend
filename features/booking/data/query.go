package data

import (
	"airbnb/features/booking"

	"gorm.io/gorm"
)

type bookingQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) booking.DataInterface {
	return &bookingQuery{
		db: db,
	}
}

func (p *bookingQuery) Insert(input booking.Core) error {
	projectGorm := CoreToGorm(input)
	tx := p.db.Create(&projectGorm)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
