package handler

import (
	"airbnb/features/booking"
	"airbnb/features/homestay"
)

type BookingResponse struct {
	ID               uint   `json:"id,omitempty"`
	UserID           uint   `json:"user_id,omitempty"`
	PenginapanID     uint   `json:"penginapan_id,omitempty"`
	NamaPenginapan   string `json:"nama_penginapan,omitempty"`
	CheckIn          string `json:"checkin,omitempty"`
	CheckOut         string `json:"checkout,omitempty"`
	TotalTransaksi   string `json:"total_transaksi,omitempty"`
	JenisTransaksi   string `json:"jenis_transaksi,omitempty"`
	StatusPembayaran string `json:"status_pembayaran,omitempty"`
}

func PaymentResponse(input booking.Core, inputHome homestay.Core) BookingResponse {
	result := BookingResponse{
		ID:             input.ID,
		UserID:         input.UserID,
		PenginapanID:   input.PenginapanID,
		NamaPenginapan: inputHome.Name,
		CheckIn:        input.CheckIn,
		CheckOut:       input.CheckOut,
		JenisTransaksi: input.JenisTransaksi,
		TotalTransaksi: input.TotalTransaksi,
	}
	return result
}

func SelectResponses(input booking.Core) BookingResponse {
	result := BookingResponse{
		ID:               input.ID,
		UserID:           input.UserID,
		PenginapanID:     input.PenginapanID,
		CheckIn:          input.CheckIn,
		CheckOut:         input.CheckOut,
		TotalTransaksi:   input.TotalTransaksi,
		StatusPembayaran: input.StatusPembayaran,
	}
	return result
}

func InvoiceResponse(input booking.Core, inputHome homestay.Core) BookingResponse {
	result := BookingResponse{
		ID:               input.ID,
		UserID:           input.UserID,
		PenginapanID:     input.PenginapanID,
		NamaPenginapan:   inputHome.Name,
		CheckIn:          input.CheckIn,
		CheckOut:         input.CheckOut,
		TotalTransaksi:   input.TotalTransaksi,
		JenisTransaksi:   input.JenisTransaksi,
		StatusPembayaran: input.StatusPembayaran,
	}
	return result
}
