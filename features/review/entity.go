package review

import (
	"io"
	"time"
)

type Core struct {
	ID           uint
	PenginapanID uint
	UserID       uint
	PesananID    uint
	Komentar     string
	Rating       uint
	Foto         string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
type DataInterface interface {
	SelectAll() ([]Core, error)
	Insert(input Core) error
	SelectById(id uint) (*Core, error)
	Delete(id uint) error
}

type ServiceInterface interface {
	GetAll() ([]Core, error)
	Create(input Core, file io.Reader, handlerFilename string) (string, error)
	GetReviews(id uint) (data *Core, err error)
	Delete(id uint) error
}
