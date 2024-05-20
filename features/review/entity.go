package review

type Core struct {
	ID           uint
	PenginapanID uint
	UserID       uint
	PesananID    uint
	Komentar     string
	Rating       uint
	Foto         string
}
type DataInterface interface {
}

type ServiceInterface interface {
}
