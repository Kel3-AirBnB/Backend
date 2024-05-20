package service

import (
	"airbnb/features/user"
	"airbnb/utils/encrypts"
	"errors"
)

type userService struct {
	userData    user.DataInterface
	hashService encrypts.HashInterface
}

func New(ud user.DataInterface, hash encrypts.HashInterface) user.ServiceInterface {
	return &userService{
		userData:    ud,
		hashService: hash,
	}
}

func (u *userService) Create(input user.Core) error {
	if input.Nama == "" || input.Email == "" || input.Password == "" {
		return errors.New("[validation] nama/email/password tidak boleh kosong")
	}
	if input.Password != "" {
		result, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return errHash
		}
		input.Password = result
	}
	err := u.userData.Insert(input)
	if err != nil {
		return err
	}
	return nil
}
