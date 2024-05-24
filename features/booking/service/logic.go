package service

import (
	"airbnb/features/booking"
	"airbnb/features/homestay"
	"airbnb/features/user"
	"airbnb/utils/helper"
	"errors"
	"fmt"
	"log"
	"strconv"
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

func (s *bookingService) Create(input booking.Core, checkin string, checkout string, price string) error {
	_, errID := s.userData.SelectById(input.UserID)
	if errID != nil {
		log.Print("Err Select By ID Service Layer", errID)
		return errID
	}

	totalHari, nil := s.helper.GetTotalDay(checkin, checkout)

	idConv, errConv := strconv.Atoi(price)
	if errConv != nil {
		fmt.Println("Error:", errConv)
		return errConv
	}
	totalTransaksi := totalHari * idConv

	input.TotalTransaksi = strconv.Itoa(totalTransaksi)

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

func (p *bookingService) Payment(id int, userid int, input booking.Core, checkin string, checkout string, price string) (string, error) {
	if id <= 0 {
		return "", errors.New("id not valid")
	}

	totalHari, nil := p.helper.GetTotalDay(checkin, checkout)

	idConv, errConv := strconv.Atoi(price)
	if errConv != nil {
		fmt.Println("Error:", errConv)
		return "", errConv
	}
	totalTransaksi := totalHari * idConv

	input.TotalTransaksi = strconv.Itoa(totalTransaksi)

	input.StatusPembayaran = "Sudah Dibayar"
	errResult := p.bookingData.Payment(id, input)
	if errResult != nil {
		return "Error", errResult
	}
	return "", nil
}

func (p *bookingService) GetAll(userid uint) ([]booking.Core, error) {
	if userid <= 0 {
		return nil, errors.New("[validation] id not valid")
	}
	return p.bookingData.SelectAll(userid)
}

func (p *bookingService) GetBookingByHomestay(id uint, userid uint) ([]booking.Core, error) {
	fmt.Println("[Service Layer] Masuk Ke GetBookingByHomestay")

	if id <= 0 {
		return nil, errors.New("[validation] id not valid")
	}

	fmt.Println("[Service Layer] id rumah", id)
	fmt.Println("[Service Layer	] userid", userid)
	dataCurrentHome, errcheckValidated := p.bookingData.ValidatedHomeById(id, userid)
	if errcheckValidated != nil {
		return nil, errors.New("[validation] id and current user not valid")
	}

	fmt.Println("[Service] dataCurrentHome", dataCurrentHome.ID)
	return p.bookingData.SelectBookingByHomestayId(dataCurrentHome.ID, dataCurrentHome.UserID)
}
