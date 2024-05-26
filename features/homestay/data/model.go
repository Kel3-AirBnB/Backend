package data

import (
	"airbnb/features/homestay"

	"gorm.io/gorm"
)

type Homestay struct {
	gorm.Model
	UserID    uint   `json:"user_id"`
	Name      string `json:"name"`
	Rating    string `json:"rating"`
	Foto      string `json:"foto"`
	Deskripsi string `json:"deskripsi"`
	Harga     string `json:"harga"`
	Alamat    string `json:"alamat"`
}

func HomestayCoreToHomestayGorm(homeStayCore homestay.Core) Homestay {
	return Homestay{
		Name:      homeStayCore.Name,
		Rating:    homeStayCore.Rating,
		Foto:      homeStayCore.Foto,
		Deskripsi: homeStayCore.Deskripsi,
		Harga:     homeStayCore.Harga,
		Alamat:    homeStayCore.Alamat,
	}
}

func HomestayGormToHomestayCore(homeStayGorm Homestay) homestay.Core {
	return homestay.Core{
		ID:        homeStayGorm.ID,
		Name:      homeStayGorm.Name,
		Rating:    homeStayGorm.Rating,
		Foto:      homeStayGorm.Foto,
		Deskripsi: homeStayGorm.Deskripsi,
		Harga:     homeStayGorm.Harga,
		Alamat:    homeStayGorm.Alamat,
		CreatedAt: homeStayGorm.CreatedAt,
		UpdatedAt: homeStayGorm.UpdatedAt,
	}
}
