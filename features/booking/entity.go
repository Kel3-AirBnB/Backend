package booking

import "time"

type Core struct {
	ID               uint
	UserID           uint
	PenginapanID     uint
	CheckIn          string
	CheckOut         string
	JenisTransaksi   string
	StatusPembayaran string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type DataInterface interface {
}

type ServiceInterface interface {
}
