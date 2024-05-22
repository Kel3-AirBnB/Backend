package data

import (
	"airbnb/features/booking"
	"time"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID           uint      `json:"user_id" form:"user_id"`
	PenginapanID     uint      `json:"penginapan_id" form:"penginapan_id"`
	CheckIn          time.Time `json:"checkin" form:"checkin"`
	CheckOut         time.Time `json:"checkout" form:"checkout"`
	JenisTransaksi   string    `json:"jenis_transaksi" form:"jenis_transaksi"`
	StatusPembayaran string    `json:"status_pembayaran" form:"status_pembayaran"`
}

func CoreToGorm(input booking.Core) Booking {
	projectGorm := Booking{
		UserID:         input.UserID,
		PenginapanID:   input.PenginapanID,
		CheckIn:        input.CheckIn,
		CheckOut:       input.CheckOut,
		JenisTransaksi: input.JenisTransaksi,
	}
	return projectGorm
}
