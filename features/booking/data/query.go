package data

import (
	"airbnb/features/booking"
	"log"

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
		log.Print("Err Select By ID Data Layer", tx.Error)
		return tx.Error
	}

	return nil
}
