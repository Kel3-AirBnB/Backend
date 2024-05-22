package handler

import (
	"airbnb/features/booking"
	"time"

	"gorm.io/gorm"
)

type BookingRequest struct {
	gorm.Model
	UserID         uint      `json:"user_id" form:"user_id"`
	PenginapanID   uint      `json:"penginapan_id" form:"penginapan_id"`
	CheckIn        time.Time `json:"checkin" form:"checkin"`
	CheckOut       time.Time `json:"checkout" form:"checkout"`
	JenisTransaksi string    `json:"jenis_transaksi" form:"jenis_transaksi"`
}

func GormToCore(input BookingRequest) booking.Core {

	exportCore := booking.Core{
		UserID:         input.UserID,
		PenginapanID:   input.PenginapanID,
		CheckIn:        input.CheckIn,
		CheckOut:       input.CheckIn,
		JenisTransaksi: input.JenisTransaksi,
	}
	return exportCore
}
