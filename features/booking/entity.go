package booking

import (
	"airbnb/features/homeStay"
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
	HomeStay         homeStay.HomeStayCore
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectById(id uint, userid uint) (*Core, error)
	SelectHomeById(id uint) (*homeStay.HomeStayCore, error)
}

type ServiceInterface interface {
	Create(input Core) error
	GetBookingById(id uint, userid uint) (data *Core, err error)
	GetHomeById(id uint) (data *homeStay.HomeStayCore, err error)
}
