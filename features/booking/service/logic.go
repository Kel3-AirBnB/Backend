package service

import (
	"airbnb/features/booking"
	"airbnb/features/homestay"
	"airbnb/features/user"
	"airbnb/utils/helper"
	"errors"
	"log"
)

type bookingService struct {
	bookingData booking.DataInterface
	userData    user.DataInterface
	helper      helper.HelperInterface
}

func New(bd booking.DataInterface, ud user.DataInterface, hp helper.HelperInterface) booking.ServiceInterface {
	return &bookingService{
		bookingData: bd,
		userData:    ud,
		helper:      hp,
	}
}

func (s *bookingService) Create(input booking.Core) error {
	_, errID := s.userData.SelectById(input.UserID)
	if errID != nil {
		log.Print("Err Select By ID Service Layer", errID)
		return errID
	}

	err := s.bookingData.Insert(input)
	if err != nil {
		log.Print("Err Insert Service Layer", err)
		return err
	}
	return nil
}

func (p *bookingService) GetHomeById(id uint) (data *homestay.Core, err error) {

	if id <= 0 {
		return nil, errors.New("[validation] id not valid")
	}
	return p.bookingData.SelectHomeById(id)
}

func (p *bookingService) GetBookingById(id uint, userid uint) (data *booking.Core, err error) {
	_, errID := p.userData.SelectById(userid)
	if errID != nil {
		return nil, errID
	}

	if id <= 0 {
		return nil, errors.New("[validation] home id not valid")
	}
	return p.bookingData.SelectById(id, userid)
}

func (p *bookingService) Payment(id uint, userid uint, input booking.Core) error {
	if id <= 0 {
		return errors.New("id not valid")
	}

	// totalHari, nil := p.helper.GetTotalDay(input.CheckIn, input.CheckOut)
	// input.TotalTransaksi = totalHari *
	// err := p.bookingData.Payment(id, input)
	// if err != nil {
	// 	return err
	// }
	return nil
}
