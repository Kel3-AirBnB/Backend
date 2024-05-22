package handler

import "airbnb/features/user"

type UserResponse struct {
	ID           uint   `json:"id,omitempty"`
	Nama         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	TanggalLahir string `json:"birth,omitempty"`
	Foto         string `json:"profilepicture,omitempty"`
	Token        string `json:"token,omitempty"`
}

func ResponseProfile(userResponse *user.Core) UserResponse {
	result := UserResponse{
		ID:           userResponse.ID,
		Nama:         userResponse.Nama,
		Email:        userResponse.Email,
		TanggalLahir: userResponse.TanggalLahir,
		Foto:         userResponse.Foto,
		Token:        userResponse.Token,
	}
	return result
}

func CoreToGorm(userGorm user.Core) UserResponse {
	userCore := UserResponse{
		ID:           userGorm.ID,
		Nama:         userGorm.Nama,
		Email:        userGorm.Email,
		TanggalLahir: userGorm.TanggalLahir,
		Foto:         userGorm.Foto,
	}

	return userCore
}

func ResponseLogin(userResponse *user.Core) UserResponse {
	result := UserResponse{
		ID:    userResponse.ID,
		Nama:  userResponse.Nama,
		Token: userResponse.Token,
	}
	return result
}
