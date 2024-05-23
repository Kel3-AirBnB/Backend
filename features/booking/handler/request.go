package handler

import (
	"airbnb/features/booking"

	"gorm.io/gorm"
)

type BookingRequest struct {
	gorm.Model
	UserID           uint   `json:"user_id" form:"user_id"`
	PenginapanID     uint   `json:"penginapan_id" form:"penginapan_id"`
	CheckIn          string `json:"checkin" form:"checkin"`
	CheckOut         string `json:"checkout" form:"checkout"`
	TotalTransaksi   string
	JenisTransaksi   string `json:"jenis_transaksi" form:"jenis_transaksi"`
	StatusPembayaran string `json:"status_pembayaran" form:"status_pembayaran"`
}

type PaymentRequest struct {
	JenisTransaksi   string `json:"jenis_transaksi" form:"jenis_transaksi"`
	StatusPembayaran string
	TotalTransaksi   string `json:"total_transaksi" form:"total_transaksi"`
}

func RequestToCore(input PaymentRequest) booking.Core {
	inputCore := booking.Core{
		JenisTransaksi:   input.JenisTransaksi,
		StatusPembayaran: input.StatusPembayaran,
		TotalTransaksi:   input.TotalTransaksi,
	}
	return inputCore
}

func GormToCore(input BookingRequest) booking.Core {

	exportCore := booking.Core{
		UserID:           input.UserID,
		PenginapanID:     input.PenginapanID,
		CheckIn:          input.CheckIn,
		CheckOut:         input.CheckIn,
		TotalTransaksi:   input.TotalTransaksi,
		JenisTransaksi:   input.JenisTransaksi,
		StatusPembayaran: input.StatusPembayaran,
	}
	return exportCore
}
