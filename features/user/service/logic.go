package service

import (
	"airbnb/features/user"
	"airbnb/utils/encrypts"
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type userService struct {
	userData    user.DataInterface
	hashService encrypts.HashInterface
	s3          *s3.S3
	s3Bucket    string
}

func New(ud user.DataInterface, hash encrypts.HashInterface, s3 *s3.S3, s3Bucket string) user.ServiceInterface {
	return &userService{
		userData:    ud,
		hashService: hash,
		s3Bucket:    s3Bucket,
		s3:          s3,
	}
}

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
	input.Foto = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", u.s3Bucket, photoFileName)
	err := u.userData.Insert(input)
	if err != nil {
		return "", err
	}
	return input.Foto, nil
}

func (u *userService) UploadFileToS3(file io.Reader, fileName string) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	_, err := u.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(u.s3Bucket),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.s3Bucket, aws.StringValue(u.s3.Config.Region), fileName), nil
}
