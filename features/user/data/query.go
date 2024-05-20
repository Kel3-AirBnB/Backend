package data

import (
	"airbnb/features/user"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func new(db *gorm.DB) user.DataInterface {
	return &userQuery{
		db: db,
	}
}

func (u *userQuery) Insert(input user.Core) error {
	userGorm := UserCoreToUserGorm(input)
	tx := u.db.Create(&userGorm)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
