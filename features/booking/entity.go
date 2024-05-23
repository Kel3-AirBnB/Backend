package booking

import (
	homestay "airbnb/features/homestay"
	"time"
)

type Core struct {
	ID               uint
	UserID           uint
	PenginapanID     uint
	CheckIn          string
	CheckOut         string
	TotalTransaksi   string
	JenisTransaksi   string
	StatusPembayaran string
	HomeStay         homestay.Core
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectById(id uint, userid uint) (*Core, error)
	SelectHomeById(id uint) (*homestay.Core, error)
}

type ServiceInterface interface {
	Create(input Core) error
	GetBookingById(id uint, userid uint) (data *Core, err error)
	GetHomeById(id uint) (data *homestay.Core, err error)
}
