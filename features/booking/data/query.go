package data

import (
	"airbnb/features/booking"
	"airbnb/features/homestay"
	"airbnb/features/homestay/data"
	"fmt"
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

func (p *bookingQuery) SelectById(id uint, userid uint) (*booking.Core, error) {
	var bookingData Booking
	tx := p.db.Where("user_id = ?", userid).First(&bookingData, id)
	fmt.Println("[Data Layer] TotalTransaksi: ", bookingData.TotalTransaksi)
	if tx.Error != nil {
		return nil, tx.Error
	}

	projectcore := GormToCore(bookingData)
	log.Print(projectcore)
	return &projectcore, nil
}

func (p *bookingQuery) SelectHomeById(id uint) (*homestay.Core, error) {
	var homestayData data.Homestay
	tx := p.db.First(&homestayData, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	homestaycore := GormToCoreHomestay(homestayData)
	log.Print(homestaycore)
	return &homestaycore, nil
}

func (p *bookingQuery) Payment(id int, input booking.Core) error {
	inputGorm := CoreToGorm(input)
	tx := p.db.Model(&Booking{}).Where("id = ?", id).Updates(&inputGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
