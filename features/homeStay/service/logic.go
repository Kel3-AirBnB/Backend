package service

import (
	homestay "airbnb/features/homeStay"
	"bytes"
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

type homestayService struct {
	homestayData homestay.DataInterface
	s3           *s3.S3
	s3BucketName string
}

func NewHomestayService(hd homestay.DataInterface, s3 *s3.S3, BucketName string) homestay.ServiceInterface {
	return &homestayService{
		homestayData: hd,
		s3:           s3,
		s3BucketName: BucketName,
	}
}

// GetAll implements homestay.ServiceInterface.
func (h *homestayService) GetAll() ([]homestay.Core, error) {
	return h.homestayData.SelectAll()
}

// GethomeStayid implements homestay.ServiceInterface.
func (h *homestayService) GethomeStayid(id uint) (data *homestay.Core, err error) {
	if id <= 0 {
		return nil, errors.New("[validation] id not valid")
	}
	return h.homestayData.SelectById(id)
}

// UploadFileToS3 implements homestay.ServiceInterface.
func (h *homestayService) UploadFileToS3(file io.Reader, fileName string) (string, error) {
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, file); err != nil {
		return "", err
	}

	_, err := h.s3.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(h.s3BucketName),
		Key:    aws.String(fileName),
		Body:   bytes.NewReader(buf.Bytes()),
		ACL:    aws.String("public-read"),
	})
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("https://%s.s3.%s.amazonaws.com/%s", h.s3BucketName, aws.StringValue(h.s3.Config.Region), fileName), err
}

// Create implements homestay.ServiceInterface.
func (h *homestayService) Create(input homestay.Core, file io.Reader, handlerFilename string) (string, error) {
	timestamp := time.Now().Unix()
	fileName := fmt.Sprintf("%d_%s", timestamp, handlerFilename)
	photoFileName, errPhoto := h.UploadFileToS3(file, fileName)
	if errPhoto != nil {
		return "", errPhoto
	}
	input.Foto = fmt.Sprintf("https://%s.s3.amazonaws.com/%s", h.s3BucketName, photoFileName)
	err := h.homestayData.Insert(input)
	if err != nil {
		return "", err
	}
	return input.Foto, nil
}

// Delete implements homestay.ServiceInterface.
func (h *homestayService) Delete(id uint) error {
	if id <= 0 {
		return errors.New("id not valid")
	}
	return h.homestayData.Delete(id)
}
