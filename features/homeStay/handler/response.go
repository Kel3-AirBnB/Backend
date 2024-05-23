package handler

import homestay "airbnb/features/homestay"

type HomeStayResponse struct {
	ID        uint   `json:"id"`
	Name      string `json:"name"`
	Rating    string `json:"rating"`
	Foto      string `json:"foto"`
	Deskripsi string `json:"deskripsi"`
	Harga     string `json:"harga"`
	Alamat    string `json:"alamat"`
}

func CoreToResponse(homestayCore homestay.Core) HomeStayResponse {
	return HomeStayResponse{
		ID:        homestayCore.ID,
		Name:      homestayCore.Name,
		Rating:    homestayCore.Rating,
		Foto:      homestayCore.Foto,
		Deskripsi: homestayCore.Deskripsi,
		Harga:     homestayCore.Harga,
		Alamat:    homestayCore.Alamat,
	}
}
