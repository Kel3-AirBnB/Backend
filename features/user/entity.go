package user

import (
	"io"
	"time"
)

type Core struct {
	ID                 uint
	Nama               string
	Email              string
	Password           string
	KetikUlangPassword string
	NomorTelepon       string
	Foto               string
	Token              string
	CreatedAt          time.Time
	UpdatedAt          time.Time
}

type DataInterface interface {
	Insert(input Core) error
	SelectByEmail(email string) (*Core, error)
	SelectById(id uint) (*Core, error)
}

type ServiceInterface interface {
	Create(input Core, file io.Reader, handlerFilename string) (string, error)
	UploadFileToS3(file io.Reader, fileName string) (string, error)
	Login(email string, password string) (data *Core, token string, err error)
	GetProfile(id uint) (data *Core, err error)
}
