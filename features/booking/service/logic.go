package service

import (
	"airbnb/features/booking"
	"airbnb/features/user"
)

type bookingService struct {
	bookingData booking.DataInterface
	userData    user.DataInterface
}

func New(bd booking.DataInterface, ud user.DataInterface) booking.ServiceInterface {
	return &bookingService{
		bookingData: bd,
		userData:    ud,
	}
}

// func (s *bookingService) Create(input booking.Core) error {
// 	_, errID := s.userData.SelectById(input.UserID)
// 	if errID != nil {
// 		return errID
// 	}

// 	err := s.bookingData.Insert(input)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
