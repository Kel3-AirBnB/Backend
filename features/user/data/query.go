package data

import (
	"airbnb/features/user"

	"github.com/aws/aws-sdk-go/service/s3"
	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
	s3 *s3.S3
}

func New(db *gorm.DB, s3 *s3.S3) user.DataInterface {
	return &userQuery{
		db: db,
		s3: s3,
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
