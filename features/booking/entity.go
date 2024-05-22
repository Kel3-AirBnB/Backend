package booking

import "time"

type Core struct {
	ID               uint
	UserID           uint
	PenginapanID     uint
	CheckIn          time.Time
	CheckOut         time.Time
	JenisTransaksi   string
	StatusPembayaran string
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type DataInterface interface {
	Insert(input Core) error
}

type ServiceInterface interface {
	Create(input Core) error
}
