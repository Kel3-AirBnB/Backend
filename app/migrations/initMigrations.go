package migrations

import (
	_dataBooking "airbnb/features/booking/data"
	_dataHomestay "airbnb/features/homestay/data"
	_dataReview "airbnb/features/review/data"
	_dataUser "airbnb/features/user/data"

	"gorm.io/gorm"
)

func InitMigrations(db *gorm.DB) {
	db.AutoMigrate(&_dataUser.User{})
	db.AutoMigrate(&_dataBooking.Booking{})
	db.AutoMigrate(&_dataReview.Review{})
	db.AutoMigrate(&_dataHomestay.Homestay{})
}
