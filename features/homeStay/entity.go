package homeStay

type HomeStayCore struct {
	ID        uint
	Name      string
	Rating    string
	Foto      string
	Deskripsi string
	Harga     string
	Alamat    string
}

type DataInterface interface {
}

type ServiceInterface interface {
}
