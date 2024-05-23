package handler

import homestay "airbnb/features/homeStay"

type HomeStayRequest struct {
	Name      string `json:"name" form:"name" validate:"required"`
	Rating    string `json:"rating" form:"rating" validate:"required"`
	Foto      string `json:"foto" form:"foto" validate:"required"`
	Deskripsi string `json:"deskripsi" form:"deskripsi" validate:"required"`
	Harga     string `json:"harga" form:"harga" validate:"required"`
	Alamat    string `json:"alamat" form:"alamat" validate:"required"`
}

func RequestToCore(input HomeStayRequest) homestay.Core {
	return homestay.Core{
		Name:      input.Name,
		Rating:    input.Rating,
		Foto:      input.Foto,
		Deskripsi: input.Deskripsi,
		Harga:     input.Harga,
		Alamat:    input.Alamat,
	}
}
