package handler

import "airbnb/features/user"

type UserResponse struct {
	ID           uint   `json:"id,omitempty"`
	Nama         string `json:"name,omitempty"`
	Email        string `json:"email,omitempty"`
	NomorTelepon string `json:"phonenumber,omitempty"`
	Foto         string `json:"profilepicture,omitempty"`
	Token        string `json:"token,omitempty"`
}

func ResponseProfile(userResponse *user.Core) UserResponse {
	result := UserResponse{
		ID:    userResponse.ID,
		Nama:  userResponse.Nama,
		Email: userResponse.Email,
		Foto:  userResponse.Foto,
		Token: userResponse.Token,
	}
	return result
}

func ResponseLogin(userResponse *user.Core) UserResponse {
	result := UserResponse{
		ID:    userResponse.ID,
		Nama:  userResponse.Nama,
		Token: userResponse.Token,
	}
	return result
}
