package service

import (
	"airbnb/app/middlewares"
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
	if input.Password != input.KetikUlangPassword {
		return "", errors.New("[validation] password dan ketik ulang password berbeda")
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
	input.Foto = photoFileName
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
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", u.s3Bucket, aws.StringValue(u.s3.Config.Region), fileName), err
}

func (u *userService) Login(email string, password string) (data *user.Core, token string, err error) {
	data, err = u.userData.SelectByEmail(email)
	if err != nil {
		return nil, "", err
	}
	isLoginValid := u.hashService.CheckPasswordHash(data.Password, password)
	if !isLoginValid {
		return nil, "", errors.New("[validation] password tidak sesuai")
	}
	token, errJWT := middlewares.CreateToken(int(data.ID))
	if errJWT != nil {
		return nil, "", errJWT
	}
	return data, token, nil
}

func (u *userService) GetProfile(id uint) (data *user.Core, err error) {
	if id <= 0 {
		return nil, errors.New("[validation] id not valid")
	}
	// redisIsFound := u.userData.ValidatedRedis()
	// log.Printf("[service layer]Rediskey- %v and err- %v\n", len(redisIsFound), err)
	// if len(redisIsFound) != 0 {
	// 	fmt.Println(redisIsFound)
	// 	log.Println("Masuk ke return redis")
	// 	return u.userData.SelectRedis()
	// } else {
	return u.userData.SelectById(id)
	// }
}

func (u *userService) UpdateById(id uint, input user.Core, file io.Reader, handlerFilename string) (string, error) {
	if id <= 0 {
		return "", errors.New("id not valid")
	} else if input.Nama == "" || input.Email == "" || input.Password == "" {
		return "", errors.New("nama/email/password tidak boleh kosong")
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

	err := u.userData.PutById(id, input)
	if err != nil {
		return "", err
	}
	return input.Foto, nil
}

func (u *userService) Delete(id uint) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	return u.userData.Delete(id)
}
