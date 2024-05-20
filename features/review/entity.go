package review

type Core struct {
	ID           uint
	PENGINAPANID uint
	USERID       uint
	PESANANID    uint
	KOMENTAR     string
	RATING       uint
	FOTO         string
}
type DataInterface interface {
}

type ServiceInterface interface {
}
