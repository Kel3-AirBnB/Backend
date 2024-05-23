package booking

import (
	"airbnb/features/homestay"
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
	Payment(id int, input Core) error
	SelectAll(userid uint) ([]Core, error)
}

type ServiceInterface interface {
	Create(input Core, checkin string, checkout string, price string) error
	GetBookingById(id uint, userid uint) (data *Core, err error)
	GetHomeById(id uint) (data *homestay.Core, err error)
	Payment(id int, userid int, input Core, checkin string, checkout string, price string) (string, error)
	GetAll(userid uint) ([]Core, error)
}
