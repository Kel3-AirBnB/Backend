package homestay

import (
	"io"
	"time"
)

type Core struct {
	ID        uint
	Name      string
	Rating    string
	Foto      string
	Deskripsi string
	Harga     string
	Alamat    string
	CreatedAt time.Time
	UpdatedAt time.Time
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
	GethomeStayid(id uint) (data *Core, err error)
	Delete(id uint) error
}
