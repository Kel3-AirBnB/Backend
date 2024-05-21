package service

import (
	"airbnb/app/configs"
	"airbnb/features/user"
	"airbnb/utils/encrypts"
	"bytes"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type userService struct {
	userData    user.DataInterface
	hashService encrypts.HashInterface
	cfg         *configs.AppConfig
	s3          *s3.S3
}

func New(ud user.DataInterface, hash encrypts.HashInterface, s3 *s3.S3, cfg *configs.AppConfig) user.ServiceInterface {
	return &userService{
		userData:    ud,
		hashService: hash,
		cfg:         cfg,
		s3:          s3,
	}
}

func (u *userService) UploadFileToS3(file io.Reader, fileName string) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	_, err := u.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(u.cfg.S3_BUCKET),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}

	log.Print("[Service Layer]File Name", fileName)
	return fileName, nil
}

// Create(user.Core, io.Reader, string) (string, error)
func (u *userService) Create(input user.Core, file io.Reader, handlerFilename string) (string, error) {
	if input.Nama == "" || input.Email == "" || input.Password == "" {
		return "", errors.New("[validation] nama/email/password tidak boleh kosong")
	}
	if input.Password != "" {
		result, errHash := u.hashService.HashPassword(input.Password)
		if errHash != nil {
			return "", errHash
		}
		input.Password = result
	}

	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)

	photoFileName, errPhoto := u.UploadFileToS3(file, fileName)
	if errPhoto != nil {
		return "", errPhoto
	}

	input.Foto = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.cfg.S3_BUCKET, photoFileName)
	log.Print("[Service Layer] input.Foto ", input.Foto)

	err := u.userData.Insert(input)
	if err != nil {
		return "", err
	}
	return input.Foto, nil
}
