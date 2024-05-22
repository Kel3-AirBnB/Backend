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
}

type ServiceInterface interface {
	GetAll() ([]Core, error)
	Create(input Core, file io.Reader, handlerFilename string) (string, error)
}
