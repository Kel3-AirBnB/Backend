package data

import (
	"airbnb/features/booking"
	homestay "airbnb/features/homeStay"
	"airbnb/features/homeStay/data"

	"gorm.io/gorm"
)

type Booking struct {
	gorm.Model
	UserID           uint          `json:"user_id" form:"user_id"`
	PenginapanID     uint          `json:"penginapan_id" form:"penginapan_id"`
	CheckIn          string        `json:"checkin" form:"checkin"`
	CheckOut         string        `json:"checkout" form:"checkout"`
	TotalTransaksi   string        `json:"total_transaksi" form:"total_transaksi"`
	JenisTransaksi   string        `json:"jenis_transaksi" form:"jenis_transaksi"`
	StatusPembayaran string        `json:"status_pembayaran" form:"status_pembayaran"`
	HomeStay         data.HomeStay `gorm:"foreignKey:PenginapanID"`
}

func CoreToGorm(input booking.Core) Booking {
	bookingGorm := Booking{
		UserID:           input.UserID,
		PenginapanID:     input.PenginapanID,
		CheckIn:          input.CheckIn,
		CheckOut:         input.CheckOut,
		JenisTransaksi:   input.JenisTransaksi,
		StatusPembayaran: input.StatusPembayaran,
	}
	return bookingGorm
}

func GormToCore(input Booking) booking.Core {
	bookingCore := booking.Core{
		ID:               input.ID,
		UserID:           input.UserID,
		PenginapanID:     input.PenginapanID,
		CheckIn:          input.CheckIn,
		CheckOut:         input.CheckOut,
		JenisTransaksi:   input.JenisTransaksi,
		StatusPembayaran: input.StatusPembayaran,
		// HomeStay:         input.HomeStay,
	}
	return bookingCore
}

func GormToCoreHomestay(input data.HomeStay) homestay.Core {
	homestayCore := homestay.Core{
		Name:  input.Name,
		Harga: input.Harga,
	}
	return homestayCore
}
